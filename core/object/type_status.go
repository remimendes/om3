package object

import (
	"strings"

	"opensvc.com/opensvc/config"
	"opensvc.com/opensvc/core/provisioned"
	"opensvc.com/opensvc/core/status"
	"opensvc.com/opensvc/util/render/tree"
)

type (
	// Status is a composite extract of different parts of
	// the cluster status.
	Status struct {
		Path      Path                        `json:"path"`
		Compat    bool                        `json:"compat"`
		Object    AggregatedStatus            `json:"service"`
		Instances map[string]InstanceStates   `json:"instances"`
		Parents   map[string]AggregatedStatus `json:"parents,omitempty"`
		Children  map[string]AggregatedStatus `json:"children,omitempty"`
		Slaves    map[string]AggregatedStatus `json:"slaves,omitempty"`
	}

	// AggregatedStatus contains the object states obtained via
	// aggregation of all instances states.
	AggregatedStatus struct {
		Avail       status.T      `json:"avail,omitempty"`
		Overall     status.T      `json:"overall,omitempty"`
		Frozen      string        `json:"frozen,omitempty"`
		Placement   string        `json:"placement,omitempty"`
		Provisioned provisioned.T `json:"provisioned,omitempty"`
	}

	// InstanceStates groups config and status of the object as seen by the daemon.
	InstanceStates struct {
		Config InstanceConfig `json:"config"`
		Status InstanceStatus `json:"status"`
	}
)

// Render returns a human friendy string representation of the type instance.
func (t Status) Render() string {
	tree := t.Tree()
	return tree.Render()
}

// Tree returns a tree loaded with the type instance.
func (t Status) Tree() *tree.Tree {
	tree := tree.New()
	t.LoadTreeNode(tree.Head())
	return tree
}

// LoadTreeNode add the tree nodes representing the type instance into another.
func (t Status) LoadTreeNode(head *tree.Node) {
	head.AddColumn().AddText(t.Path.String()).SetColor(config.Node.Color.Bold)
	head.AddColumn()
	head.AddColumn().AddText(t.Object.Avail.ColorString())
	head.AddColumn().AddText(t.DescString())
	instances := head.AddNode()
	instances.AddColumn().AddText("instances")
	for nodename, data := range t.Instances {
		data.Status.Nodename = nodename
		data.Status.Path = t.Path
		n := instances.AddNode()
		data.Status.LoadTreeNode(n)
	}
	t.loadTreeNodeParents(head)
	t.loadTreeNodeChildren(head)
	t.loadTreeNodeSlaves(head)
}

func (t Status) loadTreeNodeParents(head *tree.Node) {
	if len(t.Parents) == 0 {
		return
	}
	n := head.AddNode()
	n.AddColumn().AddText("parents")
	for p, data := range t.Parents {
		pNode := n.AddNode()
		pNode.AddColumn().AddText(p).SetColor(config.Node.Color.Bold)
		pNode.AddColumn()
		pNode.AddColumn().AddText(data.Avail.ColorString())
	}
}

func (t Status) loadTreeNodeChildren(head *tree.Node) {
	if len(t.Children) == 0 {
		return
	}
	n := head.AddNode()
	n.AddColumn().AddText("children")
	for p, data := range t.Children {
		pNode := n.AddNode()
		pNode.AddColumn().AddText(p).SetColor(config.Node.Color.Bold)
		pNode.AddColumn()
		pNode.AddColumn().AddText(data.Avail.ColorString())
	}
}

func (t Status) loadTreeNodeSlaves(head *tree.Node) {
	if len(t.Slaves) == 0 {
		return
	}
	n := head.AddNode()
	n.AddColumn().AddText("slaves")
	for p, data := range t.Slaves {
		pNode := n.AddNode()
		pNode.AddColumn().AddText(p).SetColor(config.Node.Color.Bold)
		pNode.AddColumn()
		pNode.AddColumn().AddText(data.Avail.ColorString())
	}
}

func (t Status) DescString() string {
	l := make([]string, 0)
	l = append(l, t.Object.Overall.ColorString())
	return strings.Join(l, " ")
}

// NewObjectStatus allocates and return a struct to host an objet full state dataset.
func NewObjectStatus() *Status {
	t := &Status{}
	t.Instances = make(map[string]InstanceStates)
	t.Parents = make(map[string]AggregatedStatus)
	t.Children = make(map[string]AggregatedStatus)
	t.Slaves = make(map[string]AggregatedStatus)
	return t
}
