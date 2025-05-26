package converter

import (
	"github.com/Dobefu/web2paper/internal/html_parser"
)

func (c *converter) parseHtml() {
	htmlParser := html_parser.HtmlParser{}
	_ = htmlParser.ParseHtml(c.inputData)
	c.textContent = htmlParser.GetElementData()

	c.title = htmlParser.Metadata.Title
	c.author = htmlParser.Metadata.Author
}
