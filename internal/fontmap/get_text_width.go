package fontmap

func (f *Fontmap) GetTextWidth(text string, fontSize int) (width float32) {
	for _, glyph := range text {
		glyphWidth, ok := f.CharWidths[glyph]

		if !ok {
			glyphWidth = f.CharWidths[' ']
		}

		width += float32(glyphWidth)
	}

	return width * float32(fontSize) / 1000
}
