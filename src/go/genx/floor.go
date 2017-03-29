package genx

import (
	"encoding/csv"
	"fmt"
	"strconv"
)

type Floor struct {
	Name   string
	Number int
	Rooms  []*Room
}

func NewFloor(label string, number int, rooms ... *Room) *Floor {
	return &Floor{Name: label, Number: number, Rooms: rooms}
}

func (f *Floor) Generate(w *csv.Writer, m map[int]map[string]Actuator) {
	w.Write([]string{f.Name, "", "", "1"})
	f.Lights(w, f.Rooms, m[L])
	//f.Shutter(w, f.Rooms, m[R])
	//f.Heating(w, f.Rooms, m[H])
	//f.Sockets(w, f.Rooms, m[S])
	//f.ReedContacts(w, f.Rooms, m[FK])
}

func (f *Floor) Lights(w *csv.Writer, r []*Room, a map[string]Actuator) {
	for _, r := range f.Rooms {
		j := 1
		k := 1
		w.Write([]string{"", "Licht", "", strconv.Itoa(f.Number), strconv.Itoa(L)})
		for j <= r.Lights {
			g, _ := a["AKS-2016.03"]
			k = f.WriteFunction(w, g, r.Name, j, k)
			j++
		}
	}
}

func (f *Floor) WriteFunction(w *csv.Writer, g Actuator, rName string, j, k int) int {
	for _, fun := range g.Functions() {
		label := "L_" + rName + "_" + fmt.Sprintf("%02d_", j) + fun
		main := strconv.Itoa(f.Number)
		middle := strconv.Itoa(L)
		sub := strconv.Itoa(k)
		w.Write([]string{"", "", label, main, middle, sub})
		k++
	}
	return k
}