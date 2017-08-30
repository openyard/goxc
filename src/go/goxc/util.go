package goxc

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func Append(imports []*Import, t string, namespaces map[string]string) []*Import {
	for _, v := range imports {
		if strings.Contains(t, v.Alias) {
			return imports
		}
	}
	if strings.Contains(t, ".") {
		alias := t[:strings.Index(t, ".")]
		imports = append(imports, &Import{Alias: alias, Package: packagePrefix + PackageName(namespaces[alias])})
	}
	return imports
}

func IsBaseType(t string) bool {
	for _, b := range baseTypes {
		if b.Name == Upper(t) {
			return true
		}
	}
	return false
}

func Namespace(name string, namespaces map[string]string) string {
	if strings.Contains(name, ".") {
		return namespaces[name[:strings.Index(name, ".")]] + " "
	}
	return ""
}

func Name(t string) string {
	if strings.Contains(t, ".") {
		return t[strings.Index(t, ".")+1:]
	}
	return t
}

func PackageName(ns string) string {
	if strings.Contains(ns, "http://") {
		return strings.Replace(strings.ToLower(ns[strings.LastIndex(ns, "/") + 1:]), "#", "", -1)
	}
	return strings.ToLower(ns[strings.LastIndex(ns, ":") + 1:])
}

func Parse(path, file, prefix string) (*Schema, error) {
	s, err := ioutil.ReadFile(path + file)
	if err != nil {
		return nil, logError(err)
	}
	var schema Schema
	packagePrefix = prefix
	err = xml.NewDecoder(bytes.NewReader([]byte(s))).Decode(&schema)
	if err != nil {
		return nil, logError(err)
	}
	return &schema, nil
}

func Replace(targetPrefix, t string, namespaces map[string]string) string {
	if t == "" {
		return "string"
	}
	if strings.Contains(t, targetPrefix) && len(targetPrefix) > 0 {
		return strings.Replace(t, targetPrefix + ":", "", -1)
	}

	if strings.Contains(t, ":") {
		ns, ok := namespaces[t[:strings.Index(t, ":")]]
		if ok && ns == XML_SCHEMA {
			t = strings.Replace(t, t[:strings.Index(t, ":")+1], "", -1)
		} else {
			t = strings.Replace(t, ":", ".", -1)
		}
	}
	if IsBaseType(t) {
		return "w3c." + Upper(t)
	}
	return Upper(t)
}

func Upper(t string) string {
	if strings.Contains(t, ".") || len(t) == 0 {
		return t
	}
	return strings.ToUpper(t[:1]) + t[1:]
}

func logError(err error) error {
	fmt.Printf("%#v\n", err)
	return err
}

func stdout(s string) {
	if os.Getenv("DEBUG") == "1" {
		fmt.Println(s)
	}
}

func generateStruct(data interface{}, templateName, packageName, name, kind string) {
	fileName := fmt.Sprintf("%s/%s.go", packagePrefix + packageName, name)

	if _, err := os.Stat(fileName); err == nil {
		stdout(fmt.Sprintf("file exists already %q", fileName))
		return
	}
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		stdout(fmt.Sprintf("can not create file %s - %s", fileName, err.Error()))
		return
	}
	defer f.Close()
	tmpl, err := Asset(templateName)
	if err != nil {
		stdout(fmt.Sprintf("can not open template - %s", err.Error()))
		return
	}

	stdout(fmt.Sprintf("generating %s/%s [%s]", packageName, name, kind))
	t := template.New(kind)
	t, _ = t.Parse(string(tmpl[:]))
	t.Execute(f, data)
	counter++
}