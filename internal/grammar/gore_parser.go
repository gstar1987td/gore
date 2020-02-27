// Code generated from E:/Go/src/gore/internal/grammar\Gore.g4 by ANTLR 4.8. DO NOT EDIT.

package grammar // Gore
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 237,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 44, 10, 2, 12, 2,
	14, 2, 47, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 61, 10, 4, 3, 5, 3, 5, 3, 5, 7, 5, 66, 10, 5, 12,
	5, 14, 5, 69, 11, 5, 3, 6, 3, 6, 3, 6, 7, 6, 74, 10, 6, 12, 6, 14, 6, 77,
	11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 5, 7, 91, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 99, 10,
	8, 5, 8, 101, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 107, 10, 9, 12, 9, 14,
	9, 110, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 6, 11, 123, 10, 11, 13, 11, 14, 11, 124, 3, 11, 3, 11,
	3, 11, 5, 11, 130, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 153, 10, 11, 3, 11, 3, 11, 3, 11, 5,
	11, 158, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 164, 10, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 6, 11, 171, 10, 11, 13, 11, 14, 11, 172, 3,
	11, 3, 11, 7, 11, 177, 10, 11, 12, 11, 14, 11, 180, 11, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 5, 12, 186, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 5, 13, 194, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	5, 13, 203, 10, 13, 3, 13, 3, 13, 5, 13, 207, 10, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 5, 14, 215, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16,
	5, 16, 221, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 228, 10,
	17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 5, 19, 235, 10, 19, 3, 19, 2, 3,
	20, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 2, 7, 3, 2, 38, 39, 3, 2, 23, 24, 3, 2, 26, 28, 3, 2, 29, 30, 4, 2,
	31, 31, 47, 48, 2, 259, 2, 38, 3, 2, 2, 2, 4, 48, 3, 2, 2, 2, 6, 60, 3,
	2, 2, 2, 8, 67, 3, 2, 2, 2, 10, 75, 3, 2, 2, 2, 12, 90, 3, 2, 2, 2, 14,
	92, 3, 2, 2, 2, 16, 102, 3, 2, 2, 2, 18, 113, 3, 2, 2, 2, 20, 129, 3, 2,
	2, 2, 22, 185, 3, 2, 2, 2, 24, 206, 3, 2, 2, 2, 26, 214, 3, 2, 2, 2, 28,
	216, 3, 2, 2, 2, 30, 220, 3, 2, 2, 2, 32, 227, 3, 2, 2, 2, 34, 229, 3,
	2, 2, 2, 36, 234, 3, 2, 2, 2, 38, 39, 5, 4, 3, 2, 39, 45, 5, 36, 19, 2,
	40, 41, 5, 6, 4, 2, 41, 42, 5, 36, 19, 2, 42, 44, 3, 2, 2, 2, 43, 40, 3,
	2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46,
	3, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 49, 7, 3, 2, 2, 49, 50, 7, 35, 2,
	2, 50, 51, 7, 34, 2, 2, 51, 5, 3, 2, 2, 2, 52, 53, 7, 4, 2, 2, 53, 54,
	5, 8, 5, 2, 54, 55, 7, 5, 2, 2, 55, 61, 3, 2, 2, 2, 56, 57, 7, 6, 2, 2,
	57, 58, 5, 10, 6, 2, 58, 59, 7, 5, 2, 2, 59, 61, 3, 2, 2, 2, 60, 52, 3,
	2, 2, 2, 60, 56, 3, 2, 2, 2, 61, 7, 3, 2, 2, 2, 62, 63, 5, 20, 11, 2, 63,
	64, 5, 36, 19, 2, 64, 66, 3, 2, 2, 2, 65, 62, 3, 2, 2, 2, 66, 69, 3, 2,
	2, 2, 67, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 9, 3, 2, 2, 2, 69, 67,
	3, 2, 2, 2, 70, 71, 5, 12, 7, 2, 71, 72, 5, 36, 19, 2, 72, 74, 3, 2, 2,
	2, 73, 70, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76,
	3, 2, 2, 2, 76, 11, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 91, 5, 18, 10,
	2, 79, 80, 7, 36, 2, 2, 80, 81, 7, 40, 2, 2, 81, 91, 5, 20, 11, 2, 82,
	91, 7, 7, 2, 2, 83, 91, 5, 14, 8, 2, 84, 85, 7, 8, 2, 2, 85, 86, 5, 10,
	6, 2, 86, 87, 7, 9, 2, 2, 87, 88, 5, 10, 6, 2, 88, 91, 3, 2, 2, 2, 89,
	91, 7, 46, 2, 2, 90, 78, 3, 2, 2, 2, 90, 79, 3, 2, 2, 2, 90, 82, 3, 2,
	2, 2, 90, 83, 3, 2, 2, 2, 90, 84, 3, 2, 2, 2, 90, 89, 3, 2, 2, 2, 91, 13,
	3, 2, 2, 2, 92, 93, 7, 10, 2, 2, 93, 94, 5, 20, 11, 2, 94, 100, 5, 16,
	9, 2, 95, 98, 7, 11, 2, 2, 96, 99, 5, 14, 8, 2, 97, 99, 5, 16, 9, 2, 98,
	96, 3, 2, 2, 2, 98, 97, 3, 2, 2, 2, 99, 101, 3, 2, 2, 2, 100, 95, 3, 2,
	2, 2, 100, 101, 3, 2, 2, 2, 101, 15, 3, 2, 2, 2, 102, 108, 7, 12, 2, 2,
	103, 104, 5, 12, 7, 2, 104, 105, 5, 36, 19, 2, 105, 107, 3, 2, 2, 2, 106,
	103, 3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109,
	3, 2, 2, 2, 109, 111, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 112, 7, 13,
	2, 2, 112, 17, 3, 2, 2, 2, 113, 114, 7, 36, 2, 2, 114, 115, 9, 2, 2, 2,
	115, 19, 3, 2, 2, 2, 116, 117, 8, 11, 1, 2, 117, 118, 9, 3, 2, 2, 118,
	119, 7, 18, 2, 2, 119, 122, 5, 20, 11, 2, 120, 121, 7, 19, 2, 2, 121, 123,
	5, 20, 11, 2, 122, 120, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 122, 3,
	2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 7, 20, 2,
	2, 127, 130, 3, 2, 2, 2, 128, 130, 5, 22, 12, 2, 129, 116, 3, 2, 2, 2,
	129, 128, 3, 2, 2, 2, 130, 178, 3, 2, 2, 2, 131, 132, 12, 10, 2, 2, 132,
	133, 7, 41, 2, 2, 133, 177, 5, 20, 11, 11, 134, 135, 12, 9, 2, 2, 135,
	136, 7, 42, 2, 2, 136, 177, 5, 20, 11, 10, 137, 138, 12, 8, 2, 2, 138,
	139, 7, 43, 2, 2, 139, 177, 5, 20, 11, 9, 140, 141, 12, 7, 2, 2, 141, 142,
	7, 44, 2, 2, 142, 177, 5, 20, 11, 8, 143, 144, 12, 6, 2, 2, 144, 145, 7,
	21, 2, 2, 145, 177, 5, 20, 11, 7, 146, 147, 12, 5, 2, 2, 147, 148, 7, 22,
	2, 2, 148, 177, 5, 20, 11, 6, 149, 150, 12, 13, 2, 2, 150, 152, 7, 14,
	2, 2, 151, 153, 7, 15, 2, 2, 152, 151, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2,
	153, 154, 3, 2, 2, 2, 154, 177, 7, 33, 2, 2, 155, 157, 12, 12, 2, 2, 156,
	158, 7, 15, 2, 2, 157, 156, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159,
	3, 2, 2, 2, 159, 160, 7, 16, 2, 2, 160, 177, 7, 32, 2, 2, 161, 163, 12,
	11, 2, 2, 162, 164, 7, 15, 2, 2, 163, 162, 3, 2, 2, 2, 163, 164, 3, 2,
	2, 2, 164, 165, 3, 2, 2, 2, 165, 166, 7, 17, 2, 2, 166, 167, 7, 18, 2,
	2, 167, 170, 5, 20, 11, 2, 168, 169, 7, 19, 2, 2, 169, 171, 5, 20, 11,
	2, 170, 168, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172,
	173, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 175, 7, 20, 2, 2, 175, 177,
	3, 2, 2, 2, 176, 131, 3, 2, 2, 2, 176, 134, 3, 2, 2, 2, 176, 137, 3, 2,
	2, 2, 176, 140, 3, 2, 2, 2, 176, 143, 3, 2, 2, 2, 176, 146, 3, 2, 2, 2,
	176, 149, 3, 2, 2, 2, 176, 155, 3, 2, 2, 2, 176, 161, 3, 2, 2, 2, 177,
	180, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 21, 3,
	2, 2, 2, 180, 178, 3, 2, 2, 2, 181, 182, 5, 28, 15, 2, 182, 183, 5, 22,
	12, 2, 183, 186, 3, 2, 2, 2, 184, 186, 5, 24, 13, 2, 185, 181, 3, 2, 2,
	2, 185, 184, 3, 2, 2, 2, 186, 23, 3, 2, 2, 2, 187, 207, 5, 26, 14, 2, 188,
	189, 7, 45, 2, 2, 189, 190, 7, 18, 2, 2, 190, 193, 5, 20, 11, 2, 191, 192,
	7, 19, 2, 2, 192, 194, 5, 32, 17, 2, 193, 191, 3, 2, 2, 2, 193, 194, 3,
	2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 196, 7, 20, 2, 2, 196, 207, 3, 2, 2,
	2, 197, 198, 7, 25, 2, 2, 198, 199, 7, 18, 2, 2, 199, 202, 5, 20, 11, 2,
	200, 201, 7, 19, 2, 2, 201, 203, 7, 31, 2, 2, 202, 200, 3, 2, 2, 2, 202,
	203, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 205, 7, 20, 2, 2, 205, 207,
	3, 2, 2, 2, 206, 187, 3, 2, 2, 2, 206, 188, 3, 2, 2, 2, 206, 197, 3, 2,
	2, 2, 207, 25, 3, 2, 2, 2, 208, 215, 5, 32, 17, 2, 209, 215, 5, 30, 16,
	2, 210, 211, 7, 18, 2, 2, 211, 212, 5, 20, 11, 2, 212, 213, 7, 20, 2, 2,
	213, 215, 3, 2, 2, 2, 214, 208, 3, 2, 2, 2, 214, 209, 3, 2, 2, 2, 214,
	210, 3, 2, 2, 2, 215, 27, 3, 2, 2, 2, 216, 217, 9, 4, 2, 2, 217, 29, 3,
	2, 2, 2, 218, 221, 7, 36, 2, 2, 219, 221, 7, 37, 2, 2, 220, 218, 3, 2,
	2, 2, 220, 219, 3, 2, 2, 2, 221, 31, 3, 2, 2, 2, 222, 228, 5, 34, 18, 2,
	223, 228, 7, 32, 2, 2, 224, 228, 7, 34, 2, 2, 225, 228, 9, 5, 2, 2, 226,
	228, 7, 33, 2, 2, 227, 222, 3, 2, 2, 2, 227, 223, 3, 2, 2, 2, 227, 224,
	3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 227, 226, 3, 2, 2, 2, 228, 33, 3, 2,
	2, 2, 229, 230, 9, 6, 2, 2, 230, 35, 3, 2, 2, 2, 231, 235, 7, 2, 2, 3,
	232, 235, 6, 19, 11, 2, 233, 235, 6, 19, 12, 2, 234, 231, 3, 2, 2, 2, 234,
	232, 3, 2, 2, 2, 234, 233, 3, 2, 2, 2, 235, 37, 3, 2, 2, 2, 26, 45, 60,
	67, 75, 90, 98, 100, 108, 124, 129, 152, 157, 163, 172, 176, 178, 185,
	193, 202, 206, 214, 220, 227, 234,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'rule:'", "'condition:'", "';'", "'run:'", "'return'", "'try:'", "'exception:'",
	"'if'", "'else'", "'then'", "'end'", "'is'", "'not'", "'like'", "'in'",
	"'('", "','", "')'", "'&&'", "'||'", "'max'", "'min'", "'round'", "'!'",
	"'~'", "'-'", "'true'", "'false'", "", "", "'null'", "", "", "", "", "'++'",
	"'--'", "'='", "", "", "", "", "", "'bingo'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "DECIMAL_LIT", "RAW_STRING_LIT",
	"NIL_LIT", "FLOAT_LIT", "IDENTIFIER", "VARIABLE", "FACT", "PLUS_PLUS",
	"MINUS_MINUS", "ASSIGN", "HIGH_PRIORITY_MATH", "LOW_PRIORITY_MATH", "BIT_OPER",
	"COMPARE_OPER", "CONVERT_TYPE", "BINGO", "OCTAL_LIT", "HEX_LIT", "WS",
	"COMMENT", "TERMINATOR", "LINE_COMMENT",
}

