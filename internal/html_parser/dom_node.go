package html_parser

import "golang.org/x/net/html"

type DomNode struct {
	Tag        string
	Attributes []html.Attribute
	Children   []DomNode
}
