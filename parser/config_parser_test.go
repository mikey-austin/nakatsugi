package parser

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	*BaseCONFIGListener
	Errors    int
	Variables map[string]string
}

func NewVisitor() *Visitor {
	return &Visitor{
		Errors:    0,
		Variables: make(map[string]string),
	}
}

func (this *Visitor) VisitErrorNode(node antlr.ErrorNode) {
	this.Errors += 1
}

func (this *Visitor) EnterAssignment(ctx *AssignmentContext) {
	v := ctx.GetChild(0).(antlr.ParseTree).GetText()
	h := ctx.GetChild(2).(antlr.ParseTree).GetText()
	this.Variables[v] = h
}

func TestSimpleAssignments(t *testing.T) {
	c := `a=example.com
test123=test.com
another_test=x.org
`
	input := antlr.NewInputStream(c)
	lexer := NewCONFIGLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewCONFIGParser(stream)
	tree := parser.Config()
	visitor := NewVisitor()
	antlr.ParseTreeWalkerDefault.Walk(visitor, tree)

	if visitor.Errors != 0 {
		t.Errorf("%d unexpected parse errors", visitor.Errors)
	}

	if len(visitor.Variables) != 3 {
		t.Error("Did not parse all statements")
	}

	vars := map[string]string{
		"a":            "example.com",
		"test123":      "test.com",
		"another_test": "x.org",
	}
	for v, h := range vars {
		if visitor.Variables[v] != h {
			t.Errorf("Expected %s=%s but got %s=%s", v, h, v, visitor.Variables[v])
		}
	}
}
