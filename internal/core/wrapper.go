package core

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var ActionErrorFormat = "error at at [%d.%d], msg:%s, source:%s"

func WrapActionErrMsg(msg string, ctx antlr.ParserRuleContext) error {
	return fmt.Errorf(ActionErrorFormat,
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
		msg,
		ctx.GetText())
}
