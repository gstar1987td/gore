package core

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"gore/internal/grammar"
)

type ConditionChecker struct {
	Context       antlr.ParserRuleContext
	ConditionList []Condition
}

type Condition struct {
	context *grammar.ConditionCheckContext
}

func NewCondition() *Condition {
	return &Condition{}
}

func (c *Condition) CompileCondition() error {
	return nil
}
