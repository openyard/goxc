package goxc

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"fmt"
)

const XML_SCHEMA = "http://www.w3.org/2001/XMLSchema"

type Schema struct {
	XMLName         xml.Name          `xml:"http://www.w3.org/2001/XMLSchema schema"`
	TargetNamespace string            `xml:"targetNamespace,attr"`
	XMLNS           string            `xml:"xmlns,attr"`
	Annotation      *Annotation       `xml:"annotation,omitempty"`
	Imports         []*SchemaImport   `xml:"import,omitempty"`
	Includes        []*Include        `xml:"include,omitempty"`
	SimpleTypes     []*SimpleType     `xml:"simpleType,omitempty"`
	ComplexTypes    []*ComplexType    `xml:"complexType,omitempty"`
	Elements        []*Element        `xml:"element,omitempty"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup,omitempty"`
	Namespaces      map[string]string
	TargetPrefix    string
}

type SchemaImport struct {
	XMLName        xml.Name `xml:"import"`
	Namespace      string   `xml:"namespace,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
}

type Include struct {
	XMLName        xml.Name `xml:"include"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
}

type Annotation struct {
	XMLName        xml.Name         `xml:"annotation"`
	Documentations []*Documentation `xml:"documentation"`
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

type MinInclusive struct {
	XMLName xml.Name `xml:"minInclusive"`
	Value   string   `xml:"value"`
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

type List struct {
	XMLName  xml.Name `xml:"list"`
	ItemType string   `xml:"itemType,attr"`
}

type MaxLength struct {
	XMLName xml.Name `xml:"maxLength"`
	Value   string   `xml:"value"`
}

func (schema *Schema) Generate(path string) int {
	packageName := PackageName(schema.TargetNamespace)
	os.MkdirAll("."+string(filepath.Separator)+packagePrefix+packageName, 0777)

	for prefix, ns := range schema.Namespaces {
		if ns == schema.TargetNamespace {
			schema.TargetPrefix = prefix
		}
	}

	os.MkdirAll("."+string(filepath.Separator)+packagePrefix+"w3c", 0777)
	for _, b := range baseTypes {
		b.Generate("w3c")
	}

	for _, file := range schema.Imports {
		subSchema, _ := Parse(path+"/", file.SchemaLocation, packagePrefix)
		subSchema.Generate(path)
	}
	for _, file := range schema.Includes {
		subSchema, _ := Parse(path+"/", file.SchemaLocation, packagePrefix)
		subSchema.Generate(path)
	}
	for _, s := range schema.SimpleTypes {
		s.PackageName = packageName
		s.Generate(schema.TargetPrefix, schema.Namespaces)
	}
	for _, c := range schema.ComplexTypes {
		c.PackageName = packageName
		c.Generate(schema.TargetPrefix, schema.Namespaces)
	}
	for _, e := range schema.Elements {
		e.PackageName = packageName
		e.Generate(schema.TargetPrefix, schema.Namespaces)
	}
	for _, a := range schema.AttributeGroups {
		a.PackageName = packageName
		a.Generate(schema.TargetPrefix, schema.Namespaces)
	}
	return counter
}

func (s *Schema) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s.Namespaces = map[string]string{}
	for _, attr := range start.Attr {
		if attr.Name.Space == "xmlns" {
			fmt.Printf("%s - %s - %s\n",  attr.Name.Space,  attr.Name.Local,  attr.Value)
			s.Namespaces[attr.Name.Local] = attr.Value
		}
	}
	s.Namespaces["w3c"] = "w3c"

	// Go on with unmarshalling.
	type schema Schema
	ss := (*schema)(s)
	return d.DecodeElement(ss, &start)
}
