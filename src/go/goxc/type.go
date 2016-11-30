package main

import "fmt"

type Type interface {
	Print()
}

func printTypes(types []Type) {
	//fmt.Println("\n%ss", reflect.TypeOf(types[0][0]))
	//fmt.Println(len(types))
	fmt.Println("=============================")
	for _, t := range types {
		if t != nil {
			t.Print()
		}
	}
}
