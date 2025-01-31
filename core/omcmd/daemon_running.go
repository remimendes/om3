package commands

import (
	"os"

	"github.com/opensvc/om3/core/client"
	"github.com/opensvc/om3/daemon/daemoncmd"
)

type (
	CmdDaemonRunning struct {
		OptsGlobal
		NodeSelector string
	}
)

func (t *CmdDaemonRunning) Run() error {
	cli, err := client.New(client.WithURL(t.Server))
	if err != nil {
		return err
	}
	dCli := daemoncmd.New(cli)
	dCli.SetNode(t.NodeSelector)
	if !dCli.Running() {
		os.Exit(1)
	}
	return nil
}
