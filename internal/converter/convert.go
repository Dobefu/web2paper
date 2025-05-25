package converter

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (c *converter) Convert() (err error) {
	c.creationDate = time.Now()

	c.parseHtml()

	c.outputData.Write([]byte{'%', '\200', '\201', '\202', '\203', '\n'})

	c.addObj([]string{
		"/Type",
		"/Catalog",
		fmt.Sprintf("/Pages %d 0 R", (len(c.objs) + 2)),
		fmt.Sprintf("/Metadata %d 0 R", (len(c.pages) + 3)),
	}, nil)

	c.addObj([]string{
		"/Type",
		"/Pages",
		renderPagesKidsString(c),
		fmt.Sprintf("/Count %d", len(c.pages)),
	}, nil)

	for _, page := range c.pages {
		c.addObj([]string{
			"/Type",
			"/Page",
			"/Parent 2 0 R",
			"/Resources<<>>",
			fmt.Sprintf("/MediaBox[0 0 %.2f %.2f]", page.Size.Width, page.Size.Height),
		}, nil)
	}

	c.addMetadata()
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

func renderPagesKidsString(c *converter) (output string) {
	buf := []string{}

	for pageNum := range len(c.pages) {
		buf = append(buf, fmt.Sprintf("%d 0 R", (pageNum+len(c.objs)+2)))
	}

	return fmt.Sprintf("/Kids[%s]", strings.Join(buf, " "))
}
