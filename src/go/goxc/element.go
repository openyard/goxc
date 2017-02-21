package goxc

import (
	"encoding/xml"
)

type Element struct {
	XMLName                 xml.Name     `xml:"element"`
	Name                    string       `xml:"name,attr"`
	Abstract                bool         `xml:"abstract,attr"`
	Type                    string       `xml:"type,attr"`
	Ref                     string       `xml:"ref,attr"`
	MinOccurs               string       `xml:"minOccurs,attr"`
	MaxOccurs               string       `xml:"maxOccurs,attr"`
	Annotation              *Annotation  `xml:"annotation,omitempty"`
	SimpleType              *SimpleType  `xml:"simpleType,omitempty"`
	ComplexType             *ComplexType `xml:"complexType,omitempty"`
	PackageName, Namespace  string
	Base, Parent, OmitEmpty string
	TypeName 		string
	Imports                 []*Import
}

func (e *Element) Generate(targetPrefix string, namespaces map[string]string) {
	e.TypeName = Upper(e.Name)
	if e.isOptional() {
		e.OmitEmpty = ",omitempty"
	}
	if e.isSimple() {
		e.Base = Replace(targetPrefix, e.SimpleType.Type)
		if e.SimpleType.isRestriction() {
			e.Base = Replace(targetPrefix, e.SimpleType.Restriction.Base)
		}
		e.Imports = append(e.Imports, e.SimpleType.Imports...)
	} else if e.isComplex() {
		e.ComplexType.Name = e.Parent + e.TypeName
		ce := &ComplexElement{PackageName: e.PackageName, Parent: e.Parent + e.Name, ComplexType: e.ComplexType}
		ce.Generate(targetPrefix, namespaces)
		e.Base = e.ComplexType.Name
		e.Imports = append(e.Imports, ce.ComplexType.Imports...)
	} else if e.isAbstract() {
		e.Base = e.Type
		a := &Abstract{PackageName: e.PackageName, Name: e.TypeName}
		a.Generate(targetPrefix)
	} else if e.Name != "" { // && !IsBaseType(e.Type)
		e.Base = Replace(targetPrefix, e.Type)
		e.Imports = Append(e.Imports, e.Base, namespaces)
		se := &SimpleElement{PackageName: e.PackageName, Name: e.TypeName, Type: e.Type, Imports: e.Imports}
		se.Generate(targetPrefix)
	} else {
		e.Base = Replace(targetPrefix, e.Ref)
		e.TypeName = Name(e.Base)
		//if strings.Contains(e.TypeName, ".") {
		//	e.TypeName = e.Base[strings.Index(e.Base, ".")+1:]
		//}
		e.Imports = Append(e.Imports, e.Base, namespaces)
	}
}

func (e *Element) isAbstract() bool {
	return e.Abstract
}

func (e *Element) isComplex() bool {
	return e.ComplexType != nil
}

func (e *Element) isOptional() bool {
	return e.MinOccurs == "0"
}

func (e *Element) isRef() bool {
	return e.Ref != ""
}

func (e *Element) isSimple() bool {
	return e.SimpleType != nil
}