package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootExecute(t *testing.T) {
	oldOsArgs := os.Args
	defer func() { os.Args = oldOsArgs }()

	tests := map[string]struct {
		args []string
	}{
		"success": {
			args: nil,
		},
		"invalid args": {
			args: []string{"web2paper", "--bogus"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.args != nil {
				os.Args = test.args
			}

			assert.NotPanics(t, Execute, "should not panic")
		})
	}

}
