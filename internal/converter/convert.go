package converter

import (
	"fmt"
	"os"
)

func (c *converter) Convert() (err error) {
	c.addObj("/Catalog", "/Pages 2 0 R")
	c.addObj("/Pages", "/Kids[3 0 R]", "/Count 1")
	c.addObj("/Page", "/Parent 2 0 R", "/Resources<<>>", fmt.Sprintf("/MediaBox[0 0 %.2f %.2f]", c.size.Width, c.size.Height))
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
