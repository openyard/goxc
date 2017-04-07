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
	f.Shutters(w, f.Rooms, m[R])
	//f.Heating(w, f.Rooms, m[H])
	//f.Sockets(w, f.Rooms, m[S])
	//f.ReedContacts(w, f.Rooms, m[FK])
}

func (f *Floor) Lights(w *csv.Writer, r []*Room, a map[string]Actuator) {
	k := 1
	for _, r := range f.Rooms {
		if k == 1 {
			w.Write([]string{"", "Licht", "", strconv.Itoa(f.Number), strconv.Itoa(L)})
		}
		j := 1
		for j <= r.Lights {
			g, _ := a["AKS-2016.03"]
			k = f.WriteFunction(w, g, r.Name, j, k, L)
			j++
		}
	}
}

func (f *Floor) Shutters(w *csv.Writer, r []*Room, a map[string]Actuator) {
	k := 1
	for _, r := range f.Rooms {
		if k == 1 {
			w.Write([]string{"", "Rollladen", "", strconv.Itoa(f.Number), strconv.Itoa(R)})
		}
		j := 1
		for j <= r.Shutters {
			g, _ := a["JAL-0810.02"]
			k = f.WriteFunction(w, g, r.Name, j, k, R)
			j++
		}
	}
}

func (f *Floor) WriteFunction(w *csv.Writer, g Actuator, rName string, j, k, l int) int {
	for _, fun := range g.Functions() {
		label := g.Prefix() + "_" + rName + "_" + fmt.Sprintf("%02d_", j) + fun
		main := strconv.Itoa(f.Number)
		middle := strconv.Itoa(l)
		sub := strconv.Itoa(k)
		w.Write([]string{"", "", label, main, middle, sub})
		k++
	}
	return k
}