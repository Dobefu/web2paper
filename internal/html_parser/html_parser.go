package html_parser

import "golang.org/x/net/html"

type htmlParser interface {
	ParseHtml(data []byte) (err error)
	GetElementData() (elementData []ElementData)

	collectMetadata()
}

type HtmlParser struct {
	htmlParser

	Metadata Metadata

	dom *html.Node
}
