package main

import (
	"encoding/xml"
	"fmt"
)

type Attribute struct {
	XMLName     xml.Name    `xml:"attribute"`
	Name        string      `xml:"name,attr"`
	Type        string      `xml:"type,attr"`
	Use         string      `xml:"use"`
	Default     string      `xml:"default"`
	Annotation  *Annotation `xml:"annotation"`
	PackageName string
}

func (a *Attribute) Print() {
	fmt.Printf("%25s :%s\n", a.Name, a.Type)
}

func (a *Attribute) Generate() {

}
