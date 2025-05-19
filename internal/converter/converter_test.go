package converter

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverter(t *testing.T) {
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
			t.Log(runtime.Caller(0))
			_, err := New(test.input, test.output)

			if test.shouldSucceed {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
