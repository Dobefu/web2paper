package html_parser

import (
	"golang.org/x/net/html"
)

func (p *HtmlParser) processTag(
	tokenizer *html.Tokenizer,
	tagName string,
	attrs map[string]string,
) {
	switch tagName {
	case "title":
		p.Metadata.Title = tokenizer.Token().Data
	case "meta":
		name, hasName := attrs["name"]
		content, hasContent := attrs["content"]

		if !hasName || !hasContent {
			break
		}

		p.processMetaValue(name, content)
	case "html":
	case "head":
	case "body":
		break
	default:
		p.TextContent = append(p.TextContent, tokenizer.Token().Data)
	}
}
