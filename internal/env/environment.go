package env

import (
	"gore/internal/interfaces"
)

type Environment struct {
	Loader  interfaces.DataLoader
	Checker interfaces.DataChecker
	Stack   interfaces.Stack
	Monitor interfaces.Monitor

	Variables map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{
		Variables: make(map[string]interface{}),
	}
}

func (env *Environment) Init(
	l interfaces.DataLoader,
	c interfaces.DataChecker,
	m interfaces.Monitor,
	s interfaces.Stack) {
	env.Loader = l
	env.Checker = c
	env.Monitor = m
	env.Stack = s
}

func (env *Environment) Close() {
	env.Stack.Close()
	env.Monitor.Close()
	env.Checker.Close()
	env.Loader.Close()
}
