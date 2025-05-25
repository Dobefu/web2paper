package fontmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTextWidth(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		text     string
		fontSize int
		expected int
	}{
		"success": {
			text:     "test",
			fontSize: 24,
			expected: 38,
		},
		"unknown glyph": {
			text:     "\u3333",
			fontSize: 24,
			expected: 6,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			width := Helvetica.GetTextWidth(test.text, test.fontSize)
			assert.Equal(t, width, test.expected)
		})
	}
}
