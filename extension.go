package goxc

import "encoding/xml"

type Extension struct {
	XMLName                   xml.Name        `xml:"extension"`
	Base                      string          `xml:"base,attr"`
	Sequence                  *Sequence       `xml:"sequence,omitempty"`
	Attributes                []*Attribute    `xml:"attribute,omitempty"`
	AttributeGroup            *AttributeGroup `xml:"attributeGroup,omitempty"`
	PackageName, Parent, Name string
}

func (e *Extension) Generate(targetPrefix string, namespaces map[string]string) {
	if e.Sequence != nil {
		e.Sequence.PackageName = e.PackageName
		e.Sequence.Parent = e.Parent + e.Name
		e.Sequence.Generate(targetPrefix, namespaces)
	}
	for _, attribute := range e.Attributes {
		attribute.PackageName = e.PackageName
		attribute.Parent = e.Parent + e.Name
		attribute.Generate(targetPrefix, namespaces)
	}
	if e.AttributeGroup != nil {
		e.AttributeGroup.PackageName = e.PackageName
		e.AttributeGroup.Generate(targetPrefix, namespaces)
	}
	e.Base = Replace(targetPrefix, e.Base, namespaces)
}
