package converter

import (
	"crypto/md5"
	"fmt"
)

func (c *converter) addTrailer() {
	hasher := md5.New()

	hasher.Write(c.outputData.Bytes())
	pdfId := fmt.Sprintf("%x", hasher.Sum(nil))[:16]

	_, _ = fmt.Fprintf(hasher, "%d", c.creationDate.UnixNano())
	revisionId := fmt.Sprintf("%x", hasher.Sum(nil))[:16]
	hasher.Reset()

	c.outputData.WriteString("trailer")
	c.outputData.WriteString("<</Root 1 0 R")
	c.outputData.WriteString(fmt.Sprintf("/Size %d", (len(c.objs) + 1)))
	c.outputData.WriteString(fmt.Sprintf("/Info %d 0 R", (len(c.objs))))
	c.outputData.WriteString(fmt.Sprintf("/ID[(%s)(%s)]", pdfId, revisionId))
	c.outputData.WriteString(">>\n")
}
