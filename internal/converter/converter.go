package converter

import (
	"bytes"
	"os"
)

type Converter interface {
	Convert() (err error)
}

type converter struct {
	Converter
	inputData  []byte
	outputData bytes.Buffer
	outputPath string
}

func New(input string, output string) (Converter, error) {
	data, err := os.ReadFile(input)

	if err != nil {
		return nil, err
	}

	return &converter{
		inputData:  data,
		outputData: bytes.Buffer{},
		outputPath: output,
	}, nil
}

func (c *converter) Convert() (err error) {
	c.outputData.WriteString("%PDF-2.0\n")
	c.outputData.WriteString("1 0 obj")
	c.outputData.WriteString("<</Type/Catalog")
	c.outputData.WriteString("/Pages 2 0 R")
	c.outputData.WriteString(">>")
	c.outputData.WriteString("endobj\n")
	c.outputData.WriteString("2 0 obj")
	c.outputData.WriteString("<</Type/Pages")
	c.outputData.WriteString("/Kids[3 0 R]")
	c.outputData.WriteString("/Count 1")
	c.outputData.WriteString(">>")
	c.outputData.WriteString("endobj\n")
	c.outputData.WriteString("3 0 obj")
	c.outputData.WriteString("<</Type/Page")
	c.outputData.WriteString("/Parent 2 0 R")
	c.outputData.WriteString("/Resources<<>>")
	c.outputData.WriteString("/MediaBox[0 0 612 792]")
	c.outputData.WriteString(">>")
	c.outputData.WriteString("endobj\n")
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