var ruleNames = []string{
	"sourceFile", "ruleID", "context", "conditionList", "statementList", "statement",
	"ifStmt", "thenStmt", "incDecStmt", "expression", "unaryExp", "primaryExp",
	"operand", "unaryOper", "identifier", "basicLit", "integer", "eos",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GoreParser struct {
	*antlr.BaseParser
}

func NewGoreParser(input antlr.TokenStream) *GoreParser {
	this := new(GoreParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Gore.g4"

	return this
}

// GoreParser tokens.
const (
	GoreParserEOF                = antlr.TokenEOF
	GoreParserT__0               = 1
	GoreParserT__1               = 2
	GoreParserT__2               = 3
	GoreParserT__3               = 4
	GoreParserT__4               = 5
	GoreParserT__5               = 6
	GoreParserT__6               = 7
	GoreParserT__7               = 8
	GoreParserT__8               = 9
	GoreParserT__9               = 10
	GoreParserT__10              = 11
	GoreParserT__11              = 12
	GoreParserT__12              = 13
	GoreParserT__13              = 14
	GoreParserT__14              = 15
	GoreParserT__15              = 16
	GoreParserT__16              = 17
	GoreParserT__17              = 18
	GoreParserT__18              = 19
	GoreParserT__19              = 20
	GoreParserT__20              = 21
	GoreParserT__21              = 22
	GoreParserT__22              = 23
	GoreParserT__23              = 24
	GoreParserT__24              = 25
	GoreParserT__25              = 26
	GoreParserT__26              = 27
	GoreParserT__27              = 28
	GoreParserDECIMAL_LIT        = 29
	GoreParserRAW_STRING_LIT     = 30
	GoreParserNIL_LIT            = 31
	GoreParserFLOAT_LIT          = 32
	GoreParserIDENTIFIER         = 33
	GoreParserVARIABLE           = 34
	GoreParserFACT               = 35
	GoreParserPLUS_PLUS          = 36
	GoreParserMINUS_MINUS        = 37
	GoreParserASSIGN             = 38
	GoreParserHIGH_PRIORITY_MATH = 39
	GoreParserLOW_PRIORITY_MATH  = 40
	GoreParserBIT_OPER           = 41
	GoreParserCOMPARE_OPER       = 42
	GoreParserCONVERT_TYPE       = 43
	GoreParserBINGO              = 44
	GoreParserOCTAL_LIT          = 45
	GoreParserHEX_LIT            = 46
	GoreParserWS                 = 47
	GoreParserCOMMENT            = 48
	GoreParserTERMINATOR         = 49
	GoreParserLINE_COMMENT       = 50
)

// GoreParser rules.
const (
	GoreParserRULE_sourceFile    = 0
	GoreParserRULE_ruleID        = 1
	GoreParserRULE_context       = 2
	GoreParserRULE_conditionList = 3
	GoreParserRULE_statementList = 4
	GoreParserRULE_statement     = 5
	GoreParserRULE_ifStmt        = 6
	GoreParserRULE_thenStmt      = 7
	GoreParserRULE_incDecStmt    = 8
	GoreParserRULE_expression    = 9
	GoreParserRULE_unaryExp      = 10
	GoreParserRULE_primaryExp    = 11
	GoreParserRULE_operand       = 12
	GoreParserRULE_unaryOper     = 13
	GoreParserRULE_identifier    = 14
	GoreParserRULE_basicLit      = 15
	GoreParserRULE_integer       = 16
	GoreParserRULE_eos           = 17
)

// ISourceFileContext is an interface to support dynamic dispatch.
type ISourceFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceFileContext differentiates from other interfaces.
	IsSourceFileContext()
}

type SourceFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceFileContext() *SourceFileContext {
	var p = new(SourceFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_sourceFile
	return p
}

func (*SourceFileContext) IsSourceFileContext() {}

func NewSourceFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceFileContext {
	var p = new(SourceFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_sourceFile

	return p
}

func (s *SourceFileContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceFileContext) RuleID() IRuleIDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleIDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleIDContext)
}

func (s *SourceFileContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *SourceFileContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SourceFileContext) AllContext() []IContextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IContextContext)(nil)).Elem())
	var tst = make([]IContextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IContextContext)
		}
	}

	return tst
}

