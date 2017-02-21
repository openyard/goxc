package goxc

import "encoding/xml"

type ComplexContent struct {
	XMLName                   xml.Name     `xml:"complexContent"`
	Extension                 *Extension   `xml:"extension,omitempty"`
	Restriction               *Restriction `xml:"restriction,omitempty"`
	PackageName, Parent, Name string
}

func (cc *ComplexContent) Generate(targetPrefix string, namespaces map[string]string) {
	if cc.Extension != nil {
		cc.Extension.PackageName = cc.PackageName
		cc.Extension.Parent = cc.Parent
		cc.Extension.Name = cc.Name
		cc.Extension.Generate(targetPrefix, namespaces)
	}
	if cc.Restriction != nil {
		cc.Restriction.PackageName = cc.PackageName
		cc.Restriction.Parent = cc.Parent
		cc.Restriction.Name = cc.Name
		cc.Restriction.Generate(targetPrefix, namespaces)
	}
}
