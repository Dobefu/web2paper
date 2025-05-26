package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input         string
		output        string
		shouldSucceed bool
	}{
		"edge cases": {
			input:         "testdata/000_invalid_html/index.html",
			output:        "testdata/output/convert_edge_cases.pdf",
			shouldSucceed: true,
		},
		"success 001": {
			input:         "testdata/001_empty_page/index.html",
			output:        "testdata/output/convert_success_001.pdf",
			shouldSucceed: true,
		},
		"success 002": {
			input:         "testdata/002_text_page/index.html",
			output:        "testdata/output/convert_success_002.pdf",
			shouldSucceed: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

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
