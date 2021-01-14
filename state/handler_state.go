package state

type State interface {
	Get(string) interface{}
	Set(string, interface{})
}

type handlerState struct {
	variables map[string]interface{}
}

func New() *handlerState {
	return &handlerState{variables: make(map[string]interface{})}
}

func (s *handlerState) Get(variableName string) interface{} {
	return s.variables[variableName]
}

func (s *handlerState) Set(variableName string, value interface{}) {
	s.variables[variableName] = value
}
