package converter

import (
	"bytes"
	"fmt"
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
	AddPage(size PdfSize)
	Convert() (err error)

	parseHtml()
	addObj(data []string, stream []byte)
	addMetadata()
	addXrefTable()
	addTrailer()
	addXrefOffset()
	addEOF()
}

type converter struct {
	Converter

	title string
	pages []Page

	inputData  []byte
	outputData bytes.Buffer
	outputPath string

	objs         []Obj
	xrefOffset   int
	creationDate time.Time
}

func New(input string, output string) (c Converter, err error) {
	data, err := os.ReadFile(input)

	if err != nil {
		return nil, err
	}

	conv := &converter{
		title: "",
		pages: []Page{},

		inputData:  data,
		outputData: bytes.Buffer{},
		outputPath: output,

		objs:         []Obj{},
		xrefOffset:   0,
		creationDate: time.Time{},
	}

	conv.outputData.WriteString(fmt.Sprintf("%%PDF-%s\n", version))

	return conv, nil
}