func (s *SourceFileContext) Context(i int) IContextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IContextContext)
}

func (s *SourceFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterSourceFile(s)
	}
}

func (s *SourceFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitSourceFile(s)
	}
}

func (s *SourceFileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitSourceFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) SourceFile() (localctx ISourceFileContext) {
	localctx = NewSourceFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoreParserRULE_sourceFile)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.RuleID()
	}
	{
		p.SetState(37)
		p.Eos()
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoreParserT__1 || _la == GoreParserT__3 {
		{
			p.SetState(38)
			p.Context()
		}
		{
			p.SetState(39)
			p.Eos()
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRuleIDContext is an interface to support dynamic dispatch.
type IRuleIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// GetVersion returns the version token.
	GetVersion() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// SetVersion sets the version token.
	SetVersion(antlr.Token)

	// IsRuleIDContext differentiates from other interfaces.
	IsRuleIDContext()
}

type RuleIDContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	id      antlr.Token
	version antlr.Token
}

func NewEmptyRuleIDContext() *RuleIDContext {
	var p = new(RuleIDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_ruleID
	return p
}

func (*RuleIDContext) IsRuleIDContext() {}

func NewRuleIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleIDContext {
	var p = new(RuleIDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_ruleID

	return p
}

func (s *RuleIDContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleIDContext) GetId() antlr.Token { return s.id }

func (s *RuleIDContext) GetVersion() antlr.Token { return s.version }

func (s *RuleIDContext) SetId(v antlr.Token) { s.id = v }

func (s *RuleIDContext) SetVersion(v antlr.Token) { s.version = v }

func (s *RuleIDContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GoreParserIDENTIFIER, 0)
}

func (s *RuleIDContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(GoreParserFLOAT_LIT, 0)
}

func (s *RuleIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterRuleID(s)
	}
}

func (s *RuleIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitRuleID(s)
	}
}

func (s *RuleIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitRuleID(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) RuleID() (localctx IRuleIDContext) {
	localctx = NewRuleIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoreParserRULE_ruleID)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(GoreParserT__0)
	}
	{
		p.SetState(47)

		var _m = p.Match(GoreParserIDENTIFIER)

		localctx.(*RuleIDContext).id = _m
	}
	{
		p.SetState(48)

		var _m = p.Match(GoreParserFLOAT_LIT)

		localctx.(*RuleIDContext).version = _m
	}

	return localctx
}

// IContextContext is an interface to support dynamic dispatch.
type IContextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContextContext differentiates from other interfaces.
	IsContextContext()
}

type ContextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContextContext() *ContextContext {
	var p = new(ContextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_context
	return p
}

func (*ContextContext) IsContextContext() {}

func NewContextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContextContext {
	var p = new(ContextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_context

	return p
}

func (s *ContextContext) GetParser() antlr.Parser { return s.parser }

func (s *ContextContext) CopyFrom(ctx *ContextContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ContextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConditionCheckContext struct {
	*ContextContext
}

func NewConditionCheckContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionCheckContext {
	var p = new(ConditionCheckContext)

	p.ContextContext = NewEmptyContextContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ContextContext))

	return p
}

func (s *ConditionCheckContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionCheckContext) ConditionList() IConditionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionListContext)
}

