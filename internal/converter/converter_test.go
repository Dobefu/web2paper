package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverterNew(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input         string
		output        string
		shouldSucceed bool
	}{
		"success": {
			input:         "testdata/001_empty_page/index.html",
			output:        "testdata/001_empty_page/converter_new_success.pdf",
			shouldSucceed: true,
		},
		"nonexistent file": {
			input:         "bogus.html",
			output:        "out.pdf",
			shouldSucceed: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			_, err := New(PdfSize(PdfSizeA4), test.input, test.output)

			if test.shouldSucceed {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestConverterConvert(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input         string
		output        string
		shouldSucceed bool
	}{
		"success": {
			input:         "testdata/001_empty_page/index.html",
			output:        "testdata/001_empty_page/converter_convert_success.pdf",
			shouldSucceed: true,
		},
		"invalid location": {
			input:         "testdata/001_empty_page/index.html",
			output:        "/bogus",
			shouldSucceed: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			pdfConverter, err := New(PdfSize(PdfSizeA4), test.input, test.output)
			assert.NoError(t, err)

			err = pdfConverter.Convert()

			if test.shouldSucceed {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
