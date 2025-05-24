package converter

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
)

type metadata struct {
	XMLName   xml.Name `xml:"x:xmpmeta"`
	Namespace string   `xml:"xmlns:x,attr"`

	Rdf metadataRdf
}

type metadataRdf struct {
	XMLName   xml.Name `xml:"rdf:RDF"`
	Namespace string   `xml:"xmlns:rdf,attr"`

	Description []metadataDescription
}

type metadataDescription struct {
	XMLName   xml.Name `xml:"rdf:Description"`
	Namespace string   `xml:"xmlns:dc,attr"`
	About     string   `xml:"rdf:about,attr"`

	Format string          `xml:"dc:format"`
	Title  []metadataTitle `xml:"dc:title>rdf:Alt>rdf:li"`
}

type metadataTitle struct {
	XMLName  xml.Name
	Title    string `xml:",innerxml"`
	Language string `xml:"xml:lang,attr"`
}

func (c *converter) addMetadata() (err error) {
	marshalledXml, err := xml.Marshal(
		metadata{
			Namespace: "adobe:ns:meta/",
			Rdf: metadataRdf{
				Namespace: "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
				Description: []metadataDescription{
					{
						Namespace: "http://purl.org/dc/elements/1.1/",
						About:     "",

						Format: "application/pdf",
						Title: []metadataTitle{
							{
								Title:    c.title,
								Language: "x-default",
							},
						},
					},
				},
			},
		},
	)

	if err != nil {
		return err
	}

	metadataXml := bytes.NewBuffer([]byte(strings.TrimRight(xml.Header, "\n")))
	metadataXml.WriteString("<?xpacket begin=\"\uFEFF\" id=\"W5M0MpCehiHzreSzNTczkc9d\"?>")
	metadataXml.Write(marshalledXml)
	metadataXml.WriteString("<?xpacket end=\"w\"?>")

	c.addObj(
		[]string{
			"/Type",
			"/Metadata",
			"/Subtype",
			"/XML",
			fmt.Sprintf("/Length %d", metadataXml.Len()),
		},
		metadataXml.Bytes(),
	)

	return nil
}
