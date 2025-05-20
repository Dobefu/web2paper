package converter

import (
	"os"
)

type Converter interface {
	Convert() (err error)
}

type converter struct {
	Converter
	data       []byte
	outputPath string
}

func New(input string, output string) (Converter, error) {
	data, err := os.ReadFile(input)

	if err != nil {
		return nil, err
	}

	return &converter{
		data:       data,
		outputPath: output,
	}, nil
}

func (c *converter) Convert() (err error) {
	data := []byte{}
	data = append(data, []byte("%PDF-2.0\n")...)
	data = append(data, []byte("1 0 obj")...)
	data = append(data, []byte("<</Type/Catalog")...)
	data = append(data, []byte("/Pages 2 0 R")...)
	data = append(data, []byte(">>")...)
	data = append(data, []byte("endobj\n")...)
	data = append(data, []byte("2 0 obj")...)
	data = append(data, []byte("<</Type/Pages")...)
	data = append(data, []byte("/Kids[3 0 R]")...)
	data = append(data, []byte("/Count 1")...)
	data = append(data, []byte(">>")...)
	data = append(data, []byte("endobj\n")...)
	data = append(data, []byte("3 0 obj")...)
	data = append(data, []byte("<</Type/Page")...)
	data = append(data, []byte("/Parent 2 0 R")...)
	data = append(data, []byte("/Resources<<>>")...)
	data = append(data, []byte("/MediaBox[0 0 612 792]")...)
	data = append(data, []byte(">>")...)
	data = append(data, []byte("endobj\n")...)
	data = append(data, []byte("xref\n")...)
	data = append(data, []byte("0 4\n")...)
	data = append(data, []byte("0000000000 65535 f \n")...)
	data = append(data, []byte("0000000009 00000 n \n")...)
	data = append(data, []byte("0000000052 00000 n \n")...)
	data = append(data, []byte("0000000101 00000 n \n")...)
	data = append(data, []byte("trailer")...)
	data = append(data, []byte("<</Root 1 0 R")...)
	data = append(data, []byte("/Size 4")...)
	data = append(data, []byte("/ID[(1234567890123456)(1234567890123456)]")...)
	data = append(data, []byte(">>\n")...)
	data = append(data, []byte("startxref\n")...)
	data = append(data, []byte("178\n")...)
	data = append(data, []byte("%%EOF\n")...)

	err = os.WriteFile(c.outputPath, data, 0644)
	return nil
}
