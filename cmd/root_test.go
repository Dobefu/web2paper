package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootExecute(t *testing.T) {
	assert.NotPanics(t, Execute, "should not panic")
}

func TestRootInitConfig(t *testing.T) {
	assert.NotPanics(t, initConfig, "should not panic")
}
