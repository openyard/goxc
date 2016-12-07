package goxc

type BaseType struct {
	Name, Type, PackageName string
	Imports                 []*Import
}

func (t *BaseType) Generate(packageName string) {
	t.PackageName = packageName
	generateStruct(t, "templates/base_type.tmpl", packageName, t.Name, "baseType")
}
