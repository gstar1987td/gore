package core

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"gore/internal/grammar"
)

type ConditionChecker struct {
	Context       antlr.ParserRuleContext
	ConditionList []*Condition
	activated     bool
}

func NewConditionChecker() *ConditionChecker {
	return &ConditionChecker{
		ConditionList: make([]*Condition, 0, 0),
	}
}

func (c *ConditionChecker) CompileCondition(ctx *grammar.ConditionCheckContext) error {
	if ctx == nil {
		return fmt.Errorf("empty conditon context")
	}
	compiler := NewConditionCompiler()
	antlr.ParseTreeWalkerDefault.Walk(compiler, ctx)
	if len(compiler.cond) == 0 {
		c.activated = false
		return nil
	}
	c.activated = true
	c.ConditionList = compiler.cond

	return nil
}

type Condition struct {
	context *grammar.ConditionCheckContext
}

func NewCondition() *Condition {
	return &Condition{}
}

type ConditionCompiler struct {
	*grammar.BaseGoreListener
	cond []*Condition
}

func NewConditionCompiler() *ConditionCompiler {
	return &ConditionCompiler{
		cond: make([]*Condition, 0),
	}
}

// EnterConditionList is called when production conditionList is entered.
func (c *ConditionCompiler) EnterConditionList(ctx *grammar.ConditionListContext) {
	for _, conditionExpr := range ctx.AllExpression() {

	}
}
