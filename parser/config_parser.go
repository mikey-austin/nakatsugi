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
		"", "'$'", "'='", "','", "'<'", "'>'", "'{'", "'}'", "'proto'", "'check'",
		"'http'", "'https'", "'code'", "'digest'", "'prefork'", "'priority'",
		"'retry'", "'parent'", "'ip'", "'ttl'", "'route'", "'sticky-address'",
		"'session'", "'timeout'", "'forward'", "'port'", "'interface'", "'listen'",
		"'to'", "'on'", "'host'", "'redirect'", "'disable'", "'interval'", "'table'",
	}
	staticData.SymbolicNames = []string{
		"", "DOLLA", "EQUALS", "COMMA", "LANGLE", "RANGLE", "LCURLY", "RCURLY",
		"PROTO", "CHECK", "HTTP", "HTTPS", "CODE", "DIGEST", "PREFORK", "PRIORITY",
		"RETRY", "PARENT", "IP", "TTL", "ROUTE", "STICKY_ADDRESS", "SESSION",
		"TIMEOUT", "FORWARD", "PORT", "INTERFACE", "LISTEN", "TO", "ON", "HOST",
		"REDIRECT", "DISABLE", "INTERVAL", "TABLE", "NUMBER", "NAME", "QNAME",
		"COMMENT", "WS", "CONTINUATION", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"config", "stmt", "assignment", "var", "gopts", "interval", "prefork",
		"table", "tname", "tlist", "titem", "tdisable", "topts", "toptipttl",
		"toptparent", "toptpriority", "toptretry", "listen", "port", "portrange",
		"lport", "proto", "interface", "check", "ctype", "chost", "cdigest",
		"chttpcode", "chttpdigest", "chttpscode", "chttpsdigest", "forward",
		"route", "sesstimeout", "redirect", "rname", "rstmt",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 293, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 79, 8, 0, 10, 0, 12, 0, 82, 9, 0, 1, 0, 3,
		0, 85, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 93, 8, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 104, 8, 4, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 115, 8, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 128,
		8, 9, 10, 9, 12, 9, 131, 9, 9, 1, 9, 3, 9, 134, 8, 9, 3, 9, 136, 8, 9,
		1, 10, 1, 10, 5, 10, 140, 8, 10, 10, 10, 12, 10, 143, 9, 10, 1, 10, 1,
		10, 5, 10, 147, 8, 10, 10, 10, 12, 10, 150, 9, 10, 3, 10, 152, 8, 10, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 160, 8, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 179, 8, 17, 1, 17, 1, 17, 3, 17,
		183, 8, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 3,
		20, 193, 8, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 208, 8, 24, 1, 25, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 3, 27, 219, 8, 27, 1, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 3, 28, 227, 8, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 3, 29, 234, 8, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30,
		1, 30, 3, 30, 242, 8, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3,
		31, 250, 8, 31, 1, 31, 3, 31, 253, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 3,
		32, 259, 8, 32, 1, 32, 3, 32, 262, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 273, 8, 34, 10, 34, 12, 34, 276,
		9, 34, 1, 34, 3, 34, 279, 8, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 291, 8, 36, 1, 36, 0, 0, 37, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 0,
		0, 294, 0, 80, 1, 0, 0, 0, 2, 92, 1, 0, 0, 0, 4, 94, 1, 0, 0, 0, 6, 98,
		1, 0, 0, 0, 8, 103, 1, 0, 0, 0, 10, 105, 1, 0, 0, 0, 12, 108, 1, 0, 0,
		0, 14, 111, 1, 0, 0, 0, 16, 120, 1, 0, 0, 0, 18, 135, 1, 0, 0, 0, 20, 151,
		1, 0, 0, 0, 22, 153, 1, 0, 0, 0, 24, 159, 1, 0, 0, 0, 26, 161, 1, 0, 0,
		0, 28, 165, 1, 0, 0, 0, 30, 168, 1, 0, 0, 0, 32, 171, 1, 0, 0, 0, 34, 174,
		1, 0, 0, 0, 36, 184, 1, 0, 0, 0, 38, 187, 1, 0, 0, 0, 40, 192, 1, 0, 0,
		0, 42, 194, 1, 0, 0, 0, 44, 197, 1, 0, 0, 0, 46, 200, 1, 0, 0, 0, 48, 207,
		1, 0, 0, 0, 50, 209, 1, 0, 0, 0, 52, 212, 1, 0, 0, 0, 54, 215, 1, 0, 0,
		0, 56, 223, 1, 0, 0, 0, 58, 230, 1, 0, 0, 0, 60, 238, 1, 0, 0, 0, 62, 245,
		1, 0, 0, 0, 64, 254, 1, 0, 0, 0, 66, 263, 1, 0, 0, 0, 68, 267, 1, 0, 0,
		0, 70, 282, 1, 0, 0, 0, 72, 290, 1, 0, 0, 0, 74, 79, 5, 41, 0, 0, 75, 76,
		3, 2, 1, 0, 76, 77, 5, 41, 0, 0, 77, 79, 1, 0, 0, 0, 78, 74, 1, 0, 0, 0,
		78, 75, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1,
		0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 83, 85, 3, 2, 1, 0, 84,
		83, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87, 5, 0, 0,
		1, 87, 1, 1, 0, 0, 0, 88, 93, 3, 4, 2, 0, 89, 93, 3, 14, 7, 0, 90, 93,
		3, 8, 4, 0, 91, 93, 3, 68, 34, 0, 92, 88, 1, 0, 0, 0, 92, 89, 1, 0, 0,
		0, 92, 90, 1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 3, 1, 0, 0, 0, 94, 95, 5,
		36, 0, 0, 95, 96, 5, 2, 0, 0, 96, 97, 5, 36, 0, 0, 97, 5, 1, 0, 0, 0, 98,
		99, 5, 1, 0, 0, 99, 100, 5, 36, 0, 0, 100, 7, 1, 0, 0, 0, 101, 104, 3,
		10, 5, 0, 102, 104, 3, 12, 6, 0, 103, 101, 1, 0, 0, 0, 103, 102, 1, 0,
		0, 0, 104, 9, 1, 0, 0, 0, 105, 106, 5, 33, 0, 0, 106, 107, 5, 35, 0, 0,
		107, 11, 1, 0, 0, 0, 108, 109, 5, 14, 0, 0, 109, 110, 5, 35, 0, 0, 110,
		13, 1, 0, 0, 0, 111, 112, 5, 34, 0, 0, 112, 114, 3, 16, 8, 0, 113, 115,
		3, 22, 11, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116, 1,
		0, 0, 0, 116, 117, 5, 6, 0, 0, 117, 118, 3, 18, 9, 0, 118, 119, 5, 7, 0,
		0, 119, 15, 1, 0, 0, 0, 120, 121, 5, 4, 0, 0, 121, 122, 5, 36, 0, 0, 122,
		123, 5, 5, 0, 0, 123, 17, 1, 0, 0, 0, 124, 129, 3, 20, 10, 0, 125, 126,
		5, 3, 0, 0, 126, 128, 3, 20, 10, 0, 127, 125, 1, 0, 0, 0, 128, 131, 1,
		0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 133, 1, 0, 0,
		0, 131, 129, 1, 0, 0, 0, 132, 134, 5, 3, 0, 0, 133, 132, 1, 0, 0, 0, 133,
		134, 1, 0, 0, 0, 134, 136, 1, 0, 0, 0, 135, 124, 1, 0, 0, 0, 135, 136,
		1, 0, 0, 0, 136, 19, 1, 0, 0, 0, 137, 141, 3, 6, 3, 0, 138, 140, 3, 24,
		12, 0, 139, 138, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0,
		141, 142, 1, 0, 0, 0, 142, 152, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144,
		148, 5, 36, 0, 0, 145, 147, 3, 24, 12, 0, 146, 145, 1, 0, 0, 0, 147, 150,
		1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 152, 1, 0,
		0, 0, 150, 148, 1, 0, 0, 0, 151, 137, 1, 0, 0, 0, 151, 144, 1, 0, 0, 0,
		152, 21, 1, 0, 0, 0, 153, 154, 5, 32, 0, 0, 154, 23, 1, 0, 0, 0, 155, 160,
		3, 26, 13, 0, 156, 160, 3, 28, 14, 0, 157, 160, 3, 30, 15, 0, 158, 160,
		3, 32, 16, 0, 159, 155, 1, 0, 0, 0, 159, 156, 1, 0, 0, 0, 159, 157, 1,
		0, 0, 0, 159, 158, 1, 0, 0, 0, 160, 25, 1, 0, 0, 0, 161, 162, 5, 18, 0,
		0, 162, 163, 5, 19, 0, 0, 163, 164, 5, 35, 0, 0, 164, 27, 1, 0, 0, 0, 165,
		166, 5, 17, 0, 0, 166, 167, 5, 35, 0, 0, 167, 29, 1, 0, 0, 0, 168, 169,
		5, 15, 0, 0, 169, 170, 5, 35, 0, 0, 170, 31, 1, 0, 0, 0, 171, 172, 5, 16,
		0, 0, 172, 173, 5, 35, 0, 0, 173, 33, 1, 0, 0, 0, 174, 175, 5, 27, 0, 0,
		175, 176, 5, 29, 0, 0, 176, 178, 5, 36, 0, 0, 177, 179, 3, 42, 21, 0, 178,
		177, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 182,
		3, 40, 20, 0, 181, 183, 3, 44, 22, 0, 182, 181, 1, 0, 0, 0, 182, 183, 1,
		0, 0, 0, 183, 35, 1, 0, 0, 0, 184, 185, 5, 25, 0, 0, 185, 186, 5, 35, 0,
		0, 186, 37, 1, 0, 0, 0, 187, 188, 5, 25, 0, 0, 188, 189, 5, 36, 0, 0, 189,
		39, 1, 0, 0, 0, 190, 193, 3, 38, 19, 0, 191, 193, 3, 36, 18, 0, 192, 190,
		1, 0, 0, 0, 192, 191, 1, 0, 0, 0, 193, 41, 1, 0, 0, 0, 194, 195, 5, 8,
		0, 0, 195, 196, 5, 36, 0, 0, 196, 43, 1, 0, 0, 0, 197, 198, 5, 26, 0, 0,
		198, 199, 5, 36, 0, 0, 199, 45, 1, 0, 0, 0, 200, 201, 5, 9, 0, 0, 201,
		202, 3, 48, 24, 0, 202, 47, 1, 0, 0, 0, 203, 208, 3, 54, 27, 0, 204, 208,
		3, 56, 28, 0, 205, 208, 3, 58, 29, 0, 206, 208, 3, 60, 30, 0, 207, 203,
		1, 0, 0, 0, 207, 204, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 206, 1, 0,
		0, 0, 208, 49, 1, 0, 0, 0, 209, 210, 5, 30, 0, 0, 210, 211, 5, 36, 0, 0,
		211, 51, 1, 0, 0, 0, 212, 213, 5, 13, 0, 0, 213, 214, 5, 36, 0, 0, 214,
		53, 1, 0, 0, 0, 215, 216, 5, 10, 0, 0, 216, 218, 5, 37, 0, 0, 217, 219,
		3, 50, 25, 0, 218, 217, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 220, 1,
		0, 0, 0, 220, 221, 5, 12, 0, 0, 221, 222, 5, 35, 0, 0, 222, 55, 1, 0, 0,
		0, 223, 224, 5, 10, 0, 0, 224, 226, 5, 37, 0, 0, 225, 227, 3, 50, 25, 0,
		226, 225, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228,
		229, 3, 52, 26, 0, 229, 57, 1, 0, 0, 0, 230, 231, 5, 11, 0, 0, 231, 233,
		5, 37, 0, 0, 232, 234, 3, 50, 25, 0, 233, 232, 1, 0, 0, 0, 233, 234, 1,
		0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 236, 5, 12, 0, 0, 236, 237, 5, 35,
		0, 0, 237, 59, 1, 0, 0, 0, 238, 239, 5, 11, 0, 0, 239, 241, 5, 37, 0, 0,
		240, 242, 3, 50, 25, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242,
		243, 1, 0, 0, 0, 243, 244, 3, 52, 26, 0, 244, 61, 1, 0, 0, 0, 245, 246,
		5, 24, 0, 0, 246, 247, 5, 28, 0, 0, 247, 249, 3, 16, 8, 0, 248, 250, 3,
		36, 18, 0, 249, 248, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 252, 1, 0,
		0, 0, 251, 253, 3, 46, 23, 0, 252, 251, 1, 0, 0, 0, 252, 253, 1, 0, 0,
		0, 253, 63, 1, 0, 0, 0, 254, 255, 5, 20, 0, 0, 255, 256, 5, 28, 0, 0, 256,
		258, 3, 16, 8, 0, 257, 259, 3, 36, 18, 0, 258, 257, 1, 0, 0, 0, 258, 259,
		1, 0, 0, 0, 259, 261, 1, 0, 0, 0, 260, 262, 3, 46, 23, 0, 261, 260, 1,
		0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 65, 1, 0, 0, 0, 263, 264, 5, 22, 0,
		0, 264, 265, 5, 23, 0, 0, 265, 266, 5, 35, 0, 0, 266, 67, 1, 0, 0, 0, 267,
		268, 5, 31, 0, 0, 268, 269, 3, 70, 35, 0, 269, 274, 5, 6, 0, 0, 270, 273,
		5, 41, 0, 0, 271, 273, 3, 72, 36, 0, 272, 270, 1, 0, 0, 0, 272, 271, 1,
		0, 0, 0, 273, 276, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274, 275, 1, 0, 0,
		0, 275, 278, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 277, 279, 3, 72, 36, 0,
		278, 277, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280,
		281, 5, 7, 0, 0, 281, 69, 1, 0, 0, 0, 282, 283, 5, 37, 0, 0, 283, 71, 1,
		0, 0, 0, 284, 291, 5, 32, 0, 0, 285, 291, 5, 21, 0, 0, 286, 291, 3, 66,
		33, 0, 287, 291, 3, 34, 17, 0, 288, 291, 3, 62, 31, 0, 289, 291, 3, 64,
		32, 0, 290, 284, 1, 0, 0, 0, 290, 285, 1, 0, 0, 0, 290, 286, 1, 0, 0, 0,
		290, 287, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 289, 1, 0, 0, 0, 291,
		73, 1, 0, 0, 0, 29, 78, 80, 84, 92, 103, 114, 129, 133, 135, 141, 148,
		151, 159, 178, 182, 192, 207, 218, 226, 233, 241, 249, 252, 258, 261, 272,
		274, 278, 290,
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
	CONFIGParserEOF            = antlr.TokenEOF
	CONFIGParserDOLLA          = 1
	CONFIGParserEQUALS         = 2
	CONFIGParserCOMMA          = 3
	CONFIGParserLANGLE         = 4
	CONFIGParserRANGLE         = 5
	CONFIGParserLCURLY         = 6
	CONFIGParserRCURLY         = 7
	CONFIGParserPROTO          = 8
	CONFIGParserCHECK          = 9
	CONFIGParserHTTP           = 10
	CONFIGParserHTTPS          = 11
	CONFIGParserCODE           = 12
	CONFIGParserDIGEST         = 13
	CONFIGParserPREFORK        = 14
	CONFIGParserPRIORITY       = 15
	CONFIGParserRETRY          = 16
	CONFIGParserPARENT         = 17
	CONFIGParserIP             = 18
	CONFIGParserTTL            = 19
	CONFIGParserROUTE          = 20
	CONFIGParserSTICKY_ADDRESS = 21
	CONFIGParserSESSION        = 22
	CONFIGParserTIMEOUT        = 23
	CONFIGParserFORWARD        = 24
	CONFIGParserPORT           = 25
	CONFIGParserINTERFACE      = 26
	CONFIGParserLISTEN         = 27
	CONFIGParserTO             = 28
	CONFIGParserON             = 29
	CONFIGParserHOST           = 30
	CONFIGParserREDIRECT       = 31
	CONFIGParserDISABLE        = 32
	CONFIGParserINTERVAL       = 33
	CONFIGParserTABLE          = 34
	CONFIGParserNUMBER         = 35
	CONFIGParserNAME           = 36
	CONFIGParserQNAME          = 37
	CONFIGParserCOMMENT        = 38
	CONFIGParserWS             = 39
	CONFIGParserCONTINUATION   = 40
	CONFIGParserNEWLINE        = 41
)

