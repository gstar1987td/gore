// Code generated from E:/Go/src/gore/internal/grammar\Gore.g4 by ANTLR 4.8. DO NOT EDIT.

package grammar // Gore
import "github.com/antlr/antlr4/runtime/Go/antlr"

// GoreListener is a complete listener for a parse tree produced by GoreParser.
type GoreListener interface {
	antlr.ParseTreeListener

	// EnterSourceFile is called when entering the sourceFile production.
	EnterSourceFile(c *SourceFileContext)

	// EnterRuleID is called when entering the ruleID production.
	EnterRuleID(c *RuleIDContext)

	// EnterConditionCheck is called when entering the conditionCheck production.
	EnterConditionCheck(c *ConditionCheckContext)

	// EnterRunGore is called when entering the runGore production.
	EnterRunGore(c *RunGoreContext)

	// EnterConditionList is called when entering the conditionList production.
	EnterConditionList(c *ConditionListContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterIncDec is called when entering the incDec production.
	EnterIncDec(c *IncDecContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterReturn is called when entering the return production.
	EnterReturn(c *ReturnContext)

	// EnterIf is called when entering the if production.
	EnterIf(c *IfContext)

	// EnterTry is called when entering the try production.
	EnterTry(c *TryContext)

	// EnterBingo is called when entering the bingo production.
	EnterBingo(c *BingoContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterThenStmt is called when entering the thenStmt production.
	EnterThenStmt(c *ThenStmtContext)

	// EnterIncDecStmt is called when entering the incDecStmt production.
	EnterIncDecStmt(c *IncDecStmtContext)

	// EnterLmath is called when entering the lmath production.
	EnterLmath(c *LmathContext)

	// EnterMaxmin is called when entering the maxmin production.
	EnterMaxmin(c *MaxminContext)

	// EnterCompare is called when entering the compare production.
	EnterCompare(c *CompareContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterLike is called when entering the like production.
	EnterLike(c *LikeContext)

	// EnterIn is called when entering the in production.
	EnterIn(c *InContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterIsNull is called when entering the isNull production.
	EnterIsNull(c *IsNullContext)

	// EnterUnary is called when entering the unary production.
	EnterUnary(c *UnaryContext)

	// EnterHmath is called when entering the hmath production.
	EnterHmath(c *HmathContext)

	// EnterBit is called when entering the bit production.
	EnterBit(c *BitContext)

	// EnterUnaryOp is called when entering the unaryOp production.
	EnterUnaryOp(c *UnaryOpContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterOper is called when entering the oper production.
	EnterOper(c *OperContext)

	// EnterConvert is called when entering the convert production.
	EnterConvert(c *ConvertContext)

	// EnterRound is called when entering the round production.
	EnterRound(c *RoundContext)

	// EnterBasicLiter is called when entering the basicLiter production.
	EnterBasicLiter(c *BasicLiterContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterPar is called when entering the par production.
	EnterPar(c *ParContext)

	// EnterUnaryOper is called when entering the unaryOper production.
	EnterUnaryOper(c *UnaryOperContext)

	// EnterVar is called when entering the var production.
	EnterVar(c *VarContext)

	// EnterFact is called when entering the fact production.
	EnterFact(c *FactContext)

	// EnterInt is called when entering the int production.
	EnterInt(c *IntContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterFloat is called when entering the float production.
	EnterFloat(c *FloatContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// EnterNull is called when entering the null production.
	EnterNull(c *NullContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// ExitSourceFile is called when exiting the sourceFile production.
	ExitSourceFile(c *SourceFileContext)

	// ExitRuleID is called when exiting the ruleID production.
	ExitRuleID(c *RuleIDContext)

	// ExitConditionCheck is called when exiting the conditionCheck production.
	ExitConditionCheck(c *ConditionCheckContext)

	// ExitRunGore is called when exiting the runGore production.
	ExitRunGore(c *RunGoreContext)

	// ExitConditionList is called when exiting the conditionList production.
	ExitConditionList(c *ConditionListContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitIncDec is called when exiting the incDec production.
	ExitIncDec(c *IncDecContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitReturn is called when exiting the return production.
	ExitReturn(c *ReturnContext)

	// ExitIf is called when exiting the if production.
	ExitIf(c *IfContext)

	// ExitTry is called when exiting the try production.
	ExitTry(c *TryContext)

	// ExitBingo is called when exiting the bingo production.
	ExitBingo(c *BingoContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitThenStmt is called when exiting the thenStmt production.
	ExitThenStmt(c *ThenStmtContext)

	// ExitIncDecStmt is called when exiting the incDecStmt production.
	ExitIncDecStmt(c *IncDecStmtContext)

	// ExitLmath is called when exiting the lmath production.
	ExitLmath(c *LmathContext)

	// ExitMaxmin is called when exiting the maxmin production.
	ExitMaxmin(c *MaxminContext)

	// ExitCompare is called when exiting the compare production.
	ExitCompare(c *CompareContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitLike is called when exiting the like production.
	ExitLike(c *LikeContext)

	// ExitIn is called when exiting the in production.
	ExitIn(c *InContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitIsNull is called when exiting the isNull production.
	ExitIsNull(c *IsNullContext)

	// ExitUnary is called when exiting the unary production.
	ExitUnary(c *UnaryContext)

	// ExitHmath is called when exiting the hmath production.
	ExitHmath(c *HmathContext)

	// ExitBit is called when exiting the bit production.
	ExitBit(c *BitContext)

	// ExitUnaryOp is called when exiting the unaryOp production.
	ExitUnaryOp(c *UnaryOpContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitOper is called when exiting the oper production.
	ExitOper(c *OperContext)

	// ExitConvert is called when exiting the convert production.
	ExitConvert(c *ConvertContext)

	// ExitRound is called when exiting the round production.
	ExitRound(c *RoundContext)

	// ExitBasicLiter is called when exiting the basicLiter production.
	ExitBasicLiter(c *BasicLiterContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitPar is called when exiting the par production.
	ExitPar(c *ParContext)

	// ExitUnaryOper is called when exiting the unaryOper production.
	ExitUnaryOper(c *UnaryOperContext)

	// ExitVar is called when exiting the var production.
	ExitVar(c *VarContext)

	// ExitFact is called when exiting the fact production.
	ExitFact(c *FactContext)

	// ExitInt is called when exiting the int production.
	ExitInt(c *IntContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitFloat is called when exiting the float production.
	ExitFloat(c *FloatContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)

	// ExitNull is called when exiting the null production.
	ExitNull(c *NullContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)
}
