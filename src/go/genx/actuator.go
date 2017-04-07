package genx

type Actuator interface {
	Functions() []string
	Prefix() string
}
