package goxc

type ComplexElement struct {
	PackageName, Parent string
	ComplexType         *ComplexType
}

func (ce *ComplexElement) Generate(targetPrefix string, namespaces map[string]string) {
	if ce.ComplexType.Sequence != nil || ce.ComplexType.Choice != nil {
		ce.ComplexType.PackageName = ce.PackageName
		ce.ComplexType.Generate(targetPrefix, namespaces)
		return
	}
	ce.ComplexType.Imports = make([]*Import, 0)
	ce.ComplexType.Attributes = make([]*Attribute, 0)
	ce.ComplexType.AttributeGroups = make([]*AttributeGroup, 0)
	ce.ComplexType.PackageName = ce.PackageName
	ce.ComplexType.Base = Replace(targetPrefix, ce.ComplexType.Base, namespaces)
	if ce.ComplexType.SimpleContent != nil {
		ce.ComplexType.SimpleContent.PackageName = ce.PackageName
		ce.ComplexType.SimpleContent.Parent = ce.Parent
		ce.ComplexType.SimpleContent.Generate(targetPrefix, namespaces)

		ce.ComplexType.Attributes = append(ce.ComplexType.Attributes, ce.ComplexType.SimpleContent.Extension.Attributes...)
		if ce.ComplexType.SimpleContent.Extension.AttributeGroup != nil {
			ce.ComplexType.AttributeGroups = append(ce.ComplexType.AttributeGroups, ce.ComplexType.SimpleContent.Extension.AttributeGroup)
		}
		ce.ComplexType.Base = Replace(targetPrefix, ce.ComplexType.SimpleContent.Extension.Base, namespaces)
	}
	if ce.ComplexType.ComplexContent != nil {
		ce.ComplexType.ComplexContent.PackageName = ce.PackageName
		ce.ComplexType.ComplexContent.Parent = ce.Parent
		ce.ComplexType.ComplexContent.Generate(targetPrefix, namespaces)
		ce.ComplexType.Attributes = append(ce.ComplexType.Attributes, ce.ComplexType.ComplexContent.Extension.Attributes...)
		if ce.ComplexType.ComplexContent.Extension.AttributeGroup != nil {
			ce.ComplexType.AttributeGroups = append(ce.ComplexType.AttributeGroups, ce.ComplexType.ComplexContent.Extension.AttributeGroup)
		}
		ce.ComplexType.Base = Replace(targetPrefix, ce.ComplexType.ComplexContent.Extension.Base, namespaces)
	}
	ce.ComplexType.Imports = Append(ce.ComplexType.Imports, ce.ComplexType.Base, namespaces)
	for _, a := range ce.ComplexType.Attributes {
		a.PackageName = ce.PackageName
		a.Name = Upper(a.Name)
		a.Type = Replace(targetPrefix, a.Type, namespaces)
		ce.ComplexType.Imports = Append(ce.ComplexType.Imports, a.Type, namespaces)
	}
	ce.ComplexType.Version = version
	ce.ComplexType.Rev = rev
	generateStruct(ce.ComplexType, "templates/complex_element.tmpl", ce.PackageName, ce.ComplexType.Name, "element with complexType")
}
