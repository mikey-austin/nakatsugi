package config

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/mikey-austin/nakatsugi/parser"
)

type loader struct {
	*parser.BaseCONFIGListener
	Diags       []diag
	Assignments map[string]string
	Config      *Config
}

type diag struct {
	Line     int
	Column   int
	Message  string
	Severity severity
}

type severity int

const (
	Error severity = iota
	Warning
)

func (s severity) String() string {
	switch s {
	case Error:
		return "Error"
	case Warning:
		return "Warning"
	default:
		return "Unknown"
	}
}

func (d diag) String() string {
	return fmt.Sprintf("[%v] %d:%d: %s", d.Severity, d.Line, d.Column, d.Message)
}

func newLoader() *loader {
	return &loader{
		Assignments: make(map[string]string),
		Config: &Config{
			Tables: make(map[string]*Table),
		},
	}
}

func (this *loader) addDiag(ctx antlr.ParserRuleContext, msg string, severity severity) {
	token := ctx.GetStart()
	this.Diags = append(this.Diags, diag{
		Line:     token.GetLine(),
		Column:   token.GetColumn(),
		Message:  msg,
		Severity: severity,
	})
}

func (this *loader) from(r io.Reader) error {
	rawContents, error := io.ReadAll(r)
	if error != nil {
		return error
	}

	input := antlr.NewInputStream(string(rawContents))
	lexer := parser.NewCONFIGLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewCONFIGParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(this, parser.Config())

	if len(this.Diags) > 0 {
		errors := ""
		for _, d := range this.Diags {
			errors += d.String() + "\n"
		}
		return fmt.Errorf("\n%v", errors)
	}

	return nil
}

func (this *loader) resolveVar(name string) (string, error) {
	name = strings.TrimLeft(name, "$")
	resolved, ok := this.Assignments[name]
	if !ok {
		return "", fmt.Errorf("could not resolve '%v'", name)
	}

	return resolved, nil
}

func (this *loader) ExitInterval(ctx *parser.IntervalContext) {
	value, error := strconv.Atoi(ctx.NUMBER().GetText())
	if error != nil {
		this.addDiag(ctx, error.Error(), Error)
	} else {
		this.Config.Interval = value
	}
}

func (this *loader) ExitAssignment(ctx *parser.AssignmentContext) {
	this.Assignments[ctx.NAME(0).GetText()] = ctx.NAME(1).GetText()
}

func (this *loader) ExitTable(ctx *parser.TableContext) {
	name := ctx.Tname().NAME().GetText()
	name = strings.TrimLeft(name, "<")
	name = strings.TrimRight(name, ">")

	t := &Table{
		Name:     name,
		Disabled: ctx.Tdisable() != nil,
	}

	for _, titem := range ctx.Tlist().AllTitem() {
		var itemVal string
		if titem.NAME() != nil {
			itemVal = titem.NAME().GetText()
		} else if titem.Var_() != nil {
			resolved, error := this.resolveVar(titem.Var_().GetText())
			if error != nil {
				this.addDiag(titem, error.Error(), Error)
			}
			itemVal = resolved
		}
		t.Items = append(t.Items, itemVal)
	}

	this.Config.Tables[name] = t
}

func NewConfigFromReader(r io.Reader) (*Config, error) {
	loader := newLoader()
	error := loader.from(r)
	if error != nil {
		return loader.Config, error
	}

	return loader.Config, nil
}
