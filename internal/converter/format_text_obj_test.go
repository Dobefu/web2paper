package converter

import (
	"testing"

	"github.com/Dobefu/web2paper/internal/html_parser"
	"github.com/stretchr/testify/assert"
)

func TestFormatTextObj(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input         string
		output        string
		text          string
		init          func(textOptions *_textOptions)
		shouldContain string
	}{
		"success": {
			input:         "testdata/001_empty_page/index.html",
			output:        "out.pdf",
			text:          "Test text",
			init:          nil,
			shouldContain: "",
		},
		"custom spacing": {
			input:         "testdata/001_empty_page/index.html",
			output:        "out.pdf",
			text:          "Test text",
			init:          func(textOptions *_textOptions) { textOptions.Spacing = 99 },
			shouldContain: "99 Tc",
		},
		"custom word spacing": {
			input:         "testdata/001_empty_page/index.html",
			output:        "out.pdf",
			text:          "Test text",
			init:          func(textOptions *_textOptions) { textOptions.WordSpacing = 99 },
			shouldContain: "99 Tw",
		},
		"custom scale": {
			input:         "testdata/001_empty_page/index.html",
			output:        "out.pdf",
			text:          "Test text",
			init:          func(textOptions *_textOptions) { textOptions.Scale = 99 },
			shouldContain: "99 Tz",
		},
		"custom leading": {
			input:         "testdata/001_empty_page/index.html",
			output:        "out.pdf",
			text:          "Test text",
			init:          func(textOptions *_textOptions) { textOptions.Leading = 99 },
			shouldContain: "99 TL",
		},
		"custom rendering mode": {
			input:         "testdata/001_empty_page/index.html",
			output:        "out.pdf",
			text:          "Test text",
			init:          func(textOptions *_textOptions) { textOptions.RenderingMode = renderingModeStroke },
			shouldContain: "1 Tr",
		},
		"custom rise": {
			input:         "testdata/001_empty_page/index.html",
			output:        "out.pdf",
			text:          "Test text",
			init:          func(textOptions *_textOptions) { textOptions.Rise = 99 },
			shouldContain: "99 Ts",
		},
		"halign center": {
			input:         "testdata/001_empty_page/index.html",
			output:        "out.pdf",
			text:          "Test text",
			init:          func(textOptions *_textOptions) { textOptions.Halign = html_parser.AlignCenter },
			shouldContain: "",
		},
		"halign right": {
			input:         "testdata/001_empty_page/index.html",
			output:        "out.pdf",
			text:          "Test text",
			init:          func(textOptions *_textOptions) { textOptions.Halign = html_parser.AlignEnd },
			shouldContain: "",
		},
		"valign center": {
			input:         "testdata/001_empty_page/index.html",
			output:        "out.pdf",
			text:          "Test text",
			init:          func(textOptions *_textOptions) { textOptions.Valign = html_parser.AlignCenter },
			shouldContain: "",
		},
		"valign right": {
			input:         "testdata/001_empty_page/index.html",
			output:        "out.pdf",
			text:          "Test text",
			init:          func(textOptions *_textOptions) { textOptions.Valign = html_parser.AlignEnd },
			shouldContain: "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			c, err := New(test.input, test.output)
			assert.NoError(t, err)

			textOptions := textOptionsNew()

			if test.init != nil {
				test.init(&textOptions)
			}

			textObj := c.formatTextObj(24, 0, 0, test.text, textOptions)
			assert.Contains(t, string(textObj), "BT")
			assert.Contains(t, string(textObj), test.text)
			assert.Contains(t, string(textObj), test.shouldContain)
			assert.Contains(t, string(textObj), "ET")
		})
	}
}

func BenchmarkFormatTextObj(b *testing.B) {
	c, _ := New("testdata/001_empty_page/index.html", "/dev/null")
	textOptions := textOptionsNew()

	b.ResetTimer()

	for b.Loop() {
		_ = c.formatTextObj(24, 0, 0, "Test text", textOptions)
	}
}
