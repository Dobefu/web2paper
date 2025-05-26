package html_parser

func (p *HtmlParser) processMetaValue(key string, value string) {
	switch key {
	case "author":
		p.Metadata.Author = value
	}
}
