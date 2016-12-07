package goxc

import (
	"encoding/xml"
	"strings"
)

type Attribute struct {
	XMLName           xml.Name    `xml:"attribute"`
	Name              string      `xml:"name,attr"`
	Type              string      `xml:"type,attr"`
	Use               string      `xml:"use,attr"`
	Default           string      `xml:"default,attr"`
	Annotation        *Annotation `xml:"annotation,omitempty"`
	SimpleType        *SimpleType `xml:"simpleType,omitempty"`
	PackageName       string
	Parent, OmitEmpty string
	TypeName 	  string
	Imports           []*Import
}

func (a *Attribute) Generate(targetPrefix string, namespaces map[string]string) {
	a.TypeName = Upper(a.Name)
	if a.hasSimpleType() {
		a.prepareSimpleType()
		a.SimpleType.Generate(targetPrefix, namespaces)
		a.finish()
	} else {
		a.prepareMe(targetPrefix, namespaces)
	}
	generateStruct(a, "templates/attribute.tmpl", a.PackageName, a.TypeName, "attribute")
}

func (a *Attribute) hasParent() bool {
	return a.Parent != ""
}

func (a *Attribute) hasSimpleType() bool {
	return a.SimpleType != nil
}

func (a *Attribute) finish() {
	a.Imports = a.SimpleType.Imports
	a.Type = a.SimpleType.Type
	if a.hasParent() {
		a.Type = a.SimpleType.Name
	}
}

func (a *Attribute) prepareMe(targetPrefix string, namespaces map[string]string) {
	a.Type = Replace(targetPrefix, a.Type)
	if strings.Contains(a.Type, ".") && !strings.Contains(a.Type, targetPrefix) {
		a.Imports = Append(a.Imports, a.Type, namespaces)
	}
	if a.Use == "optional" {
		a.OmitEmpty = ",omitempty"
	}
}

func (a *Attribute) prepareSimpleType() {
	a.SimpleType.PackageName = a.PackageName
	a.SimpleType.Parent = a.Parent + a.TypeName
	if a.SimpleType.Name == "" {
		a.SimpleType.Name = a.Parent + a.TypeName
	}
}