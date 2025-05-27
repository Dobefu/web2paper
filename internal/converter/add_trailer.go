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

	fmt.Fprintf(c.outputData, "trailer")
	fmt.Fprintf(c.outputData, "<</Root 1 0 R")
	fmt.Fprintf(c.outputData, "/Size %d", (len(c.objs) + 1))
	fmt.Fprintf(c.outputData, "/Info %d 0 R", (len(c.objs)))
	fmt.Fprintf(c.outputData, "/ID[(%s)(%s)]", pdfId, revisionId)
	fmt.Fprintf(c.outputData, ">>\n")
}
