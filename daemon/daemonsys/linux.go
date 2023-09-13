//go:build linux

package daemonsys

import (
	"context"
	"os"

	sddaemon "github.com/coreos/go-systemd/daemon"
	"github.com/coreos/go-systemd/v22/dbus"

	"github.com/opensvc/om3/util/command"
)

type (
	// T handle dbus.Conn for systemd
	T struct {
		conn *dbus.Conn
	}
)

const (
	name = "opensvc-agent.service"
)

// Activated detects if opensvc unit is activated
func (t *T) Activated(ctx context.Context) (bool, error) {
	prop, err := t.conn.GetUnitPropertyContext(ctx, name, "ActiveState")
	if err != nil {
		return false, err
	}
	if prop == nil {
		return false, nil
	}
	return prop.Value.String() == "\"active\"", nil
}

// CalledFromManager detects if current process as been launched by systemd
func (t *T) CalledFromManager() bool {
	return os.Getenv("INVOCATION_ID") != ""
}

// Close closes systemd dbus connection
func (t *T) Close() error {
	if t.conn != nil {
		t.conn.Close()
	}
	return nil
}

// Defined verify if opensvc systemd unit exists
func (t *T) Defined(ctx context.Context) (bool, error) {
	units, err := t.conn.ListUnitsByNamesContext(ctx, []string{name})
	if err != nil {
		return false, err
	}
	for _, v := range units {
		if v.LoadState == "loaded" {
			return true, nil
		}
	}
	return false, nil
}

// NotifyWatchdog sends watch dog notify to systemd
func (t *T) NotifyWatchdog() (bool, error) {
	if t.conn == nil {
		return false, nil
	}
	return sddaemon.SdNotify(false, sddaemon.SdNotifyWatchdog)
}

// Restart restarts the opensvc systemd unit
//
// restart calls systemd-run systemctl restart opensvc-agent. This allows
// the command to be attached on another control group and prevent systemd
// warnings during 'om daemon restart' such as:
//
//	systemd[1]: Stopping OpenSVC agent...
//	systemd[1]: opensvc-agent.service: Succeeded.
//	systemd[1]: Stopped OpenSVC agent.
//	systemd[1]: opensvc-agent.service: Found left-over process 2899690 (om) in control group while starting unit. Ignoring.
//	systemd[1]: This usually indicates unclean termination of a previous run, or service implementation deficiencies.
//	systemd[1]: opensvc-agent.service: Found left-over process 2899697 (systemctl) in control group while starting unit. Ignoring.
//	systemd[1]: This usually indicates unclean termination of a previous run, or service implementation deficiencies.
//	systemd[1]: Starting OpenSVC agent...
//	systemd[1]: Started OpenSVC agent.
func (t *T) Restart() error {
	return command.New(
		command.WithName("systemd-run"),
		command.WithVarArgs("systemctl", "restart", name),
	).Run()
}

// Start starts the opensvc systemd unit
func (t *T) Start(ctx context.Context) error {
	c := make(chan string)
	_, err := t.conn.StartUnitContext(ctx, name, "replace", c)
	if err != nil {
		return err
	}
	<-c
	return nil
}

// Stop stops the opensvc systemd unit
func (t *T) Stop(ctx context.Context) error {
	c := make(chan string)
	_, err := t.conn.StopUnitContext(ctx, name, "replace", c)
	if err != nil {
		return err
	}
	<-c
	return nil
}

// New provides a connected object to dbus systemd that implement following interfaces:
//
//	Activated(ctx context.Context) (bool, error)
//	CalledFromManager() bool
//	Close() error
//	NotifyWatchdog() (bool, error)
//	Start(ctx context.Context) error
//	Stop(ctx context.Context) error
func New(ctx context.Context) (*T, error) {
	c, err := dbus.NewSystemdConnectionContext(ctx)
	if err != nil {
		return nil, err
	}
	return &T{conn: c}, nil
}
