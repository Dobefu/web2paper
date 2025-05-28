package converter

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func (c *converter) addTrailer() {
	var pdfId [32]byte
	var revisionId [32]byte

	hasher := md5.New()

	hasher.Write(c.outputData.Bytes())
	hex.Encode(pdfId[:], hasher.Sum(nil))

	_, _ = fmt.Fprintf(hasher, "%d", c.creationDate.UnixNano())
	hex.Encode(revisionId[:], hasher.Sum(nil))
	hasher.Reset()

	fmt.Fprintf(c.outputData, "trailer")
	fmt.Fprintf(c.outputData, "<</Root 1 0 R")
	fmt.Fprintf(c.outputData, "/Size %d", (len(c.objs) + 1))
	fmt.Fprintf(c.outputData, "/Info %d 0 R", (len(c.objs)))
	fmt.Fprintf(c.outputData, "/ID[(%s)(%s)]", pdfId, revisionId)
	fmt.Fprintf(c.outputData, ">>\n")
}
