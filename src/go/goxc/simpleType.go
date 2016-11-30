package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type SimpleType struct {
	XMLName           xml.Name     `xml:"simpleType"`
	Name              string       `xml:"name,attr"`
	Annotation        *Annotation  `xml:"annotation,omitempty"`
	Restriction       *Restriction `xml:"restriction,omitempty"`
	Union             *Union       `xml:"union,omitempty"`
	List              *List        `xml:"list,omitempty"`
	PackageName, Type string
}

func (s *SimpleType) Print() {
	if s.Restriction != nil {
		fmt.Printf("%25s :%s\n", s.Name, s.Restriction.Base)
		s.Type = s.Restriction.Base
	}
	if s.Union != nil {
		fmt.Printf("%25s :%s\n", s.Name, s.Union.MemberTypes)
		s.Type = s.Union.MemberTypes
	}
}

func (s *SimpleType) Generate() {
	tmpl, err := Asset("templates/simpleType.tmpl")
	if err != nil {
		log.Fatalf("can not open template - %s", err.Error())
	}
	fileName := fmt.Sprintf("%s/%s.go", s.PackageName, s.Name)

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalf("can not create file %s - %s", fileName, err.Error())
	}
	defer f.Close()
	t := template.New("simpleType")
	fmt.Println(string(tmpl[:]))
	t, err = t.Parse(string(tmpl[:]))
	//fmt.Printf("%s %s %s", s.PackageName, s.Name, s.Type)
	t.Execute(f, s)
}
