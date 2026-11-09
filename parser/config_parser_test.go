package parser

import (
	"testing"

	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*BaseCONFIGListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func TestSimpleParse(t *testing.T) {
	c := `a=example.com
test123=test.com
another_test=x.org
`
	input := antlr.NewInputStream(c)
	lexer := NewCONFIGLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewCONFIGParser(stream)
	tree := parser.Config()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
