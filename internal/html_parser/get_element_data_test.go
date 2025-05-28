package html_parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetElementData(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input      string
		pageWidths []float32
	}{
		"text page": {
			input:      "../converter/testdata/002_text_page/index.html",
			pageWidths: []float32{768},
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

			elementData := parser.GetElementData(test.pageWidths)
			assert.NotEmpty(t, elementData)
		})
	}
}

func BenchmarkGetElementData(b *testing.B) {
	data, _ := os.ReadFile("../converter/testdata/002_text_page/index.html")

	p := HtmlParser{}
	_ = p.ParseHtml(data)
	b.ResetTimer()

	for b.Loop() {
		p.GetElementData([]float32{768})
	}
}
