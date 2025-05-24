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
	AddPage(size PdfSize)

	addObj(data ...string)
	addXrefTable()
	addTrailer()
	addXrefOffset()
	addEOF()
	Convert() (err error)
}

type converter struct {
	Converter

	pages []Page

	inputData  []byte
	outputData bytes.Buffer
	outputPath string

	objs       []Obj
	xrefOffset int
}

func New(input string, output string) (c Converter, err error) {
	data, err := os.ReadFile(input)

	if err != nil {
		return nil, err
	}

	conv := &converter{
		pages: []Page{},

		inputData:  data,
		outputData: bytes.Buffer{},
		outputPath: output,

		objs:       []Obj{},
		xrefOffset: 0,
	}

	conv.outputData.WriteString(fmt.Sprintf("%%PDF-%s\n", version))

	return conv, nil
}
