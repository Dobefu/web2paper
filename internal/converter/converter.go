package converter

import (
	"bytes"
	"fmt"
	"os"
)

var (
	version = "2.0"
)

type Obj struct {
	buf    bytes.Buffer
	offset int
}

type Converter interface {
	AddObj(data ...string)
	AddXrefTable()
	Convert() (err error)
}

type converter struct {
	Converter
	inputData  []byte
	outputData bytes.Buffer
	outputPath string

	objs []Obj
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

		objs: []Obj{},
	}

	conv.outputData.WriteString(fmt.Sprintf("%%PDF-%s\n", version))

	return conv, nil
}

func (c *converter) AddObj(data ...string) {
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

func (c *converter) AddXrefTable() {
	c.outputData.WriteString("xref\n")
	c.outputData.WriteString("0 4\n")
	c.outputData.WriteString("0000000000 65535 f \n")

	for _, obj := range c.objs {
		c.outputData.WriteString(fmt.Sprintf("%010d 00000 n \n", obj.offset))
	}
}

func (c *converter) Convert() (err error) {
	c.AddObj("/Catalog", "/Pages 2 0 R")
	c.AddObj("/Pages", "/Kids[3 0 R]", "/Count 1")
	c.AddObj("/Page", "/Parent 2 0 R", "/Resources<<>>", "/MediaBox[0 0 612 792]")
	c.AddXrefTable()
	c.outputData.WriteString("trailer")
	c.outputData.WriteString("<</Root 1 0 R")
	c.outputData.WriteString("/Size 4")
	c.outputData.WriteString("/ID[(1234567890123456)(1234567890123456)]")
	c.outputData.WriteString(">>\n")
	c.outputData.WriteString("startxref\n")
	c.outputData.WriteString("178\n")
	c.outputData.WriteString("%%EOF\n")

	err = os.WriteFile(c.outputPath, c.outputData.Bytes(), 0644)

	if err != nil {
		return err
	}

	return nil
}
