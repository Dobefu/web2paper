package converter

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"hash"
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
	addObj(data ...string)
	addXrefTable()
	addTrailer()
	addXrefOffset()
	addEOF()
	Convert() (err error)
}

type converter struct {
	Converter

	size PdfSize

	inputData  []byte
	outputData bytes.Buffer
	outputPath string
	idHasher   hash.Hash

	objs       []Obj
	xrefOffset int
}

func New(
	size PdfSize,
	input string,
	output string,
) (c Converter, err error) {
	data, err := os.ReadFile(input)

	if err != nil {
		return nil, err
	}

	conv := &converter{
		size: size,

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
