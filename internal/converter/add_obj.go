package converter

import (
	"bytes"
	"fmt"
)

func (c *converter) addObj(data []string, stream []byte) {
	obj := &bytes.Buffer{}

	fmt.Fprintf(obj, "%d 0 obj", (len(c.objs) + 1))
	obj.WriteString("<<")

	for _, item := range data {
		obj.WriteString(item)
	}

	obj.WriteString(">>")

	if stream != nil {
		obj.WriteString("\nstream\n")
		obj.Write(stream)
		obj.WriteString("\nendstream\n")
	}

	obj.WriteString("endobj\n")

	c.objs = append(c.objs, Obj{
		buf:    *obj,
		offset: len(c.outputData.Bytes()),
	})

	c.outputData.Write(obj.Bytes())
}
