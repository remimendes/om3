package discover

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/opensvc/om3/core/client"
	"github.com/opensvc/om3/core/cluster"
	"github.com/opensvc/om3/core/freeze"
	"github.com/opensvc/om3/core/instance"
	"github.com/opensvc/om3/core/naming"
	"github.com/opensvc/om3/core/node"
	"github.com/opensvc/om3/core/object"
	"github.com/opensvc/om3/daemon/daemonenv"
	"github.com/opensvc/om3/daemon/daemonlogctx"
	"github.com/opensvc/om3/daemon/icfg"
	"github.com/opensvc/om3/daemon/msgbus"
	"github.com/opensvc/om3/daemon/remoteconfig"
	"github.com/opensvc/om3/util/file"
	"github.com/opensvc/om3/util/hostname"
	"github.com/opensvc/om3/util/pubsub"
)

var (
	// SubscriptionQueueSizeCfg is size of "discover.cfg" subscription
	SubscriptionQueueSizeCfg = 30000
)

func (d *discover) startSubscriptions() *pubsub.Subscription {
	bus := pubsub.BusFromContext(d.ctx)
	sub := bus.Sub("discover.cfg", pubsub.WithQueueSize(SubscriptionQueueSizeCfg))
	sub.AddFilter(&msgbus.InstanceConfigUpdated{})
	sub.AddFilter(&msgbus.InstanceConfigDeleted{})
	sub.AddFilter(&msgbus.ConfigFileUpdated{})
	sub.AddFilter(&msgbus.ClusterConfigUpdated{})
	sub.AddFilter(&msgbus.ObjectStatusUpdated{})
	sub.AddFilter(&msgbus.ObjectStatusDeleted{})
	sub.Start()
	return sub
}

func (d *discover) cfg(started chan<- bool) {
	d.log.Info().Msg("daemon: discover: cfg started")
	defer d.log.Info().Msg("daemon: discover: cfg stopped")
	defer func() {
		d.log.Debug().Msg("daemon: discover: flushing the command bus message queue")
		defer d.log.Debug().Msg("daemon: discover: flushed the command bus message queue")
		t := time.NewTicker(d.drainDuration)
		defer t.Stop()
		for {
			select {
			case <-d.ctx.Done():
				return
			case <-t.C:
				return
			case <-d.cfgCmdC:
			}
		}
	}()
	sub := d.startSubscriptions()
	defer func() {
		if err := sub.Stop(); err != nil {
			d.log.Error().Err(err).Msgf("daemon: discover: subscription stop: %s", err)
		}
	}()
	if last := cluster.ConfigData.Get(); last != nil {
		msg := &msgbus.ClusterConfigUpdated{Value: *last}
		d.onClusterConfigUpdated(msg)
	}
	started <- true
	for {
		select {
		case <-d.ctx.Done():
			return
		case i := <-sub.C:
			switch c := i.(type) {
			case *msgbus.InstanceConfigUpdated:
				d.onInstanceConfigUpdated(c)
			case *msgbus.InstanceConfigDeleted:
				d.onInstanceConfigDeleted(c)
			case *msgbus.ConfigFileUpdated:
				d.onConfigFileUpdated(c)
			case *msgbus.ClusterConfigUpdated:
				d.onClusterConfigUpdated(c)
			case *msgbus.ObjectStatusUpdated:
				d.onObjectStatusUpdated(c)
			case *msgbus.ObjectStatusDeleted:
				d.onObjectStatusDeleted(c)
			}
		case i := <-d.cfgCmdC:
			switch c := i.(type) {
			case *msgbus.RemoteFileConfig:
				d.onRemoteConfigFetched(c)
			case *msgbus.InstanceConfigManagerDone:
				d.onMonConfigDone(c)
			default:
				d.log.Error().Msgf("daemon: discover: unsupported command bus message type: %#v", i)
			}
		case nfo := <-d.objectList.InfoC:
			d.log.Info().Msg("daemon: discover: " + nfo)
		case err := <-d.objectList.ErrC:
			d.log.Info().Err(err).Msgf("daemon: discover: %s", err)
		case nfo := <-d.nodeList.InfoC:
			d.log.Info().Msg("daemon: discover: " + nfo)
		case err := <-d.nodeList.ErrC:
			d.log.Info().Err(err).Msgf("daemon: discover: %s", err)
		}
	}
}

func (d *discover) onClusterConfigUpdated(c *msgbus.ClusterConfigUpdated) {
	d.clusterConfig = c.Value
	d.nodeList.Add(c.NodesAdded...)
	d.nodeList.Del(c.NodesRemoved...)
}

func (d *discover) onObjectStatusUpdated(c *msgbus.ObjectStatusUpdated) {
	d.objectList.Add(c.Path.String())
}

