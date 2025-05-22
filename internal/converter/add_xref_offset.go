package converter

import (
	"fmt"
)

func (c *converter) addXrefOffset() {
	c.outputData.WriteString("startxref\n")
	c.outputData.WriteString(fmt.Sprintf("%d\n", c.xrefOffset))
}