// CONFIGParser rules.
const (
	CONFIGParserRULE_config       = 0
	CONFIGParserRULE_stmt         = 1
	CONFIGParserRULE_assignment   = 2
	CONFIGParserRULE_var          = 3
	CONFIGParserRULE_gopts        = 4
	CONFIGParserRULE_interval     = 5
	CONFIGParserRULE_prefork      = 6
	CONFIGParserRULE_table        = 7
	CONFIGParserRULE_tname        = 8
	CONFIGParserRULE_tlist        = 9
	CONFIGParserRULE_titem        = 10
	CONFIGParserRULE_tdisable     = 11
	CONFIGParserRULE_topts        = 12
	CONFIGParserRULE_toptipttl    = 13
	CONFIGParserRULE_toptparent   = 14
	CONFIGParserRULE_toptpriority = 15
	CONFIGParserRULE_toptretry    = 16
	CONFIGParserRULE_listen       = 17
	CONFIGParserRULE_port         = 18
	CONFIGParserRULE_portrange    = 19
	CONFIGParserRULE_lport        = 20
	CONFIGParserRULE_proto        = 21
	CONFIGParserRULE_interface    = 22
	CONFIGParserRULE_check        = 23
	CONFIGParserRULE_ctype        = 24
	CONFIGParserRULE_chost        = 25
	CONFIGParserRULE_cdigest      = 26
	CONFIGParserRULE_chttpcode    = 27
	CONFIGParserRULE_chttpdigest  = 28
	CONFIGParserRULE_chttpscode   = 29
	CONFIGParserRULE_chttpsdigest = 30
	CONFIGParserRULE_forward      = 31
	CONFIGParserRULE_route        = 32
	CONFIGParserRULE_sesstimeout  = 33
	CONFIGParserRULE_redirect     = 34
	CONFIGParserRULE_rname        = 35
	CONFIGParserRULE_rstmt        = 36
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
	p.SetState(80)
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
			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case CONFIGParserNEWLINE:
				{
					p.SetState(74)
					p.Match(CONFIGParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case CONFIGParserPREFORK, CONFIGParserREDIRECT, CONFIGParserINTERVAL, CONFIGParserTABLE, CONFIGParserNAME:
				{
					p.SetState(75)
					p.Stmt()
				}
				{
					p.SetState(76)
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
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&96636780544) != 0 {
		{
			p.SetState(83)
			p.Stmt()
		}

	}
	{
		p.SetState(86)
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
	Gopts() IGoptsContext
	Redirect() IRedirectContext

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

func (s *StmtContext) Gopts() IGoptsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoptsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoptsContext)
}

func (s *StmtContext) Redirect() IRedirectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRedirectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRedirectContext)
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
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Assignment()
		}

	case CONFIGParserTABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Table()
		}

	case CONFIGParserPREFORK, CONFIGParserINTERVAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Gopts()
		}

	case CONFIGParserREDIRECT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.Redirect()
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
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode
	EQUALS() antlr.TerminalNode

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