func (d *discover) onObjectStatusDeleted(c *msgbus.ObjectStatusDeleted) {
	d.objectList.Del(c.Path.String())
}

func (d *discover) onConfigFileUpdated(c *msgbus.ConfigFileUpdated) {
	if c.Path.Kind == naming.KindInvalid {
		// may be node.conf
		return
	}
	s := c.Path.String()
	mtime := file.ModTime(c.File)
	if mtime.IsZero() {
		d.log.Info().Msgf("daemon: discover: config file %s mtime is zero", c.File)
		return
	}
	if _, ok := d.cfgMTime[s]; !ok {
		if err := icfg.Start(d.ctx, c.Path, c.File, d.cfgCmdC); err != nil {
			return
		}
	}
	d.cfgMTime[s] = mtime
}

// cmdLocalConfigDeleted starts a new icfg when a local configuration file exists
func (d *discover) onMonConfigDone(c *msgbus.InstanceConfigManagerDone) {
	filename := c.File
	p := c.Path
	s := p.String()

	delete(d.cfgMTime, s)
	mtime := file.ModTime(filename)
	if mtime.IsZero() {
		return
	}
	if err := icfg.Start(d.ctx, p, filename, d.cfgCmdC); err != nil {
		return
	}
	d.cfgMTime[s] = mtime
}

func (d *discover) onInstanceConfigUpdated(c *msgbus.InstanceConfigUpdated) {
	if c.Node == d.localhost {
		return
	}
	d.onRemoteConfigUpdated(c.Path, c.Node, c.Value)
}

func (d *discover) onRemoteConfigUpdated(p naming.Path, node string, remoteInstanceConfig instance.Config) {
	s := p.String()

	localUpdated := file.ModTime(p.ConfigFile())

	// Never drop local cluster config, ignore remote config older that local
	if !p.Equal(naming.Cluster) && remoteInstanceConfig.UpdatedAt.After(localUpdated) && !d.inScope(&remoteInstanceConfig) {
		d.cancelFetcher(s)
		cfgFile := p.ConfigFile()
		if file.Exists(cfgFile) {
			d.log.Info().Msgf("daemon: discover: remove local config %s (localnode not in node %s config scope)", s, node)
			if err := os.Remove(cfgFile); err != nil {
				d.log.Debug().Err(err).Msgf("daemon: discover: remove %s: %s", cfgFile, err)
			}
		}
		return
	}
	if mtime, ok := d.cfgMTime[s]; ok {
		if !remoteInstanceConfig.UpdatedAt.After(mtime) {
			// our version is more recent than remote one
			return
		}
	} else if !remoteInstanceConfig.UpdatedAt.After(localUpdated) {
		// Not yet started icfg, but file exists
		return
	}
	if remoteFetcherUpdated, ok := d.fetcherUpdated[s]; ok {
		// fetcher in progress for s, verify if new fetcher is required
		if remoteInstanceConfig.UpdatedAt.After(remoteFetcherUpdated) {
			d.log.Warn().Msgf("daemon: discover: cancel pending remote cfg fetcher, a more recent %s config is available on node %s", s, node)
			d.cancelFetcher(s)
		} else {
			// let running fetcher does its job
			return
		}
	}
	d.log.Info().Msgf("daemon: discover: fetch %s config from node %s", s, node)
	d.fetchConfigFromRemote(p, node, remoteInstanceConfig)
}

func (d *discover) onInstanceConfigDeleted(c *msgbus.InstanceConfigDeleted) {
	if c.Node == "" || c.Node == d.localhost {
		return
	}
	s := c.Path.String()
	if fetchFrom, ok := d.fetcherFrom[s]; ok {
		if fetchFrom == c.Node {
			d.log.Info().Msgf("daemon: discover: cancel pending remote cfg fetcher, instance %s@%s is no longer present", s, c.Node)
			d.cancelFetcher(s)
		}
	}
}

func (d *discover) onRemoteConfigFetched(c *msgbus.RemoteFileConfig) {

	freezeIfOrchestrateHA := func(confFile string) error {
		if !c.Freeze {
			return nil
		}
		if err := freeze.Freeze(c.Path.FrozenFile()); err != nil {
			d.log.Error().Err(err).Msgf("daemon: discover: can't freeze instance before installing %s config fetched from node %s: %s", c.Path, c.Node, err)
			return err
		}
		d.log.Info().Msgf("daemon: discover: freeze instance before installing %s config fetched from node %s", c.Path, c.Node)
		return nil
	}

	defer d.cancelFetcher(c.Path.String())
	select {
	case <-c.Ctx.Done():
		c.Err <- nil
	default:
		confFile := c.Path.ConfigFile()
		if err := freezeIfOrchestrateHA(confFile); err != nil {
			c.Err <- err
			return
		}
		if err := os.Rename(c.File, confFile); err != nil {
			d.log.Error().Err(err).Msgf("daemon: discover: can't install %s config fetched from node %s to %s: %s", c.Path, c.Node, confFile, err)
			c.Err <- err
		} else {
			d.log.Info().Msgf("daemon: discover: install %s config fetched from node %s", c.Path, c.Node)
		}
		c.Err <- nil
	}
}

