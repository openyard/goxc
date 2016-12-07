package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"go/goxc"
	"strings"
)

var (
	version = "TRUNK"
	rev     = "unknown"
	built   = "19990101-235959"
)

func main() {
	var schemaFile = flag.String("schema", "", "the XSD schema file to parse")
	var schemaPrefix = flag.String("prefix", "", "the XSD schema prefix, e.g. github.com/foo/bar")
	var printVersion = flag.Bool("v", false, "print the goxc version")
	flag.Parse()

	if *printVersion {
		fmt.Printf("goxc %s\n", version)
		fmt.Printf(" rev     : %s\n", rev)
		fmt.Printf(" built   : %s\n\n", built)
		os.Exit(0)
	}
	file := *schemaFile
	path, err := filepath.Abs(filepath.Dir(file))
	fmt.Printf("using %s/%s\n", path, file)
	exitOnError(err, 1)

	if !strings.HasSuffix(*schemaPrefix, "/") {
		*schemaPrefix = *schemaPrefix + "/"
	}

	schema, err := goxc.Parse("", file, *schemaPrefix)
	exitOnError(err, 2)

	fmt.Printf("prefix %s\n", *schemaPrefix)
	i := schema.Generate(path)
	fmt.Printf("\nGenerated %d structs .... have fun\n", i)
}

func exitOnError(err error, code int) {
	if err != nil {
		fmt.Printf("exit on error: %d", code)
		os.Exit(code)
	}
}
