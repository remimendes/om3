package daemondata

import (
	"context"
	"runtime"
	"time"

	"github.com/rs/zerolog"

	"opensvc.com/opensvc/core/cluster"
	"opensvc.com/opensvc/daemon/daemonctx"
	"opensvc.com/opensvc/daemon/daemonlogctx"
	"opensvc.com/opensvc/util/callcount"
	"opensvc.com/opensvc/util/jsondelta"
)

type (
	caller interface {
		call(*data)
	}

	data struct {
		committed       *cluster.Status // pending dataset committed
		pending         *cluster.Status
		pendingOps      []jsondelta.Operation // local data pending operations not yet in patchQueue
		patchQueue      patchQueue            // local data patch queue for remotes
		gen             uint64                // gen of local NodeStatus
		mergedFromPeer  gens                  // remote dateset gen merged locally
		mergedOnPeer    gens                  // local dataset gen merged remotely
		remotesNeedFull map[string]bool
		localNode       string
		counterCmd      chan<- interface{}
		log             zerolog.Logger
		pubSub          chan<- interface{}
	}

	gens       map[string]uint64
	patchQueue map[string]jsondelta.Patch
)

func run(ctx context.Context, cmdC <-chan interface{}) {
	counterCmd, cancel := callcount.Start(ctx, idToName)
	defer cancel()
	d := newData(counterCmd)
	d.log = daemonlogctx.Logger(ctx).With().Str("name", "daemon-data").Logger()
	d.log.Info().Msg("starting")
	d.pubSub = daemonctx.DaemonPubSubCmd(ctx)

	defer d.log.Info().Msg("stopped")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			cancel()
			return
		case <-ticker.C:
			d.pending.Monitor.Routines = runtime.NumGoroutine()
		case cmd := <-cmdC:
			if c, ok := cmd.(caller); ok {
				c.call(d)
			} else {
				counterCmd <- idUndef
			}
		}
	}
}
