package core

import (
	"gore/internal/grammar"
)

type Runner struct {
	context *grammar.RunGoreContext
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) CompileRunner() error {
	return nil
}
