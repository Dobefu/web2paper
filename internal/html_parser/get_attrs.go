package html_parser

import "golang.org/x/net/html"

func (p *HtmlParser) getAttrs(tokenizer *html.Tokenizer) (attrs map[string]string) {
	attrs = make(map[string]string)

	for {
		key, val, moreAttr := tokenizer.TagAttr()

		if string(key) == "" {
			break
		}

		attrs[string(key)] = string(val)

		if !moreAttr {
			break
		}
	}

	return attrs
}
