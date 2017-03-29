package genx

import "go/genx/mdt"

type Manufactorer int

const (
	MDT  = 1 + iota
	GIRA
)

var Actuators = map[Manufactorer]map[int]map[string]Actuator{
	MDT: {
		L: {
			"AKS-2016.03": &mdt.AKS201603{},
		},
		R: {
		},
		H: {
		},
		FK: {
		},
		S: {
		},
		U: {
		},
	},
	GIRA: {},
}
