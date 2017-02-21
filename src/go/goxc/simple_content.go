package goxc

import "encoding/xml"

type SimpleContent struct {
	XMLName           xml.Name   `xml:"simpleContent"`
	Extension         *Extension `xml:"extension"`
	Name, PackageName string
	Parent            string
}

func (sc *SimpleContent) Generate(targetPrefix string, namespaces map[string]string) {
	sc.Extension.PackageName = sc.PackageName
	sc.Extension.Parent = sc.Parent
	sc.Extension.Name = sc.Name
	sc.Extension.Generate(targetPrefix, namespaces)
}
