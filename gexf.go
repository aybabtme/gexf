package gexf

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"time"
)

func Encode(w io.Writer, g *Graph) error {
	gx := gexf{
		Namespace: "http://www.gexf.net/1.2draft",
		Version:   "1.2",
		Meta: &meta{
			LastModified: time.Now().Format("2006-01-02"),
			Creator:      "webscale!",
			Desc:         "so fast!",
		},
		Graph: g,
	}

	data, err := xml.MarshalIndent(gx, "", "    ")
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(data)
	_, err = io.Copy(w, buf)
	return err
	// return xml.NewEncoder(w).Encode(gx)
}

type Attr struct {
	Title   string
	Type    GEXFType
	Default interface{}
}

type AttrValue struct {
	Title string
	Value interface{}
}

type Graph struct {
	XMLName xml.Name `xml:"graph"`

	Mode     string `xml:"mode,attr,omitempty"`
	EdgeType string `xml:"defaultedgetype,attr"`

	Attrs *attributes `xml:"attributes"`
	Nodes []node      `xml:"nodes>node"`
	Edges []edge      `xml:"edges>edge"`

	attrCount     int
	attrTitleToID map[string]string
}

func NewGraph() *Graph {
	return &Graph{
		Mode:          "static",
		EdgeType:      "directed",
		attrTitleToID: make(map[string]string),
	}
}

func (g *Graph) SetNodeAttrs(attrs []Attr) error {
	g.Attrs = &attributes{
		Class: "node",
	}
	for _, a := range attrs {
		if _, ok := g.attrTitleToID[a.Title]; ok {
			return fmt.Errorf("attr '%s' defined multiple times", a.Title)
		}

		id := len(g.attrTitleToID)
		attr := attribute{
			ID:      strconv.Itoa(id),
			Title:   a.Title,
			Type:    string(a.Type),
			Default: a.Default,
		}

		g.Attrs.Attrs = append(g.Attrs.Attrs, attr)
		g.attrTitleToID[attr.Title] = attr.ID
	}
	return nil
}

func (g *Graph) AddNode(label string, attr []AttrValue) string {
	n := node{
		ID:    strconv.Itoa(len(g.Nodes)),
		Label: label,
	}

	var values []attrValue
	for _, a := range attr {
		av := attrValue{
			For:   g.attrTitleToID[a.Title],
			Value: a.Value,
		}
		values = append(values, av)
	}

	if len(values) > 0 {
		n.Attr = &values
	}

	g.Nodes = append(g.Nodes, n)
	return n.ID
}

func (g *Graph) AddEdge(from, to string) {
	e := edge{
		ID:     strconv.Itoa(len(g.Edges)),
		Source: from,
		Target: to,
	}
	g.Edges = append(g.Edges, e)
}
