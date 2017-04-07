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
			"AKS-2016.03": &mdt.AKSxx1603{},
		},
		R: {
			"JAL-0810.02": &mdt.JAL0x1002{},
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
