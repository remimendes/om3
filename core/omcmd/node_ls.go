package commands

import (
	"github.com/opensvc/om3/core/client"
	"github.com/opensvc/om3/core/nodeselector"
	"github.com/opensvc/om3/core/output"
	"github.com/opensvc/om3/core/rawconfig"
)

type (
	CmdNodeLs struct {
		OptsGlobal
		NodeSelector string
	}
)

func (t *CmdNodeLs) Run() error {
	var (
		err      error
		selector string
	)
	c, err := client.New(client.WithURL(t.Server))
	if err != nil {
		return err
	}
	if t.NodeSelector == "" {
		selector = "*"
	} else {
		selector = t.NodeSelector
	}
	nodes, err := nodeselector.New(selector, nodeselector.WithClient(c)).Expand()
	if err != nil {
		return err
	}
	output.Renderer{
		Output: t.Output,
		Color:  t.Color,
		Data:   nodes,
		HumanRenderer: func() string {
			s := ""
			for _, e := range nodes {
				s += e + "\n"
			}
			return s
		},
		Colorize: rawconfig.Colorize,
	}.Print()
	return nil
}
