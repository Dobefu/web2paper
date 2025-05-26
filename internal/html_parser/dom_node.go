package html_parser

type DomNode struct {
	Tag        string
	Attributes []Attribute
	Children   []DomNode
}
