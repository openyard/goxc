package goxc

import "encoding/xml"

type Choice struct {
	XMLName             xml.Name   `xml:"choice"`
	MaxOccurs           string     `xml:"maxOccurs,attr,omitempty"`
	Elements            []*Element `xml:"element"`
	Sequence            *Sequence  `xml:"sequence,omitempty"`
	Any                 *Any       `xml:"any,omitempty"`
	PackageName, Parent string
	Imports             []*Import
}

func (c *Choice) Generate(targetPrefix string, namespaces map[string]string) {
	for _, element := range c.Elements {
		element.PackageName = c.PackageName
		element.Parent = c.Parent
		element.Generate(targetPrefix, namespaces)
		c.Imports = Append(c.Imports, element.Base, namespaces)
	}
	if c.Sequence != nil {
		c.Sequence.PackageName = c.PackageName
		c.Sequence.Parent = c.Parent
		c.Sequence.Generate(targetPrefix, namespaces)
		c.Imports = c.Sequence.Imports
		c.AppendElement(c.Sequence.Elements)
		if c.Sequence.Any != nil {
			c.Any = c.Sequence.Any
			c.Imports = Append(c.Imports, "w3c.Any", namespaces)
		}
	}
	if c.Any != nil {
		c.Any.Generate()
		c.Imports = append(c.Imports, &Import{"w3c", "w3c"})
	}
}

func (c *Choice) AppendElement(elements []*Element) {
	for _, element := range elements {
		if ok := c.contains(element); ok {
			continue
		}
		c.Elements = append(c.Elements, element)
	}
}

func (c *Choice) contains(e *Element) bool {
	for _, v := range c.Elements {
		if v.Name == e.Name {
			return true
		}
	}
	return false
}