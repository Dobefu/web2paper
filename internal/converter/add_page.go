package converter

func (c *converter) AddPage(size PdfSize) {
	c.pages = append(c.pages, Page{
		Size: size,
	})
}