func (s *ConditionCheckContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterConditionCheck(s)
	}
}

func (s *ConditionCheckContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitConditionCheck(s)
	}
}

func (s *ConditionCheckContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitConditionCheck(s)

	default:
		return t.VisitChildren(s)
	}
}

type RunGoreContext struct {
	*ContextContext
}

func NewRunGoreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RunGoreContext {
	var p = new(RunGoreContext)

	p.ContextContext = NewEmptyContextContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ContextContext))

	return p
}

func (s *RunGoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunGoreContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *RunGoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterRunGore(s)
	}
}

func (s *RunGoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitRunGore(s)
	}
}

func (s *RunGoreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitRunGore(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) Context() (localctx IContextContext) {
	localctx = NewContextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoreParserRULE_context)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(58)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoreParserT__1:
		localctx = NewConditionCheckContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(GoreParserT__1)
		}
		{
			p.SetState(51)
			p.ConditionList()
		}
		{
			p.SetState(52)
			p.Match(GoreParserT__2)
		}

	case GoreParserT__3:
		localctx = NewRunGoreContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Match(GoreParserT__3)
		}
		{
			p.SetState(55)
			p.StatementList()
		}
		{
			p.SetState(56)
			p.Match(GoreParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConditionListContext is an interface to support dynamic dispatch.
type IConditionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionListContext differentiates from other interfaces.
	IsConditionListContext()
}

type ConditionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionListContext() *ConditionListContext {
	var p = new(ConditionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_conditionList
	return p
}

func (*ConditionListContext) IsConditionListContext() {}

func NewConditionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionListContext {
	var p = new(ConditionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_conditionList

	return p
}

func (s *ConditionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionListContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *ConditionListContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ConditionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterConditionList(s)
	}
}

func (s *ConditionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitConditionList(s)
	}
}

func (s *ConditionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitConditionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) ConditionList() (localctx IConditionListContext) {
	localctx = NewConditionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoreParserRULE_conditionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(GoreParserT__15-16))|(1<<(GoreParserT__20-16))|(1<<(GoreParserT__21-16))|(1<<(GoreParserT__22-16))|(1<<(GoreParserT__23-16))|(1<<(GoreParserT__24-16))|(1<<(GoreParserT__25-16))|(1<<(GoreParserT__26-16))|(1<<(GoreParserT__27-16))|(1<<(GoreParserDECIMAL_LIT-16))|(1<<(GoreParserRAW_STRING_LIT-16))|(1<<(GoreParserNIL_LIT-16))|(1<<(GoreParserFLOAT_LIT-16))|(1<<(GoreParserVARIABLE-16))|(1<<(GoreParserFACT-16))|(1<<(GoreParserCONVERT_TYPE-16))|(1<<(GoreParserOCTAL_LIT-16))|(1<<(GoreParserHEX_LIT-16)))) != 0 {
		{
			p.SetState(60)
			p.expression(0)
		}
		{
			p.SetState(61)
			p.Eos()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *StatementListContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoreParserRULE_statementList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(68)
				p.Statement()
			}
			{
				p.SetState(69)
				p.Eos()
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BingoContext struct {
	*StatementContext
}

func NewBingoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BingoContext {
	var p = new(BingoContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *BingoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BingoContext) BINGO() antlr.TerminalNode {
	return s.GetToken(GoreParserBINGO, 0)
}

func (s *BingoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterBingo(s)
	}
}

func (s *BingoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitBingo(s)
	}
}

func (s *BingoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitBingo(s)

	default:
		return t.VisitChildren(s)
	}
}

type TryContext struct {
	*StatementContext
}

func NewTryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TryContext {
	var p = new(TryContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *TryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryContext) AllStatementList() []IStatementListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementListContext)(nil)).Elem())
	var tst = make([]IStatementListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementListContext)
		}
	}

	return tst
}

func (s *TryContext) StatementList(i int) IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *TryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterTry(s)
	}
}

func (s *TryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitTry(s)
	}
}

func (s *TryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitTry(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncDecContext struct {
	*StatementContext
}

func NewIncDecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncDecContext {
	var p = new(IncDecContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IncDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncDecContext) IncDecStmt() IIncDecStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncDecStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncDecStmtContext)
}

func (s *IncDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterIncDec(s)
	}
}

func (s *IncDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitIncDec(s)
	}
}

func (s *IncDecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitIncDec(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfContext struct {
	*StatementContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) IfStmt() IIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitIf(s)
	}
}

func (s *IfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnContext struct {
	*StatementContext
}

func NewReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnContext {
	var p = new(ReturnContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (s *ReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignContext struct {
	*StatementContext
	key   antlr.Token
	value IExpressionContext
}

func NewAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignContext {
	var p = new(AssignContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *AssignContext) GetKey() antlr.Token { return s.key }

func (s *AssignContext) SetKey(v antlr.Token) { s.key = v }

func (s *AssignContext) GetValue() IExpressionContext { return s.value }

func (s *AssignContext) SetValue(v IExpressionContext) { s.value = v }

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoreParserASSIGN, 0)
}

func (s *AssignContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(GoreParserVARIABLE, 0)
}

func (s *AssignContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoreParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIncDecContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.IncDecStmt()
		}

	case 2:
		localctx = NewAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)

			var _m = p.Match(GoreParserVARIABLE)

			localctx.(*AssignContext).key = _m
		}
		{
			p.SetState(78)
			p.Match(GoreParserASSIGN)
		}
		{
			p.SetState(79)

			var _x = p.expression(0)

			localctx.(*AssignContext).value = _x
		}

	case 3:
		localctx = NewReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.Match(GoreParserT__4)
		}

	case 4:
		localctx = NewIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(81)
			p.IfStmt()
		}

	case 5:
		localctx = NewTryContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(82)
			p.Match(GoreParserT__5)
		}
		{
			p.SetState(83)
			p.StatementList()
		}
		{
			p.SetState(84)
			p.Match(GoreParserT__6)
		}
		{
			p.SetState(85)
			p.StatementList()
		}

	case 6:
		localctx = NewBingoContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(87)
			p.Match(GoreParserBINGO)
		}

	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionContext)

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   IExpressionContext
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) GetCond() IExpressionContext { return s.cond }

func (s *IfStmtContext) SetCond(v IExpressionContext) { s.cond = v }

func (s *IfStmtContext) AllThenStmt() []IThenStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IThenStmtContext)(nil)).Elem())
	var tst = make([]IThenStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IThenStmtContext)
		}
	}

	return tst
}

