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
	return nil
}
