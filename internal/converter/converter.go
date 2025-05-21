package converter

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"hash"
	"os"
	"time"
)

var (
	version = "2.0"
)

type Obj struct {
	buf    bytes.Buffer
	offset int
}

type Converter interface {
	addObj(data ...string)
	addXrefTable()
	addTrailer()
	addXrefOffset()
	addEOF()
	Convert() (err error)
}

type converter struct {
	Converter
	inputData  []byte
	outputData bytes.Buffer
	outputPath string
	idHasher   hash.Hash

	objs       []Obj
	xrefOffset int
}

func New(input string, output string) (c Converter, err error) {
	data, err := os.ReadFile(input)

	if err != nil {
		return nil, err
	}

	conv := &converter{
		inputData:  data,
		outputData: bytes.Buffer{},
		outputPath: output,
		idHasher:   md5.New(),

		objs:       []Obj{},
		xrefOffset: 0,
	}

	conv.outputData.WriteString(fmt.Sprintf("%%PDF-%s\n", version))

	return conv, nil
}

func (c *converter) addObj(data ...string) {
	obj := bytes.Buffer{}

	obj.WriteString(fmt.Sprintf("%d 0 obj", (len(c.objs) + 1)))
	obj.WriteString("<<")
	obj.WriteString("/Type")

	for _, item := range data {
		obj.WriteString(item)
	}

	obj.WriteString(">>")
	obj.WriteString("endobj\n")

	c.objs = append(c.objs, Obj{
		buf:    obj,
		offset: len(c.outputData.Bytes()),
	})

	c.outputData.Write(obj.Bytes())
}

func (c *converter) addXrefTable() {
	c.xrefOffset = c.outputData.Len()

	c.outputData.WriteString("xref\n")
	c.outputData.WriteString("0 4\n")
	c.outputData.WriteString("0000000000 65535 f \n")

	for _, obj := range c.objs {
		c.outputData.WriteString(fmt.Sprintf("%010d 00000 n \n", obj.offset))
	}
}

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

func (c *converter) addXrefOffset() {
	c.outputData.WriteString("startxref\n")
	c.outputData.WriteString(fmt.Sprintf("%d\n", c.xrefOffset))
}

func (c *converter) addEOF() {
	c.outputData.WriteString("%%EOF\n")
}

func (c *converter) Convert() (err error) {
	c.addObj("/Catalog", "/Pages 2 0 R")
	c.addObj("/Pages", "/Kids[3 0 R]", "/Count 1")
	c.addObj("/Page", "/Parent 2 0 R", "/Resources<<>>", "/MediaBox[0 0 612 792]")
	c.addXrefTable()
	c.addTrailer()
	c.addXrefOffset()

	c.addEOF()

	err = os.WriteFile(c.outputPath, c.outputData.Bytes(), 0644)

	if err != nil {
		return err
	}

	return nil
}
