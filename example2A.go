// example2A.go
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/go-sqlparser/goadvent-antlr/parser"
)

type calcListener struct {
	*parser.BaseCalcListener
	parser *parser.CalcParser
}

// ExitEveryRule is called when any rule is exited.
func (l *calcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	count := ctx.GetChildCount()
	// count == 1, NUMBER
	ch := ctx.GetChild(0)
	if count == 3 {
		// operation
		ch = ctx.GetChild(1)
	}
	q := ch.(antlr.Tree)
	t, ok := q.(antlr.TerminalNode)
	if ok {
		s := t.GetSymbol()
		x := s.GetTokenType()
		y := l.parser.GetSymbolicNames()
		fmt.Println(y[x], ctx.GetText())
	}
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("12 * (34 - 56)")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	listener := calcListener{}
	listener.parser = p
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
}

/*

$ go run example2A.go
NUMBER 12
NUMBER 34
NUMBER 56
SUB 34-56
MUL 12*(34-56)

*/
