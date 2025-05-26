package html_parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetElementData(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
	}{
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

			elementData := parser.GetElementData()
			assert.NotEmpty(t, elementData)
		})
	}
}
