package html_parser

import "golang.org/x/net/html"

type htmlParser interface {
	ParseHtml(data []byte)
	getAttrs(tokenizer *html.Tokenizer) (attrs map[string]string)
	processTag(tokenizer *html.Tokenizer, tagName string, attrs map[string]string)
	processMetaValue(key string, value string)
}

type HtmlParser struct {
	htmlParser

	Metadata    Metadata
	TextContent []string

	dom DomNode
}
