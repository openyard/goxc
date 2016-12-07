package goxc

import "encoding/xml"

type Sequence struct {
	XMLName             xml.Name   `xml:"sequence"`
	Choices             []*Choice  `xml:"choice"`
	Elements            []*Element `xml:"element"`
	Any                 *Any       `xml:"any,omitempty"`
	PackageName, Parent string
	Imports             []*Import
}

func (s *Sequence) Generate(targetPrefix string, namespaces map[string]string) {
	s.Imports = make([]*Import, 0)
	for _, choice := range s.Choices {
		choice.PackageName = s.PackageName
		choice.Parent = s.Parent
		choice.Generate(targetPrefix, namespaces)
		s.Imports = append(s.Imports, choice.Imports...)
	}
	for _, element := range s.Elements {
		element.PackageName = s.PackageName
		element.Parent = s.Parent
		element.Generate(targetPrefix, namespaces)
		s.Imports = Append(s.Imports, element.Base, namespaces)
	}
	if s.Any != nil {
		s.Any.Generate()
		s.Imports = append(s.Imports, &Import{"w3c", "w3c"})
	}
}
