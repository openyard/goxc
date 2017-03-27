package main

import (
	"encoding/csv"
	"os"
	"flag"
	"fmt"
	"go/genx"
	"strconv"
)

func main() {
	var source = flag.String("s", "", "source file name")
	var dest = flag.String("d", "com_objects.csv", "destination file name")
	flag.Parse()

	f, err := os.Open(*source)
	if err != nil {
		fmt.Printf("could not open source file: %s", err.Error())
		os.Exit(1)
	}
	defer f.Close()
	d, err := os.OpenFile(*dest, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("could not create destination file: %s", err.Error())
		os.Exit(2)
	}
	defer d.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.FieldsPerRecord = 6

	w := csv.NewWriter(d)
	w.Comma = ';'

	records, err := r.ReadAll()
	for _, r := range records {
		lights, err := strconv.Atoi(r[1])
		shutters, err := strconv.Atoi(r[2])
		heating, err := strconv.Atoi(r[3])
		sockets, err := strconv.Atoi(r[4])
		reedContacts, err := strconv.Atoi(r[5])
		if err != nil {
			fmt.Printf("could not read input: %s", err.Error())
			os.Exit(3)
		}
		room := genx.NewRoom(r[0], lights, shutters, heating, sockets, reedContacts)
		room.Generate(w)
	}

}