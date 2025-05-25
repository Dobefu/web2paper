package fontmap

type fontmap struct {
	CharHeight int
	CharWidths map[rune]int
}

func (f *fontmap) GetTextWidth(text string, fontSize int) (width int) {
	for _, glyph := range text {
		width += f.CharWidths[glyph]
	}

	return width * fontSize / 1000
}
