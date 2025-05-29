package fontmap

type Fontmap struct {
	Ascent     int
	Descent    int
	CapHeight  int
	XHeight    int
	CharWidths map[rune]int
}
