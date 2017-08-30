package goxc

import (
	"encoding/xml"
	"strings"
)

type SimpleType struct {
	Class
	XMLName                   xml.Name     `xml:"simpleType"`
	Name                      string       `xml:"name,attr"`
	Annotation                *Annotation  `xml:"annotation,omitempty"`
	Restriction               *Restriction `xml:"restriction,omitempty"`
	Union                     *Union       `xml:"union,omitempty"`
	List                      *List        `xml:"list,omitempty"`
	Imports                   []*Import
	PackageName, Parent, Type string
}

func (s *SimpleType) Generate(targetPrefix string, namespaces map[string]string) {
	s.Version = version
	s.Rev = rev
	if s.isRestriction() {
		if s.Restriction.isEnumeration() {
			s.prepareRestriction(targetPrefix, namespaces)
			s.Restriction.Generate(targetPrefix, namespaces)
		} else {
			s.Type = Replace(targetPrefix, s.Restriction.Base, namespaces)
			s.Imports = Append(s.Imports, s.Type, namespaces)
			generateStruct(s, "templates/simple_type.tmpl", s.PackageName, s.Name, "simpleType")
		}
	}
	if s.isUnion() {
		s.Type = Replace(targetPrefix, strings.Replace(s.Union.MemberTypes, " ", "\n", -1), namespaces)
		generateStruct(s, "templates/union_type.tmpl", s.PackageName, s.Name, "simpleType (unionType)")
	}
	if s.isList() {
		s.Type = Replace(targetPrefix, s.List.ItemType, namespaces)
		generateStruct(s, "templates/list_type.tmpl", s.PackageName, s.Name, "simpleType (listType)")
	}
}

func (s *SimpleType) isRestriction() bool {
	return s.Restriction != nil
}

func (s *SimpleType) isUnion() bool {
	return s.Union != nil
}

func (s *SimpleType) isList() bool {
	return s.List != nil
}

func (s *SimpleType) hasParent() bool {
	return s.Parent != ""
}

func (s *SimpleType) prepareRestriction(targetPrefix string, namespaces map[string]string) {
	s.Restriction.PackageName = s.PackageName
	s.Restriction.Name = s.Name

	if strings.Contains(s.Restriction.Base, ":") {
		s.Type = Replace(targetPrefix, s.Restriction.Base, namespaces)
	} else {
		s.Type = Upper(s.Restriction.Base)
	}

	if s.hasParent() && strings.Contains(s.Restriction.Base, ":") {
		s.Restriction.Name = s.Type
	}
}