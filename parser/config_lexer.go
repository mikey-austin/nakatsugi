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
		"", "'='", "'$'", "'{'", "'}'", "'<'", "'>'", "','", "'disable'", "'interval'",
		"'table'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "DISABLE", "INTERVAL", "TABLE", "NUMBER",
		"NAME", "HOST", "COMMENT", "WS", "CONTINUATION", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "DISABLE", "INTERVAL",
		"TABLE", "NUMBER", "NAME", "HOST", "COMMENT", "WS", "CONTINUATION",
		"NEWLINE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 118, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 4, 10, 74, 8, 10, 11, 10, 12, 10, 75,
		1, 11, 1, 11, 5, 11, 80, 8, 11, 10, 11, 12, 11, 83, 9, 11, 1, 12, 4, 12,
		86, 8, 12, 11, 12, 12, 12, 87, 1, 13, 1, 13, 5, 13, 92, 8, 13, 10, 13,
		12, 13, 95, 9, 13, 1, 13, 1, 13, 1, 14, 4, 14, 100, 8, 14, 11, 14, 12,
		14, 101, 1, 14, 1, 14, 1, 15, 1, 15, 3, 15, 108, 8, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 3, 16, 115, 8, 16, 1, 16, 1, 16, 0, 0, 17, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 1, 0, 6, 1, 0, 48, 57, 1, 0, 97, 122,
		3, 0, 48, 57, 95, 95, 97, 122, 4, 0, 45, 46, 48, 57, 95, 95, 97, 122, 2,
		0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 124, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		1, 35, 1, 0, 0, 0, 3, 37, 1, 0, 0, 0, 5, 39, 1, 0, 0, 0, 7, 41, 1, 0, 0,
		0, 9, 43, 1, 0, 0, 0, 11, 45, 1, 0, 0, 0, 13, 47, 1, 0, 0, 0, 15, 49, 1,
		0, 0, 0, 17, 57, 1, 0, 0, 0, 19, 66, 1, 0, 0, 0, 21, 73, 1, 0, 0, 0, 23,
		77, 1, 0, 0, 0, 25, 85, 1, 0, 0, 0, 27, 89, 1, 0, 0, 0, 29, 99, 1, 0, 0,
		0, 31, 105, 1, 0, 0, 0, 33, 114, 1, 0, 0, 0, 35, 36, 5, 61, 0, 0, 36, 2,
		1, 0, 0, 0, 37, 38, 5, 36, 0, 0, 38, 4, 1, 0, 0, 0, 39, 40, 5, 123, 0,
		0, 40, 6, 1, 0, 0, 0, 41, 42, 5, 125, 0, 0, 42, 8, 1, 0, 0, 0, 43, 44,
		5, 60, 0, 0, 44, 10, 1, 0, 0, 0, 45, 46, 5, 62, 0, 0, 46, 12, 1, 0, 0,
		0, 47, 48, 5, 44, 0, 0, 48, 14, 1, 0, 0, 0, 49, 50, 5, 100, 0, 0, 50, 51,
		5, 105, 0, 0, 51, 52, 5, 115, 0, 0, 52, 53, 5, 97, 0, 0, 53, 54, 5, 98,
		0, 0, 54, 55, 5, 108, 0, 0, 55, 56, 5, 101, 0, 0, 56, 16, 1, 0, 0, 0, 57,
		58, 5, 105, 0, 0, 58, 59, 5, 110, 0, 0, 59, 60, 5, 116, 0, 0, 60, 61, 5,
		101, 0, 0, 61, 62, 5, 114, 0, 0, 62, 63, 5, 118, 0, 0, 63, 64, 5, 97, 0,
		0, 64, 65, 5, 108, 0, 0, 65, 18, 1, 0, 0, 0, 66, 67, 5, 116, 0, 0, 67,
		68, 5, 97, 0, 0, 68, 69, 5, 98, 0, 0, 69, 70, 5, 108, 0, 0, 70, 71, 5,
		101, 0, 0, 71, 20, 1, 0, 0, 0, 72, 74, 7, 0, 0, 0, 73, 72, 1, 0, 0, 0,
		74, 75, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 22, 1,
		0, 0, 0, 77, 81, 7, 1, 0, 0, 78, 80, 7, 2, 0, 0, 79, 78, 1, 0, 0, 0, 80,
		83, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 24, 1, 0, 0,
		0, 83, 81, 1, 0, 0, 0, 84, 86, 7, 3, 0, 0, 85, 84, 1, 0, 0, 0, 86, 87,
		1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 26, 1, 0, 0, 0,
		89, 93, 5, 35, 0, 0, 90, 92, 8, 4, 0, 0, 91, 90, 1, 0, 0, 0, 92, 95, 1,
		0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 96, 1, 0, 0, 0, 95,
		93, 1, 0, 0, 0, 96, 97, 6, 13, 0, 0, 97, 28, 1, 0, 0, 0, 98, 100, 7, 5,
		0, 0, 99, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101,
		102, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 6, 14, 0, 0, 104, 30,
		1, 0, 0, 0, 105, 107, 5, 92, 0, 0, 106, 108, 5, 13, 0, 0, 107, 106, 1,
		0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 110, 5, 10, 0,
		0, 110, 111, 1, 0, 0, 0, 111, 112, 6, 15, 0, 0, 112, 32, 1, 0, 0, 0, 113,
		115, 5, 13, 0, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 117, 5, 10, 0, 0, 117, 34, 1, 0, 0, 0, 8, 0, 75, 81, 87,
		93, 101, 107, 114, 1, 6, 0, 0,
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
	CONFIGLexerT__0         = 1
	CONFIGLexerT__1         = 2
	CONFIGLexerT__2         = 3
	CONFIGLexerT__3         = 4
	CONFIGLexerT__4         = 5
	CONFIGLexerT__5         = 6
	CONFIGLexerT__6         = 7
	CONFIGLexerDISABLE      = 8
	CONFIGLexerINTERVAL     = 9
	CONFIGLexerTABLE        = 10
	CONFIGLexerNUMBER       = 11
	CONFIGLexerNAME         = 12
	CONFIGLexerHOST         = 13
	CONFIGLexerCOMMENT      = 14
	CONFIGLexerWS           = 15
	CONFIGLexerCONTINUATION = 16
	CONFIGLexerNEWLINE      = 17
)
