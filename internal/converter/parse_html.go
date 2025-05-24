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

		if token != html.StartTagToken && token != html.EndTagToken {
			continue
		}

		tagName, _ := tokenizer.TagName()

		if token == html.StartTagToken {
			depth++
			token = tokenizer.Next()

			if token != html.TextToken {
				continue
			}

			processTag(c, tokenizer, string(tagName))

			continue
		}

		depth--
	}
}

func processTag(c *converter, tokenizer *html.Tokenizer, tagName string) {
	switch tagName {
	case "title":
		c.title = tokenizer.Token().Data
	}
}
