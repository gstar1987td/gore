package core

import (
	"context"
	"gore/internal/env"
)

type Executable interface {
	Execute(ctx context.Context, environment env.Environment) (interface{}, error)
}
