package engine

import environment "gore/internal/env"

type Worker struct {
	env *environment.Environment
}

func NewWorker(e *environment.Environment) *Worker {
	return &Worker{env: e}
}
