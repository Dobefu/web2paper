package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPage(t *testing.T) {
	conv := converter{}
	assert.Equal(t, len(conv.pages), 0)

	conv.AddPage(PdfSize(PdfSizeA4))
	assert.Equal(t, len(conv.pages), 1)
}