func (s *AssignmentContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(CONFIGParserNAME)
}

func (s *AssignmentContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, i)
}

func (s *AssignmentContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(CONFIGParserEQUALS, 0)
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
		p.SetState(94)
		p.Match(CONFIGParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(CONFIGParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
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

// IVarContext is an interface to support dynamic dispatch.
type IVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOLLA() antlr.TerminalNode
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

func (s *VarContext) DOLLA() antlr.TerminalNode {
	return s.GetToken(CONFIGParserDOLLA, 0)
}

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
	p.EnterRule(localctx, 6, CONFIGParserRULE_var)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(CONFIGParserDOLLA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
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

// IGoptsContext is an interface to support dynamic dispatch.
type IGoptsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Interval() IIntervalContext
	Prefork() IPreforkContext

	// IsGoptsContext differentiates from other interfaces.
	IsGoptsContext()
}

type GoptsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoptsContext() *GoptsContext {
	var p = new(GoptsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_gopts
	return p
}

func InitEmptyGoptsContext(p *GoptsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_gopts
}

func (*GoptsContext) IsGoptsContext() {}

func NewGoptsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoptsContext {
	var p = new(GoptsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_gopts

	return p
}

func (s *GoptsContext) GetParser() antlr.Parser { return s.parser }

func (s *GoptsContext) Interval() IIntervalContext {
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

func (s *GoptsContext) Prefork() IPreforkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPreforkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPreforkContext)
}

func (s *GoptsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoptsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoptsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterGopts(s)
	}
}

func (s *GoptsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitGopts(s)
	}
}

func (p *CONFIGParser) Gopts() (localctx IGoptsContext) {
	localctx = NewGoptsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CONFIGParserRULE_gopts)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserINTERVAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Interval()
		}

	case CONFIGParserPREFORK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Prefork()
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
	p.EnterRule(localctx, 10, CONFIGParserRULE_interval)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(CONFIGParserINTERVAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
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

// IPreforkContext is an interface to support dynamic dispatch.
type IPreforkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PREFORK() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsPreforkContext differentiates from other interfaces.
	IsPreforkContext()
}

type PreforkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreforkContext() *PreforkContext {
	var p = new(PreforkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_prefork
	return p
}

func InitEmptyPreforkContext(p *PreforkContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_prefork
}

func (*PreforkContext) IsPreforkContext() {}

func NewPreforkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreforkContext {
	var p = new(PreforkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_prefork

	return p
}

func (s *PreforkContext) GetParser() antlr.Parser { return s.parser }

func (s *PreforkContext) PREFORK() antlr.TerminalNode {
	return s.GetToken(CONFIGParserPREFORK, 0)
}

func (s *PreforkContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNUMBER, 0)
}

func (s *PreforkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreforkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreforkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterPrefork(s)
	}
}

func (s *PreforkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitPrefork(s)
	}
}

