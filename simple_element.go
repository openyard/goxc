package goxc

type SimpleElement struct {
	PackageName, Name, Type string
	Imports                 []*Import
}

func (se *SimpleElement) Generate(targetPrefix string) {
	se.Type = Replace(targetPrefix, se.Type)
	generateStruct(se, "templates/simple_type.tmpl", se.PackageName, se.Name, "element with simpleElement")
}
