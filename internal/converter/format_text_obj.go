package converter

import (
	"bytes"
	"fmt"

	"github.com/Dobefu/web2paper/internal/fontmap"
	"github.com/Dobefu/web2paper/internal/html_parser"
)

type renderingMode byte

const (
	renderingModeFill = iota
	renderingModeStroke
	renderingModeFillStroke
	renderingModeNone
	renderingModeFillClip
	renderingModeStrokeClip
	renderingModeFillStrokeClip
	renderingModeClip
)

type _textOptions struct {
	Font          fontmap.Fontmap
	Spacing       int
	WordSpacing   int
	Scale         int
	Leading       int
	RenderingMode renderingMode
	Rise          int
	Halign        html_parser.Align
	Valign        html_parser.Align
}

func textOptionsNew() _textOptions {
	return _textOptions{
		Font:          fontmap.Helvetica,
		Spacing:       0,
		WordSpacing:   0,
		Scale:         100,
		Leading:       0,
		RenderingMode: renderingModeFill,
		Rise:          0,
		Halign:        html_parser.AlignStart,
		Valign:        html_parser.AlignStart,
	}
}

func (c *converter) formatTextObj(
	fontSize int,
	x float32,
	y float32,
	width float32,
	height float32,
	text string,
	options _textOptions,
) (textObj []byte) {
	textOptionsDefaults := textOptionsNew()

	if options.Halign == html_parser.AlignCenter {
		x -= width / 2
	}

	if options.Halign == html_parser.AlignEnd {
		x -= width
	}

	y -= height

	if options.Valign == html_parser.AlignCenter {
		y += height / 2
	}

	if options.Valign == html_parser.AlignEnd {
		y += height
	}

	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "BT\n")                     // "Begin Text".
	fmt.Fprintf(buf, "/F1 %d Tf\n", fontSize)    // Font and font size.
	fmt.Fprintf(buf, "1 0 0 1 %f %f Tm\n", x, y) // Transformation matrix.

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
