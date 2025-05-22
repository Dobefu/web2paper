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
			output:        "testdata/001_empty_page/convert_success.pdf",
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
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
