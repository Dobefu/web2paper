package converter

import (
	"bytes"
	"fmt"
	"os"
)

var (
	version = "2.0"
)

type Converter interface {
	AddObj(data ...string)
	Convert() (err error)
}

type converter struct {
	Converter
	inputData  []byte
	outputData bytes.Buffer
	outputPath string

	objCount uint
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

		objCount: 0,
	}

	conv.outputData.WriteString(fmt.Sprintf("%%PDF-%s\n", version))

	return conv, nil
}

func (c *converter) AddObj(data ...string) {
	c.objCount += 1

	c.outputData.WriteString(fmt.Sprintf("%d 0 obj", c.objCount))
	c.outputData.WriteString("<<")
	c.outputData.WriteString("/Type")

	for _, item := range data {
		c.outputData.WriteString(item)
	}

	c.outputData.WriteString(">>")
	c.outputData.WriteString("endobj\n")
}

func (c *converter) Convert() (err error) {
	c.AddObj("/Catalog", "/Pages 2 0 R")
	c.AddObj("/Pages", "/Kids[3 0 R]", "/Count 1")
	c.AddObj("/Page", "/Parent 2 0 R", "/Resources<<>>", "/MediaBox[0 0 612 792]")
	c.outputData.WriteString("xref\n")
	c.outputData.WriteString("0 4\n")
	c.outputData.WriteString("0000000000 65535 f \n")
	c.outputData.WriteString("0000000009 00000 n \n")
	c.outputData.WriteString("0000000052 00000 n \n")
	c.outputData.WriteString("0000000101 00000 n \n")
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
