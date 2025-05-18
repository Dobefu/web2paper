package converter

import "testing"

func TestConverter(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			New(test.input, test.output)
		})
	}
}
