package genx

type Room struct {
	Name string
	Lights int
	Shutters int
	Heating int
	Sockets int
	ReedContacts int
}

func NewRoom(label string, lights, shutters, heating, sockets, reedContacts int) *Room {
	return &Room{
		Name: label,
		Lights: lights,
		Shutters: shutters,
		Heating: heating,
		Sockets: sockets,
		ReedContacts: reedContacts,
	}
}