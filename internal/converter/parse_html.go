package converter

import (
	"strings"

	"golang.org/x/net/html"
)

func (c *converter) parseHtml() {
	reader := strings.NewReader(string(c.inputData))
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

			attrs := getAttrs(tokenizer)
			token = tokenizer.Next()

			if token != html.TextToken {
				continue
			}

			processTag(c, tokenizer, string(tagName), attrs)

			continue
		}

		depth--
	}
}

func getAttrs(tokenizer *html.Tokenizer) (attrs map[string]string) {
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

func processTag(
	c *converter,
	tokenizer *html.Tokenizer,
	tagName string,
	attrs map[string]string,
) {
	switch tagName {
	case "title":
		c.title = tokenizer.Token().Data
	case "meta":
		name, hasName := attrs["name"]
		content, hasContent := attrs["content"]

		if !hasName || !hasContent {
			break
		}

		processMetaValue(c, name, content)
	}
}

func processMetaValue(c *converter, key string, value string) {
	switch key {
	case "author":
		c.author = value
	}
}
