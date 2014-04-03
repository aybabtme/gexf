package gexf

import (
	"encoding/xml"
)

type attrValue struct {
	XMLName xml.Name    `xml:"attvalue"`
	For     string      `xml:"for,attr"`
	Value   interface{} `xml:"value,attr"`
}

type attribute struct {
	XMLName xml.Name    `xml:"attribute"`
	ID      string      `xml:"id,attr"`
	Title   string      `xml:"title,attr"`
	Type    string      `xml:"type,attr"`
	Default interface{} `xml:"default"`
}

type attributes struct {
	XMLName xml.Name    `xml:"attributes"`
	Class   string      `xml:"class,attr"`
	Attrs   []attribute `xml:"attribute"`
}

type node struct {
	XMLName xml.Name     `xml:"node"`
	ID      string       `xml:"id,attr"`
	Label   string       `xml:"label,attr"`
	Attr    *[]attrValue `xml:"attvalues>attvalue"`
}

type edge struct {
	XMLName xml.Name `xml:"edge"`
	ID      string   `xml:"id,attr"`
	Source  string   `xml:"source,attr"`
	Target  string   `xml:"target,attr"`
}

type meta struct {
	LastModified string `xml:"lastmodifieddate,attr"`
	Creator      string `xml:"creator"`
	Desc         string `xml:"description"`
}

type gexf struct {
	XMLName   xml.Name `xml:"gexf"`
	Namespace string   `xml:"xmlns,attr"`
	Version   string   `xml:"version,attr"`
	Meta      *meta    `xml:"meta"`
	Graph     *Graph   `xml:"graph"`
}
