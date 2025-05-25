package converter

import (
	"bytes"
	"fmt"

	"github.com/Dobefu/web2paper/internal/fontmap"
)

type renderingMode byte
type align byte

const (
	renderingModeFill = iota
	renderingModeStroke
	renderingModeFillStroke
	renderingModeNone
	renderingModeFillClip
	renderingModeStrokeClip
	renderingModeFillStrokeClip
	renderingModeClip

	alignLeft = iota
	alignCenter
	alignRight
)

type _textOptions struct {
	Spacing       int
	WordSpacing   int
	Scale         int
	Leading       int
	RenderingMode renderingMode
	Rise          int
	Halign        align
}

func textOptionsNew() _textOptions {
	return _textOptions{
		Spacing:       0,
		WordSpacing:   0,
		Scale:         100,
		Leading:       0,
		RenderingMode: renderingModeFill,
		Rise:          0,
		Halign:        alignLeft,
	}
}

func (c *converter) formatTextObj(
	fontSize int,
	x int,
	y int,
	text string,
	options _textOptions,
) (textObj []byte) {
	textOptionsDefaults := textOptionsNew()

	fm := fontmap.Helvetica

	if options.Halign == alignCenter {
		x = x - (fm.GetTextWidth(text, fontSize) / 2)
	}

	if options.Halign == alignRight {
		x = x - fm.GetTextWidth(text, fontSize)
	}

	buf := bytes.NewBuffer([]byte("BT\n"))       // "Begin Text".
	fmt.Fprintf(buf, "F1 %d Tf\n", fontSize)     // Font and font size.
	fmt.Fprintf(buf, "1 0 0 1 %d %d Tm\n", x, y) // Transformation matrix.

	if options.Spacing != textOptionsDefaults.Spacing {
		fmt.Fprintf(buf, "%d Tc\n", options.Spacing)
	}

	if options.WordSpacing != textOptionsDefaults.WordSpacing {
		fmt.Fprintf(buf, "%d Tw\n", options.WordSpacing)
	}

	if options.Scale != textOptionsDefaults.Scale {
		fmt.Fprintf(buf, "%d Tz\n", options.Scale)
	}

	if options.Leading != textOptionsDefaults.Leading {
		fmt.Fprintf(buf, "%d TL\n", options.Leading)
	}

	if options.RenderingMode != textOptionsDefaults.RenderingMode {
		fmt.Fprintf(buf, "%d Tr\n", options.RenderingMode)
	}
	if options.Rise != textOptionsDefaults.Rise {
		fmt.Fprintf(buf, "%d Ts\n", options.Rise)
	}

	fmt.Fprintf(buf, "(%s) Tj\n", text) // Text content.
	fmt.Fprintf(buf, "ET\n")            // "End Text".

	return buf.Bytes()
}
