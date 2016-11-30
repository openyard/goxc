package main

import (
	"encoding/xml"
	"strings"
)

type Schema struct {
	XMLName         xml.Name       `xml:"schema"`
	TargetNamespace string         `xml:"targetNamespace,attr"`
	Annotation      *Annotation    `xml:"annotation,omitempty"`
	Imports         []*Import      `xml:"import"`
	SimpleTypes     []*SimpleType  `xml:"simpleType"`
	ComplexTypes    []*ComplexType `xml:"complexType"`
	Elements        []*Element     `xml:"element"`
}

type Import struct {
	XMLName        xml.Name `xml:"import"`
	Namespace      string   `xml:"namespace,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
}

type Annotation struct {
	XMLName        xml.Name         `xml:"annotation"`
	Documentations []*Documentation `xml:"documentation"`
}

type Restriction struct {
	XMLName        xml.Name        `xml:"restriction"`
	Base           string          `xml:"base,attr"`
	Length         *Length         `xml:"length,omitempty"`
	Pattern        *Pattern        `xml:"pattern,omitempty"`
	MaxInclusive   *MaxInclusive   `xml:"maxInclusive,omitempty"`
	TotalDigits    *TotalDigits    `xml:"totalDigits,omitempty"`
	FractionDigits *FractionDigits `xml:"fractionDigits,omitempty"`
	MaxLength      *MaxLength      `xml:"maxLength,omitempty"`
	Enumerations   []*Enumeration  `xml:"enumeration"`
}

type Union struct {
	XMLName     xml.Name `xml:"union"`
	MemberTypes string   `xml:"memberTypes,attr"`
}

type Documentation struct {
	XMLName  xml.Name `xml:"documentation"`
	Language string   `xml:"lang,attr"`
	Value    string   `xml:",chardata"`
}

type Length struct {
	XMLName xml.Name `xml:"length"`
	Value   string   `xml:"value,attr"`
}

type Pattern struct {
	XMLName xml.Name `xml:"pattern"`
	Value   string   `xml:"value,attr"`
}

type MaxInclusive struct {
	XMLName xml.Name `xml:"maxInclusive"`
	Value   string   `xml:"value"`
}

type TotalDigits struct {
	XMLName xml.Name `xml:"totalDigits"`
	Value   string   `xml:"value"`
}

type FractionDigits struct {
	XMLName xml.Name `xml:"fractionDigits"`
	Value   string   `xml:"value"`
}

type Enumeration struct {
	XMLName    xml.Name    `xml:"enumeration"`
	Value      string      `xml:"value,attr"`
	Annotation *Annotation `xml:"annotation,omitempty"`
}

type SimpleContent struct {
	XMLName   xml.Name   `xml:"simpleContent"`
	Extension *Extension `xml:"extension"`
}

type Extension struct {
	XMLName    xml.Name     `xml:"extension"`
	Base       string       `xml:"base,attr"`
	Sequence   *Sequence    `xml:"sequence,omitempty"`
	Attributes []*Attribute `xml:"attribute,omitempty"`
}

type List struct {
	XMLName  xml.Name `xml:"list"`
	ItemType string   `xml:"itemType,attr"`
}

type MaxLength struct {
	XMLName xml.Name `xml:"maxLength"`
	Value   string   `xml:"value"`
}

type Sequence struct {
	XMLName  xml.Name   `xml:"sequence"`
	Choices  []*Choice  `xml:"choice"`
	Elements []*Element `xml:"element"`
}

type Choice struct {
	XMLName   xml.Name   `xml:"choice"`
	MaxOccurs string     `xml:"maxOccurs,attr,omitempty"`
	Elements  []*Element `xml:"element"`
}

type ComplexContent struct {
	Extension *Extension `xml:"extension"`
}

func (s *Schema) Print() {
	var out []Type
	for _, t := range s.SimpleTypes {
		out = append(out, t)
	}
	for _, t := range s.ComplexTypes {
		out = append(out, t)
	}
	for _, t := range s.Elements {
		out = append(out, t)
	}
	//printTypes(s.SimpleTypes, s.ComplexTypes)
	printTypes(out)
}

func (s *Schema) Generate() {
	packageName := strings.ToLower(s.TargetNamespace[strings.LastIndex(s.TargetNamespace, ":")+1:])
	for _, s := range s.SimpleTypes {
		s.PackageName = packageName
		s.Generate()
	}
	for _, c := range s.ComplexTypes {
		c.PackageName = packageName
		c.Generate()
	}
	for _, e := range s.Elements {
		e.PackageName = packageName
		e.Generate()
	}
}
