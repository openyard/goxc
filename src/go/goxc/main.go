package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	version = "TRUNK"
	rev     = "unknown"
	built   = "19990101-235959"
)

func main() {
	var schemaFile = flag.String("schema", "", "the XSD schema file to parse")
	var printVersion = flag.Bool("v", false, "print the goxc version")
	flag.Parse()

	if *printVersion {
		fmt.Printf("goxc v%s\n", version)
		fmt.Printf(" rev  : %s\n", rev)
		fmt.Printf(" built: %s\n\n", built)
		os.Exit(0)
	}
	file := *schemaFile
	s, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	var schema Schema
	err = xml.NewDecoder(bytes.NewReader([]byte(s))).Decode(&schema)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	//output, _ := xml.MarshalIndent(schema, "", "  ")
	//os.Stdout.Write(output)

	fmt.Println("")
	//fmt.Printf("%#v", schema)

	fmt.Printf("%25s :%s\n", "XMLName", schema.XMLName.Local)
	fmt.Printf("%25s :%s\n", "Namespace", schema.TargetNamespace)
	for _, doc := range schema.Annotation.Documentations {
		fmt.Printf("%25s :%s\n", doc.Language, doc.Value)
	}
	for _, imp := range schema.Imports {
		fmt.Printf("%25s :%s\n", imp.SchemaLocation, imp.Namespace)
	}
	schema.Print()
	schema.Generate()
}
