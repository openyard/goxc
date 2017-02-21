package goxc

import "encoding/xml"

type AttributeGroup struct {
	Class
	XMLName     xml.Name     `xml:"attributeGroup"`
	Name        string       `xml:"name,attr,omitempty"`
	Ref         string       `xml:"ref,attr,omitempty"`
	Attributes  []*Attribute `xml:"attribute,omitempty"`
	PackageName string
	Imports     []*Import
}

func (ag *AttributeGroup) Generate(targetPrefix string, namespaces map[string]string) {
	for _, a := range ag.Attributes {
		a.PackageName = ag.PackageName
		a.Generate(targetPrefix, namespaces)
		ag.Imports = Append(ag.Imports, a.Type, namespaces)
	}
	if ag.Name != "" {
		ag.Version = version
		ag.Rev = rev
		generateStruct(ag, "templates/attribute_group.tmpl", ag.PackageName, ag.Name, "attribute group")
	}
	ag.Ref = Replace(targetPrefix, ag.Ref)
}
