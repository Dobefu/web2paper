package cmd

import (
	"testing"

	"github.com/Dobefu/web2paper/internal/converter"
	"github.com/stretchr/testify/assert"
)

type mockConverter struct {
	converter.Converter
	convertErr error
}

func (c *mockConverter) Convert() (err error) {
	return c.convertErr
}

func TestRunConvertCmd(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args                []string
		shouldSucceed       bool
		converterNewErr     error
		converterConvertErr error
	}{
		"success": {
			args:                []string{"-i", "in.html", "-o", "out.pdf"},
			converterNewErr:     nil,
			converterConvertErr: nil,
			shouldSucceed:       true,
		},
		"converter constructor err": {
			args:                []string{"-i", "in.html", "-o", "out.pdf"},
			converterNewErr:     assert.AnError,
			converterConvertErr: nil,
			shouldSucceed:       false,
		},
		"converter convert err": {
			args:                []string{"-i", "in.html", "-o", "out.pdf"},
			converterNewErr:     nil,
			converterConvertErr: assert.AnError,
			shouldSucceed:       false,
		},
		"input missing": {
			args:                []string{"-o", "out.pdf"},
			converterNewErr:     nil,
			converterConvertErr: nil,
			shouldSucceed:       false,
		},
		"output missing": {
			args:                []string{"-i", "in.html"},
			converterNewErr:     nil,
			converterConvertErr: nil,
			shouldSucceed:       false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {

			converterNew = func(_ converter.PdfSize, _ string, _ string) (converter.Converter, error) {
				return &mockConverter{convertErr: test.converterConvertErr}, test.converterNewErr
			}

			defer func() { converterNew = converter.New }()

			cmd := NewConvertCmd()
			cmd.SetArgs(append([]string{"convert"}, test.args...))

			err := cmd.Execute()

			if test.shouldSucceed {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)

				if test.converterNewErr != nil {
					assert.EqualError(t, err, test.converterNewErr.Error())
				}

				if test.converterConvertErr != nil {
					assert.EqualError(t, err, test.converterConvertErr.Error())
				}
			}
		})
	}
}
