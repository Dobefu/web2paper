package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	tests := map[string]struct {
		input         string
		output        string
		shouldSucceed bool
	}{
		"success": {
			input:         "testdata/001_empty_page/index.html",
			output:        "testdata/output/convert_success.pdf",
			shouldSucceed: true,
		},
		"edge cases": {
			input:         "testdata/000_invalid_html/index.html",
			output:        "testdata/output/convert_edge_cases.pdf",
			shouldSucceed: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c, err := New(test.input, test.output)
			assert.NoError(t, err)

			c.AddPage(PdfSize(PdfSizeA4))
			err = c.Convert()

			if test.shouldSucceed {
				assert.NoError(t, err, "should not return an error")
			} else {
				assert.Error(t, err, "should return an error")
			}
		})
	}
}
