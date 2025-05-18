package converter

import "testing"

func TestConverter(t *testing.T) {
	tests := map[string]struct {
		input  string
		output string
	}{
		"success": {
			input:  "in.html",
			output: "out.pdf",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			New(test.input, test.output)
		})
	}
}
