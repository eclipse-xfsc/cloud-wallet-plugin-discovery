package types

var env = new(Environment)

type Environment struct {
}

func (t *Environment) IsHealthy() bool {
	return true
}

func GetEnvironment() *Environment {
	return env
}
