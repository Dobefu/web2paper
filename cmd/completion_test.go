package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCompletionCmd(t *testing.T) {
	var tests = map[string]struct {
		shell       string
		isSupported bool
	}{
		"bash": {
			shell:       "bash",
			isSupported: true,
		},
		"zsh": {
			shell:       "zsh",
			isSupported: true,
		},
		"fish": {
			shell:       "fish",
			isSupported: true,
		},
		"powershell": {
			shell:       "powershell",
			isSupported: true,
		},
		"invalid": {
			shell:       "bogus",
			isSupported: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := runCompletionCmd(completionCmd, []string{test.shell})

			if test.isSupported {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, fmt.Sprintf("unsupported shell type %q", test.shell))
			}
		})
	}
}