func (p *CONFIGParser) Prefork() (localctx IPreforkContext) {
	localctx = NewPreforkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CONFIGParserRULE_prefork)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(CONFIGParserPREFORK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
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

// ITableContext is an interface to support dynamic dispatch.
type ITableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TABLE() antlr.TerminalNode
	Tname() ITnameContext
	LCURLY() antlr.TerminalNode
	Tlist() ITlistContext
	RCURLY() antlr.TerminalNode
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

func (s *TableContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(CONFIGParserLCURLY, 0)
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

func (s *TableContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(CONFIGParserRCURLY, 0)
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
	p.EnterRule(localctx, 14, CONFIGParserRULE_table)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(CONFIGParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Tname()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserDISABLE {
		{
			p.SetState(113)
			p.Tdisable()
		}

	}
	{
		p.SetState(116)
		p.Match(CONFIGParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Tlist()
	}
	{
		p.SetState(118)
		p.Match(CONFIGParserRCURLY)
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
	LANGLE() antlr.TerminalNode
	NAME() antlr.TerminalNode
	RANGLE() antlr.TerminalNode

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

func (s *TnameContext) LANGLE() antlr.TerminalNode {
	return s.GetToken(CONFIGParserLANGLE, 0)
}

func (s *TnameContext) NAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, 0)
}

func (s *TnameContext) RANGLE() antlr.TerminalNode {
	return s.GetToken(CONFIGParserRANGLE, 0)
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
	p.EnterRule(localctx, 16, CONFIGParserRULE_tname)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(CONFIGParserLANGLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(CONFIGParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(CONFIGParserRANGLE)
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
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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

func (s *TlistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CONFIGParserCOMMA)
}

func (s *TlistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CONFIGParserCOMMA, i)
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
	p.EnterRule(localctx, 18, CONFIGParserRULE_tlist)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserDOLLA || _la == CONFIGParserNAME {
		{
			p.SetState(124)
			p.Titem()
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(125)
					p.Match(CONFIGParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(126)
					p.Titem()
				}

			}
			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CONFIGParserCOMMA {
			{
				p.SetState(132)
				p.Match(CONFIGParserCOMMA)
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
	AllTopts() []IToptsContext
	Topts(i int) IToptsContext
	NAME() antlr.TerminalNode

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

func (s *TitemContext) AllTopts() []IToptsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IToptsContext); ok {
			len++
		}
	}

	tst := make([]IToptsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IToptsContext); ok {
			tst[i] = t.(IToptsContext)
			i++
		}
	}

	return tst
}

func (s *TitemContext) Topts(i int) IToptsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IToptsContext); ok {
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

	return t.(IToptsContext)
}

func (s *TitemContext) NAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, 0)
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
	p.EnterRule(localctx, 20, CONFIGParserRULE_titem)
	var _la int

	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserDOLLA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Var_()
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&491520) != 0 {
			{
				p.SetState(138)
				p.Topts()
			}

			p.SetState(143)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case CONFIGParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.Match(CONFIGParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&491520) != 0 {
			{
				p.SetState(145)
				p.Topts()
			}

			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 22, CONFIGParserRULE_tdisable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
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

// IToptsContext is an interface to support dynamic dispatch.
type IToptsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Toptipttl() IToptipttlContext
	Toptparent() IToptparentContext
	Toptpriority() IToptpriorityContext
	Toptretry() IToptretryContext

	// IsToptsContext differentiates from other interfaces.
	IsToptsContext()
}

type ToptsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyToptsContext() *ToptsContext {
	var p = new(ToptsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_topts
	return p
}

func InitEmptyToptsContext(p *ToptsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_topts
}

func (*ToptsContext) IsToptsContext() {}

func NewToptsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ToptsContext {
	var p = new(ToptsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_topts

	return p
}

func (s *ToptsContext) GetParser() antlr.Parser { return s.parser }

func (s *ToptsContext) Toptipttl() IToptipttlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IToptipttlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IToptipttlContext)
}

func (s *ToptsContext) Toptparent() IToptparentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IToptparentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IToptparentContext)
}

func (s *ToptsContext) Toptpriority() IToptpriorityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IToptpriorityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IToptpriorityContext)
}

func (s *ToptsContext) Toptretry() IToptretryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IToptretryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IToptretryContext)
}

func (s *ToptsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToptsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ToptsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterTopts(s)
	}
}

func (s *ToptsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitTopts(s)
	}
}

func (p *CONFIGParser) Topts() (localctx IToptsContext) {
	localctx = NewToptsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CONFIGParserRULE_topts)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserIP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Toptipttl()
		}

	case CONFIGParserPARENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Toptparent()
		}

	case CONFIGParserPRIORITY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.Toptpriority()
		}

	case CONFIGParserRETRY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(158)
			p.Toptretry()
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

// IToptipttlContext is an interface to support dynamic dispatch.
type IToptipttlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IP() antlr.TerminalNode
	TTL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsToptipttlContext differentiates from other interfaces.
	IsToptipttlContext()
}

type ToptipttlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyToptipttlContext() *ToptipttlContext {
	var p = new(ToptipttlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_toptipttl
	return p
}

func InitEmptyToptipttlContext(p *ToptipttlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_toptipttl
}

func (*ToptipttlContext) IsToptipttlContext() {}

func NewToptipttlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ToptipttlContext {
	var p = new(ToptipttlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_toptipttl

	return p
}

func (s *ToptipttlContext) GetParser() antlr.Parser { return s.parser }

func (s *ToptipttlContext) IP() antlr.TerminalNode {
	return s.GetToken(CONFIGParserIP, 0)
}

func (s *ToptipttlContext) TTL() antlr.TerminalNode {
	return s.GetToken(CONFIGParserTTL, 0)
}

func (s *ToptipttlContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNUMBER, 0)
}

func (s *ToptipttlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToptipttlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ToptipttlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterToptipttl(s)
	}
}

func (s *ToptipttlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitToptipttl(s)
	}
}

func (p *CONFIGParser) Toptipttl() (localctx IToptipttlContext) {
	localctx = NewToptipttlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CONFIGParserRULE_toptipttl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(CONFIGParserIP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(CONFIGParserTTL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
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

// IToptparentContext is an interface to support dynamic dispatch.
type IToptparentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARENT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsToptparentContext differentiates from other interfaces.
	IsToptparentContext()
}

type ToptparentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyToptparentContext() *ToptparentContext {
	var p = new(ToptparentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_toptparent
	return p
}

func InitEmptyToptparentContext(p *ToptparentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_toptparent
}

func (*ToptparentContext) IsToptparentContext() {}

func NewToptparentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ToptparentContext {
	var p = new(ToptparentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_toptparent

	return p
}

func (s *ToptparentContext) GetParser() antlr.Parser { return s.parser }

func (s *ToptparentContext) PARENT() antlr.TerminalNode {
	return s.GetToken(CONFIGParserPARENT, 0)
}

func (s *ToptparentContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNUMBER, 0)
}

func (s *ToptparentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToptparentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ToptparentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterToptparent(s)
	}
}

func (s *ToptparentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitToptparent(s)
	}
}

func (p *CONFIGParser) Toptparent() (localctx IToptparentContext) {
	localctx = NewToptparentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CONFIGParserRULE_toptparent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(CONFIGParserPARENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
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

// IToptpriorityContext is an interface to support dynamic dispatch.
type IToptpriorityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRIORITY() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsToptpriorityContext differentiates from other interfaces.
	IsToptpriorityContext()
}

type ToptpriorityContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyToptpriorityContext() *ToptpriorityContext {
	var p = new(ToptpriorityContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_toptpriority
	return p
}

func InitEmptyToptpriorityContext(p *ToptpriorityContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_toptpriority
}

func (*ToptpriorityContext) IsToptpriorityContext() {}

func NewToptpriorityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ToptpriorityContext {
	var p = new(ToptpriorityContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_toptpriority

	return p
}

func (s *ToptpriorityContext) GetParser() antlr.Parser { return s.parser }

func (s *ToptpriorityContext) PRIORITY() antlr.TerminalNode {
	return s.GetToken(CONFIGParserPRIORITY, 0)
}

func (s *ToptpriorityContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNUMBER, 0)
}

func (s *ToptpriorityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToptpriorityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ToptpriorityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterToptpriority(s)
	}
}

func (s *ToptpriorityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitToptpriority(s)
	}
}

func (p *CONFIGParser) Toptpriority() (localctx IToptpriorityContext) {
	localctx = NewToptpriorityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CONFIGParserRULE_toptpriority)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(CONFIGParserPRIORITY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
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

// IToptretryContext is an interface to support dynamic dispatch.
type IToptretryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETRY() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsToptretryContext differentiates from other interfaces.
	IsToptretryContext()
}

type ToptretryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyToptretryContext() *ToptretryContext {
	var p = new(ToptretryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_toptretry
	return p
}

func InitEmptyToptretryContext(p *ToptretryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_toptretry
}

func (*ToptretryContext) IsToptretryContext() {}

func NewToptretryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ToptretryContext {
	var p = new(ToptretryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_toptretry

	return p
}

func (s *ToptretryContext) GetParser() antlr.Parser { return s.parser }

func (s *ToptretryContext) RETRY() antlr.TerminalNode {
	return s.GetToken(CONFIGParserRETRY, 0)
}

func (s *ToptretryContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNUMBER, 0)
}

func (s *ToptretryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToptretryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ToptretryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterToptretry(s)
	}
}

func (s *ToptretryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitToptretry(s)
	}
}

