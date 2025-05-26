package html_parser

import (
	"golang.org/x/net/html"
)

func (p *HtmlParser) GetElementData() (elementData []string) {
	for descendant := range p.dom.Descendants() {
		if descendant.Type == html.TextNode {
			switch descendant.Parent.Data {
			case "h1", "p":
				elementData = append(elementData, descendant.Data)
			}
		}
	}

	return elementData
}
