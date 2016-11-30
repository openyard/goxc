package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"os"
)

type Element struct {
	XMLName     xml.Name     `xml:"element"`
	Name        string       `xml:"name,attr"`
	Type        string       `xml:"type,attr"`
	MinOccurs   string       `xml:"minOccurs,attr"`
	MaxOccurs   string       `xml:"maxOccurs,attr"`
	Annotation  *Annotation  `xml:"annotation,omitempty"`
	ComplexType *ComplexType `xml:"complexType,omitempty"`
	PackageName string
}

func (e *Element) Print() {
	if e.ComplexType != nil {
		e.ComplexType.Name = e.Name
		e.ComplexType.Print()
	} else {
		fmt.Printf("%25s :%s\n", e.Name, e.Type)
	}
}

func (e *Element) Generate() {
	tmpl, err := Asset("templates/element.tmpl")
	if err != nil {
		log.Fatalf("can not open template - %s", err.Error())
	}
	fileName := fmt.Sprintf("%s/%s.go", e.PackageName, e.Name)

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalf("can not create file %s - %s", fileName, err.Error())
	}
	defer f.Close()
	t := template.New("element")
	fmt.Println(string(tmpl[:]))
	t, err = t.Parse(string(tmpl[:]))
	//fmt.Printf("%s %s %s", s.PackageName, s.Name, s.Type)
	t.Execute(f, e)
}
