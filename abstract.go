package goxc

type Abstract struct {
	Class
	PackageName, Name string
}

func (a *Abstract) Generate(targetPrefix string) {
	a.Version = version
	a.Rev = rev
	generateStruct(a, "templates/abstract.tmpl", a.PackageName, a.Name, "abstract type")
}
