package config

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

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
			Interval:  time.Duration(10) * time.Second,
			Timeout:   time.Duration(200) * time.Millisecond,
			Tables:    make(map[string]*Table),
			Redirects: make(map[string]*Redirect),
		},
	}
}

func (this *loader) addErr(ctx antlr.ParserRuleContext, msgFmt string, args ...string) {
	this.addDiag(ctx, fmt.Sprintf(msgFmt, args), Error)
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

	// Uncomment for debugging the token stream
	/*
		for {
			token := lexer.NextToken()
			tokenName := lexer.SymbolicNames[token.GetTokenType()]
			fmt.Printf("%v %q\n", tokenName, token.GetText())

			if token.GetTokenType() == antlr.TokenEOF {
				break
			}
		}
	*/
	// End debugging

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
	if _, ok := ctx.GetParent().(*parser.GoptsContext); !ok {
		// only top-level intervals here
		return
	}

	value, error := strconv.Atoi(ctx.NUMBER().GetText())
	if error != nil {
		this.addErr(ctx, error.Error())
	} else {
		this.Config.Interval = time.Duration(value) * time.Second
	}
}

func (this *loader) ExitTimeout(ctx *parser.TimeoutContext) {
	if _, ok := ctx.GetParent().(*parser.GoptsContext); !ok {
		// only top-level timeouts here
		return
	}

	value, error := strconv.Atoi(ctx.NUMBER().GetText())
	if error != nil {
		this.addErr(ctx, error.Error())
	} else {
		this.Config.Timeout = time.Duration(value) * time.Millisecond
	}
}

func (this *loader) ExitAssignment(ctx *parser.AssignmentContext) {
	varName := ctx.NAME(0).GetText()
	if !IsValidVarName(varName) {
		this.addErr(ctx, "invalid var name: %v", varName)
	} else {
		this.Assignments[varName] = ctx.NAME(1).GetText()
	}
}

func (this *loader) ExitTable(ctx *parser.TableContext) {
	name := StripTableName(ctx.Tname().NAME().GetText())
	if !IsValidName(name) {
		this.addErr(ctx, "invalid table name", name)
		return
	}

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
				this.addErr(titem, error.Error())
			}
			itemVal = resolved
		}
		item, error := NewTItemFromParseCtx(itemVal, titem)
		if error != nil {
			this.addErr(titem, "error parsing table item: %v", error.Error())
		} else {
			t.Items = append(t.Items, item)
		}
	}

	this.Config.Tables[name] = t
}

func (this *loader) ExitRedirect(ctx *parser.RedirectContext) {
	var redirect Redirect

	name := ctx.Rname().QNAME().GetText()
	name = strings.Trim(name, "\"")
	if !IsValidName(name) {
		this.addErr(ctx, "invalid redirect name", name)
		return
	}
	redirect.Name = name
	redirect.Disabled = false
	redirect.SessionTimeout = time.Duration(600) * time.Second

	var defaultPort uint16 = 0
	for _, rstmt := range ctx.AllRstmt() {
		lastListenPort, err := this.setRstmt(rstmt, &redirect, defaultPort)
		if err != nil {
			this.addErr(rstmt, "error parsing redirect statement: %v", err.Error())
		}
		if lastListenPort != 0 {
			defaultPort = lastListenPort
		}
	}

	this.Config.Redirects[name] = &redirect
}

func (this *loader) setRstmt(ctx parser.IRstmtContext, redirect *Redirect, defaultPort uint16) (uint16, error) {
	var seenPort uint16 = 0

	if ctx.DISABLE() != nil {
		redirect.Disabled = true
	}

	if ctx.STICKY_ADDRESS() != nil {
		redirect.StickyAddress = true
	}

	if ctx.Sesstimeout() != nil {
		parsedTimeout, err := strconv.Atoi(ctx.Sesstimeout().NUMBER().GetText())
		if err != nil {
			return seenPort, err
		}
		redirect.SessionTimeout = time.Duration(parsedTimeout) * time.Second
	}

	if ctx.Listen() != nil {
		listen, err := NewListen(ctx.Listen())
		if err != nil {
			return seenPort, err
		}
		redirect.Listens = append(redirect.Listens, listen)
		seenPort = listen.MinPort
	}

	if ctx.Forward() != nil {
		forward, err := NewForward(ctx.Forward(), this.resolveTable, defaultPort)
		if err != nil {
			return seenPort, err
		}
		redirect.Forwards = append(redirect.Forwards, forward)
	}

	if ctx.Route() != nil {
		route, err := NewRoute(ctx.Route(), this.resolveTable, defaultPort)
		if err != nil {
			return seenPort, err
		}
		redirect.Routes = append(redirect.Routes, route)
	}

	return seenPort, nil
}

func (this *loader) resolveTable(ctx parser.ITnameContext) (*Table, error) {
	name := StripTableName(ctx.GetText())
	table, ok := this.Config.Tables[name]
	if !ok {
		return nil, fmt.Errorf("unknown table referenced: %v", name)
	}
	return table, nil
}

func NewConfigFromReader(r io.Reader) (*Config, error) {
	loader := newLoader()
	error := loader.from(r)
	if error != nil {
		return loader.Config, error
	}

	return loader.Config, nil
}
