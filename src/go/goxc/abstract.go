package goxc

type Abstract struct {
	PackageName, Name string
}

func (a *Abstract) Generate(targetPrefix string) {
	generateStruct(a, "templates/abstract.tmpl", a.PackageName, a.Name, "abstract type")
}