func (p *CONFIGParser) Toptretry() (localctx IToptretryContext) {
	localctx = NewToptretryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CONFIGParserRULE_toptretry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(CONFIGParserRETRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
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

// IListenContext is an interface to support dynamic dispatch.
type IListenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LISTEN() antlr.TerminalNode
	ON() antlr.TerminalNode
	NAME() antlr.TerminalNode
	Lport() ILportContext
	Proto() IProtoContext
	Interface_() IInterfaceContext

	// IsListenContext differentiates from other interfaces.
	IsListenContext()
}

type ListenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListenContext() *ListenContext {
	var p = new(ListenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_listen
	return p
}

func InitEmptyListenContext(p *ListenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_listen
}

func (*ListenContext) IsListenContext() {}

func NewListenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListenContext {
	var p = new(ListenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_listen

	return p
}

func (s *ListenContext) GetParser() antlr.Parser { return s.parser }

func (s *ListenContext) LISTEN() antlr.TerminalNode {
	return s.GetToken(CONFIGParserLISTEN, 0)
}

func (s *ListenContext) ON() antlr.TerminalNode {
	return s.GetToken(CONFIGParserON, 0)
}

func (s *ListenContext) NAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, 0)
}

func (s *ListenContext) Lport() ILportContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILportContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILportContext)
}

func (s *ListenContext) Proto() IProtoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProtoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProtoContext)
}

func (s *ListenContext) Interface_() IInterfaceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceContext)
}

func (s *ListenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterListen(s)
	}
}

func (s *ListenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitListen(s)
	}
}

func (p *CONFIGParser) Listen() (localctx IListenContext) {
	localctx = NewListenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CONFIGParserRULE_listen)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(CONFIGParserLISTEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(CONFIGParserON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(CONFIGParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserPROTO {
		{
			p.SetState(177)
			p.Proto()
		}

	}
	{
		p.SetState(180)
		p.Lport()
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserINTERFACE {
		{
			p.SetState(181)
			p.Interface_()
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

// IPortContext is an interface to support dynamic dispatch.
type IPortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PORT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsPortContext differentiates from other interfaces.
	IsPortContext()
}

type PortContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortContext() *PortContext {
	var p = new(PortContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_port
	return p
}

func InitEmptyPortContext(p *PortContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_port
}

func (*PortContext) IsPortContext() {}

func NewPortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortContext {
	var p = new(PortContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_port

	return p
}

func (s *PortContext) GetParser() antlr.Parser { return s.parser }

func (s *PortContext) PORT() antlr.TerminalNode {
	return s.GetToken(CONFIGParserPORT, 0)
}

func (s *PortContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNUMBER, 0)
}

func (s *PortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterPort(s)
	}
}

func (s *PortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitPort(s)
	}
}

func (p *CONFIGParser) Port() (localctx IPortContext) {
	localctx = NewPortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CONFIGParserRULE_port)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(CONFIGParserPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
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

// IPortrangeContext is an interface to support dynamic dispatch.
type IPortrangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PORT() antlr.TerminalNode
	NAME() antlr.TerminalNode

	// IsPortrangeContext differentiates from other interfaces.
	IsPortrangeContext()
}

type PortrangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortrangeContext() *PortrangeContext {
	var p = new(PortrangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_portrange
	return p
}

func InitEmptyPortrangeContext(p *PortrangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_portrange
}

func (*PortrangeContext) IsPortrangeContext() {}

func NewPortrangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortrangeContext {
	var p = new(PortrangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_portrange

	return p
}

func (s *PortrangeContext) GetParser() antlr.Parser { return s.parser }

func (s *PortrangeContext) PORT() antlr.TerminalNode {
	return s.GetToken(CONFIGParserPORT, 0)
}

func (s *PortrangeContext) NAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, 0)
}

func (s *PortrangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortrangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortrangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterPortrange(s)
	}
}

func (s *PortrangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitPortrange(s)
	}
}

func (p *CONFIGParser) Portrange() (localctx IPortrangeContext) {
	localctx = NewPortrangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CONFIGParserRULE_portrange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(CONFIGParserPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
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

// ILportContext is an interface to support dynamic dispatch.
type ILportContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Portrange() IPortrangeContext
	Port() IPortContext

	// IsLportContext differentiates from other interfaces.
	IsLportContext()
}

type LportContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLportContext() *LportContext {
	var p = new(LportContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_lport
	return p
}

func InitEmptyLportContext(p *LportContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_lport
}

func (*LportContext) IsLportContext() {}

func NewLportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LportContext {
	var p = new(LportContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_lport

	return p
}

func (s *LportContext) GetParser() antlr.Parser { return s.parser }

func (s *LportContext) Portrange() IPortrangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPortrangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPortrangeContext)
}

func (s *LportContext) Port() IPortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPortContext)
}

func (s *LportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterLport(s)
	}
}

func (s *LportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitLport(s)
	}
}

func (p *CONFIGParser) Lport() (localctx ILportContext) {
	localctx = NewLportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CONFIGParserRULE_lport)
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.Portrange()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.Port()
		}

	case antlr.ATNInvalidAltNumber:
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

// IProtoContext is an interface to support dynamic dispatch.
type IProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROTO() antlr.TerminalNode
	NAME() antlr.TerminalNode

	// IsProtoContext differentiates from other interfaces.
	IsProtoContext()
}

type ProtoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtoContext() *ProtoContext {
	var p = new(ProtoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_proto
	return p
}

func InitEmptyProtoContext(p *ProtoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_proto
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) PROTO() antlr.TerminalNode {
	return s.GetToken(CONFIGParserPROTO, 0)
}

func (s *ProtoContext) NAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, 0)
}

func (s *ProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterProto(s)
	}
}

func (s *ProtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitProto(s)
	}
}

func (p *CONFIGParser) Proto() (localctx IProtoContext) {
	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CONFIGParserRULE_proto)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(CONFIGParserPROTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
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

// IInterfaceContext is an interface to support dynamic dispatch.
type IInterfaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTERFACE() antlr.TerminalNode
	NAME() antlr.TerminalNode

	// IsInterfaceContext differentiates from other interfaces.
	IsInterfaceContext()
}

type InterfaceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceContext() *InterfaceContext {
	var p = new(InterfaceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_interface
	return p
}

func InitEmptyInterfaceContext(p *InterfaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_interface
}

func (*InterfaceContext) IsInterfaceContext() {}

func NewInterfaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceContext {
	var p = new(InterfaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_interface

	return p
}

func (s *InterfaceContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(CONFIGParserINTERFACE, 0)
}

func (s *InterfaceContext) NAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, 0)
}

func (s *InterfaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterInterface(s)
	}
}

func (s *InterfaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitInterface(s)
	}
}

