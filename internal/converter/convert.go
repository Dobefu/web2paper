package converter

import (
	"fmt"
	"os"
)

func (c *converter) Convert() (err error) {
	c.parseHtml()

	c.addObj("/Type", "/Catalog", "/Pages 2 0 R")
	c.addObj("/Type", "/Pages", "/Kids[3 0 R]", fmt.Sprintf("/Count %d", len(c.pages)))

	for _, page := range c.pages {
		c.addObj(
			"/Type",
			"/Page",
			"/Parent 2 0 R",
			"/Resources<<>>",
			fmt.Sprintf("/MediaBox[0 0 %.2f %.2f]", page.Size.Width, page.Size.Height),
		)
	}

	c.addXrefTable()
	c.addTrailer()
	c.addXrefOffset()

	c.addEOF()

	err = os.WriteFile(c.outputPath, c.outputData.Bytes(), 0644)

	if err != nil {
		return err
	}

	return nil
}
