package core

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const (
	OpEqual          = "=="
	OpLargerThen     = ">"
	OpLessThen       = "<"
	OpLargerAndEqual = ">="
	OpLessAndEqual   = "<="
	OpNotEqual       = "!="

	OpAnd = "&&"
	OpOr  = "||"

	OpMultiply = "*"
	OpDived    = "/"
	OpMod      = "%"
	OpPlus     = "+"
	OpMinos    = "-"

	OpLeftShift  = "<<"
	OpRightShift = ">>"
	OpBitAnd     = "&"
	OpBitOr      = "|"

	OpInvert = "-"
	OpNot    = "!"

	OpIncrease = "++"
	OpDecrease = "--"
)

const (
	TypeInt    = "int"
	TypeFloat  = "float"
	TypeString = "string"
	TypeBool   = "bool"
	TypeTs     = "timestamp"
)

// IdentifierTypeSet use for verify identifiers.
// TODO：if you want to use other identifier prefix,add here
var IdentifierTypeSet = map[string]interface{}{
	"var":  nil,
	"fact": nil,
}

type Null struct{}

func (n *Null) String() string {
	return "null"
}

func getLocation(ctx antlr.ParserRuleContext) string {
	return fmt.Sprintf("%d.%d", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

type DataTypeNull struct{}
