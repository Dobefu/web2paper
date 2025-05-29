package html_parser

import (
	"github.com/Dobefu/web2paper/internal/fontmap"
	"golang.org/x/net/html"
)

func (p *HtmlParser) GetElementData(pageWidths []float32) (elementData []ElementData) {
	for descendant := range p.dom.Descendants() {
		var display Display = DisplayBlock

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
		default:
			continue
		}

		el.Width = el.Font.GetTextWidth(string(el.Content), el.FontSize)
		el.Height = float32(el.Font.Ascent+el.Font.Descent) * float32(el.FontSize) / 1000

		positionElement(elementData, pageWidths[0], &el, display)

		if el.Content != "" {
			elementData = append(elementData, el)
		}
	}

	return elementData
}

func positionElement(
	elementData []ElementData,
	pageWidth float32,
	node *ElementData,
	display Display,
) {
	if len(elementData) > 0 {
		node.X = elementData[len(elementData)-1].X
		node.Y = elementData[len(elementData)-1].Y
	}

	for _, element := range elementData {
		elWidth := element.Width

		if display == DisplayBlock {
			elWidth = pageWidth
		}

		if node.X >= element.X && node.X < element.X+elWidth &&
			node.Y >= element.Y && node.Y < element.Y+element.Height {
			node.X += elWidth
		}

		if node.X+node.Width > pageWidth {
			node.X = 0
			node.Y += element.Height
		}
	}
}
