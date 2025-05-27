package converter

import (
	"fmt"
)

func (c *converter) addXrefOffset() {
	fmt.Fprintf(c.outputData, "startxref\n")
	fmt.Fprintf(c.outputData, "%d\n", c.xrefOffset)
}
