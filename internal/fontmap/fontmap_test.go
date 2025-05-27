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
		expected float32
	}{
		"success": {
			text:     "test",
			fontSize: 24,
			expected: 38.688,
		},
		"unknown glyph": {
			text:     "\u3333",
			fontSize: 24,
			expected: 6.672,
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

func BenchmarkGetTextWidth(b *testing.B) {
	f := Helvetica
	b.ResetTimer()

	for b.Loop() {
		_ = f.GetTextWidth("Test string", 16)
	}
}
