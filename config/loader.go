package config

import (
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/mikey-austin/nakatsugi/parser"
)

type loader struct {
	*parser.BaseCONFIGListener
	Errors      int
	Assignments map[string]string
	Config      *Config
}

func newLoader() *loader {
	return &loader{
		Errors:      0,
		Assignments: make(map[string]string),
		Config: &Config{
			Tables: make(map[string]*Table),
		},
	}
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

	return nil
}

func (this *loader) EnterInterval(ctx *parser.IntervalContext) {
	value, error := strconv.Atoi(ctx.GetChild(1).(antlr.ParseTree).GetText())
	if error != nil {
		// TODO: find out how to add an error
	} else {
		this.Config.Interval = value
	}
}

func (this *loader) EnterAssignment(ctx *parser.AssignmentContext) {
	name := ctx.GetChild(0).(antlr.ParseTree).GetText()
	value := ctx.GetChild(2).(antlr.ParseTree).GetText()
	this.Assignments[name] = value
}

func (this *loader) EnterTable(ctx *parser.TableContext) {
	// Main table metadata
	name := ctx.GetChild(1).(antlr.ParseTree).GetText()
	name = strings.TrimLeft(name, "<")
	name = strings.TrimRight(name, ">")

	disableTag := ctx.GetChild(2).(antlr.ParseTree).GetText()
	disable := false
	itemsIndex := 3
	if disableTag == "disable" {
		disable = true
		itemsIndex = 4
	}

	t := &Table{
		Name:     name,
		Disabled: disable,
	}

	// Extract table items
	for _, titem := range ctx.GetChild(itemsIndex).GetChildren() {
		ti := titem.(antlr.ParseTree).GetText()
		if ti == "," {
			continue
		}

		t.Items = append(t.Items, ti)
	}

	// Save table in the configuration
	this.Config.Tables[name] = t
}

func (this *loader) config() (*Config, error) {
	if this.Errors > 0 {
		return nil, errors.New("errors parsing config")
	}
	return this.Config, nil
}

func NewConfigFromReader(r io.Reader) (*Config, error) {
	loader := newLoader()
	error := loader.from(r)
	if error != nil {
		return nil, error
	}

	return loader.config()
}
