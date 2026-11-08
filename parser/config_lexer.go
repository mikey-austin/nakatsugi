// Code generated from CONFIG.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 6, 29, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 3, 3,
	3, 6, 3, 16, 10, 3, 13, 3, 14, 3, 17, 3, 4, 6, 4, 21, 10, 4, 13, 4, 14,
	4, 22, 3, 5, 6, 5, 26, 10, 5, 13, 5, 14, 5, 27, 2, 2, 6, 3, 3, 5, 4, 7,
	5, 9, 6, 3, 2, 6, 3, 2, 99, 124, 5, 2, 50, 59, 97, 97, 99, 124, 5, 2, 48,
	48, 50, 59, 99, 124, 4, 2, 12, 12, 15, 15, 2, 31, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 3, 11, 3, 2, 2, 2, 5,
	13, 3, 2, 2, 2, 7, 20, 3, 2, 2, 2, 9, 25, 3, 2, 2, 2, 11, 12, 7, 63, 2,
	2, 12, 4, 3, 2, 2, 2, 13, 15, 9, 2, 2, 2, 14, 16, 9, 3, 2, 2, 15, 14, 3,
	2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18,
	6, 3, 2, 2, 2, 19, 21, 9, 4, 2, 2, 20, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2,
	2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 8, 3, 2, 2, 2, 24, 26, 9,
	5, 2, 2, 25, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27,
	28, 3, 2, 2, 2, 28, 10, 3, 2, 2, 2, 6, 2, 17, 22, 27, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "VAR", "HOST", "NEWLINE",
}

var lexerRuleNames = []string{
	"T__0", "VAR", "HOST", "NEWLINE",
}

type CONFIGLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewCONFIGLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *CONFIGLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCONFIGLexer(input antlr.CharStream) *CONFIGLexer {
	l := new(CONFIGLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
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
