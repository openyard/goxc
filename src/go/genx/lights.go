package genx

type SwitchActuatorFunction int

type Lights interface {
	Functions(m Manufactorer) []SwitchActuatorFunction
}