func (s *IfStmtContext) ThenStmt(i int) IThenStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThenStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IThenStmtContext)
}

func (s *IfStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) IfStmt() IIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoreParserRULE_ifStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(GoreParserT__7)
	}
	{
		p.SetState(91)

		var _x = p.expression(0)

		localctx.(*IfStmtContext).cond = _x
	}
	{
		p.SetState(92)
		p.ThenStmt()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(93)
			p.Match(GoreParserT__8)
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GoreParserT__7:
			{
				p.SetState(94)
				p.IfStmt()
			}

		case GoreParserT__9:
			{
				p.SetState(95)
				p.ThenStmt()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IThenStmtContext is an interface to support dynamic dispatch.
type IThenStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThenStmtContext differentiates from other interfaces.
	IsThenStmtContext()
}

type ThenStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThenStmtContext() *ThenStmtContext {
	var p = new(ThenStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_thenStmt
	return p
}

func (*ThenStmtContext) IsThenStmtContext() {}

func NewThenStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThenStmtContext {
	var p = new(ThenStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_thenStmt

	return p
}

func (s *ThenStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ThenStmtContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ThenStmtContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ThenStmtContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *ThenStmtContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ThenStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThenStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThenStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterThenStmt(s)
	}
}

func (s *ThenStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitThenStmt(s)
	}
}

func (s *ThenStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitThenStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) ThenStmt() (localctx IThenStmtContext) {
	localctx = NewThenStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoreParserRULE_thenStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(GoreParserT__9)
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoreParserT__4)|(1<<GoreParserT__5)|(1<<GoreParserT__7))) != 0) || _la == GoreParserVARIABLE || _la == GoreParserBINGO {
		{
			p.SetState(101)
			p.Statement()
		}
		{
			p.SetState(102)
			p.Eos()
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(109)
		p.Match(GoreParserT__10)
	}

	return localctx
}

// IIncDecStmtContext is an interface to support dynamic dispatch.
type IIncDecStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncDecStmtContext differentiates from other interfaces.
	IsIncDecStmtContext()
}

type IncDecStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncDecStmtContext() *IncDecStmtContext {
	var p = new(IncDecStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_incDecStmt
	return p
}

func (*IncDecStmtContext) IsIncDecStmtContext() {}

func NewIncDecStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncDecStmtContext {
	var p = new(IncDecStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_incDecStmt

	return p
}

func (s *IncDecStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IncDecStmtContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(GoreParserVARIABLE, 0)
}

func (s *IncDecStmtContext) PLUS_PLUS() antlr.TerminalNode {
	return s.GetToken(GoreParserPLUS_PLUS, 0)
}

func (s *IncDecStmtContext) MINUS_MINUS() antlr.TerminalNode {
	return s.GetToken(GoreParserMINUS_MINUS, 0)
}

func (s *IncDecStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncDecStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncDecStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterIncDecStmt(s)
	}
}

func (s *IncDecStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitIncDecStmt(s)
	}
}

func (s *IncDecStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitIncDecStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) IncDecStmt() (localctx IIncDecStmtContext) {
	localctx = NewIncDecStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoreParserRULE_incDecStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(GoreParserVARIABLE)
	}
	{
		p.SetState(112)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoreParserPLUS_PLUS || _la == GoreParserMINUS_MINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LmathContext struct {
	*ExpressionContext
	l IExpressionContext
	r IExpressionContext
}

func NewLmathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LmathContext {
	var p = new(LmathContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LmathContext) GetL() IExpressionContext { return s.l }

func (s *LmathContext) GetR() IExpressionContext { return s.r }

func (s *LmathContext) SetL(v IExpressionContext) { s.l = v }

func (s *LmathContext) SetR(v IExpressionContext) { s.r = v }

func (s *LmathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LmathContext) LOW_PRIORITY_MATH() antlr.TerminalNode {
	return s.GetToken(GoreParserLOW_PRIORITY_MATH, 0)
}

func (s *LmathContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LmathContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LmathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterLmath(s)
	}
}

func (s *LmathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitLmath(s)
	}
}

func (s *LmathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitLmath(s)

	default:
		return t.VisitChildren(s)
	}
}

type MaxminContext struct {
	*ExpressionContext
}

func NewMaxminContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MaxminContext {
	var p = new(MaxminContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MaxminContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaxminContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MaxminContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MaxminContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterMaxmin(s)
	}
}

func (s *MaxminContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitMaxmin(s)
	}
}

func (s *MaxminContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitMaxmin(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareContext struct {
	*ExpressionContext
	l IExpressionContext
	r IExpressionContext
}

func NewCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareContext {
	var p = new(CompareContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CompareContext) GetL() IExpressionContext { return s.l }

func (s *CompareContext) GetR() IExpressionContext { return s.r }

func (s *CompareContext) SetL(v IExpressionContext) { s.l = v }

func (s *CompareContext) SetR(v IExpressionContext) { s.r = v }

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) COMPARE_OPER() antlr.TerminalNode {
	return s.GetToken(GoreParserCOMPARE_OPER, 0)
}

func (s *CompareContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CompareContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterCompare(s)
	}
}

func (s *CompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitCompare(s)
	}
}

func (s *CompareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitCompare(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrContext struct {
	*ExpressionContext
	l IExpressionContext
	r IExpressionContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OrContext) GetL() IExpressionContext { return s.l }

func (s *OrContext) GetR() IExpressionContext { return s.r }

func (s *OrContext) SetL(v IExpressionContext) { s.l = v }

func (s *OrContext) SetR(v IExpressionContext) { s.r = v }

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OrContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LikeContext struct {
	*ExpressionContext
	n antlr.Token
}

func NewLikeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LikeContext {
	var p = new(LikeContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LikeContext) GetN() antlr.Token { return s.n }

func (s *LikeContext) SetN(v antlr.Token) { s.n = v }

func (s *LikeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LikeContext) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoreParserRAW_STRING_LIT, 0)
}

func (s *LikeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterLike(s)
	}
}

func (s *LikeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitLike(s)
	}
}

func (s *LikeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitLike(s)

	default:
		return t.VisitChildren(s)
	}
}

type InContext struct {
	*ExpressionContext
	n antlr.Token
}

func NewInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InContext {
	var p = new(InContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InContext) GetN() antlr.Token { return s.n }

func (s *InContext) SetN(v antlr.Token) { s.n = v }

func (s *InContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *InContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterIn(s)
	}
}

func (s *InContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitIn(s)
	}
}

