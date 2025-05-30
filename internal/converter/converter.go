package converter

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/Dobefu/web2paper/internal/html_parser"
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

	formatTextObj(
		fontSize int,
		x float32,
		y float32,
		width float32,
		height float32,
		text string,
		options _textOptions,
	) (textObj []byte)
}

type converter struct {
	Converter

	title       string
	author      string
	pages       []Page
	elementData []html_parser.ElementData

	inputData  []byte
	outputData *bytes.Buffer
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
		title:  "",
		author: "",
		pages:  []Page{},

		inputData:  data,
		outputData: &bytes.Buffer{},
		outputPath: output,

		objs:         []Obj{},
		xrefOffset:   0,
		creationDate: time.Time{},
	}

	fmt.Fprintf(conv.outputData, "%%PDF-%s\n", version)

	return conv, nil
}
