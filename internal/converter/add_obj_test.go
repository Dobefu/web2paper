package converter

import (
	"testing"
)

func BenchmarkAddObj(b *testing.B) {
	c, _ := New("testdata/001_empty_page/index.html", "/dev/null")
	b.ResetTimer()

	for b.Loop() {
		c.addObj([]string{
			"/Type",
			"/Font",
			"/Subtype",
			"/Type1",
			"/BaseFont",
			"/Helvetica",
		}, []byte("Test stream data"))
	}
}
