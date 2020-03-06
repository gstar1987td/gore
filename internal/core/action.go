package core

import (
	"context"
	"gore/common/constants"
	"gore/internal/grammar"
	"gore/internal/interfaces"
)

type Executable interface {
	Execute(ctx context.Context, env interfaces.Runtime) (interface{}, error)
}

/*
	expression
    : expression 'is' (n='not')? 'null'                             #isNull
    | expression (n='not')? 'like' RAW_STRING_LIT                   #like
    | expression (n='not')? 'in' '('expression (',' expression)+')' #in
    | l=expression HIGH_PRIORITY_MATH r=expression                  #hmath
    | l=expression LOW_PRIORITY_MATH r=expression                   #lmath
    | l=expression BIT_OPER r=expression                            #bit
    | l=expression COMPARE_OPER r=expression                        #compare
    | l=expression '&&' r=expression                                #and
    | l=expression '||' r=expression                                #or
    | ('max'|'min') '('expression (',' expression)+ ')'             #maxmin
    | unaryExp                                                      #unary
    ;
*/

/*
	isNull
*/
type IsNullExecutable struct {
	src   *grammar.IsNullContext
	where string
	not   bool
}

func NewIsNullExecutable(ctx *grammar.IsNullContext) (Executable, error) {
	exe := &IsNullExecutable{
		src:   ctx,
		where: getLocation(ctx),
	}
	if ctx.GetN() != nil {
		exe.not = true
	}
	return exe, nil
}

func (exe *IsNullExecutable) Execute(ctx context.Context, runtime interfaces.Runtime) (interface{}, error) {
	preStack, ok := runtime.GetStack().Pop()
	result := false
	if !ok {
		return false, WrapActionErrMsg("empty stack", exe.src)
	}
	if _, ok := preStack.(*DataTypeNull); ok {
		result = true
	}
	if exe.not {
		result = !result
	}
	runtime.GetMonitor().AddFeature(constants.MonitorPrefixIsNull, exe.where, result, nil, nil)
	return result, nil
}
