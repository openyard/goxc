package main

import (
	"encoding/xml"
	"fmt"
)

type ComplexType struct {
	XMLName        xml.Name        `xml:"complexType"`
	Name           string          `xml:"name,attr"`
	Annotation     *Annotation     `xml:"annotation,omitempty"`
	SimpleContent  *SimpleContent  `xml:"simpleContent,omitempty"`
	ComplexContent *ComplexContent `xml:"complexContent,omitempty"`
	Sequence       *Sequence       `xml:"sequence,omitempty"`
	Attributes     []*Attribute    `xml:"attribute,omitempty"`
	PackageName    string
}

func (c *ComplexType) Print() {
	if c.SimpleContent != nil {
		fmt.Printf("%25s :%s\n", c.Name, c.SimpleContent.Extension.Base)
		for _, attribute := range c.SimpleContent.Extension.Attributes {
			attribute.Print()
		}
	}
	if c.ComplexContent != nil {
		fmt.Printf("%25s :%s\n", c.Name, c.ComplexContent.Extension.Base)
		for _, element := range c.ComplexContent.Extension.Sequence.Elements {
			element.Print()
		}
	}
	if c.Sequence != nil {
		fmt.Printf("%25s\n", c.Name)
		for _, choice := range c.Sequence.Choices {
			for _, element := range choice.Elements {
				element.Print()
			}
		}
		for _, element := range c.Sequence.Elements {
			element.Print()
		}
	}
	for _, attribute := range c.Attributes {
		attribute.Print()
	}
}

func (c *ComplexType) Generate() {

}
