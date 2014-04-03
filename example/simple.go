package main

import (
	"github.com/aybabtme/gexf"
	"os"
)

var (
	nodes = []struct {
		id    string
		label string
		attr  []gexf.AttrValue
		edges []string
	}{
		{
			label: "Hello",
			attr: []gexf.AttrValue{
				{Title: "foo", Value: true},
			},
			edges: []string{"1"},
		},

		{label: "Hello"},
	}

	attrs = []gexf.Attr{
		{Title: "foo", Type: gexf.Boolean, Default: false},
	}
)

func main() {

	g := gexf.NewGraph()
	if err := g.SetNodeAttrs(attrs); err != nil {
		panic(err)
	}

	for _, node := range nodes {
		id := g.AddNode(node.label, node.attr)
		for _, toID := range node.edges {
			g.AddEdge(id, toID)
		}
	}

	if err := gexf.Encode(os.Stdout, g); err != nil {
		panic(err)
	}

}
