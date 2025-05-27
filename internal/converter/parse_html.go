package converter

import (
	"github.com/Dobefu/web2paper/internal/html_parser"
)

func (c *converter) parseHtml() {
	htmlParser := html_parser.HtmlParser{}
	_ = htmlParser.ParseHtml(c.inputData)

	var pageWidths []float32

	for _, page := range c.pages {
		pageWidths = append(pageWidths, page.Size.Width)
	}

	c.elementData = htmlParser.GetElementData(pageWidths)

	c.title = htmlParser.Metadata.Title
	c.author = htmlParser.Metadata.Author
}
