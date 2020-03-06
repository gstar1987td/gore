package rule

import (
	"context"
	"gore/internal/core"
)

type GoreRule struct {
	*core.Rule
}

func NewGoreRuleFromInput(ctx context.Context, input string) (*GoreRule, error) {
	var compiler = core.NewCompiler()
	rule, compileErr := compiler.CompileInput(ctx, input)
	if compileErr != nil {
		return nil, compileErr
	}
	gore := &GoreRule{
		rule,
	}
	return gore, nil
}

func NewGoreRuleFromFile(ctx context.Context, file string) (*GoreRule, error) {
	var compiler = core.NewCompiler()
	rule, compileErr := compiler.CompileFile(ctx, file)
	if compileErr != nil {
		return nil, compileErr
	}
	gore := &GoreRule{
		rule,
	}
	return gore, nil
}