func (s *InContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitIn(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	*ExpressionContext
	l IExpressionContext
	r IExpressionContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndContext) GetL() IExpressionContext { return s.l }

func (s *AndContext) GetR() IExpressionContext { return s.r }

func (s *AndContext) SetL(v IExpressionContext) { s.l = v }

func (s *AndContext) SetR(v IExpressionContext) { s.r = v }

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AndContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsNullContext struct {
	*ExpressionContext
	n antlr.Token
}

func NewIsNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNullContext {
	var p = new(IsNullContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IsNullContext) GetN() antlr.Token { return s.n }

func (s *IsNullContext) SetN(v antlr.Token) { s.n = v }

func (s *IsNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsNullContext) NIL_LIT() antlr.TerminalNode {
	return s.GetToken(GoreParserNIL_LIT, 0)
}

func (s *IsNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterIsNull(s)
	}
}

func (s *IsNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitIsNull(s)
	}
}

func (s *IsNullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitIsNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryContext struct {
	*ExpressionContext
}

func NewUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryContext {
	var p = new(UnaryContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) UnaryExp() IUnaryExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpContext)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitUnary(s)
	}
}

func (s *UnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

type HmathContext struct {
	*ExpressionContext
	l IExpressionContext
	r IExpressionContext
}

func NewHmathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HmathContext {
	var p = new(HmathContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *HmathContext) GetL() IExpressionContext { return s.l }

func (s *HmathContext) GetR() IExpressionContext { return s.r }

func (s *HmathContext) SetL(v IExpressionContext) { s.l = v }

func (s *HmathContext) SetR(v IExpressionContext) { s.r = v }

func (s *HmathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HmathContext) HIGH_PRIORITY_MATH() antlr.TerminalNode {
	return s.GetToken(GoreParserHIGH_PRIORITY_MATH, 0)
}

func (s *HmathContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *HmathContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HmathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterHmath(s)
	}
}

func (s *HmathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitHmath(s)
	}
}

func (s *HmathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitHmath(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitContext struct {
	*ExpressionContext
	l IExpressionContext
	r IExpressionContext
}

func NewBitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitContext {
	var p = new(BitContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BitContext) GetL() IExpressionContext { return s.l }

func (s *BitContext) GetR() IExpressionContext { return s.r }

func (s *BitContext) SetL(v IExpressionContext) { s.l = v }

func (s *BitContext) SetR(v IExpressionContext) { s.r = v }

func (s *BitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitContext) BIT_OPER() antlr.TerminalNode {
	return s.GetToken(GoreParserBIT_OPER, 0)
}

func (s *BitContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BitContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterBit(s)
	}
}

func (s *BitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitBit(s)
	}
}

func (s *BitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitBit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *GoreParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, GoreParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoreParserT__20, GoreParserT__21:
		localctx = NewMaxminContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(115)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GoreParserT__20 || _la == GoreParserT__21) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(116)
			p.Match(GoreParserT__15)
		}
		{
			p.SetState(117)
			p.expression(0)
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == GoreParserT__16 {
			{
				p.SetState(118)
				p.Match(GoreParserT__16)
			}
			{
				p.SetState(119)
				p.expression(0)
			}

			p.SetState(122)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(124)
			p.Match(GoreParserT__17)
		}

	case GoreParserT__15, GoreParserT__22, GoreParserT__23, GoreParserT__24, GoreParserT__25, GoreParserT__26, GoreParserT__27, GoreParserDECIMAL_LIT, GoreParserRAW_STRING_LIT, GoreParserNIL_LIT, GoreParserFLOAT_LIT, GoreParserVARIABLE, GoreParserFACT, GoreParserCONVERT_TYPE, GoreParserOCTAL_LIT, GoreParserHEX_LIT:
		localctx = NewUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(126)
			p.UnaryExp()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(174)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewHmathContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*HmathContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GoreParserRULE_expression)
				p.SetState(129)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(130)
					p.Match(GoreParserHIGH_PRIORITY_MATH)
				}
				{
					p.SetState(131)

					var _x = p.expression(9)

					localctx.(*HmathContext).r = _x
				}

			case 2:
				localctx = NewLmathContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*LmathContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GoreParserRULE_expression)
				p.SetState(132)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(133)
					p.Match(GoreParserLOW_PRIORITY_MATH)
				}
				{
					p.SetState(134)

					var _x = p.expression(8)

					localctx.(*LmathContext).r = _x
				}

			case 3:
				localctx = NewBitContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BitContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GoreParserRULE_expression)
				p.SetState(135)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(136)
					p.Match(GoreParserBIT_OPER)
				}
				{
					p.SetState(137)

					var _x = p.expression(7)

					localctx.(*BitContext).r = _x
				}

			case 4:
				localctx = NewCompareContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*CompareContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GoreParserRULE_expression)
				p.SetState(138)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(139)
					p.Match(GoreParserCOMPARE_OPER)
				}
				{
					p.SetState(140)

					var _x = p.expression(6)

					localctx.(*CompareContext).r = _x
				}

			case 5:
				localctx = NewAndContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*AndContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GoreParserRULE_expression)
				p.SetState(141)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(142)
					p.Match(GoreParserT__18)
				}
				{
					p.SetState(143)

					var _x = p.expression(5)

					localctx.(*AndContext).r = _x
				}

			case 6:
				localctx = NewOrContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*OrContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GoreParserRULE_expression)
				p.SetState(144)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(145)
					p.Match(GoreParserT__19)
				}
				{
					p.SetState(146)

					var _x = p.expression(4)

					localctx.(*OrContext).r = _x
				}

			case 7:
				localctx = NewIsNullContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoreParserRULE_expression)
				p.SetState(147)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(148)
					p.Match(GoreParserT__11)
				}
				p.SetState(150)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == GoreParserT__12 {
					{
						p.SetState(149)

						var _m = p.Match(GoreParserT__12)

						localctx.(*IsNullContext).n = _m
					}

				}
				{
					p.SetState(152)
					p.Match(GoreParserNIL_LIT)
				}

			case 8:
				localctx = NewLikeContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoreParserRULE_expression)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				p.SetState(155)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == GoreParserT__12 {
					{
						p.SetState(154)

						var _m = p.Match(GoreParserT__12)

						localctx.(*LikeContext).n = _m
					}

				}
				{
					p.SetState(157)
					p.Match(GoreParserT__13)
				}
				{
					p.SetState(158)
					p.Match(GoreParserRAW_STRING_LIT)
				}

			case 9:
				localctx = NewInContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoreParserRULE_expression)
				p.SetState(159)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				p.SetState(161)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == GoreParserT__12 {
					{
						p.SetState(160)

						var _m = p.Match(GoreParserT__12)

						localctx.(*InContext).n = _m
					}

				}
				{
					p.SetState(163)
					p.Match(GoreParserT__14)
				}
				{
					p.SetState(164)
					p.Match(GoreParserT__15)
				}
				{
					p.SetState(165)
					p.expression(0)
				}
				p.SetState(168)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for ok := true; ok; ok = _la == GoreParserT__16 {
					{
						p.SetState(166)
						p.Match(GoreParserT__16)
					}
					{
						p.SetState(167)
						p.expression(0)
					}

					p.SetState(170)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(172)
					p.Match(GoreParserT__17)
				}

			}

		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryExpContext is an interface to support dynamic dispatch.
type IUnaryExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExpContext differentiates from other interfaces.
	IsUnaryExpContext()
}

type UnaryExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpContext() *UnaryExpContext {
	var p = new(UnaryExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_unaryExp
	return p
}

func (*UnaryExpContext) IsUnaryExpContext() {}

func NewUnaryExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpContext {
	var p = new(UnaryExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_unaryExp

	return p
}

func (s *UnaryExpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpContext) CopyFrom(ctx *UnaryExpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *UnaryExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnaryOpContext struct {
	*UnaryExpContext
}

func NewUnaryOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryOpContext {
	var p = new(UnaryOpContext)

	p.UnaryExpContext = NewEmptyUnaryExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnaryExpContext))

	return p
}

func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) UnaryOper() IUnaryOperContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOperContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOperContext)
}

func (s *UnaryOpContext) UnaryExp() IUnaryExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpContext)
}

func (s *UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterUnaryOp(s)
	}
}

func (s *UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitUnaryOp(s)
	}
}

func (s *UnaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitUnaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryContext struct {
	*UnaryExpContext
}

func NewPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryContext {
	var p = new(PrimaryContext)

	p.UnaryExpContext = NewEmptyUnaryExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnaryExpContext))

	return p
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) PrimaryExp() IPrimaryExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpContext)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) UnaryExp() (localctx IUnaryExpContext) {
	localctx = NewUnaryExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoreParserRULE_unaryExp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(183)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoreParserT__23, GoreParserT__24, GoreParserT__25:
		localctx = NewUnaryOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.UnaryOper()
		}
		{
			p.SetState(180)
			p.UnaryExp()
		}

	case GoreParserT__15, GoreParserT__22, GoreParserT__26, GoreParserT__27, GoreParserDECIMAL_LIT, GoreParserRAW_STRING_LIT, GoreParserNIL_LIT, GoreParserFLOAT_LIT, GoreParserVARIABLE, GoreParserFACT, GoreParserCONVERT_TYPE, GoreParserOCTAL_LIT, GoreParserHEX_LIT:
		localctx = NewPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(182)
			p.PrimaryExp()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimaryExpContext is an interface to support dynamic dispatch.
type IPrimaryExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpContext differentiates from other interfaces.
	IsPrimaryExpContext()
}

type PrimaryExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpContext() *PrimaryExpContext {
	var p = new(PrimaryExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_primaryExp
	return p
}

func (*PrimaryExpContext) IsPrimaryExpContext() {}

func NewPrimaryExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpContext {
	var p = new(PrimaryExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_primaryExp

	return p
}

func (s *PrimaryExpContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpContext) CopyFrom(ctx *PrimaryExpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PrimaryExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RoundContext struct {
	*PrimaryExpContext
	prec antlr.Token
}

func NewRoundContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoundContext {
	var p = new(RoundContext)

	p.PrimaryExpContext = NewEmptyPrimaryExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpContext))

	return p
}

func (s *RoundContext) GetPrec() antlr.Token { return s.prec }

func (s *RoundContext) SetPrec(v antlr.Token) { s.prec = v }

func (s *RoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RoundContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(GoreParserDECIMAL_LIT, 0)
}

func (s *RoundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterRound(s)
	}
}

func (s *RoundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitRound(s)
	}
}

func (s *RoundContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitRound(s)

	default:
		return t.VisitChildren(s)
	}
}

type OperContext struct {
	*PrimaryExpContext
}

func NewOperContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperContext {
	var p = new(OperContext)

	p.PrimaryExpContext = NewEmptyPrimaryExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpContext))

	return p
}

func (s *OperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperContext) Operand() IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *OperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterOper(s)
	}
}

func (s *OperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitOper(s)
	}
}

func (s *OperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitOper(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConvertContext struct {
	*PrimaryExpContext
}

func NewConvertContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConvertContext {
	var p = new(ConvertContext)

	p.PrimaryExpContext = NewEmptyPrimaryExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpContext))

	return p
}

func (s *ConvertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConvertContext) CONVERT_TYPE() antlr.TerminalNode {
	return s.GetToken(GoreParserCONVERT_TYPE, 0)
}

func (s *ConvertContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConvertContext) BasicLit() IBasicLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicLitContext)
}

func (s *ConvertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterConvert(s)
	}
}

func (s *ConvertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitConvert(s)
	}
}

func (s *ConvertContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitConvert(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) PrimaryExp() (localctx IPrimaryExpContext) {
	localctx = NewPrimaryExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoreParserRULE_primaryExp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(204)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoreParserT__15, GoreParserT__26, GoreParserT__27, GoreParserDECIMAL_LIT, GoreParserRAW_STRING_LIT, GoreParserNIL_LIT, GoreParserFLOAT_LIT, GoreParserVARIABLE, GoreParserFACT, GoreParserOCTAL_LIT, GoreParserHEX_LIT:
		localctx = NewOperContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Operand()
		}

	case GoreParserCONVERT_TYPE:
		localctx = NewConvertContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.Match(GoreParserCONVERT_TYPE)
		}
		{
			p.SetState(187)
			p.Match(GoreParserT__15)
		}
		{
			p.SetState(188)
			p.expression(0)
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GoreParserT__16 {
			{
				p.SetState(189)
				p.Match(GoreParserT__16)
			}
			{
				p.SetState(190)
				p.BasicLit()
			}

		}
		{
			p.SetState(193)
			p.Match(GoreParserT__17)
		}

	case GoreParserT__22:
		localctx = NewRoundContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(195)
			p.Match(GoreParserT__22)
		}
		{
			p.SetState(196)
			p.Match(GoreParserT__15)
		}
		{
			p.SetState(197)
			p.expression(0)
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GoreParserT__16 {
			{
				p.SetState(198)
				p.Match(GoreParserT__16)
			}
			{
				p.SetState(199)

				var _m = p.Match(GoreParserDECIMAL_LIT)

				localctx.(*RoundContext).prec = _m
			}

		}
		{
			p.SetState(202)
			p.Match(GoreParserT__17)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) CopyFrom(ctx *OperandContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParContext struct {
	*OperandContext
}

func NewParContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParContext {
	var p = new(ParContext)

	p.OperandContext = NewEmptyOperandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperandContext))

	return p
}

