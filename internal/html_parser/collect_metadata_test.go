package html_parser

import (
	"os"
	"testing"
)

func BenchmarkCollectMetadata(b *testing.B) {
	data, _ := os.ReadFile("../converter/testdata/002_text_page/index.html")

	p := HtmlParser{}
	_ = p.ParseHtml(data)

	for b.Loop() {
		p.collectMetadata()
	}
}
