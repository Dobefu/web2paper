package converter

import (
	"bytes"
	"fmt"
)

func (c *converter) addObj(data ...string) {
	obj := bytes.Buffer{}

	obj.WriteString(fmt.Sprintf("%d 0 obj", (len(c.objs) + 1)))
	obj.WriteString("<<")
	obj.WriteString("/Type")

	for _, item := range data {
		obj.WriteString(item)
	}

	obj.WriteString(">>")
	obj.WriteString("endobj\n")

	c.objs = append(c.objs, Obj{
		buf:    obj,
		offset: len(c.outputData.Bytes()),
	})

	c.outputData.Write(obj.Bytes())
}
