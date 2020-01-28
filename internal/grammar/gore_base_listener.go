// Code generated from C:/Users/yangy/go/src/gstar1987td/gore/internal/grammar\Gore.g4 by ANTLR 4.8. DO NOT EDIT.

package grammar // Gore
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGoreListener is a complete listener for a parse tree produced by GoreParser.
type BaseGoreListener struct{}

var _ GoreListener = &BaseGoreListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoreListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoreListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoreListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoreListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *BaseGoreListener) EnterSourceFile(ctx *SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *BaseGoreListener) ExitSourceFile(ctx *SourceFileContext) {}

// EnterRuleID is called when production ruleID is entered.
func (s *BaseGoreListener) EnterRuleID(ctx *RuleIDContext) {}

// ExitRuleID is called when production ruleID is exited.
func (s *BaseGoreListener) ExitRuleID(ctx *RuleIDContext) {}

// EnterConditionCheck is called when production conditionCheck is entered.
func (s *BaseGoreListener) EnterConditionCheck(ctx *ConditionCheckContext) {}

// ExitConditionCheck is called when production conditionCheck is exited.
func (s *BaseGoreListener) ExitConditionCheck(ctx *ConditionCheckContext) {}

// EnterRunGore is called when production runGore is entered.
func (s *BaseGoreListener) EnterRunGore(ctx *RunGoreContext) {}

// ExitRunGore is called when production runGore is exited.
func (s *BaseGoreListener) ExitRunGore(ctx *RunGoreContext) {}

// EnterConditionList is called when production conditionList is entered.
func (s *BaseGoreListener) EnterConditionList(ctx *ConditionListContext) {}

// ExitConditionList is called when production conditionList is exited.
func (s *BaseGoreListener) ExitConditionList(ctx *ConditionListContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseGoreListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseGoreListener) ExitStatementList(ctx *StatementListContext) {}

// EnterIncDec is called when production incDec is entered.
func (s *BaseGoreListener) EnterIncDec(ctx *IncDecContext) {}

// ExitIncDec is called when production incDec is exited.
func (s *BaseGoreListener) ExitIncDec(ctx *IncDecContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseGoreListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseGoreListener) ExitAssign(ctx *AssignContext) {}

