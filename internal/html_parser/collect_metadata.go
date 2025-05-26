package html_parser

import "golang.org/x/net/html"

func (p *HtmlParser) collectMetadata() {
	for descendant := range p.dom.Descendants() {
		if descendant.Data == "title" && descendant.FirstChild != nil {
			p.Metadata.Title = descendant.FirstChild.Data
		}

		if descendant.Data != "meta" {
			continue
		}

		name, content := getAttributesNameAndContent(descendant)

		if name == "" || content == "" {
			continue
		}

		switch name {
		case "author":
			p.Metadata.Author = content
		}
	}
}

func getAttributesNameAndContent(descendant *html.Node) (name string, content string) {
	for _, attr := range descendant.Attr {
		switch attr.Key {
		case "name":
			name = attr.Val
		case "content":
			content = attr.Val
		}
	}

	return name, content
}
