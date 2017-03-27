package genx

import (
	"io"
	"strconv"
)

type Room struct {
	name string
	lights int
	shutters int
	heating int
	sockets int
	reedContacts int
}

func NewRoom(label string, lights, shutters, heating, sockets, reedContacts int) *Room {
	return &Room{
		name: label,
		lights: lights,
		shutters: shutters,
		heating: heating,
		sockets: sockets,
		reedContacts: reedContacts,
	}
}

func (r *Room) Generate(w io.Writer, m Manufactorer) {
	i := 1
	for i <= r.lights {
		channel := m.LightFunctions("")
		for _, f := range channel {
			s := r.name + "_" + f.String() + strconv.Itoa(L) + "\n"
			w.Write([]byte(s))
		}
		i++
	}
}