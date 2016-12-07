package goxc

import "encoding/xml"

type Enumeration struct {
	XMLName        xml.Name    `xml:"enumeration"`
	Value          string      `xml:"value,attr"`
	Annotation     *Annotation `xml:"annotation,omitempty"`
	Name, TypeName string
}
