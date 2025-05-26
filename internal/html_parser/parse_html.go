package html_parser

import (
	"strings"

	"golang.org/x/net/html"
)

func (p *HtmlParser) ParseHtml(data []byte) (err error) {
	reader := strings.NewReader(string(data))
	rootNode, _ := html.Parse(reader)

	p.dom = rootNode
	p.collectMetadata()

	return nil
}
