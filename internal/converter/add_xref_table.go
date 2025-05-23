package converter

import (
	"fmt"
)

func (c *converter) addXrefTable() {
	c.xrefOffset = c.outputData.Len()

	c.outputData.WriteString("xref\n")
	c.outputData.WriteString(fmt.Sprintf("0 %d\n", (len(c.objs) + 1)))
	c.outputData.WriteString("0000000000 65535 f \n")

	for _, obj := range c.objs {
		c.outputData.WriteString(fmt.Sprintf("%010d 00000 n \n", obj.offset))
	}
}
