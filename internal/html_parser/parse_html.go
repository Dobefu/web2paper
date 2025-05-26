package html_parser

import (
	"strings"

	"golang.org/x/net/html"
)

func (p *HtmlParser) ParseHtml(data []byte) {
	reader := strings.NewReader(string(data))
	tokenizer := html.NewTokenizer(reader)

	depth := 0

	for {
		token := tokenizer.Next()

		if token == html.ErrorToken {
			return
		}

		tagName, _ := tokenizer.TagName()

		if token == html.StartTagToken || token == html.SelfClosingTagToken {
			depth++

			attrs := p.getAttrs(tokenizer)
			token = tokenizer.Next()

			if token != html.TextToken {
				continue
			}

			p.processTag(tokenizer, string(tagName), attrs)

			continue
		}

		depth--
	}
}
