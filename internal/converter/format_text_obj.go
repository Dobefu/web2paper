package converter

import (
	"bytes"
	"fmt"
)

func (c *converter) formatTextObj(
	fontSize int,
	x int,
	y int,
	text string,
) (textObj []byte) {
	buf := bytes.NewBuffer([]byte("BT\n"))   // "Begin Text".
	fmt.Fprintf(buf, "F1 %d Tf\n", fontSize) // Font and font size.
	fmt.Fprintf(buf, "%d %d Td\n", x, y)     // Coordinates.
	fmt.Fprintf(buf, "0 Tc\n")               // Text spacing.
	fmt.Fprintf(buf, "0 Tw\n")               // Text word spacing.
	fmt.Fprintf(buf, "100 Tz\n")             // Text scale % (horizontal).
	fmt.Fprintf(buf, "0 TL\n")               // Text leading.
	fmt.Fprintf(buf, "0 Tr\n")               // Text rendering mode.
	fmt.Fprintf(buf, "0 Ts\n")               // Text rise.
	fmt.Fprintf(buf, "(%s) Tj\n", text)      // Text content.
	buf.WriteString("ET")                    // "End Text".

	return buf.Bytes()
}
