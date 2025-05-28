package html_parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHtml(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		input string
	}{
		"invalid html": {
			input: "../converter/testdata/000_invalid_html/index.html",
		},
		"empty page": {
			input: "../converter/testdata/001_empty_page/index.html",
		},
		"text page": {
			input: "../converter/testdata/002_text_page/index.html",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			data, err := os.ReadFile(test.input)
			assert.NoError(t, err)

			parser := HtmlParser{}
			err = parser.ParseHtml(data)
			assert.NoError(t, err)
		})
	}
}

func BenchmarkParseHtml(b *testing.B) {
	data, _ := os.ReadFile("../converter/testdata/002_text_page/index.html")

	p := HtmlParser{}
	b.ResetTimer()

	for b.Loop() {
		_ = p.ParseHtml(data)
	}
}
