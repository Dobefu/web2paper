package html_parser

func (p *HtmlParser) collectMetadata() {
	for descendant := range p.dom.Descendants() {
		if descendant.Data == "title" && descendant.FirstChild != nil {
			p.Metadata.Title = descendant.FirstChild.Data
		}

		if descendant.Data == "meta" {
			name := ""
			content := ""

			for _, attr := range descendant.Attr {
				switch attr.Key {
				case "name":
					name = attr.Val
				case "content":
					content = attr.Val
				}
			}

			if name == "" || content == "" {
				continue
			}

			switch name {
			case "author":
				p.Metadata.Author = content
			}
		}
	}
}
