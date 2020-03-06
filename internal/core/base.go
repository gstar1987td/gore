package core

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"gore/internal/grammar"
)

type BaseActionCompiler struct {
	*grammar.BaseGoreListener
	Executables []*Executable
}

func NewBaseActionCompiler() *BaseActionCompiler {
	return &BaseActionCompiler{
		Executables: make([]*Executable, 0),
	}
}

func (c *BaseActionCompiler) Compile(ctx antlr.ParserRuleContext) {
	antlr.ParseTreeWalkerDefault.Walk(c, ctx)
}

// for test
func (c *BaseActionCompiler) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("base enter:", ctx.GetText())
}
