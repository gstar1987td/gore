package exception

import (
	"context"
	"fmt"
	"runtime"
)

func HandlePanic(ctx context.Context, recoverError interface{}) error {
	stack := make([]byte, 1024)
	stack = stack[:runtime.Stack(stack, false)]
	return fmt.Errorf("panic [%s] recovered. stack trace:%s", recoverError, string(stack))
}
