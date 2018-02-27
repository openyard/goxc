package goxc

import "encoding/xml"

type SimpleContent struct {
	XMLName           xml.Name     `xml:"simpleContent"`
	Extension         *Extension   `xml:"extension,omitempty"`
	Restriction       *Restriction `xml:"restriction,omitempty"`
	Name, PackageName string
	Parent            string
}

func (sc *SimpleContent) Generate(targetPrefix string, namespaces map[string]string) {
	if sc.Extension != nil {
		sc.Extension.PackageName = sc.PackageName
		sc.Extension.Parent = sc.Parent
		sc.Extension.Name = sc.Name
		sc.Extension.Generate(targetPrefix, namespaces)
	} else if sc.Restriction != nil {
		sc.Restriction.PackageName = sc.PackageName
		sc.Restriction.Parent = sc.PackageName
		sc.Restriction.Name = sc.Name
		sc.Restriction.Generate(targetPrefix, namespaces)
	}

}
