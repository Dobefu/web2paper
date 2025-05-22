package converter

func (c *converter) addEOF() {
	c.outputData.WriteString("%%EOF\n")
}
