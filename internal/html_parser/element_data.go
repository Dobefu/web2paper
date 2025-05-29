package html_parser

import "github.com/Dobefu/web2paper/internal/fontmap"

type ElementData struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	Content  string
	Font     fontmap.Fontmap
	FontSize int
	Halign   Align
	Valign   Align
}
