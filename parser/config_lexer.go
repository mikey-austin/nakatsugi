// Code generated from CONFIG.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CONFIGLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CONFIGLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func configlexerLexerInit() {
	staticData := &CONFIGLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "VAR", "HOST", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"T__0", "VAR", "HOST", "NEWLINE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 28, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 1, 1, 1, 5, 1, 14, 8, 1, 10, 1, 12, 1, 17, 9, 1, 1, 2, 4, 2,
		20, 8, 2, 11, 2, 12, 2, 21, 1, 3, 4, 3, 25, 8, 3, 11, 3, 12, 3, 26, 0,
		0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1, 0, 4, 1, 0, 97, 122, 3, 0, 48, 57, 95,
		95, 97, 122, 3, 0, 46, 46, 48, 57, 97, 122, 2, 0, 10, 10, 13, 13, 30, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1,
		9, 1, 0, 0, 0, 3, 11, 1, 0, 0, 0, 5, 19, 1, 0, 0, 0, 7, 24, 1, 0, 0, 0,
		9, 10, 5, 61, 0, 0, 10, 2, 1, 0, 0, 0, 11, 15, 7, 0, 0, 0, 12, 14, 7, 1,
		0, 0, 13, 12, 1, 0, 0, 0, 14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16,
		1, 0, 0, 0, 16, 4, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 18, 20, 7, 2, 0, 0,
		19, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1,
		0, 0, 0, 22, 6, 1, 0, 0, 0, 23, 25, 7, 3, 0, 0, 24, 23, 1, 0, 0, 0, 25,
		26, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 8, 1, 0, 0,
		0, 4, 0, 15, 21, 26, 0,
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

// CONFIGLexerInit initializes any static state used to implement CONFIGLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCONFIGLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CONFIGLexerInit() {
	staticData := &CONFIGLexerLexerStaticData
	staticData.once.Do(configlexerLexerInit)
}

// NewCONFIGLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCONFIGLexer(input antlr.CharStream) *CONFIGLexer {
	CONFIGLexerInit()
	l := new(CONFIGLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CONFIGLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "CONFIG.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CONFIGLexer tokens.
const (
	CONFIGLexerT__0    = 1
	CONFIGLexerVAR     = 2
	CONFIGLexerHOST    = 3
	CONFIGLexerNEWLINE = 4
)