// EnterReturn is called when production return is entered.
func (s *BaseGoreListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production return is exited.
func (s *BaseGoreListener) ExitReturn(ctx *ReturnContext) {}

// EnterIf is called when production if is entered.
func (s *BaseGoreListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production if is exited.
func (s *BaseGoreListener) ExitIf(ctx *IfContext) {}

// EnterTry is called when production try is entered.
func (s *BaseGoreListener) EnterTry(ctx *TryContext) {}

// ExitTry is called when production try is exited.
func (s *BaseGoreListener) ExitTry(ctx *TryContext) {}

// EnterBingo is called when production bingo is entered.
func (s *BaseGoreListener) EnterBingo(ctx *BingoContext) {}

// ExitBingo is called when production bingo is exited.
func (s *BaseGoreListener) ExitBingo(ctx *BingoContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseGoreListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseGoreListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterThenStmt is called when production thenStmt is entered.
func (s *BaseGoreListener) EnterThenStmt(ctx *ThenStmtContext) {}

// ExitThenStmt is called when production thenStmt is exited.
func (s *BaseGoreListener) ExitThenStmt(ctx *ThenStmtContext) {}

// EnterIncDecStmt is called when production incDecStmt is entered.
func (s *BaseGoreListener) EnterIncDecStmt(ctx *IncDecStmtContext) {}

// ExitIncDecStmt is called when production incDecStmt is exited.
func (s *BaseGoreListener) ExitIncDecStmt(ctx *IncDecStmtContext) {}

// EnterLmath is called when production lmath is entered.
func (s *BaseGoreListener) EnterLmath(ctx *LmathContext) {}

// ExitLmath is called when production lmath is exited.
func (s *BaseGoreListener) ExitLmath(ctx *LmathContext) {}

// EnterMaxmin is called when production maxmin is entered.
func (s *BaseGoreListener) EnterMaxmin(ctx *MaxminContext) {}

// ExitMaxmin is called when production maxmin is exited.
func (s *BaseGoreListener) ExitMaxmin(ctx *MaxminContext) {}

// EnterCompare is called when production compare is entered.
func (s *BaseGoreListener) EnterCompare(ctx *CompareContext) {}

// ExitCompare is called when production compare is exited.
func (s *BaseGoreListener) ExitCompare(ctx *CompareContext) {}

// EnterOr is called when production or is entered.
func (s *BaseGoreListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production or is exited.
func (s *BaseGoreListener) ExitOr(ctx *OrContext) {}

// EnterLike is called when production like is entered.
func (s *BaseGoreListener) EnterLike(ctx *LikeContext) {}

// ExitLike is called when production like is exited.
func (s *BaseGoreListener) ExitLike(ctx *LikeContext) {}

// EnterIn is called when production in is entered.
func (s *BaseGoreListener) EnterIn(ctx *InContext) {}

// ExitIn is called when production in is exited.
func (s *BaseGoreListener) ExitIn(ctx *InContext) {}

// EnterAnd is called when production and is entered.
func (s *BaseGoreListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BaseGoreListener) ExitAnd(ctx *AndContext) {}

// EnterIsNull is called when production isNull is entered.
func (s *BaseGoreListener) EnterIsNull(ctx *IsNullContext) {}

// ExitIsNull is called when production isNull is exited.
func (s *BaseGoreListener) ExitIsNull(ctx *IsNullContext) {}

// EnterUnary is called when production unary is entered.
func (s *BaseGoreListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production unary is exited.
func (s *BaseGoreListener) ExitUnary(ctx *UnaryContext) {}

// EnterHmath is called when production hmath is entered.
func (s *BaseGoreListener) EnterHmath(ctx *HmathContext) {}

// ExitHmath is called when production hmath is exited.
func (s *BaseGoreListener) ExitHmath(ctx *HmathContext) {}

// EnterBit is called when production bit is entered.
func (s *BaseGoreListener) EnterBit(ctx *BitContext) {}

// ExitBit is called when production bit is exited.
func (s *BaseGoreListener) ExitBit(ctx *BitContext) {}

// EnterUnaryOp is called when production unaryOp is entered.
func (s *BaseGoreListener) EnterUnaryOp(ctx *UnaryOpContext) {}

// ExitUnaryOp is called when production unaryOp is exited.
func (s *BaseGoreListener) ExitUnaryOp(ctx *UnaryOpContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseGoreListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseGoreListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterOper is called when production oper is entered.
func (s *BaseGoreListener) EnterOper(ctx *OperContext) {}

// ExitOper is called when production oper is exited.
func (s *BaseGoreListener) ExitOper(ctx *OperContext) {}

// EnterConvert is called when production convert is entered.
func (s *BaseGoreListener) EnterConvert(ctx *ConvertContext) {}

// ExitConvert is called when production convert is exited.
func (s *BaseGoreListener) ExitConvert(ctx *ConvertContext) {}

// EnterRound is called when production round is entered.
func (s *BaseGoreListener) EnterRound(ctx *RoundContext) {}

// ExitRound is called when production round is exited.
func (s *BaseGoreListener) ExitRound(ctx *RoundContext) {}

// EnterBasicLiter is called when production basicLiter is entered.
func (s *BaseGoreListener) EnterBasicLiter(ctx *BasicLiterContext) {}

// ExitBasicLiter is called when production basicLiter is exited.
func (s *BaseGoreListener) ExitBasicLiter(ctx *BasicLiterContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseGoreListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseGoreListener) ExitIdent(ctx *IdentContext) {}

// EnterPar is called when production par is entered.
func (s *BaseGoreListener) EnterPar(ctx *ParContext) {}

// ExitPar is called when production par is exited.
func (s *BaseGoreListener) ExitPar(ctx *ParContext) {}

// EnterUnaryOper is called when production unaryOper is entered.
func (s *BaseGoreListener) EnterUnaryOper(ctx *UnaryOperContext) {}

// ExitUnaryOper is called when production unaryOper is exited.
func (s *BaseGoreListener) ExitUnaryOper(ctx *UnaryOperContext) {}

// EnterVar is called when production var is entered.
func (s *BaseGoreListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production var is exited.
func (s *BaseGoreListener) ExitVar(ctx *VarContext) {}

// EnterFact is called when production fact is entered.
func (s *BaseGoreListener) EnterFact(ctx *FactContext) {}

// ExitFact is called when production fact is exited.
func (s *BaseGoreListener) ExitFact(ctx *FactContext) {}

// EnterInt is called when production int is entered.
func (s *BaseGoreListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BaseGoreListener) ExitInt(ctx *IntContext) {}

// EnterString is called when production string is entered.
func (s *BaseGoreListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseGoreListener) ExitString(ctx *StringContext) {}

// EnterFloat is called when production float is entered.
func (s *BaseGoreListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production float is exited.
func (s *BaseGoreListener) ExitFloat(ctx *FloatContext) {}

// EnterBool is called when production bool is entered.
func (s *BaseGoreListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BaseGoreListener) ExitBool(ctx *BoolContext) {}

// EnterNull is called when production null is entered.
func (s *BaseGoreListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BaseGoreListener) ExitNull(ctx *NullContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseGoreListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseGoreListener) ExitInteger(ctx *IntegerContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseGoreListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseGoreListener) ExitEos(ctx *EosContext) {}
