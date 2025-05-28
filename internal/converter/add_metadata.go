package converter

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
)

const (
	producer = "Web2Paper (https://github.com/Dobefu/web2paper)"
)

type metadata struct {
	XMLName   xml.Name `xml:"x:xmpmeta"`
	Namespace string   `xml:"xmlns:x,attr"`

	Rdf metadataRdf
}

type metadataRdf struct {
	XMLName   xml.Name `xml:"rdf:RDF"`
	Namespace string   `xml:"xmlns:rdf,attr"`

	Description []metadataDescription `xml:"rdf:Description"`
}

type metadataDescription struct {
	NamespaceDc  string `xml:"xmlns:dc,attr,omitempty"`
	NamespacePdf string `xml:"xmlns:pdf,attr,omitempty"`
	About        string `xml:"rdf:about,attr"`

	Format   string          `xml:"dc:format,omitempty"`
	Title    []metadataTitle `xml:"dc:title,omitempty"`
	Producer string          `xml:"pdf:Producer,omitempty"`
}

type metadataTitle struct {
	Languages []metadataTitleAlt `xml:"rdf:Alt,omitempty"`
}

type metadataTitleAlt struct {
	Li metadataTitleLi `xml:"rdf:li,omitempty"`
}

type metadataTitleLi struct {
	Title    string `xml:",innerxml"`
	Language string `xml:"xml:lang,attr,omitempty"`
}

func (c *converter) addMetadata() {
	marshalledXml, _ := xml.Marshal(
		metadata{
			Namespace: "adobe:ns:meta/",
			Rdf: metadataRdf{
				Namespace: "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
				Description: []metadataDescription{
					{
						NamespaceDc: "http://purl.org/dc/elements/1.1/",
						About:       "",

						Format: "application/pdf",
						Title: []metadataTitle{
							{
								Languages: []metadataTitleAlt{
									{
										Li: metadataTitleLi{
											Title:    c.title,
											Language: "x-default",
										},
									},
								},
							},
						},
					},
					{
						NamespacePdf: "http://ns.adobe.com/pdf/1.3/",
						About:        "",

						Title:    nil,
						Producer: producer,
					},
				},
			},
		},
	)

	metadataXml := &bytes.Buffer{}
	metadataXml.Grow((len(xml.Header) - 1) + 53 + len(marshalledXml) + 19)
	metadataXml.WriteString(strings.TrimRight(xml.Header, "\n"))
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

	date := c.creationDate.Format("20060102150405-07")
	metaObjData := []string{
		fmt.Sprintf("/Title(%s)", c.title),
		fmt.Sprintf("/Producer(%s)", producer),
		fmt.Sprintf("/CreationDate(D:%s'00')", date),
		fmt.Sprintf("/ModDate(D:%s'00')", date),
	}

	if c.author != "" {
		metaObjData = append(metaObjData, fmt.Sprintf("/Author(%s)", c.author))
	}

	c.addObj(metaObjData, nil)
}
