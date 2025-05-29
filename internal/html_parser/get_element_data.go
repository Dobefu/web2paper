package html_parser

import (
	"github.com/Dobefu/web2paper/internal/fontmap"
	"golang.org/x/net/html"
)

func (p *HtmlParser) GetElementData(pageWidths []float32) (elementData []ElementData) {
	for descendant := range p.dom.Descendants() {
		el := ElementData{
			X:        0,
			Y:        0,
			Content:  "",
			Font:     fontmap.Helvetica,
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

		el.Width = el.Font.GetTextWidth(string(el.Content), el.FontSize)
		el.Height = float32(el.Font.Ascent+el.Font.Descent) * float32(el.FontSize) / 1000

		if el.Content != "" {
			elementData = append(elementData, el)
		}
	}

	return elementData
}