func (d *discover) inScope(cfg *instance.Config) bool {
	localhost := d.localhost
	for _, s := range cfg.Scope {
		if s == localhost {
			return true
		}
	}
	return false
}

func (d *discover) cancelFetcher(s string) {
	if cancel, ok := d.fetcherCancel[s]; ok {
		d.log.Debug().Msgf("daemon: discover: cancelFetcher %s", s)
		cancel()
		peer := d.fetcherFrom[s]
		delete(d.fetcherCancel, s)
		delete(d.fetcherNodeCancel[peer], s)
		delete(d.fetcherUpdated, s)
		delete(d.fetcherFrom, s)
	}
}

func (d *discover) fetchConfigFromRemote(p naming.Path, peer string, remoteInstanceConfig instance.Config) {
	s := p.String()
	if n, ok := d.fetcherFrom[s]; ok {
		d.log.Error().Msgf("daemon: discover: fetcher already in progress for %s from node %s", s, n)
		return
	}
	ctx, cancel := context.WithCancel(d.ctx)
	d.fetcherCancel[s] = cancel
	d.fetcherFrom[s] = peer
	d.fetcherUpdated[s] = remoteInstanceConfig.UpdatedAt
	if _, ok := d.fetcherNodeCancel[peer]; ok {
		d.fetcherNodeCancel[peer][s] = cancel
	} else {
		d.fetcherNodeCancel[peer] = make(map[string]context.CancelFunc)
	}

	peerPort := fmt.Sprintf("%d", daemonenv.HttpPort)
	if lsnr := node.LsnrData.Get(peer); lsnr != nil {
		peerPort = lsnr.Port
	}
	cli, err := d.newDaemonClient(peer, peerPort)
	if err != nil {
		d.log.Error().Msgf("daemon: discover: can't create newDaemonClient to fetch %s from node %s", p, peer)
		return
	}
	go fetch(ctx, cli, p, peer, d.cfgCmdC, remoteInstanceConfig)
}

func (d *discover) newDaemonClient(node, port string) (*client.T, error) {
	// TODO add WithRootCa to avoid send password to wrong url ?
	return client.New(
		client.WithURL(daemonenv.UrlHttpNodeAndPort(node, port)),
		client.WithUsername(hostname.Hostname()),
		client.WithPassword(d.clusterConfig.Secret()),
		client.WithCertificate(daemonenv.CertChainFile()),
	)
}

func fetch(ctx context.Context, cli *client.T, p naming.Path, node string, cmdC chan<- any, remoteInstanceConfig instance.Config) {
	id := p.String() + "@" + node
	log := daemonlogctx.Logger(ctx).With().Str("_pkg", "cfg.fetch").Str("id", id).Logger()

	tmpFilename, updated, err := remoteconfig.FetchObjectFile(cli, p)
	if err != nil {
		log.Info().Err(err).Msgf("daemon: discover: FetchObjectFile %s: %s", id, err)
		return
	}
	defer func() {
		log.Debug().Msgf("daemon: discover: done fetcher routine for instance %s@%s", p, node)
		_ = os.Remove(tmpFilename)
	}()
	configure, err := object.NewConfigurer(p, object.WithConfigFile(tmpFilename), object.WithVolatile(true))
	if err != nil {
		log.Error().Err(err).Msgf("daemon: discover: configure error for %s: %s", p, err)
		return
	}
	nodes, err := configure.Config().Referrer.Nodes()
	if err != nil {
		log.Error().Err(err).Msgf("daemon: discover: nodes eval error for %s: %s", p, err)
		return
	}
	validScope := false
	for _, n := range nodes {
		if n == hostname.Hostname() {
			validScope = true
			break
		}
	}
	if !validScope {
		log.Info().Msgf("daemon: discover: invalid scope %s", nodes)
		return
	}
	var freeze bool
	if remoteInstanceConfig.Orchestrate == "ha" && len(remoteInstanceConfig.Scope) > 1 {
		freeze = true
	}
	select {
	case <-ctx.Done():
		log.Info().Msgf("daemon: discover: abort fetch config %s", id)
		return
	default:
		err := make(chan error)
		cmdC <- &msgbus.RemoteFileConfig{
			Path:      p,
			Node:      node,
			File:      tmpFilename,
			Freeze:    freeze,
			UpdatedAt: updated,
			Ctx:       ctx,
			Err:       err,
		}
		<-err
	}
}
