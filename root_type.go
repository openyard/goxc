package goxc

type RootNamespace struct {
	Name   string
	Prefix string
}

type RootType struct {
	Class
	Name, XMLName, Type string
	PackageName         string
	TargetNamespace     string
	Imports             []*Import
}

func (s *RootType) Generate(targetPrefix string, namespaces map[string]string) {
	s.Version = version
	s.Rev = rev
	s.Type = Replace(targetPrefix, s.Type, namespaces)
	generateStruct(s, "templates/root_type.tmpl", s.PackageName, s.Name, "rootType")
}
