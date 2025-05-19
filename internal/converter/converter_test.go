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
			output:        "out.pdf",
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

			_, err := New(test.input, test.output)

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
			output:        "out.pdf",
			shouldSucceed: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			pdfConverter, err := New(test.input, test.output)
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