func (p *CONFIGParser) Interface_() (localctx IInterfaceContext) {
	localctx = NewInterfaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CONFIGParserRULE_interface)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(CONFIGParserINTERFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
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

// ICheckContext is an interface to support dynamic dispatch.
type ICheckContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CHECK() antlr.TerminalNode
	Ctype() ICtypeContext

	// IsCheckContext differentiates from other interfaces.
	IsCheckContext()
}

type CheckContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCheckContext() *CheckContext {
	var p = new(CheckContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_check
	return p
}

func InitEmptyCheckContext(p *CheckContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_check
}

func (*CheckContext) IsCheckContext() {}

func NewCheckContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CheckContext {
	var p = new(CheckContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_check

	return p
}

func (s *CheckContext) GetParser() antlr.Parser { return s.parser }

func (s *CheckContext) CHECK() antlr.TerminalNode {
	return s.GetToken(CONFIGParserCHECK, 0)
}

func (s *CheckContext) Ctype() ICtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICtypeContext)
}

func (s *CheckContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CheckContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CheckContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterCheck(s)
	}
}

func (s *CheckContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitCheck(s)
	}
}

func (p *CONFIGParser) Check() (localctx ICheckContext) {
	localctx = NewCheckContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CONFIGParserRULE_check)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(CONFIGParserCHECK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Ctype()
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

// ICtypeContext is an interface to support dynamic dispatch.
type ICtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Chttpcode() IChttpcodeContext
	Chttpdigest() IChttpdigestContext
	Chttpscode() IChttpscodeContext
	Chttpsdigest() IChttpsdigestContext

	// IsCtypeContext differentiates from other interfaces.
	IsCtypeContext()
}

type CtypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCtypeContext() *CtypeContext {
	var p = new(CtypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_ctype
	return p
}

func InitEmptyCtypeContext(p *CtypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_ctype
}

func (*CtypeContext) IsCtypeContext() {}

func NewCtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CtypeContext {
	var p = new(CtypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_ctype

	return p
}

func (s *CtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CtypeContext) Chttpcode() IChttpcodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChttpcodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChttpcodeContext)
}

func (s *CtypeContext) Chttpdigest() IChttpdigestContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChttpdigestContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChttpdigestContext)
}

func (s *CtypeContext) Chttpscode() IChttpscodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChttpscodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChttpscodeContext)
}

func (s *CtypeContext) Chttpsdigest() IChttpsdigestContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChttpsdigestContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChttpsdigestContext)
}

func (s *CtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterCtype(s)
	}
}

func (s *CtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitCtype(s)
	}
}

func (p *CONFIGParser) Ctype() (localctx ICtypeContext) {
	localctx = NewCtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CONFIGParserRULE_ctype)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.Chttpcode()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(204)
			p.Chttpdigest()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(205)
			p.Chttpscode()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(206)
			p.Chttpsdigest()
		}

	case antlr.ATNInvalidAltNumber:
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

// IChostContext is an interface to support dynamic dispatch.
type IChostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HOST() antlr.TerminalNode
	NAME() antlr.TerminalNode

	// IsChostContext differentiates from other interfaces.
	IsChostContext()
}

type ChostContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChostContext() *ChostContext {
	var p = new(ChostContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_chost
	return p
}

func InitEmptyChostContext(p *ChostContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_chost
}

func (*ChostContext) IsChostContext() {}

func NewChostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChostContext {
	var p = new(ChostContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_chost

	return p
}

func (s *ChostContext) GetParser() antlr.Parser { return s.parser }

func (s *ChostContext) HOST() antlr.TerminalNode {
	return s.GetToken(CONFIGParserHOST, 0)
}

func (s *ChostContext) NAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, 0)
}

func (s *ChostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterChost(s)
	}
}

func (s *ChostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitChost(s)
	}
}

func (p *CONFIGParser) Chost() (localctx IChostContext) {
	localctx = NewChostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CONFIGParserRULE_chost)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(CONFIGParserHOST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
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

// ICdigestContext is an interface to support dynamic dispatch.
type ICdigestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DIGEST() antlr.TerminalNode
	NAME() antlr.TerminalNode

	// IsCdigestContext differentiates from other interfaces.
	IsCdigestContext()
}

type CdigestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCdigestContext() *CdigestContext {
	var p = new(CdigestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_cdigest
	return p
}

func InitEmptyCdigestContext(p *CdigestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_cdigest
}

func (*CdigestContext) IsCdigestContext() {}

func NewCdigestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CdigestContext {
	var p = new(CdigestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_cdigest

	return p
}

func (s *CdigestContext) GetParser() antlr.Parser { return s.parser }

func (s *CdigestContext) DIGEST() antlr.TerminalNode {
	return s.GetToken(CONFIGParserDIGEST, 0)
}

func (s *CdigestContext) NAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, 0)
}

func (s *CdigestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CdigestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CdigestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterCdigest(s)
	}
}

func (s *CdigestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitCdigest(s)
	}
}

func (p *CONFIGParser) Cdigest() (localctx ICdigestContext) {
	localctx = NewCdigestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CONFIGParserRULE_cdigest)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(CONFIGParserDIGEST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
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

// IChttpcodeContext is an interface to support dynamic dispatch.
type IChttpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HTTP() antlr.TerminalNode
	QNAME() antlr.TerminalNode
	CODE() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	Chost() IChostContext

	// IsChttpcodeContext differentiates from other interfaces.
	IsChttpcodeContext()
}

type ChttpcodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChttpcodeContext() *ChttpcodeContext {
	var p = new(ChttpcodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_chttpcode
	return p
}

func InitEmptyChttpcodeContext(p *ChttpcodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_chttpcode
}

func (*ChttpcodeContext) IsChttpcodeContext() {}

func NewChttpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChttpcodeContext {
	var p = new(ChttpcodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_chttpcode

	return p
}

func (s *ChttpcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *ChttpcodeContext) HTTP() antlr.TerminalNode {
	return s.GetToken(CONFIGParserHTTP, 0)
}

func (s *ChttpcodeContext) QNAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserQNAME, 0)
}

func (s *ChttpcodeContext) CODE() antlr.TerminalNode {
	return s.GetToken(CONFIGParserCODE, 0)
}

func (s *ChttpcodeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNUMBER, 0)
}

func (s *ChttpcodeContext) Chost() IChostContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChostContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChostContext)
}

func (s *ChttpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChttpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChttpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterChttpcode(s)
	}
}

func (s *ChttpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitChttpcode(s)
	}
}

func (p *CONFIGParser) Chttpcode() (localctx IChttpcodeContext) {
	localctx = NewChttpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CONFIGParserRULE_chttpcode)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(CONFIGParserHTTP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Match(CONFIGParserQNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserHOST {
		{
			p.SetState(217)
			p.Chost()
		}

	}
	{
		p.SetState(220)
		p.Match(CONFIGParserCODE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
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

// IChttpdigestContext is an interface to support dynamic dispatch.
type IChttpdigestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HTTP() antlr.TerminalNode
	QNAME() antlr.TerminalNode
	Cdigest() ICdigestContext
	Chost() IChostContext

	// IsChttpdigestContext differentiates from other interfaces.
	IsChttpdigestContext()
}

type ChttpdigestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChttpdigestContext() *ChttpdigestContext {
	var p = new(ChttpdigestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_chttpdigest
	return p
}

func InitEmptyChttpdigestContext(p *ChttpdigestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_chttpdigest
}

func (*ChttpdigestContext) IsChttpdigestContext() {}

func NewChttpdigestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChttpdigestContext {
	var p = new(ChttpdigestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_chttpdigest

	return p
}

func (s *ChttpdigestContext) GetParser() antlr.Parser { return s.parser }

func (s *ChttpdigestContext) HTTP() antlr.TerminalNode {
	return s.GetToken(CONFIGParserHTTP, 0)
}

func (s *ChttpdigestContext) QNAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserQNAME, 0)
}

func (s *ChttpdigestContext) Cdigest() ICdigestContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICdigestContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICdigestContext)
}

func (s *ChttpdigestContext) Chost() IChostContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChostContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChostContext)
}

