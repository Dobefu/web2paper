package converter

import (
	"github.com/Dobefu/web2paper/internal/html_parser"
)

func (c *converter) parseHtml() {
	htmlParser := html_parser.HtmlParser{}
	htmlParser.ParseHtml(c.inputData)

	c.title = htmlParser.Metadata.Title
	c.author = htmlParser.Metadata.Author
	c.textContent = htmlParser.TextContent
}
