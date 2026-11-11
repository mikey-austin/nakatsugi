// Code generated from CONFIG.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CONFIG
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CONFIGParser struct {
	*antlr.BaseParser
}

var CONFIGParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func configParserInit() {
	staticData := &CONFIGParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'$'", "'{'", "'}'", "'<'", "'>'", "','", "'disable'", "'interval'",
		"'table'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "DISABLE", "INTERVAL", "TABLE", "NUMBER",
		"NAME", "HOST", "COMMENT", "WS", "CONTINUATION", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"config", "stmt", "assignment", "interval", "var", "table", "tname",
		"tlist", "titem", "tdisable",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 82, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 0, 5, 0, 25, 8, 0, 10, 0, 12, 0, 28, 9, 0, 1, 0, 3, 0, 31,
		8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 38, 8, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 53, 8, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7,
		66, 8, 7, 10, 7, 12, 7, 69, 9, 7, 1, 7, 3, 7, 72, 8, 7, 3, 7, 74, 8, 7,
		1, 8, 1, 8, 3, 8, 78, 8, 8, 1, 9, 1, 9, 1, 9, 0, 0, 10, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 0, 0, 81, 0, 26, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4,
		39, 1, 0, 0, 0, 6, 43, 1, 0, 0, 0, 8, 46, 1, 0, 0, 0, 10, 49, 1, 0, 0,
		0, 12, 58, 1, 0, 0, 0, 14, 73, 1, 0, 0, 0, 16, 77, 1, 0, 0, 0, 18, 79,
		1, 0, 0, 0, 20, 25, 5, 17, 0, 0, 21, 22, 3, 2, 1, 0, 22, 23, 5, 17, 0,
		0, 23, 25, 1, 0, 0, 0, 24, 20, 1, 0, 0, 0, 24, 21, 1, 0, 0, 0, 25, 28,
		1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 30, 1, 0, 0, 0,
		28, 26, 1, 0, 0, 0, 29, 31, 3, 2, 1, 0, 30, 29, 1, 0, 0, 0, 30, 31, 1,
		0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 33, 5, 0, 0, 1, 33, 1, 1, 0, 0, 0, 34,
		38, 3, 4, 2, 0, 35, 38, 3, 10, 5, 0, 36, 38, 3, 6, 3, 0, 37, 34, 1, 0,
		0, 0, 37, 35, 1, 0, 0, 0, 37, 36, 1, 0, 0, 0, 38, 3, 1, 0, 0, 0, 39, 40,
		5, 12, 0, 0, 40, 41, 5, 1, 0, 0, 41, 42, 5, 13, 0, 0, 42, 5, 1, 0, 0, 0,
		43, 44, 5, 9, 0, 0, 44, 45, 5, 11, 0, 0, 45, 7, 1, 0, 0, 0, 46, 47, 5,
		2, 0, 0, 47, 48, 5, 12, 0, 0, 48, 9, 1, 0, 0, 0, 49, 50, 5, 10, 0, 0, 50,
		52, 3, 12, 6, 0, 51, 53, 3, 18, 9, 0, 52, 51, 1, 0, 0, 0, 52, 53, 1, 0,
		0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 5, 3, 0, 0, 55, 56, 3, 14, 7, 0, 56,
		57, 5, 4, 0, 0, 57, 11, 1, 0, 0, 0, 58, 59, 5, 5, 0, 0, 59, 60, 5, 12,
		0, 0, 60, 61, 5, 6, 0, 0, 61, 13, 1, 0, 0, 0, 62, 67, 3, 16, 8, 0, 63,
		64, 5, 7, 0, 0, 64, 66, 3, 16, 8, 0, 65, 63, 1, 0, 0, 0, 66, 69, 1, 0,
		0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67,
		1, 0, 0, 0, 70, 72, 5, 7, 0, 0, 71, 70, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0,
		72, 74, 1, 0, 0, 0, 73, 62, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 15, 1,
		0, 0, 0, 75, 78, 3, 8, 4, 0, 76, 78, 5, 13, 0, 0, 77, 75, 1, 0, 0, 0, 77,
		76, 1, 0, 0, 0, 78, 17, 1, 0, 0, 0, 79, 80, 5, 8, 0, 0, 80, 19, 1, 0, 0,
		0, 9, 24, 26, 30, 37, 52, 67, 71, 73, 77,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CONFIGParserInit initializes any static state used to implement CONFIGParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCONFIGParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CONFIGParserInit() {
	staticData := &CONFIGParserStaticData
	staticData.once.Do(configParserInit)
}

// NewCONFIGParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCONFIGParser(input antlr.TokenStream) *CONFIGParser {
	CONFIGParserInit()
	this := new(CONFIGParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CONFIGParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CONFIG.g4"

	return this
}

// CONFIGParser tokens.
const (
	CONFIGParserEOF          = antlr.TokenEOF
	CONFIGParserT__0         = 1
	CONFIGParserT__1         = 2
	CONFIGParserT__2         = 3
	CONFIGParserT__3         = 4
	CONFIGParserT__4         = 5
	CONFIGParserT__5         = 6
	CONFIGParserT__6         = 7
	CONFIGParserDISABLE      = 8
	CONFIGParserINTERVAL     = 9
	CONFIGParserTABLE        = 10
	CONFIGParserNUMBER       = 11
	CONFIGParserNAME         = 12
	CONFIGParserHOST         = 13
	CONFIGParserCOMMENT      = 14
	CONFIGParserWS           = 15
	CONFIGParserCONTINUATION = 16
	CONFIGParserNEWLINE      = 17
)

// CONFIGParser rules.
const (
	CONFIGParserRULE_config     = 0
	CONFIGParserRULE_stmt       = 1
	CONFIGParserRULE_assignment = 2
	CONFIGParserRULE_interval   = 3
	CONFIGParserRULE_var        = 4
	CONFIGParserRULE_table      = 5
	CONFIGParserRULE_tname      = 6
	CONFIGParserRULE_tlist      = 7
	CONFIGParserRULE_titem      = 8
	CONFIGParserRULE_tdisable   = 9
)

// IConfigContext is an interface to support dynamic dispatch.
type IConfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsConfigContext differentiates from other interfaces.
	IsConfigContext()
}

type ConfigContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigContext() *ConfigContext {
	var p = new(ConfigContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_config
	return p
}

func InitEmptyConfigContext(p *ConfigContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_config
}

func (*ConfigContext) IsConfigContext() {}

func NewConfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigContext {
	var p = new(ConfigContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_config

	return p
}

func (s *ConfigContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigContext) EOF() antlr.TerminalNode {
	return s.GetToken(CONFIGParserEOF, 0)
}

func (s *ConfigContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(CONFIGParserNEWLINE)
}

func (s *ConfigContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(CONFIGParserNEWLINE, i)
}

func (s *ConfigContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ConfigContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ConfigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterConfig(s)
	}
}

func (s *ConfigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitConfig(s)
	}
}

func (p *CONFIGParser) Config() (localctx IConfigContext) {
	localctx = NewConfigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CONFIGParserRULE_config)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(24)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case CONFIGParserNEWLINE:
				{
					p.SetState(20)
					p.Match(CONFIGParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case CONFIGParserINTERVAL, CONFIGParserTABLE, CONFIGParserNAME:
				{
					p.SetState(21)
					p.Stmt()
				}
				{
					p.SetState(22)
					p.Match(CONFIGParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5632) != 0 {
		{
			p.SetState(29)
			p.Stmt()
		}

	}
	{
		p.SetState(32)
		p.Match(CONFIGParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	Table() ITableContext
	Interval() IIntervalContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StmtContext) Table() ITableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableContext)
}

func (s *StmtContext) Interval() IIntervalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntervalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntervalContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *CONFIGParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CONFIGParserRULE_stmt)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Assignment()
		}

	case CONFIGParserTABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Table()
		}

	case CONFIGParserINTERVAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Interval()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	HOST() antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) NAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, 0)
}

