package goxc

type SimpleElement struct {
	Class
	PackageName, Name, Type string
	Imports                 []*Import
}

func (se *SimpleElement) Generate(targetPrefix string, namespaces map[string]string) {
	se.Type = Replace(targetPrefix, se.Type, namespaces)
	se.Version = version
	se.Rev = rev
	generateStruct(se, "templates/simple_type.tmpl", se.PackageName, se.Name, "element with simpleElement")
}
