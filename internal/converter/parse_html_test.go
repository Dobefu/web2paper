package converter

import "testing"

func BenchmarkParseHtml(b *testing.B) {
	c, _ := New("testdata/001_empty_page/index.html", "/dev/null")
	b.ResetTimer()

	for b.Loop() {
		c.parseHtml()
	}
}
