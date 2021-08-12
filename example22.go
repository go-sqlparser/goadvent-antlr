// example2.go
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
func (s *calcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	rc := ctx.GetRuleContext()
	ri := rc.GetRuleIndex()
	fmt.Printf("%s (%s)\n", s.parser.RuleNames[ri], ctx.GetText())
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("12 * (34 - 56)")
	//is := antlr.NewInputStream("1 + 2 * 3")

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

$ go run example22.go
expression (12)
expression (34)
expression (56)
expression (34-56)
expression ((34-56))
expression (12*(34-56))
start (12*(34-56)<EOF>)

*/
