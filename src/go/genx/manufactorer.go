package genx

type Manufactorer int

const (
	MDT  = 1 + iota
	GIRA
)

var actuators = map[Manufactorer]map[int][]string{
	MDT: {
		L: {
			"AKS-2016.03",
			"AKS-1216.03",
		},
		R: {
			"JAL-0810.02",
		},
		H: {
			"AKH-0800.02",
			"AKH-0400.02",
		},
		FK: {
			"BE-08000.01",
		},
		S: {
			"AKS-2016.03",
			"AKS-1216.03",
		},
		U: {
			"AKU-1616.01",
		},
	},
	GIRA: {},
}