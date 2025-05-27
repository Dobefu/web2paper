package html_parser

import (
	"golang.org/x/net/html"
)

func (p *HtmlParser) GetElementData(pageWidths []float32) (elementData []ElementData) {
	for descendant := range p.dom.Descendants() {
		el := ElementData{
			X:        0,
			Y:        0,
			Content:  "",
			FontSize: 12,
			Halign:   AlignStart,
			Valign:   AlignStart,
		}

		if descendant.Type != html.TextNode {
			continue
		}

		switch descendant.Parent.Data {
		case "h1":
			el.Content = descendant.Data
			el.FontSize = 24
		case "p":
			el.Content = descendant.Data
		}

		if el.Content != "" {
			elementData = append(elementData, el)
		}
	}

	return elementData
}