func (s *ParContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterPar(s)
	}
}

func (s *ParContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitPar(s)
	}
}

func (s *ParContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitPar(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentContext struct {
	*OperandContext
}

func NewIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentContext {
	var p = new(IdentContext)

	p.OperandContext = NewEmptyOperandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperandContext))

	return p
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type BasicLiterContext struct {
	*OperandContext
}

func NewBasicLiterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BasicLiterContext {
	var p = new(BasicLiterContext)

	p.OperandContext = NewEmptyOperandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperandContext))

	return p
}

func (s *BasicLiterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLiterContext) BasicLit() IBasicLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicLitContext)
}

func (s *BasicLiterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterBasicLiter(s)
	}
}

func (s *BasicLiterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitBasicLiter(s)
	}
}

func (s *BasicLiterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitBasicLiter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoreParserRULE_operand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(212)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoreParserT__26, GoreParserT__27, GoreParserDECIMAL_LIT, GoreParserRAW_STRING_LIT, GoreParserNIL_LIT, GoreParserFLOAT_LIT, GoreParserOCTAL_LIT, GoreParserHEX_LIT:
		localctx = NewBasicLiterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.BasicLit()
		}

	case GoreParserVARIABLE, GoreParserFACT:
		localctx = NewIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
			p.Identifier()
		}

	case GoreParserT__15:
		localctx = NewParContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(208)
			p.Match(GoreParserT__15)
		}
		{
			p.SetState(209)
			p.expression(0)
		}
		{
			p.SetState(210)
			p.Match(GoreParserT__17)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnaryOperContext is an interface to support dynamic dispatch.
type IUnaryOperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperContext differentiates from other interfaces.
	IsUnaryOperContext()
}

type UnaryOperContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperContext() *UnaryOperContext {
	var p = new(UnaryOperContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_unaryOper
	return p
}

func (*UnaryOperContext) IsUnaryOperContext() {}

func NewUnaryOperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperContext {
	var p = new(UnaryOperContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_unaryOper

	return p
}

func (s *UnaryOperContext) GetParser() antlr.Parser { return s.parser }
func (s *UnaryOperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterUnaryOper(s)
	}
}

func (s *UnaryOperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitUnaryOper(s)
	}
}

func (s *UnaryOperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitUnaryOper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) UnaryOper() (localctx IUnaryOperContext) {
	localctx = NewUnaryOperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoreParserRULE_unaryOper)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoreParserT__23)|(1<<GoreParserT__24)|(1<<GoreParserT__25))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) CopyFrom(ctx *IdentifierContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FactContext struct {
	*IdentifierContext
}

func NewFactContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FactContext {
	var p = new(FactContext)

	p.IdentifierContext = NewEmptyIdentifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentifierContext))

	return p
}

func (s *FactContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactContext) FACT() antlr.TerminalNode {
	return s.GetToken(GoreParserFACT, 0)
}

func (s *FactContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterFact(s)
	}
}

func (s *FactContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitFact(s)
	}
}

func (s *FactContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitFact(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarContext struct {
	*IdentifierContext
}

func NewVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarContext {
	var p = new(VarContext)

	p.IdentifierContext = NewEmptyIdentifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentifierContext))

	return p
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(GoreParserVARIABLE, 0)
}

func (s *VarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterVar(s)
	}
}

func (s *VarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitVar(s)
	}
}

func (s *VarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoreParserRULE_identifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(218)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoreParserVARIABLE:
		localctx = NewVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.Match(GoreParserVARIABLE)
		}

	case GoreParserFACT:
		localctx = NewFactContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(GoreParserFACT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBasicLitContext is an interface to support dynamic dispatch.
type IBasicLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicLitContext differentiates from other interfaces.
	IsBasicLitContext()
}

type BasicLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicLitContext() *BasicLitContext {
	var p = new(BasicLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_basicLit
	return p
}

func (*BasicLitContext) IsBasicLitContext() {}

func NewBasicLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicLitContext {
	var p = new(BasicLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_basicLit

	return p
}

func (s *BasicLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicLitContext) CopyFrom(ctx *BasicLitContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BasicLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringContext struct {
	*BasicLitContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.BasicLitContext = NewEmptyBasicLitContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BasicLitContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoreParserRAW_STRING_LIT, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolContext struct {
	*BasicLitContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	p.BasicLitContext = NewEmptyBasicLitContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BasicLitContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitBool(s)
	}
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullContext struct {
	*BasicLitContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	p.BasicLitContext = NewEmptyBasicLitContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BasicLitContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NIL_LIT() antlr.TerminalNode {
	return s.GetToken(GoreParserNIL_LIT, 0)
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitNull(s)
	}
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatContext struct {
	*BasicLitContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.BasicLitContext = NewEmptyBasicLitContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BasicLitContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(GoreParserFLOAT_LIT, 0)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	*BasicLitContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.BasicLitContext = NewEmptyBasicLitContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BasicLitContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) BasicLit() (localctx IBasicLitContext) {
	localctx = NewBasicLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoreParserRULE_basicLit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(225)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoreParserDECIMAL_LIT, GoreParserOCTAL_LIT, GoreParserHEX_LIT:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Integer()
		}

	case GoreParserRAW_STRING_LIT:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(221)
			p.Match(GoreParserRAW_STRING_LIT)
		}

	case GoreParserFLOAT_LIT:
		localctx = NewFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(222)
			p.Match(GoreParserFLOAT_LIT)
		}

	case GoreParserT__26, GoreParserT__27:
		localctx = NewBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(223)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GoreParserT__26 || _la == GoreParserT__27) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case GoreParserNIL_LIT:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(224)
			p.Match(GoreParserNIL_LIT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(GoreParserDECIMAL_LIT, 0)
}

func (s *IntegerContext) OCTAL_LIT() antlr.TerminalNode {
	return s.GetToken(GoreParserOCTAL_LIT, 0)
}

func (s *IntegerContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(GoreParserHEX_LIT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoreParserRULE_integer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(GoreParserDECIMAL_LIT-29))|(1<<(GoreParserOCTAL_LIT-29))|(1<<(GoreParserHEX_LIT-29)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoreParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoreParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(GoreParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoreListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoreVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoreParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoreParserRULE_eos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Match(GoreParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(230)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "p.lineTerminatorAhead()", ""))
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(231)

		if !(p.checkPreviousTokenText("}")) {
			panic(antlr.NewFailedPredicateException(p, "p.checkPreviousTokenText(\"}\")", ""))
		}

	}

	return localctx
}

func (p *GoreParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 17:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GoreParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GoreParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.lineTerminatorAhead()

	case 10:
		return p.checkPreviousTokenText("}")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
