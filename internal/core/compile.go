package core

import (
	"context"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"gore/exception"
	"gore/internal/grammar"
	"strings"
)

type Compiler struct {
	*grammar.BaseGoreListener
	rule *Rule
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) Init() {
	c.rule = NewRule()
}

func (c *Compiler) CompileFile(ctx context.Context, filename string) (*Rule, error) {
	c.Init()
	inputStream, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, fmt.Errorf("compile input file error:%s", err.Error())
	}
	lexer := grammar.NewGoreLexer(inputStream)
	lexer.AddErrorListener(NewCompileErrorListener())

	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := grammar.NewGoreParser(tokenStream)
	parser.AddErrorListener(NewCompileErrorListener())
	parser.BuildParseTrees = true
	return c.compile(ctx, parser)
}

func (c *Compiler) CompileInput(ctx context.Context, src string) (*Rule, error) {
	c.Init()
	input := antlr.NewInputStream(src)
	lexer := grammar.NewGoreLexer(input)
	lexer.AddErrorListener(&CompileErrorListener{})
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := grammar.NewGoreParser(tokenStream)
	parser.AddErrorListener(&CompileErrorListener{})
	parser.BuildParseTrees = true
	return c.compile(ctx, parser)
}

func (c *Compiler) compile(ctx context.Context, parser *grammar.GoreParser) (rule *Rule, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = exception.HandlePanic(ctx, rec)
		}
	}()
	antlr.ParseTreeWalkerDefault.Walk(c, parser.SourceFile())
	return c.rule, nil
}

// for test
func (c *Compiler) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("Enter:", ctx.GetText())
}

// EnterRuleID is called when production ruleID is entered.
func (c *Compiler) EnterRuleID(ctx *grammar.RuleIDContext) {
	c.rule.ID = ctx.GetId().GetText()
	c.rule.Version = ctx.GetVersion().GetText()
}

// ExitConditionCheck is called when production conditionCheck is exited.
func (c *Compiler) ExitConditionCheck(ctx *grammar.ConditionCheckContext) {
	if compileErr := c.rule.Condition.CompileCondition(ctx); compileErr != nil {
		panic(fmt.Sprintf("compile condition error:%s", compileErr.Error()))
	}
}

// ExitRunGore is called when production runGore is exited.
func (c *Compiler) ExitRunGore(ctx *grammar.RunGoreContext) {
	if compileErr := c.rule.Runner.CompileRunner(ctx); compileErr != nil {
		panic(fmt.Sprintf("compile run error:%s", compileErr.Error()))
	}
}

// ExitVar is called when production var is exited.
func (c *Compiler) ExitVar(ctx *grammar.VarContext) {
	c.rule.SymbolTable.AddIdentifier(ctx.GetText())
}

// ExitFact is called when production fact is exited.
func (c *Compiler) ExitFact(ctx *grammar.FactContext) {
	c.rule.SymbolTable.AddIdentifier(ctx.GetText())
}

type SymbolTable struct {
	symbols map[string]Symbol
}

type Symbol struct {
	Type  string
	Field string
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbols: make(map[string]Symbol),
	}
}

func (st *SymbolTable) GetIdentifier(ident string) Symbol {
	if _, ok := st.symbols[ident]; ok {
		return st.symbols[ident]
	}
	return Symbol{"", ""}
}

func (st *SymbolTable) AddIdentifier(ident string) {
	if _, ok := st.symbols[ident]; ok {
		return
	}
	prefix, field := parseSymbol(ident)
	if !verifyIdentifier(prefix) {
		panic(fmt.Sprintf("invalid identifier:%s", ident))
	}
	st.symbols[ident] = Symbol{Type: prefix, Field: field}
}

func parseSymbol(symbol string) (string, string) {
	sp := strings.Split(symbol, ".")
	if len(sp) != 2 {
		panic(fmt.Sprintf("invalid symbol:%s", symbol))
	}
	return sp[0], sp[1]
}

func verifyIdentifier(iType string) bool {
	if _, ok := IdentifierTypeSet[iType]; !ok {
		return false
	}
	return true
}

type CompileErrorListener struct {
	*antlr.DefaultErrorListener
}

func NewCompileErrorListener() *CompileErrorListener {
	return &CompileErrorListener{}
}

func (el *CompileErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Errorf("syntax error: %d,%d, %s", line, column, msg))
}

func (el *CompileErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	el.DefaultErrorListener.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
	panic(fmt.Errorf("ambiguity at %s", ambigAlts.String()))
}