func (s *ChttpdigestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChttpdigestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChttpdigestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterChttpdigest(s)
	}
}

func (s *ChttpdigestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitChttpdigest(s)
	}
}

func (p *CONFIGParser) Chttpdigest() (localctx IChttpdigestContext) {
	localctx = NewChttpdigestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CONFIGParserRULE_chttpdigest)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(CONFIGParserHTTP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(CONFIGParserQNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserHOST {
		{
			p.SetState(225)
			p.Chost()
		}

	}
	{
		p.SetState(228)
		p.Cdigest()
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

// IChttpscodeContext is an interface to support dynamic dispatch.
type IChttpscodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HTTPS() antlr.TerminalNode
	QNAME() antlr.TerminalNode
	CODE() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	Chost() IChostContext

	// IsChttpscodeContext differentiates from other interfaces.
	IsChttpscodeContext()
}

type ChttpscodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChttpscodeContext() *ChttpscodeContext {
	var p = new(ChttpscodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_chttpscode
	return p
}

func InitEmptyChttpscodeContext(p *ChttpscodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_chttpscode
}

func (*ChttpscodeContext) IsChttpscodeContext() {}

func NewChttpscodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChttpscodeContext {
	var p = new(ChttpscodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_chttpscode

	return p
}

func (s *ChttpscodeContext) GetParser() antlr.Parser { return s.parser }

func (s *ChttpscodeContext) HTTPS() antlr.TerminalNode {
	return s.GetToken(CONFIGParserHTTPS, 0)
}

func (s *ChttpscodeContext) QNAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserQNAME, 0)
}

func (s *ChttpscodeContext) CODE() antlr.TerminalNode {
	return s.GetToken(CONFIGParserCODE, 0)
}

func (s *ChttpscodeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNUMBER, 0)
}

func (s *ChttpscodeContext) Chost() IChostContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChostContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChostContext)
}

func (s *ChttpscodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChttpscodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChttpscodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterChttpscode(s)
	}
}

func (s *ChttpscodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitChttpscode(s)
	}
}

func (p *CONFIGParser) Chttpscode() (localctx IChttpscodeContext) {
	localctx = NewChttpscodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CONFIGParserRULE_chttpscode)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(CONFIGParserHTTPS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.Match(CONFIGParserQNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserHOST {
		{
			p.SetState(232)
			p.Chost()
		}

	}
	{
		p.SetState(235)
		p.Match(CONFIGParserCODE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
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

// IChttpsdigestContext is an interface to support dynamic dispatch.
type IChttpsdigestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HTTPS() antlr.TerminalNode
	QNAME() antlr.TerminalNode
	Cdigest() ICdigestContext
	Chost() IChostContext

	// IsChttpsdigestContext differentiates from other interfaces.
	IsChttpsdigestContext()
}

type ChttpsdigestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChttpsdigestContext() *ChttpsdigestContext {
	var p = new(ChttpsdigestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_chttpsdigest
	return p
}

func InitEmptyChttpsdigestContext(p *ChttpsdigestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_chttpsdigest
}

func (*ChttpsdigestContext) IsChttpsdigestContext() {}

func NewChttpsdigestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChttpsdigestContext {
	var p = new(ChttpsdigestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_chttpsdigest

	return p
}

func (s *ChttpsdigestContext) GetParser() antlr.Parser { return s.parser }

func (s *ChttpsdigestContext) HTTPS() antlr.TerminalNode {
	return s.GetToken(CONFIGParserHTTPS, 0)
}

func (s *ChttpsdigestContext) QNAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserQNAME, 0)
}

func (s *ChttpsdigestContext) Cdigest() ICdigestContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICdigestContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICdigestContext)
}

func (s *ChttpsdigestContext) Chost() IChostContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChostContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChostContext)
}

func (s *ChttpsdigestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChttpsdigestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChttpsdigestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterChttpsdigest(s)
	}
}

func (s *ChttpsdigestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitChttpsdigest(s)
	}
}

func (p *CONFIGParser) Chttpsdigest() (localctx IChttpsdigestContext) {
	localctx = NewChttpsdigestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CONFIGParserRULE_chttpsdigest)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(CONFIGParserHTTPS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Match(CONFIGParserQNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserHOST {
		{
			p.SetState(240)
			p.Chost()
		}

	}
	{
		p.SetState(243)
		p.Cdigest()
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

// IForwardContext is an interface to support dynamic dispatch.
type IForwardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FORWARD() antlr.TerminalNode
	TO() antlr.TerminalNode
	Tname() ITnameContext
	Port() IPortContext
	Check() ICheckContext

	// IsForwardContext differentiates from other interfaces.
	IsForwardContext()
}

type ForwardContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForwardContext() *ForwardContext {
	var p = new(ForwardContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_forward
	return p
}

func InitEmptyForwardContext(p *ForwardContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_forward
}

func (*ForwardContext) IsForwardContext() {}

func NewForwardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForwardContext {
	var p = new(ForwardContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_forward

	return p
}

func (s *ForwardContext) GetParser() antlr.Parser { return s.parser }

func (s *ForwardContext) FORWARD() antlr.TerminalNode {
	return s.GetToken(CONFIGParserFORWARD, 0)
}

func (s *ForwardContext) TO() antlr.TerminalNode {
	return s.GetToken(CONFIGParserTO, 0)
}

func (s *ForwardContext) Tname() ITnameContext {
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

func (s *ForwardContext) Port() IPortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPortContext)
}

func (s *ForwardContext) Check() ICheckContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICheckContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICheckContext)
}

func (s *ForwardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForwardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForwardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterForward(s)
	}
}

func (s *ForwardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitForward(s)
	}
}

