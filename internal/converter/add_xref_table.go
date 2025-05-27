package converter

import (
	"fmt"
)

func (c *converter) addXrefTable() {
	c.xrefOffset = c.outputData.Len()

	fmt.Fprintf(c.outputData, "xref\n")
	fmt.Fprintf(c.outputData, "0 %d\n", (len(c.objs) + 1))
	fmt.Fprintf(c.outputData, "0000000000 65535 f \n")

	for _, obj := range c.objs {
		fmt.Fprintf(c.outputData, "%010d 00000 n \n", obj.offset)
	}
}
