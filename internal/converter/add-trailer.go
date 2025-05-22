package converter

import (
	"fmt"
	"time"
)

func (c *converter) addTrailer() {
	c.idHasher.Write(c.outputData.Bytes())
	pdfId := fmt.Sprintf("%x", c.idHasher.Sum(nil))[:16]

	_, _ = fmt.Fprintf(c.idHasher, "%d", time.Now().UnixNano())
	revisionId := fmt.Sprintf("%x", c.idHasher.Sum(nil))[:16]
	c.idHasher.Reset()

	c.outputData.WriteString("trailer")
	c.outputData.WriteString("<</Root 1 0 R")
	c.outputData.WriteString(fmt.Sprintf("/Size %d", (len(c.objs) + 1)))
	c.outputData.WriteString(fmt.Sprintf("/ID[(%s)(%s)]", pdfId, revisionId))
	c.outputData.WriteString(">>\n")
}
