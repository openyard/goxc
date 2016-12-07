package goxc

import "encoding/xml"

type Any struct {
	XMLName         xml.Name `xml:"any"`
	Namespace       string   `xml:"namespace,attr"`
	ProcessContents string   `xml:"processContents,attr"`
	MinOccurs       string   `xml:"minOccurs,attr,omitempty"`
	MaxOccurs       string   `xml:"maxOccurs,attr,omitempty"`
	PackageName     string
}

func (a *Any) Generate() {
	a.PackageName = "w3c"
	generateStruct(a, "templates/any.tmpl", a.PackageName, "Any", "any type")
}
