// Code generated from E:/Go/src/gore/internal/grammar\Gore.g4 by ANTLR 4.8. DO NOT EDIT.

package grammar // Gore
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GoreParser.
type GoreVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GoreParser#sourceFile.
	VisitSourceFile(ctx *SourceFileContext) interface{}

	// Visit a parse tree produced by GoreParser#ruleID.
	VisitRuleID(ctx *RuleIDContext) interface{}

	// Visit a parse tree produced by GoreParser#conditionCheck.
	VisitConditionCheck(ctx *ConditionCheckContext) interface{}

	// Visit a parse tree produced by GoreParser#runGore.
	VisitRunGore(ctx *RunGoreContext) interface{}

	// Visit a parse tree produced by GoreParser#conditionList.
	VisitConditionList(ctx *ConditionListContext) interface{}

	// Visit a parse tree produced by GoreParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by GoreParser#incDec.
	VisitIncDec(ctx *IncDecContext) interface{}

	// Visit a parse tree produced by GoreParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by GoreParser#return.
	VisitReturn(ctx *ReturnContext) interface{}

	// Visit a parse tree produced by GoreParser#if.
	VisitIf(ctx *IfContext) interface{}

	// Visit a parse tree produced by GoreParser#try.
	VisitTry(ctx *TryContext) interface{}

	// Visit a parse tree produced by GoreParser#bingo.
	VisitBingo(ctx *BingoContext) interface{}

	// Visit a parse tree produced by GoreParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by GoreParser#thenStmt.
	VisitThenStmt(ctx *ThenStmtContext) interface{}

	// Visit a parse tree produced by GoreParser#incDecStmt.
	VisitIncDecStmt(ctx *IncDecStmtContext) interface{}

	// Visit a parse tree produced by GoreParser#lmath.
	VisitLmath(ctx *LmathContext) interface{}

	// Visit a parse tree produced by GoreParser#maxmin.
	VisitMaxmin(ctx *MaxminContext) interface{}

	// Visit a parse tree produced by GoreParser#compare.
	VisitCompare(ctx *CompareContext) interface{}

	// Visit a parse tree produced by GoreParser#or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by GoreParser#like.
	VisitLike(ctx *LikeContext) interface{}

	// Visit a parse tree produced by GoreParser#in.
	VisitIn(ctx *InContext) interface{}

	// Visit a parse tree produced by GoreParser#and.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by GoreParser#isNull.
	VisitIsNull(ctx *IsNullContext) interface{}

	// Visit a parse tree produced by GoreParser#unary.
	VisitUnary(ctx *UnaryContext) interface{}

	// Visit a parse tree produced by GoreParser#hmath.
	VisitHmath(ctx *HmathContext) interface{}

	// Visit a parse tree produced by GoreParser#bit.
	VisitBit(ctx *BitContext) interface{}

	// Visit a parse tree produced by GoreParser#unaryOp.
	VisitUnaryOp(ctx *UnaryOpContext) interface{}

	// Visit a parse tree produced by GoreParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by GoreParser#oper.
	VisitOper(ctx *OperContext) interface{}

	// Visit a parse tree produced by GoreParser#convert.
	VisitConvert(ctx *ConvertContext) interface{}

	// Visit a parse tree produced by GoreParser#round.
	VisitRound(ctx *RoundContext) interface{}

	// Visit a parse tree produced by GoreParser#basicLiter.
	VisitBasicLiter(ctx *BasicLiterContext) interface{}

	// Visit a parse tree produced by GoreParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by GoreParser#par.
	VisitPar(ctx *ParContext) interface{}

	// Visit a parse tree produced by GoreParser#unaryOper.
	VisitUnaryOper(ctx *UnaryOperContext) interface{}

	// Visit a parse tree produced by GoreParser#var.
	VisitVar(ctx *VarContext) interface{}

	// Visit a parse tree produced by GoreParser#fact.
	VisitFact(ctx *FactContext) interface{}

	// Visit a parse tree produced by GoreParser#int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by GoreParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by GoreParser#float.
	VisitFloat(ctx *FloatContext) interface{}

	// Visit a parse tree produced by GoreParser#bool.
	VisitBool(ctx *BoolContext) interface{}

	// Visit a parse tree produced by GoreParser#null.
	VisitNull(ctx *NullContext) interface{}

	// Visit a parse tree produced by GoreParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by GoreParser#eos.
	VisitEos(ctx *EosContext) interface{}
}