func (s *AssignmentContext) HOST() antlr.TerminalNode {
	return s.GetToken(CONFIGParserHOST, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *CONFIGParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CONFIGParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(CONFIGParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(CONFIGParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.Match(CONFIGParserHOST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntervalContext is an interface to support dynamic dispatch.
type IIntervalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTERVAL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsIntervalContext differentiates from other interfaces.
	IsIntervalContext()
}

type IntervalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalContext() *IntervalContext {
	var p = new(IntervalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_interval
	return p
}

func InitEmptyIntervalContext(p *IntervalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_interval
}

func (*IntervalContext) IsIntervalContext() {}

func NewIntervalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalContext {
	var p = new(IntervalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_interval

	return p
}

func (s *IntervalContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalContext) INTERVAL() antlr.TerminalNode {
	return s.GetToken(CONFIGParserINTERVAL, 0)
}

func (s *IntervalContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNUMBER, 0)
}

func (s *IntervalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterInterval(s)
	}
}

func (s *IntervalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitInterval(s)
	}
}

func (p *CONFIGParser) Interval() (localctx IIntervalContext) {
	localctx = NewIntervalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CONFIGParserRULE_interval)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(CONFIGParserINTERVAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(CONFIGParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarContext is an interface to support dynamic dispatch.
type IVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsVarContext differentiates from other interfaces.
	IsVarContext()
}

type VarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarContext() *VarContext {
	var p = new(VarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_var
	return p
}

func InitEmptyVarContext(p *VarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_var
}

func (*VarContext) IsVarContext() {}

func NewVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarContext {
	var p = new(VarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_var

	return p
}

func (s *VarContext) GetParser() antlr.Parser { return s.parser }

func (s *VarContext) NAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, 0)
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterVar(s)
	}
}

func (s *VarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitVar(s)
	}
}

func (p *CONFIGParser) Var_() (localctx IVarContext) {
	localctx = NewVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CONFIGParserRULE_var)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(CONFIGParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(CONFIGParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableContext is an interface to support dynamic dispatch.
type ITableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TABLE() antlr.TerminalNode
	Tname() ITnameContext
	Tlist() ITlistContext
	Tdisable() ITdisableContext

	// IsTableContext differentiates from other interfaces.
	IsTableContext()
}

type TableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableContext() *TableContext {
	var p = new(TableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_table
	return p
}

func InitEmptyTableContext(p *TableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_table
}

func (*TableContext) IsTableContext() {}

func NewTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableContext {
	var p = new(TableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_table

	return p
}

func (s *TableContext) GetParser() antlr.Parser { return s.parser }

func (s *TableContext) TABLE() antlr.TerminalNode {
	return s.GetToken(CONFIGParserTABLE, 0)
}

func (s *TableContext) Tname() ITnameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITnameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITnameContext)
}

func (s *TableContext) Tlist() ITlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITlistContext)
}

func (s *TableContext) Tdisable() ITdisableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITdisableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITdisableContext)
}

func (s *TableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterTable(s)
	}
}

func (s *TableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitTable(s)
	}
}

