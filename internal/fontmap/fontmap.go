package fontmap

type fontmap struct {
	CharHeight int
	CharWidths map[rune]int
}

func (f *fontmap) GetTextWidth(text string, fontSize int) (width int) {
	for _, glyph := range text {
		glyphWidth, ok := f.CharWidths[glyph]

		if !ok {
			glyphWidth = f.CharWidths[' ']
		}

		width += glyphWidth
	}

	return width * fontSize / 1000
}
