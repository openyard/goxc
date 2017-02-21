package goxc

type BaseType struct {
	Class
	Name, Type, PackageName string
	Imports                 []*Import
}

func (t *BaseType) Generate(packageName string) {
	t.PackageName = packageName
	t.Version = version
	t.Rev = rev
	generateStruct(t, "templates/base_type.tmpl", packageName, t.Name, "baseType")
}
