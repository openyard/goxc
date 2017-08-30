package goxc

import "encoding/xml"

type ComplexType struct {
	Class
	XMLName         xml.Name          `xml:"complexType"`
	Name            string            `xml:"name,attr"`
	Annotation      *Annotation       `xml:"annotation,omitempty"`
	SimpleContent   *SimpleContent    `xml:"simpleContent,omitempty"`
	ComplexContent  *ComplexContent   `xml:"complexContent,omitempty"`
	Sequence        *Sequence         `xml:"sequence,omitempty"`
	Choice          *Choice           `xml:"choice,omitempty"`
	Attributes      []*Attribute      `xml:"attribute,omitempty"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup,omitempty"`
	Any             []*Any            `xml:"any,omitempty"`
	Elements        []*Element
	Refs            []*Element
	Imports         []*Import
	PackageName     string
	Base, Parent    string
}

func (c *ComplexType) Generate(targetPrefix string, namespaces map[string]string) {
	for _, attribute := range c.Attributes {
		attribute.PackageName = c.PackageName
		attribute.Generate(targetPrefix, namespaces)
		c.Imports = Append(c.Imports, attribute.Type, namespaces)
	}
	if c.SimpleContent != nil {
		c.SimpleContent.Name = c.Name
		c.SimpleContent.PackageName = c.PackageName
		c.Base = Replace(targetPrefix, c.SimpleContent.Extension.Base, namespaces)
		c.Imports = Append(c.Imports, c.Base, namespaces)
		c.Attributes = c.SimpleContent.Extension.Attributes
		if c.SimpleContent.Extension.AttributeGroup != nil {
			c.AttributeGroups = append(c.AttributeGroups, c.SimpleContent.Extension.AttributeGroup)
		}
		c.SimpleContent.Generate(targetPrefix, namespaces)
		for _, a := range c.SimpleContent.Extension.Attributes {
			c.Imports = Append(c.Imports, Replace(targetPrefix, a.Type, namespaces), namespaces)
		}
	}
	if c.ComplexContent != nil {
		if c.ComplexContent.Extension != nil {
			c.ComplexContent.Name = c.Name
			c.ComplexContent.PackageName = c.PackageName
			c.Base = Replace(targetPrefix, c.ComplexContent.Extension.Base, namespaces)
			c.Imports = Append(c.Imports, c.Base, namespaces)
			c.Attributes = c.ComplexContent.Extension.Attributes
			if c.ComplexContent.Extension.AttributeGroup != nil {
				c.AttributeGroups = append(c.AttributeGroups, c.ComplexContent.Extension.AttributeGroup)
			}
			c.ComplexContent.Generate(targetPrefix, namespaces)
		}
		if c.ComplexContent.Restriction != nil {
			c.ComplexContent.Name = c.Name
			c.ComplexContent.PackageName = c.PackageName
			c.Base = Replace(targetPrefix, c.ComplexContent.Restriction.Base, namespaces)
			c.Imports = Append(c.Imports, c.Base, namespaces)
			c.ComplexContent.Generate(targetPrefix, namespaces)
		}
	}
	if c.Sequence != nil {
		c.Sequence.PackageName = c.PackageName
		c.Sequence.Parent = c.Name
		c.Sequence.Generate(targetPrefix, namespaces)
		for _, element := range c.Sequence.Elements {
			c.Append(element, namespaces)
			c.Imports = Append(c.Imports, element.Base, namespaces)
		}
		for _, choice := range c.Sequence.Choices {
			for _, element := range choice.Elements {
				c.Append(element, namespaces)
			}
		}
		for _, a := range c.AttributeGroups {
			a.PackageName = c.PackageName
			a.Generate(targetPrefix, namespaces)
		}
		if c.Sequence.Any != nil {
			c.Any = append(c.Any, c.Sequence.Any)
			c.Imports = Append(c.Imports, "w3c.Any", namespaces)
		}
	}
	if c.Choice != nil {
		c.Choice.PackageName = c.PackageName
		c.Choice.Parent = c.Name
		c.Choice.Generate(targetPrefix, namespaces)
		for _, element := range c.Choice.Elements {
			c.Append(element, namespaces)
			c.Imports = Append(c.Imports, element.Base, namespaces)
		}
		if c.Choice.Any != nil {
			c.Any = append(c.Any, c.Choice.Any)
			c.Imports = Append(c.Imports, "w3c.Any", namespaces)
		}
	}
	c.Version = version
	c.Rev = rev
	generateStruct(c, "templates/complex_type.tmpl", c.PackageName, c.Name, "complexType")
}

func (c *ComplexType) Append(element *Element, namespaces map[string]string) {
	c.Imports = Append(c.Imports, element.Base, namespaces)
	if element.isRef() {
		element.Namespace = Namespace(element.Base, namespaces)
		element.Name = Name(element.Base)
		c.Refs = append(c.Refs, element)
	} else {
		c.Elements = append(c.Elements, element)
	}
}
