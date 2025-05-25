package converter

import (
	"bytes"
	"fmt"
)

func (c *converter) formatTextObj(fontSize int, x int, y int, text string) (textObj []byte) {
	buf := bytes.NewBuffer([]byte("BT\n"))
	fmt.Fprintf(buf, "F1 %d Tf\n", fontSize)
	fmt.Fprintf(buf, "%d %d Td\n", x, y)
	fmt.Fprintf(buf, "(%s) Tj\n", text)
	buf.WriteString("ET")

	return buf.Bytes()
}
