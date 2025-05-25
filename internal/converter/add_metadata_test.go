package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMetadata(t *testing.T) {
	tests := map[string]struct {
		input  string
		output string
		title  string
		author string
	}{
		"success": {
			input:  "testdata/001_empty_page/index.html",
			output: "testdata/output/add_metadata_success.pdf",
			title:  "PDF Title",
			author: "Some Author",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c, err := New(test.input, test.output)
			assert.NoError(t, err)

			c.(*converter).title = test.title
			c.(*converter).author = test.author

			c.addMetadata()
			assert.Contains(t, c.(*converter).outputData.String(), test.title)
			assert.Contains(t, c.(*converter).outputData.String(), test.author)
		})
	}
}
