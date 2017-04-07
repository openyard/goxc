package genx

import "go/genx/mdt"

type Manufactorer int

const (
	MDT = 1 + iota
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
			"AKH-0800.02": &mdt.AKH0x0002{},
		},
		FK: {
			"BE-1600.01": &mdt.BExx0001{},
		},
		S: {
			"AKS-2016.03": &mdt.AKSxx1603{Socket: true},
		},
		U: {
			"AKU-2016.01": &mdt.AKUxx1601{},
		},
	},
	GIRA: {},
}
