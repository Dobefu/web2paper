package cmd

import (
	"testing"

	"github.com/Dobefu/web2paper/internal/converter"
	"github.com/stretchr/testify/assert"
)

func TestRunConvertCmd(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args                  []string
		err                   error
		shouldCreateConverter bool
	}{
		"success": {
			args:                  []string{"-i", "in.html", "-o", "out.pdf"},
			err:                   nil,
			shouldCreateConverter: true,
		},
		"error": {
			args:                  []string{"-i", "in.html", "-o", "out.pdf"},
			err:                   assert.AnError,
			shouldCreateConverter: false,
		},
		"input missing": {
			args:                  []string{"-o", "out.pdf"},
			err:                   nil,
			shouldCreateConverter: false,
		},
		"output missing": {
			args:                  []string{"-i", "in.html"},
			err:                   nil,
			shouldCreateConverter: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			isConverterCreated := false

			converterNew = func(_ string, _ string) (converter.Converter, error) {
				if test.err == nil {
					isConverterCreated = true
				}

				return nil, test.err
			}

			defer func() { converterNew = converter.New }()

			cmd := NewConvertCmd()
			cmd.SetArgs(append([]string{"convert"}, test.args...))

			err := cmd.Execute()

			if test.shouldCreateConverter {
				assert.NoError(t, err)
				assert.True(t, isConverterCreated, "converter should be created")
			} else {
				assert.Error(t, err)
				assert.False(t, isConverterCreated, "converter should not be created")

				if test.err != nil {
					assert.EqualError(t, err, test.err.Error())
				}
			}
		})
	}
}
