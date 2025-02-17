// Code generated from Calc.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 43, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 6, 8, 33, 10, 8, 13, 8, 14, 8, 34, 3,
	9, 6, 9, 38, 10, 9, 13, 9, 14, 9, 39, 3, 9, 3, 9, 2, 2, 10, 3, 3, 5, 4,
	7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 3, 2, 4, 3, 2, 50, 59, 5, 2, 11,
	12, 15, 15, 34, 34, 2, 44, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5, 21, 3, 2, 2, 2, 7,
	23, 3, 2, 2, 2, 9, 25, 3, 2, 2, 2, 11, 27, 3, 2, 2, 2, 13, 29, 3, 2, 2,
	2, 15, 32, 3, 2, 2, 2, 17, 37, 3, 2, 2, 2, 19, 20, 7, 42, 2, 2, 20, 4,
	3, 2, 2, 2, 21, 22, 7, 43, 2, 2, 22, 6, 3, 2, 2, 2, 23, 24, 7, 44, 2, 2,
	24, 8, 3, 2, 2, 2, 25, 26, 7, 49, 2, 2, 26, 10, 3, 2, 2, 2, 27, 28, 7,
	45, 2, 2, 28, 12, 3, 2, 2, 2, 29, 30, 7, 47, 2, 2, 30, 14, 3, 2, 2, 2,
	31, 33, 9, 2, 2, 2, 32, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 32, 3,
	2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 16, 3, 2, 2, 2, 36, 38, 9, 3, 2, 2, 37,
	36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2,
	2, 40, 41, 3, 2, 2, 2, 41, 42, 8, 9, 2, 2, 42, 18, 3, 2, 2, 2, 5, 2, 34,
	39, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'*'", "'/'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "OPP", "CLP", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE",
}

var lexerRuleNames = []string{
	"OPP", "CLP", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE",
}

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *CalcLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCalcLexer(input antlr.CharStream) *CalcLexer {
	l := new(CalcLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Calc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerOPP        = 1
	CalcLexerCLP        = 2
	CalcLexerMUL        = 3
	CalcLexerDIV        = 4
	CalcLexerADD        = 5
	CalcLexerSUB        = 6
	CalcLexerNUMBER     = 7
	CalcLexerWHITESPACE = 8
)
