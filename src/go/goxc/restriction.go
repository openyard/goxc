package goxc

import (
	"encoding/xml"
	"strings"
)

type Restriction struct {
	Class
	XMLName                   xml.Name        `xml:"restriction"`
	Base                      string          `xml:"base,attr"`
	Length                    *Length         `xml:"length,omitempty"`
	Pattern                   *Pattern        `xml:"pattern,omitempty"`
	MaxInclusive              *MaxInclusive   `xml:"maxInclusive,omitempty"`
	TotalDigits               *TotalDigits    `xml:"totalDigits,omitempty"`
	FractionDigits            *FractionDigits `xml:"fractionDigits,omitempty"`
	MaxLength                 *MaxLength      `xml:"maxLength,omitempty"`
	Enumerations              []*Enumeration  `xml:"enumeration"`
	Sequence                  *Sequence       `xml:"sequence,omitempty"`
	Any                       *Any            `xml:"any,omitempty"`
	Imports                   []*Import
	PackageName, Parent, Name string
}

func (r *Restriction) Generate(targetPrefix string, namespaces map[string]string) {
	r.Base = Replace(targetPrefix, r.Base, namespaces)
	r.Imports = Append(r.Imports, r.Base, namespaces)
	if r.isEnumeration() {
		r.prepareEnumeration()
		r.generateEnumeration()
	}
	if r.Sequence != nil {
		r.Sequence.PackageName = r.PackageName
		r.Sequence.Parent = r.Parent
		r.Sequence.Generate(targetPrefix, namespaces)
		r.Imports = append(r.Imports, r.Sequence.Imports...)
	}
	if r.Any != nil {
		r.Any.Generate()
		r.Imports = append(r.Imports, &Import{"w3c", "w3c"})
	}
}

func (r *Restriction) isEnumeration() bool {
	return r.Enumerations != nil && len(r.Enumerations) > 0
}

func (r *Restriction) prepareEnumeration() {
	for _, e := range r.Enumerations {
		e.Name = r.Name
		e.TypeName = strings.ToUpper(e.Value)
	}
}

func (r *Restriction) generateEnumeration() {
	r.Version = version
	r.Rev = rev
	generateStruct(r, "templates/enumeration.tmpl", r.PackageName, r.Parent+r.Name, "enumeration")
}