func (p *CONFIGParser) Table() (localctx ITableContext) {
	localctx = NewTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CONFIGParserRULE_table)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(CONFIGParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Tname()
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserDISABLE {
		{
			p.SetState(51)
			p.Tdisable()
		}

	}
	{
		p.SetState(54)
		p.Match(CONFIGParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Tlist()
	}
	{
		p.SetState(56)
		p.Match(CONFIGParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITnameContext is an interface to support dynamic dispatch.
type ITnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsTnameContext differentiates from other interfaces.
	IsTnameContext()
}

type TnameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTnameContext() *TnameContext {
	var p = new(TnameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_tname
	return p
}

func InitEmptyTnameContext(p *TnameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_tname
}

func (*TnameContext) IsTnameContext() {}

func NewTnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TnameContext {
	var p = new(TnameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_tname

	return p
}

func (s *TnameContext) GetParser() antlr.Parser { return s.parser }

func (s *TnameContext) NAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, 0)
}

func (s *TnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterTname(s)
	}
}

func (s *TnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitTname(s)
	}
}

func (p *CONFIGParser) Tname() (localctx ITnameContext) {
	localctx = NewTnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CONFIGParserRULE_tname)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(CONFIGParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(CONFIGParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(CONFIGParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITlistContext is an interface to support dynamic dispatch.
type ITlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTitem() []ITitemContext
	Titem(i int) ITitemContext

	// IsTlistContext differentiates from other interfaces.
	IsTlistContext()
}

type TlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTlistContext() *TlistContext {
	var p = new(TlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_tlist
	return p
}

func InitEmptyTlistContext(p *TlistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_tlist
}

func (*TlistContext) IsTlistContext() {}

func NewTlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TlistContext {
	var p = new(TlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_tlist

	return p
}

func (s *TlistContext) GetParser() antlr.Parser { return s.parser }

func (s *TlistContext) AllTitem() []ITitemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITitemContext); ok {
			len++
		}
	}

	tst := make([]ITitemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITitemContext); ok {
			tst[i] = t.(ITitemContext)
			i++
		}
	}

	return tst
}

func (s *TlistContext) Titem(i int) ITitemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITitemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITitemContext)
}

func (s *TlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterTlist(s)
	}
}

func (s *TlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitTlist(s)
	}
}

func (p *CONFIGParser) Tlist() (localctx ITlistContext) {
	localctx = NewTlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CONFIGParserRULE_tlist)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserT__1 || _la == CONFIGParserHOST {
		{
			p.SetState(62)
			p.Titem()
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(63)
					p.Match(CONFIGParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(64)
					p.Titem()
				}

			}
			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CONFIGParserT__6 {
			{
				p.SetState(70)
				p.Match(CONFIGParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITitemContext is an interface to support dynamic dispatch.
type ITitemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_() IVarContext
	HOST() antlr.TerminalNode

	// IsTitemContext differentiates from other interfaces.
	IsTitemContext()
}

type TitemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTitemContext() *TitemContext {
	var p = new(TitemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_titem
	return p
}

func InitEmptyTitemContext(p *TitemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_titem
}

func (*TitemContext) IsTitemContext() {}

func NewTitemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TitemContext {
	var p = new(TitemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_titem

	return p
}

func (s *TitemContext) GetParser() antlr.Parser { return s.parser }

func (s *TitemContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *TitemContext) HOST() antlr.TerminalNode {
	return s.GetToken(CONFIGParserHOST, 0)
}

func (s *TitemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TitemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TitemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterTitem(s)
	}
}

func (s *TitemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitTitem(s)
	}
}

func (p *CONFIGParser) Titem() (localctx ITitemContext) {
	localctx = NewTitemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CONFIGParserRULE_titem)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Var_()
		}

	case CONFIGParserHOST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Match(CONFIGParserHOST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITdisableContext is an interface to support dynamic dispatch.
type ITdisableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DISABLE() antlr.TerminalNode

	// IsTdisableContext differentiates from other interfaces.
	IsTdisableContext()
}

type TdisableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTdisableContext() *TdisableContext {
	var p = new(TdisableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_tdisable
	return p
}

func InitEmptyTdisableContext(p *TdisableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_tdisable
}

func (*TdisableContext) IsTdisableContext() {}

func NewTdisableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TdisableContext {
	var p = new(TdisableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_tdisable

	return p
}

func (s *TdisableContext) GetParser() antlr.Parser { return s.parser }

func (s *TdisableContext) DISABLE() antlr.TerminalNode {
	return s.GetToken(CONFIGParserDISABLE, 0)
}

func (s *TdisableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TdisableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TdisableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterTdisable(s)
	}
}

func (s *TdisableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitTdisable(s)
	}
}

func (p *CONFIGParser) Tdisable() (localctx ITdisableContext) {
	localctx = NewTdisableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CONFIGParserRULE_tdisable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(CONFIGParserDISABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
