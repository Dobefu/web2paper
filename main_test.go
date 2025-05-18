package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, main, "should not panic")
}