func (p *CONFIGParser) Forward() (localctx IForwardContext) {
	localctx = NewForwardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CONFIGParserRULE_forward)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(CONFIGParserFORWARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Match(CONFIGParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.Tname()
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserPORT {
		{
			p.SetState(248)
			p.Port()
		}

	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserCHECK {
		{
			p.SetState(251)
			p.Check()
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

// IRouteContext is an interface to support dynamic dispatch.
type IRouteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ROUTE() antlr.TerminalNode
	TO() antlr.TerminalNode
	Tname() ITnameContext
	Port() IPortContext
	Check() ICheckContext

	// IsRouteContext differentiates from other interfaces.
	IsRouteContext()
}

type RouteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRouteContext() *RouteContext {
	var p = new(RouteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_route
	return p
}

func InitEmptyRouteContext(p *RouteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_route
}

func (*RouteContext) IsRouteContext() {}

func NewRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RouteContext {
	var p = new(RouteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_route

	return p
}

func (s *RouteContext) GetParser() antlr.Parser { return s.parser }

func (s *RouteContext) ROUTE() antlr.TerminalNode {
	return s.GetToken(CONFIGParserROUTE, 0)
}

func (s *RouteContext) TO() antlr.TerminalNode {
	return s.GetToken(CONFIGParserTO, 0)
}

func (s *RouteContext) Tname() ITnameContext {
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

func (s *RouteContext) Port() IPortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPortContext)
}

func (s *RouteContext) Check() ICheckContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICheckContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICheckContext)
}

func (s *RouteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RouteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RouteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterRoute(s)
	}
}

func (s *RouteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitRoute(s)
	}
}

func (p *CONFIGParser) Route() (localctx IRouteContext) {
	localctx = NewRouteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CONFIGParserRULE_route)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(CONFIGParserROUTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.Match(CONFIGParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.Tname()
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserPORT {
		{
			p.SetState(257)
			p.Port()
		}

	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserCHECK {
		{
			p.SetState(260)
			p.Check()
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

// ISesstimeoutContext is an interface to support dynamic dispatch.
type ISesstimeoutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SESSION() antlr.TerminalNode
	TIMEOUT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsSesstimeoutContext differentiates from other interfaces.
	IsSesstimeoutContext()
}

type SesstimeoutContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySesstimeoutContext() *SesstimeoutContext {
	var p = new(SesstimeoutContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_sesstimeout
	return p
}

func InitEmptySesstimeoutContext(p *SesstimeoutContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_sesstimeout
}

func (*SesstimeoutContext) IsSesstimeoutContext() {}

func NewSesstimeoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SesstimeoutContext {
	var p = new(SesstimeoutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_sesstimeout

	return p
}

func (s *SesstimeoutContext) GetParser() antlr.Parser { return s.parser }

func (s *SesstimeoutContext) SESSION() antlr.TerminalNode {
	return s.GetToken(CONFIGParserSESSION, 0)
}

func (s *SesstimeoutContext) TIMEOUT() antlr.TerminalNode {
	return s.GetToken(CONFIGParserTIMEOUT, 0)
}

func (s *SesstimeoutContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNUMBER, 0)
}

func (s *SesstimeoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SesstimeoutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SesstimeoutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterSesstimeout(s)
	}
}

func (s *SesstimeoutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitSesstimeout(s)
	}
}

func (p *CONFIGParser) Sesstimeout() (localctx ISesstimeoutContext) {
	localctx = NewSesstimeoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CONFIGParserRULE_sesstimeout)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(CONFIGParserSESSION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.Match(CONFIGParserTIMEOUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
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

// IRedirectContext is an interface to support dynamic dispatch.
type IRedirectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REDIRECT() antlr.TerminalNode
	Rname() IRnameContext
	LCURLY() antlr.TerminalNode
	RCURLY() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllRstmt() []IRstmtContext
	Rstmt(i int) IRstmtContext

	// IsRedirectContext differentiates from other interfaces.
	IsRedirectContext()
}

type RedirectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRedirectContext() *RedirectContext {
	var p = new(RedirectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_redirect
	return p
}

func InitEmptyRedirectContext(p *RedirectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_redirect
}

func (*RedirectContext) IsRedirectContext() {}

func NewRedirectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RedirectContext {
	var p = new(RedirectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_redirect

	return p
}

func (s *RedirectContext) GetParser() antlr.Parser { return s.parser }

func (s *RedirectContext) REDIRECT() antlr.TerminalNode {
	return s.GetToken(CONFIGParserREDIRECT, 0)
}

func (s *RedirectContext) Rname() IRnameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRnameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRnameContext)
}

func (s *RedirectContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(CONFIGParserLCURLY, 0)
}

func (s *RedirectContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(CONFIGParserRCURLY, 0)
}

func (s *RedirectContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(CONFIGParserNEWLINE)
}

func (s *RedirectContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(CONFIGParserNEWLINE, i)
}

func (s *RedirectContext) AllRstmt() []IRstmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRstmtContext); ok {
			len++
		}
	}

	tst := make([]IRstmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRstmtContext); ok {
			tst[i] = t.(IRstmtContext)
			i++
		}
	}

	return tst
}

func (s *RedirectContext) Rstmt(i int) IRstmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRstmtContext); ok {
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

	return t.(IRstmtContext)
}

func (s *RedirectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RedirectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RedirectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterRedirect(s)
	}
}

func (s *RedirectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitRedirect(s)
	}
}

func (p *CONFIGParser) Redirect() (localctx IRedirectContext) {
	localctx = NewRedirectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CONFIGParserRULE_redirect)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(CONFIGParserREDIRECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Rname()
	}
	{
		p.SetState(269)
		p.Match(CONFIGParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(272)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case CONFIGParserNEWLINE:
				{
					p.SetState(270)
					p.Match(CONFIGParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case CONFIGParserROUTE, CONFIGParserSTICKY_ADDRESS, CONFIGParserSESSION, CONFIGParserFORWARD, CONFIGParserLISTEN, CONFIGParserDISABLE:
				{
					p.SetState(271)
					p.Rstmt()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4453302272) != 0 {
		{
			p.SetState(277)
			p.Rstmt()
		}

	}

	{
		p.SetState(280)
		p.Match(CONFIGParserRCURLY)
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

// IRnameContext is an interface to support dynamic dispatch.
type IRnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QNAME() antlr.TerminalNode

	// IsRnameContext differentiates from other interfaces.
	IsRnameContext()
}

type RnameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRnameContext() *RnameContext {
	var p = new(RnameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_rname
	return p
}

func InitEmptyRnameContext(p *RnameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_rname
}

func (*RnameContext) IsRnameContext() {}

func NewRnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RnameContext {
	var p = new(RnameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_rname

	return p
}

func (s *RnameContext) GetParser() antlr.Parser { return s.parser }

func (s *RnameContext) QNAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserQNAME, 0)
}

func (s *RnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterRname(s)
	}
}

func (s *RnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitRname(s)
	}
}

func (p *CONFIGParser) Rname() (localctx IRnameContext) {
	localctx = NewRnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, CONFIGParserRULE_rname)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(CONFIGParserQNAME)
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

// IRstmtContext is an interface to support dynamic dispatch.
type IRstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DISABLE() antlr.TerminalNode
	STICKY_ADDRESS() antlr.TerminalNode
	Sesstimeout() ISesstimeoutContext
	Listen() IListenContext
	Forward() IForwardContext
	Route() IRouteContext

	// IsRstmtContext differentiates from other interfaces.
	IsRstmtContext()
}

type RstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRstmtContext() *RstmtContext {
	var p = new(RstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_rstmt
	return p
}

func InitEmptyRstmtContext(p *RstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_rstmt
}

func (*RstmtContext) IsRstmtContext() {}

func NewRstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RstmtContext {
	var p = new(RstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_rstmt

	return p
}

func (s *RstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RstmtContext) DISABLE() antlr.TerminalNode {
	return s.GetToken(CONFIGParserDISABLE, 0)
}

func (s *RstmtContext) STICKY_ADDRESS() antlr.TerminalNode {
	return s.GetToken(CONFIGParserSTICKY_ADDRESS, 0)
}

func (s *RstmtContext) Sesstimeout() ISesstimeoutContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISesstimeoutContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISesstimeoutContext)
}

func (s *RstmtContext) Listen() IListenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListenContext)
}

func (s *RstmtContext) Forward() IForwardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForwardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForwardContext)
}

func (s *RstmtContext) Route() IRouteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRouteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRouteContext)
}

func (s *RstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterRstmt(s)
	}
}

func (s *RstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitRstmt(s)
	}
}

func (p *CONFIGParser) Rstmt() (localctx IRstmtContext) {
	localctx = NewRstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CONFIGParserRULE_rstmt)
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserDISABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.Match(CONFIGParserDISABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CONFIGParserSTICKY_ADDRESS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.Match(CONFIGParserSTICKY_ADDRESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CONFIGParserSESSION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(286)
			p.Sesstimeout()
		}

	case CONFIGParserLISTEN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(287)
			p.Listen()
		}

	case CONFIGParserFORWARD:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(288)
			p.Forward()
		}

	case CONFIGParserROUTE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(289)
			p.Route()
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
