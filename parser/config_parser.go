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
		"", "'$'", "'='", "','", "'<'", "'>'", "'{'", "'}'", "'hash'", "'mode'",
		"'demote'", "'proto'", "'check'", "'http'", "'https'", "'code'", "'digest'",
		"'prefork'", "'priority'", "'retry'", "'parent'", "'ip'", "'ttl'", "'route'",
		"'source-hash'", "'roundrobin'", "'random'", "'loadbalance'", "'least-states'",
		"'sticky-address'", "'session'", "'timeout'", "'forward'", "'socket'",
		"'port'", "'interface'", "'listen'", "'to'", "'on'", "'host'", "'redirect'",
		"'disable'", "'interval'", "'table'",
	}
	staticData.SymbolicNames = []string{
		"", "DOLLA", "EQUALS", "COMMA", "LANGLE", "RANGLE", "LCURLY", "RCURLY",
		"HASH", "MODE", "DEMOTE", "PROTO", "CHECK", "HTTP", "HTTPS", "CODE",
		"DIGEST", "PREFORK", "PRIORITY", "RETRY", "PARENT", "IP", "TTL", "ROUTE",
		"SOURCE_HASH", "ROUND_ROBIN", "RANDOM", "LOAD_BALANCE", "LEAST_STATES",
		"STICKY_ADDRESS", "SESSION", "TIMEOUT", "FORWARD", "SOCKET", "PORT",
		"INTERFACE", "LISTEN", "TO", "ON", "HOST", "REDIRECT", "DISABLE", "INTERVAL",
		"TABLE", "NUMBER", "NAME", "QNAME", "COMMENT", "WS", "CONTINUATION",
		"NEWLINE",
	}
	staticData.RuleNames = []string{
		"config", "stmt", "assignment", "var", "gopts", "timeout", "interval",
		"prefork", "socket", "table", "tname", "tlist", "titem", "tdisable",
		"topts", "toptipttl", "toptparent", "toptpriority", "toptretry", "listen",
		"port", "portrange", "lport", "proto", "interface", "check", "ctype",
		"chost", "cdigest", "chttpcode", "chttpdigest", "chttpscode", "chttpsdigest",
		"balancemode", "mtype", "mhash", "mleaststates", "mloadbalance", "mrandom",
		"mroundrobin", "msourcehash", "demote", "dopts", "destination", "forward",
		"route", "sesstimeout", "redirect", "rname", "rstmt",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 363, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 105, 8,
		0, 10, 0, 12, 0, 108, 9, 0, 1, 0, 3, 0, 111, 8, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 119, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 132, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 149,
		8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 5, 11, 162, 8, 11, 10, 11, 12, 11, 165, 9, 11, 1, 11, 3, 11, 168,
		8, 11, 3, 11, 170, 8, 11, 1, 12, 1, 12, 5, 12, 174, 8, 12, 10, 12, 12,
		12, 177, 9, 12, 1, 12, 1, 12, 5, 12, 181, 8, 12, 10, 12, 12, 12, 184, 9,
		12, 3, 12, 186, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14,
		194, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 213,
		8, 19, 1, 19, 1, 19, 3, 19, 217, 8, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 3, 22, 227, 8, 22, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 242,
		8, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 3,
		29, 253, 8, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 3, 30, 261, 8,
		30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 3, 31, 268, 8, 31, 1, 31, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 32, 3, 32, 276, 8, 32, 1, 32, 1, 32, 1, 33, 1,
		33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 289, 8, 34,
		1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1,
		39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 3, 42, 314, 8, 42, 1, 43, 1, 43, 3, 43, 318, 8, 43, 1, 43,
		5, 43, 321, 8, 43, 10, 43, 12, 43, 324, 9, 43, 1, 44, 1, 44, 1, 44, 1,
		44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 5, 47, 343, 8, 47, 10, 47, 12, 47, 346, 9, 47, 1,
		47, 3, 47, 349, 8, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 3, 49, 361, 8, 49, 1, 49, 0, 0, 50, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
		82, 84, 86, 88, 90, 92, 94, 96, 98, 0, 0, 360, 0, 106, 1, 0, 0, 0, 2, 118,
		1, 0, 0, 0, 4, 120, 1, 0, 0, 0, 6, 124, 1, 0, 0, 0, 8, 131, 1, 0, 0, 0,
		10, 133, 1, 0, 0, 0, 12, 136, 1, 0, 0, 0, 14, 139, 1, 0, 0, 0, 16, 142,
		1, 0, 0, 0, 18, 145, 1, 0, 0, 0, 20, 154, 1, 0, 0, 0, 22, 169, 1, 0, 0,
		0, 24, 185, 1, 0, 0, 0, 26, 187, 1, 0, 0, 0, 28, 193, 1, 0, 0, 0, 30, 195,
		1, 0, 0, 0, 32, 199, 1, 0, 0, 0, 34, 202, 1, 0, 0, 0, 36, 205, 1, 0, 0,
		0, 38, 208, 1, 0, 0, 0, 40, 218, 1, 0, 0, 0, 42, 221, 1, 0, 0, 0, 44, 226,
		1, 0, 0, 0, 46, 228, 1, 0, 0, 0, 48, 231, 1, 0, 0, 0, 50, 234, 1, 0, 0,
		0, 52, 241, 1, 0, 0, 0, 54, 243, 1, 0, 0, 0, 56, 246, 1, 0, 0, 0, 58, 249,
		1, 0, 0, 0, 60, 257, 1, 0, 0, 0, 62, 264, 1, 0, 0, 0, 64, 272, 1, 0, 0,
		0, 66, 279, 1, 0, 0, 0, 68, 288, 1, 0, 0, 0, 70, 290, 1, 0, 0, 0, 72, 293,
		1, 0, 0, 0, 74, 295, 1, 0, 0, 0, 76, 298, 1, 0, 0, 0, 78, 300, 1, 0, 0,
		0, 80, 302, 1, 0, 0, 0, 82, 305, 1, 0, 0, 0, 84, 313, 1, 0, 0, 0, 86, 315,
		1, 0, 0, 0, 88, 325, 1, 0, 0, 0, 90, 329, 1, 0, 0, 0, 92, 333, 1, 0, 0,
		0, 94, 337, 1, 0, 0, 0, 96, 352, 1, 0, 0, 0, 98, 360, 1, 0, 0, 0, 100,
		105, 5, 50, 0, 0, 101, 102, 3, 2, 1, 0, 102, 103, 5, 50, 0, 0, 103, 105,
		1, 0, 0, 0, 104, 100, 1, 0, 0, 0, 104, 101, 1, 0, 0, 0, 105, 108, 1, 0,
		0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0,
		108, 106, 1, 0, 0, 0, 109, 111, 3, 2, 1, 0, 110, 109, 1, 0, 0, 0, 110,
		111, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 5, 0, 0, 1, 113, 1, 1,
		0, 0, 0, 114, 119, 3, 4, 2, 0, 115, 119, 3, 18, 9, 0, 116, 119, 3, 8, 4,
		0, 117, 119, 3, 94, 47, 0, 118, 114, 1, 0, 0, 0, 118, 115, 1, 0, 0, 0,
		118, 116, 1, 0, 0, 0, 118, 117, 1, 0, 0, 0, 119, 3, 1, 0, 0, 0, 120, 121,
		5, 45, 0, 0, 121, 122, 5, 2, 0, 0, 122, 123, 5, 45, 0, 0, 123, 5, 1, 0,
		0, 0, 124, 125, 5, 1, 0, 0, 125, 126, 5, 45, 0, 0, 126, 7, 1, 0, 0, 0,
		127, 132, 3, 12, 6, 0, 128, 132, 3, 10, 5, 0, 129, 132, 3, 14, 7, 0, 130,
		132, 3, 16, 8, 0, 131, 127, 1, 0, 0, 0, 131, 128, 1, 0, 0, 0, 131, 129,
		1, 0, 0, 0, 131, 130, 1, 0, 0, 0, 132, 9, 1, 0, 0, 0, 133, 134, 5, 31,
		0, 0, 134, 135, 5, 44, 0, 0, 135, 11, 1, 0, 0, 0, 136, 137, 5, 42, 0, 0,
		137, 138, 5, 44, 0, 0, 138, 13, 1, 0, 0, 0, 139, 140, 5, 17, 0, 0, 140,
		141, 5, 44, 0, 0, 141, 15, 1, 0, 0, 0, 142, 143, 5, 33, 0, 0, 143, 144,
		5, 46, 0, 0, 144, 17, 1, 0, 0, 0, 145, 146, 5, 43, 0, 0, 146, 148, 3, 20,
		10, 0, 147, 149, 3, 26, 13, 0, 148, 147, 1, 0, 0, 0, 148, 149, 1, 0, 0,
		0, 149, 150, 1, 0, 0, 0, 150, 151, 5, 6, 0, 0, 151, 152, 3, 22, 11, 0,
		152, 153, 5, 7, 0, 0, 153, 19, 1, 0, 0, 0, 154, 155, 5, 4, 0, 0, 155, 156,
		5, 45, 0, 0, 156, 157, 5, 5, 0, 0, 157, 21, 1, 0, 0, 0, 158, 163, 3, 24,
		12, 0, 159, 160, 5, 3, 0, 0, 160, 162, 3, 24, 12, 0, 161, 159, 1, 0, 0,
		0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164,
		167, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 168, 5, 3, 0, 0, 167, 166,
		1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 170, 1, 0, 0, 0, 169, 158, 1, 0,
		0, 0, 169, 170, 1, 0, 0, 0, 170, 23, 1, 0, 0, 0, 171, 175, 3, 6, 3, 0,
		172, 174, 3, 28, 14, 0, 173, 172, 1, 0, 0, 0, 174, 177, 1, 0, 0, 0, 175,
		173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 186, 1, 0, 0, 0, 177, 175,
		1, 0, 0, 0, 178, 182, 5, 45, 0, 0, 179, 181, 3, 28, 14, 0, 180, 179, 1,
		0, 0, 0, 181, 184, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0,
		0, 183, 186, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 185, 171, 1, 0, 0, 0, 185,
		178, 1, 0, 0, 0, 186, 25, 1, 0, 0, 0, 187, 188, 5, 41, 0, 0, 188, 27, 1,
		0, 0, 0, 189, 194, 3, 30, 15, 0, 190, 194, 3, 32, 16, 0, 191, 194, 3, 34,
		17, 0, 192, 194, 3, 36, 18, 0, 193, 189, 1, 0, 0, 0, 193, 190, 1, 0, 0,
		0, 193, 191, 1, 0, 0, 0, 193, 192, 1, 0, 0, 0, 194, 29, 1, 0, 0, 0, 195,
		196, 5, 21, 0, 0, 196, 197, 5, 22, 0, 0, 197, 198, 5, 44, 0, 0, 198, 31,
		1, 0, 0, 0, 199, 200, 5, 20, 0, 0, 200, 201, 5, 44, 0, 0, 201, 33, 1, 0,
		0, 0, 202, 203, 5, 18, 0, 0, 203, 204, 5, 44, 0, 0, 204, 35, 1, 0, 0, 0,
		205, 206, 5, 19, 0, 0, 206, 207, 5, 44, 0, 0, 207, 37, 1, 0, 0, 0, 208,
		209, 5, 36, 0, 0, 209, 210, 5, 38, 0, 0, 210, 212, 5, 45, 0, 0, 211, 213,
		3, 46, 23, 0, 212, 211, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 1,
		0, 0, 0, 214, 216, 3, 44, 22, 0, 215, 217, 3, 48, 24, 0, 216, 215, 1, 0,
		0, 0, 216, 217, 1, 0, 0, 0, 217, 39, 1, 0, 0, 0, 218, 219, 5, 34, 0, 0,
		219, 220, 5, 44, 0, 0, 220, 41, 1, 0, 0, 0, 221, 222, 5, 34, 0, 0, 222,
		223, 5, 45, 0, 0, 223, 43, 1, 0, 0, 0, 224, 227, 3, 42, 21, 0, 225, 227,
		3, 40, 20, 0, 226, 224, 1, 0, 0, 0, 226, 225, 1, 0, 0, 0, 227, 45, 1, 0,
		0, 0, 228, 229, 5, 11, 0, 0, 229, 230, 5, 45, 0, 0, 230, 47, 1, 0, 0, 0,
		231, 232, 5, 35, 0, 0, 232, 233, 5, 45, 0, 0, 233, 49, 1, 0, 0, 0, 234,
		235, 5, 12, 0, 0, 235, 236, 3, 52, 26, 0, 236, 51, 1, 0, 0, 0, 237, 242,
		3, 58, 29, 0, 238, 242, 3, 60, 30, 0, 239, 242, 3, 62, 31, 0, 240, 242,
		3, 64, 32, 0, 241, 237, 1, 0, 0, 0, 241, 238, 1, 0, 0, 0, 241, 239, 1,
		0, 0, 0, 241, 240, 1, 0, 0, 0, 242, 53, 1, 0, 0, 0, 243, 244, 5, 39, 0,
		0, 244, 245, 5, 45, 0, 0, 245, 55, 1, 0, 0, 0, 246, 247, 5, 16, 0, 0, 247,
		248, 5, 46, 0, 0, 248, 57, 1, 0, 0, 0, 249, 250, 5, 13, 0, 0, 250, 252,
		5, 46, 0, 0, 251, 253, 3, 54, 27, 0, 252, 251, 1, 0, 0, 0, 252, 253, 1,
		0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 5, 15, 0, 0, 255, 256, 5, 44,
		0, 0, 256, 59, 1, 0, 0, 0, 257, 258, 5, 13, 0, 0, 258, 260, 5, 46, 0, 0,
		259, 261, 3, 54, 27, 0, 260, 259, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261,
		262, 1, 0, 0, 0, 262, 263, 3, 56, 28, 0, 263, 61, 1, 0, 0, 0, 264, 265,
		5, 14, 0, 0, 265, 267, 5, 46, 0, 0, 266, 268, 3, 54, 27, 0, 267, 266, 1,
		0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 270, 5, 15, 0,
		0, 270, 271, 5, 44, 0, 0, 271, 63, 1, 0, 0, 0, 272, 273, 5, 14, 0, 0, 273,
		275, 5, 46, 0, 0, 274, 276, 3, 54, 27, 0, 275, 274, 1, 0, 0, 0, 275, 276,
		1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 3, 56, 28, 0, 278, 65, 1, 0,
		0, 0, 279, 280, 5, 9, 0, 0, 280, 281, 3, 68, 34, 0, 281, 67, 1, 0, 0, 0,
		282, 289, 3, 70, 35, 0, 283, 289, 3, 72, 36, 0, 284, 289, 3, 74, 37, 0,
		285, 289, 3, 76, 38, 0, 286, 289, 3, 78, 39, 0, 287, 289, 3, 80, 40, 0,
		288, 282, 1, 0, 0, 0, 288, 283, 1, 0, 0, 0, 288, 284, 1, 0, 0, 0, 288,
		285, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 287, 1, 0, 0, 0, 289, 69, 1,
		0, 0, 0, 290, 291, 5, 8, 0, 0, 291, 292, 5, 46, 0, 0, 292, 71, 1, 0, 0,
		0, 293, 294, 5, 28, 0, 0, 294, 73, 1, 0, 0, 0, 295, 296, 5, 27, 0, 0, 296,
		297, 5, 46, 0, 0, 297, 75, 1, 0, 0, 0, 298, 299, 5, 26, 0, 0, 299, 77,
		1, 0, 0, 0, 300, 301, 5, 25, 0, 0, 301, 79, 1, 0, 0, 0, 302, 303, 5, 24,
		0, 0, 303, 304, 5, 46, 0, 0, 304, 81, 1, 0, 0, 0, 305, 306, 5, 10, 0, 0,
		306, 307, 5, 45, 0, 0, 307, 83, 1, 0, 0, 0, 308, 314, 3, 50, 25, 0, 309,
		314, 3, 66, 33, 0, 310, 314, 3, 12, 6, 0, 311, 314, 3, 10, 5, 0, 312, 314,
		3, 82, 41, 0, 313, 308, 1, 0, 0, 0, 313, 309, 1, 0, 0, 0, 313, 310, 1,
		0, 0, 0, 313, 311, 1, 0, 0, 0, 313, 312, 1, 0, 0, 0, 314, 85, 1, 0, 0,
		0, 315, 317, 3, 20, 10, 0, 316, 318, 3, 40, 20, 0, 317, 316, 1, 0, 0, 0,
		317, 318, 1, 0, 0, 0, 318, 322, 1, 0, 0, 0, 319, 321, 3, 84, 42, 0, 320,
		319, 1, 0, 0, 0, 321, 324, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 322, 323,
		1, 0, 0, 0, 323, 87, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 325, 326, 5, 32,
		0, 0, 326, 327, 5, 37, 0, 0, 327, 328, 3, 86, 43, 0, 328, 89, 1, 0, 0,
		0, 329, 330, 5, 23, 0, 0, 330, 331, 5, 37, 0, 0, 331, 332, 3, 86, 43, 0,
		332, 91, 1, 0, 0, 0, 333, 334, 5, 30, 0, 0, 334, 335, 5, 31, 0, 0, 335,
		336, 5, 44, 0, 0, 336, 93, 1, 0, 0, 0, 337, 338, 5, 40, 0, 0, 338, 339,
		3, 96, 48, 0, 339, 344, 5, 6, 0, 0, 340, 343, 5, 50, 0, 0, 341, 343, 3,
		98, 49, 0, 342, 340, 1, 0, 0, 0, 342, 341, 1, 0, 0, 0, 343, 346, 1, 0,
		0, 0, 344, 342, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 348, 1, 0, 0, 0,
		346, 344, 1, 0, 0, 0, 347, 349, 3, 98, 49, 0, 348, 347, 1, 0, 0, 0, 348,
		349, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 351, 5, 7, 0, 0, 351, 95, 1,
		0, 0, 0, 352, 353, 5, 46, 0, 0, 353, 97, 1, 0, 0, 0, 354, 361, 5, 41, 0,
		0, 355, 361, 5, 29, 0, 0, 356, 361, 3, 92, 46, 0, 357, 361, 3, 38, 19,
		0, 358, 361, 3, 88, 44, 0, 359, 361, 3, 90, 45, 0, 360, 354, 1, 0, 0, 0,
		360, 355, 1, 0, 0, 0, 360, 356, 1, 0, 0, 0, 360, 357, 1, 0, 0, 0, 360,
		358, 1, 0, 0, 0, 360, 359, 1, 0, 0, 0, 361, 99, 1, 0, 0, 0, 29, 104, 106,
		110, 118, 131, 148, 163, 167, 169, 175, 182, 185, 193, 212, 216, 226, 241,
		252, 260, 267, 275, 288, 313, 317, 322, 342, 344, 348, 360,
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
	CONFIGParserHASH           = 8
	CONFIGParserMODE           = 9
	CONFIGParserDEMOTE         = 10
	CONFIGParserPROTO          = 11
	CONFIGParserCHECK          = 12
	CONFIGParserHTTP           = 13
	CONFIGParserHTTPS          = 14
	CONFIGParserCODE           = 15
	CONFIGParserDIGEST         = 16
	CONFIGParserPREFORK        = 17
	CONFIGParserPRIORITY       = 18
	CONFIGParserRETRY          = 19
	CONFIGParserPARENT         = 20
	CONFIGParserIP             = 21
	CONFIGParserTTL            = 22
	CONFIGParserROUTE          = 23
	CONFIGParserSOURCE_HASH    = 24
	CONFIGParserROUND_ROBIN    = 25
	CONFIGParserRANDOM         = 26
	CONFIGParserLOAD_BALANCE   = 27
	CONFIGParserLEAST_STATES   = 28
	CONFIGParserSTICKY_ADDRESS = 29
	CONFIGParserSESSION        = 30
	CONFIGParserTIMEOUT        = 31
	CONFIGParserFORWARD        = 32
	CONFIGParserSOCKET         = 33
	CONFIGParserPORT           = 34
	CONFIGParserINTERFACE      = 35
	CONFIGParserLISTEN         = 36
	CONFIGParserTO             = 37
	CONFIGParserON             = 38
	CONFIGParserHOST           = 39
	CONFIGParserREDIRECT       = 40
	CONFIGParserDISABLE        = 41
	CONFIGParserINTERVAL       = 42
	CONFIGParserTABLE          = 43
	CONFIGParserNUMBER         = 44
	CONFIGParserNAME           = 45
	CONFIGParserQNAME          = 46
	CONFIGParserCOMMENT        = 47
	CONFIGParserWS             = 48
	CONFIGParserCONTINUATION   = 49
	CONFIGParserNEWLINE        = 50
)

// CONFIGParser rules.
const (
	CONFIGParserRULE_config       = 0
	CONFIGParserRULE_stmt         = 1
	CONFIGParserRULE_assignment   = 2
	CONFIGParserRULE_var          = 3
	CONFIGParserRULE_gopts        = 4
	CONFIGParserRULE_timeout      = 5
	CONFIGParserRULE_interval     = 6
	CONFIGParserRULE_prefork      = 7
	CONFIGParserRULE_socket       = 8
	CONFIGParserRULE_table        = 9
	CONFIGParserRULE_tname        = 10
	CONFIGParserRULE_tlist        = 11
	CONFIGParserRULE_titem        = 12
	CONFIGParserRULE_tdisable     = 13
	CONFIGParserRULE_topts        = 14
	CONFIGParserRULE_toptipttl    = 15
	CONFIGParserRULE_toptparent   = 16
	CONFIGParserRULE_toptpriority = 17
	CONFIGParserRULE_toptretry    = 18
	CONFIGParserRULE_listen       = 19
	CONFIGParserRULE_port         = 20
	CONFIGParserRULE_portrange    = 21
	CONFIGParserRULE_lport        = 22
	CONFIGParserRULE_proto        = 23
	CONFIGParserRULE_interface    = 24
	CONFIGParserRULE_check        = 25
	CONFIGParserRULE_ctype        = 26
	CONFIGParserRULE_chost        = 27
	CONFIGParserRULE_cdigest      = 28
	CONFIGParserRULE_chttpcode    = 29
	CONFIGParserRULE_chttpdigest  = 30
	CONFIGParserRULE_chttpscode   = 31
	CONFIGParserRULE_chttpsdigest = 32
	CONFIGParserRULE_balancemode  = 33
	CONFIGParserRULE_mtype        = 34
	CONFIGParserRULE_mhash        = 35
	CONFIGParserRULE_mleaststates = 36
	CONFIGParserRULE_mloadbalance = 37
	CONFIGParserRULE_mrandom      = 38
	CONFIGParserRULE_mroundrobin  = 39
	CONFIGParserRULE_msourcehash  = 40
	CONFIGParserRULE_demote       = 41
	CONFIGParserRULE_dopts        = 42
	CONFIGParserRULE_destination  = 43
	CONFIGParserRULE_forward      = 44
	CONFIGParserRULE_route        = 45
	CONFIGParserRULE_sesstimeout  = 46
	CONFIGParserRULE_redirect     = 47
	CONFIGParserRULE_rname        = 48
	CONFIGParserRULE_rstmt        = 49
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
	p.SetState(106)
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
			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case CONFIGParserNEWLINE:
				{
					p.SetState(100)
					p.Match(CONFIGParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case CONFIGParserPREFORK, CONFIGParserTIMEOUT, CONFIGParserSOCKET, CONFIGParserREDIRECT, CONFIGParserINTERVAL, CONFIGParserTABLE, CONFIGParserNAME:
				{
					p.SetState(101)
					p.Stmt()
				}
				{
					p.SetState(102)
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
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&49488760799232) != 0 {
		{
			p.SetState(109)
			p.Stmt()
		}

	}
	{
		p.SetState(112)
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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Assignment()
		}

	case CONFIGParserTABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Table()
		}

	case CONFIGParserPREFORK, CONFIGParserTIMEOUT, CONFIGParserSOCKET, CONFIGParserINTERVAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(116)
			p.Gopts()
		}

	case CONFIGParserREDIRECT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(117)
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
		p.SetState(120)
		p.Match(CONFIGParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(CONFIGParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
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
		p.SetState(124)
		p.Match(CONFIGParserDOLLA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
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
	Timeout() ITimeoutContext
	Prefork() IPreforkContext
	Socket() ISocketContext

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

func (s *GoptsContext) Timeout() ITimeoutContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeoutContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeoutContext)
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

func (s *GoptsContext) Socket() ISocketContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISocketContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISocketContext)
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
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserINTERVAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Interval()
		}

	case CONFIGParserTIMEOUT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Timeout()
		}

	case CONFIGParserPREFORK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.Prefork()
		}

	case CONFIGParserSOCKET:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)
			p.Socket()
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

// ITimeoutContext is an interface to support dynamic dispatch.
type ITimeoutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TIMEOUT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsTimeoutContext differentiates from other interfaces.
	IsTimeoutContext()
}

type TimeoutContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeoutContext() *TimeoutContext {
	var p = new(TimeoutContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_timeout
	return p
}

func InitEmptyTimeoutContext(p *TimeoutContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_timeout
}

func (*TimeoutContext) IsTimeoutContext() {}

func NewTimeoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeoutContext {
	var p = new(TimeoutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_timeout

	return p
}

func (s *TimeoutContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeoutContext) TIMEOUT() antlr.TerminalNode {
	return s.GetToken(CONFIGParserTIMEOUT, 0)
}

func (s *TimeoutContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNUMBER, 0)
}

func (s *TimeoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeoutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeoutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterTimeout(s)
	}
}

func (s *TimeoutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitTimeout(s)
	}
}

func (p *CONFIGParser) Timeout() (localctx ITimeoutContext) {
	localctx = NewTimeoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CONFIGParserRULE_timeout)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(CONFIGParserTIMEOUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
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
	p.EnterRule(localctx, 12, CONFIGParserRULE_interval)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(CONFIGParserINTERVAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
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
	p.EnterRule(localctx, 14, CONFIGParserRULE_prefork)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(CONFIGParserPREFORK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
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

// ISocketContext is an interface to support dynamic dispatch.
type ISocketContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SOCKET() antlr.TerminalNode
	QNAME() antlr.TerminalNode

	// IsSocketContext differentiates from other interfaces.
	IsSocketContext()
}

type SocketContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySocketContext() *SocketContext {
	var p = new(SocketContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_socket
	return p
}

func InitEmptySocketContext(p *SocketContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_socket
}

func (*SocketContext) IsSocketContext() {}

func NewSocketContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SocketContext {
	var p = new(SocketContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_socket

	return p
}

func (s *SocketContext) GetParser() antlr.Parser { return s.parser }

func (s *SocketContext) SOCKET() antlr.TerminalNode {
	return s.GetToken(CONFIGParserSOCKET, 0)
}

func (s *SocketContext) QNAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserQNAME, 0)
}

func (s *SocketContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SocketContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SocketContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterSocket(s)
	}
}

func (s *SocketContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitSocket(s)
	}
}

func (p *CONFIGParser) Socket() (localctx ISocketContext) {
	localctx = NewSocketContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CONFIGParserRULE_socket)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(CONFIGParserSOCKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
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
	p.EnterRule(localctx, 18, CONFIGParserRULE_table)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(CONFIGParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Tname()
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserDISABLE {
		{
			p.SetState(147)
			p.Tdisable()
		}

	}
	{
		p.SetState(150)
		p.Match(CONFIGParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Tlist()
	}
	{
		p.SetState(152)
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
	p.EnterRule(localctx, 20, CONFIGParserRULE_tname)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(CONFIGParserLANGLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(CONFIGParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
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
	p.EnterRule(localctx, 22, CONFIGParserRULE_tlist)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserDOLLA || _la == CONFIGParserNAME {
		{
			p.SetState(158)
			p.Titem()
		}
		p.SetState(163)
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
					p.SetState(159)
					p.Match(CONFIGParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(160)
					p.Titem()
				}

			}
			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CONFIGParserCOMMA {
			{
				p.SetState(166)
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
	p.EnterRule(localctx, 24, CONFIGParserRULE_titem)
	var _la int

	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserDOLLA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Var_()
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3932160) != 0 {
			{
				p.SetState(172)
				p.Topts()
			}

			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case CONFIGParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
			p.Match(CONFIGParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3932160) != 0 {
			{
				p.SetState(179)
				p.Topts()
			}

			p.SetState(184)
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
	p.EnterRule(localctx, 26, CONFIGParserRULE_tdisable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
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
	p.EnterRule(localctx, 28, CONFIGParserRULE_topts)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserIP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Toptipttl()
		}

	case CONFIGParserPARENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Toptparent()
		}

	case CONFIGParserPRIORITY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.Toptpriority()
		}

	case CONFIGParserRETRY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(192)
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
	p.EnterRule(localctx, 30, CONFIGParserRULE_toptipttl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(CONFIGParserIP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Match(CONFIGParserTTL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
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
	p.EnterRule(localctx, 32, CONFIGParserRULE_toptparent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(CONFIGParserPARENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
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
	p.EnterRule(localctx, 34, CONFIGParserRULE_toptpriority)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(CONFIGParserPRIORITY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
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
	p.EnterRule(localctx, 36, CONFIGParserRULE_toptretry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(CONFIGParserRETRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
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
	p.EnterRule(localctx, 38, CONFIGParserRULE_listen)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(CONFIGParserLISTEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(CONFIGParserON)
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
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserPROTO {
		{
			p.SetState(211)
			p.Proto()
		}

	}
	{
		p.SetState(214)
		p.Lport()
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserINTERFACE {
		{
			p.SetState(215)
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
	p.EnterRule(localctx, 40, CONFIGParserRULE_port)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(CONFIGParserPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
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
	p.EnterRule(localctx, 42, CONFIGParserRULE_portrange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(CONFIGParserPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
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
	p.EnterRule(localctx, 44, CONFIGParserRULE_lport)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(224)
			p.Portrange()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(225)
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
	p.EnterRule(localctx, 46, CONFIGParserRULE_proto)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(CONFIGParserPROTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
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
	p.EnterRule(localctx, 48, CONFIGParserRULE_interface)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(CONFIGParserINTERFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
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
	p.EnterRule(localctx, 50, CONFIGParserRULE_check)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(CONFIGParserCHECK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
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
	p.EnterRule(localctx, 52, CONFIGParserRULE_ctype)
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(237)
			p.Chttpcode()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.Chttpdigest()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(239)
			p.Chttpscode()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(240)
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
	p.EnterRule(localctx, 54, CONFIGParserRULE_chost)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(CONFIGParserHOST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
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
	QNAME() antlr.TerminalNode

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

func (s *CdigestContext) QNAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserQNAME, 0)
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
	p.EnterRule(localctx, 56, CONFIGParserRULE_cdigest)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(CONFIGParserDIGEST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
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
	p.EnterRule(localctx, 58, CONFIGParserRULE_chttpcode)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(CONFIGParserHTTP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Match(CONFIGParserQNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserHOST {
		{
			p.SetState(251)
			p.Chost()
		}

	}
	{
		p.SetState(254)
		p.Match(CONFIGParserCODE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
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
	p.EnterRule(localctx, 60, CONFIGParserRULE_chttpdigest)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(CONFIGParserHTTP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(CONFIGParserQNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserHOST {
		{
			p.SetState(259)
			p.Chost()
		}

	}
	{
		p.SetState(262)
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
	p.EnterRule(localctx, 62, CONFIGParserRULE_chttpscode)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(CONFIGParserHTTPS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
		p.Match(CONFIGParserQNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserHOST {
		{
			p.SetState(266)
			p.Chost()
		}

	}
	{
		p.SetState(269)
		p.Match(CONFIGParserCODE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
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
	p.EnterRule(localctx, 64, CONFIGParserRULE_chttpsdigest)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(CONFIGParserHTTPS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Match(CONFIGParserQNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserHOST {
		{
			p.SetState(274)
			p.Chost()
		}

	}
	{
		p.SetState(277)
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

// IBalancemodeContext is an interface to support dynamic dispatch.
type IBalancemodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MODE() antlr.TerminalNode
	Mtype() IMtypeContext

	// IsBalancemodeContext differentiates from other interfaces.
	IsBalancemodeContext()
}

type BalancemodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBalancemodeContext() *BalancemodeContext {
	var p = new(BalancemodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_balancemode
	return p
}

func InitEmptyBalancemodeContext(p *BalancemodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_balancemode
}

func (*BalancemodeContext) IsBalancemodeContext() {}

func NewBalancemodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BalancemodeContext {
	var p = new(BalancemodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_balancemode

	return p
}

func (s *BalancemodeContext) GetParser() antlr.Parser { return s.parser }

func (s *BalancemodeContext) MODE() antlr.TerminalNode {
	return s.GetToken(CONFIGParserMODE, 0)
}

func (s *BalancemodeContext) Mtype() IMtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMtypeContext)
}

func (s *BalancemodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BalancemodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BalancemodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterBalancemode(s)
	}
}

func (s *BalancemodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitBalancemode(s)
	}
}

func (p *CONFIGParser) Balancemode() (localctx IBalancemodeContext) {
	localctx = NewBalancemodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CONFIGParserRULE_balancemode)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(CONFIGParserMODE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.Mtype()
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

// IMtypeContext is an interface to support dynamic dispatch.
type IMtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mhash() IMhashContext
	Mleaststates() IMleaststatesContext
	Mloadbalance() IMloadbalanceContext
	Mrandom() IMrandomContext
	Mroundrobin() IMroundrobinContext
	Msourcehash() IMsourcehashContext

	// IsMtypeContext differentiates from other interfaces.
	IsMtypeContext()
}

type MtypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMtypeContext() *MtypeContext {
	var p = new(MtypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_mtype
	return p
}

func InitEmptyMtypeContext(p *MtypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_mtype
}

func (*MtypeContext) IsMtypeContext() {}

func NewMtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MtypeContext {
	var p = new(MtypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_mtype

	return p
}

func (s *MtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MtypeContext) Mhash() IMhashContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMhashContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMhashContext)
}

func (s *MtypeContext) Mleaststates() IMleaststatesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMleaststatesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMleaststatesContext)
}

func (s *MtypeContext) Mloadbalance() IMloadbalanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMloadbalanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMloadbalanceContext)
}

func (s *MtypeContext) Mrandom() IMrandomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMrandomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMrandomContext)
}

func (s *MtypeContext) Mroundrobin() IMroundrobinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMroundrobinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMroundrobinContext)
}

func (s *MtypeContext) Msourcehash() IMsourcehashContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMsourcehashContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMsourcehashContext)
}

func (s *MtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterMtype(s)
	}
}

func (s *MtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitMtype(s)
	}
}

func (p *CONFIGParser) Mtype() (localctx IMtypeContext) {
	localctx = NewMtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CONFIGParserRULE_mtype)
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserHASH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(282)
			p.Mhash()
		}

	case CONFIGParserLEAST_STATES:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)
			p.Mleaststates()
		}

	case CONFIGParserLOAD_BALANCE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(284)
			p.Mloadbalance()
		}

	case CONFIGParserRANDOM:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(285)
			p.Mrandom()
		}

	case CONFIGParserROUND_ROBIN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(286)
			p.Mroundrobin()
		}

	case CONFIGParserSOURCE_HASH:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(287)
			p.Msourcehash()
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

// IMhashContext is an interface to support dynamic dispatch.
type IMhashContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HASH() antlr.TerminalNode
	QNAME() antlr.TerminalNode

	// IsMhashContext differentiates from other interfaces.
	IsMhashContext()
}

type MhashContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMhashContext() *MhashContext {
	var p = new(MhashContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_mhash
	return p
}

func InitEmptyMhashContext(p *MhashContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_mhash
}

func (*MhashContext) IsMhashContext() {}

func NewMhashContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MhashContext {
	var p = new(MhashContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_mhash

	return p
}

func (s *MhashContext) GetParser() antlr.Parser { return s.parser }

func (s *MhashContext) HASH() antlr.TerminalNode {
	return s.GetToken(CONFIGParserHASH, 0)
}

func (s *MhashContext) QNAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserQNAME, 0)
}

func (s *MhashContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MhashContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MhashContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterMhash(s)
	}
}

func (s *MhashContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitMhash(s)
	}
}

func (p *CONFIGParser) Mhash() (localctx IMhashContext) {
	localctx = NewMhashContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, CONFIGParserRULE_mhash)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(CONFIGParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
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

// IMleaststatesContext is an interface to support dynamic dispatch.
type IMleaststatesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEAST_STATES() antlr.TerminalNode

	// IsMleaststatesContext differentiates from other interfaces.
	IsMleaststatesContext()
}

type MleaststatesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMleaststatesContext() *MleaststatesContext {
	var p = new(MleaststatesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_mleaststates
	return p
}

func InitEmptyMleaststatesContext(p *MleaststatesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_mleaststates
}

func (*MleaststatesContext) IsMleaststatesContext() {}

func NewMleaststatesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MleaststatesContext {
	var p = new(MleaststatesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_mleaststates

	return p
}

func (s *MleaststatesContext) GetParser() antlr.Parser { return s.parser }

func (s *MleaststatesContext) LEAST_STATES() antlr.TerminalNode {
	return s.GetToken(CONFIGParserLEAST_STATES, 0)
}

func (s *MleaststatesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MleaststatesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MleaststatesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterMleaststates(s)
	}
}

func (s *MleaststatesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitMleaststates(s)
	}
}

func (p *CONFIGParser) Mleaststates() (localctx IMleaststatesContext) {
	localctx = NewMleaststatesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CONFIGParserRULE_mleaststates)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(CONFIGParserLEAST_STATES)
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

// IMloadbalanceContext is an interface to support dynamic dispatch.
type IMloadbalanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOAD_BALANCE() antlr.TerminalNode
	QNAME() antlr.TerminalNode

	// IsMloadbalanceContext differentiates from other interfaces.
	IsMloadbalanceContext()
}

type MloadbalanceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMloadbalanceContext() *MloadbalanceContext {
	var p = new(MloadbalanceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_mloadbalance
	return p
}

func InitEmptyMloadbalanceContext(p *MloadbalanceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_mloadbalance
}

func (*MloadbalanceContext) IsMloadbalanceContext() {}

func NewMloadbalanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MloadbalanceContext {
	var p = new(MloadbalanceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_mloadbalance

	return p
}

func (s *MloadbalanceContext) GetParser() antlr.Parser { return s.parser }

func (s *MloadbalanceContext) LOAD_BALANCE() antlr.TerminalNode {
	return s.GetToken(CONFIGParserLOAD_BALANCE, 0)
}

func (s *MloadbalanceContext) QNAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserQNAME, 0)
}

func (s *MloadbalanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MloadbalanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MloadbalanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterMloadbalance(s)
	}
}

func (s *MloadbalanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitMloadbalance(s)
	}
}

func (p *CONFIGParser) Mloadbalance() (localctx IMloadbalanceContext) {
	localctx = NewMloadbalanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, CONFIGParserRULE_mloadbalance)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Match(CONFIGParserLOAD_BALANCE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
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

// IMrandomContext is an interface to support dynamic dispatch.
type IMrandomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RANDOM() antlr.TerminalNode

	// IsMrandomContext differentiates from other interfaces.
	IsMrandomContext()
}

type MrandomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMrandomContext() *MrandomContext {
	var p = new(MrandomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_mrandom
	return p
}

func InitEmptyMrandomContext(p *MrandomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_mrandom
}

func (*MrandomContext) IsMrandomContext() {}

func NewMrandomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MrandomContext {
	var p = new(MrandomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_mrandom

	return p
}

func (s *MrandomContext) GetParser() antlr.Parser { return s.parser }

func (s *MrandomContext) RANDOM() antlr.TerminalNode {
	return s.GetToken(CONFIGParserRANDOM, 0)
}

func (s *MrandomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MrandomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MrandomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterMrandom(s)
	}
}

func (s *MrandomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitMrandom(s)
	}
}

func (p *CONFIGParser) Mrandom() (localctx IMrandomContext) {
	localctx = NewMrandomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, CONFIGParserRULE_mrandom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(CONFIGParserRANDOM)
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

// IMroundrobinContext is an interface to support dynamic dispatch.
type IMroundrobinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ROUND_ROBIN() antlr.TerminalNode

	// IsMroundrobinContext differentiates from other interfaces.
	IsMroundrobinContext()
}

type MroundrobinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMroundrobinContext() *MroundrobinContext {
	var p = new(MroundrobinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_mroundrobin
	return p
}

func InitEmptyMroundrobinContext(p *MroundrobinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_mroundrobin
}

func (*MroundrobinContext) IsMroundrobinContext() {}

func NewMroundrobinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MroundrobinContext {
	var p = new(MroundrobinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_mroundrobin

	return p
}

func (s *MroundrobinContext) GetParser() antlr.Parser { return s.parser }

func (s *MroundrobinContext) ROUND_ROBIN() antlr.TerminalNode {
	return s.GetToken(CONFIGParserROUND_ROBIN, 0)
}

func (s *MroundrobinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MroundrobinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MroundrobinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterMroundrobin(s)
	}
}

func (s *MroundrobinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitMroundrobin(s)
	}
}

func (p *CONFIGParser) Mroundrobin() (localctx IMroundrobinContext) {
	localctx = NewMroundrobinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, CONFIGParserRULE_mroundrobin)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(CONFIGParserROUND_ROBIN)
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

// IMsourcehashContext is an interface to support dynamic dispatch.
type IMsourcehashContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SOURCE_HASH() antlr.TerminalNode
	QNAME() antlr.TerminalNode

	// IsMsourcehashContext differentiates from other interfaces.
	IsMsourcehashContext()
}

type MsourcehashContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsourcehashContext() *MsourcehashContext {
	var p = new(MsourcehashContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_msourcehash
	return p
}

func InitEmptyMsourcehashContext(p *MsourcehashContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_msourcehash
}

func (*MsourcehashContext) IsMsourcehashContext() {}

func NewMsourcehashContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsourcehashContext {
	var p = new(MsourcehashContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_msourcehash

	return p
}

func (s *MsourcehashContext) GetParser() antlr.Parser { return s.parser }

func (s *MsourcehashContext) SOURCE_HASH() antlr.TerminalNode {
	return s.GetToken(CONFIGParserSOURCE_HASH, 0)
}

func (s *MsourcehashContext) QNAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserQNAME, 0)
}

func (s *MsourcehashContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsourcehashContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsourcehashContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterMsourcehash(s)
	}
}

func (s *MsourcehashContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitMsourcehash(s)
	}
}

func (p *CONFIGParser) Msourcehash() (localctx IMsourcehashContext) {
	localctx = NewMsourcehashContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, CONFIGParserRULE_msourcehash)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Match(CONFIGParserSOURCE_HASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
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

// IDemoteContext is an interface to support dynamic dispatch.
type IDemoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEMOTE() antlr.TerminalNode
	NAME() antlr.TerminalNode

	// IsDemoteContext differentiates from other interfaces.
	IsDemoteContext()
}

type DemoteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDemoteContext() *DemoteContext {
	var p = new(DemoteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_demote
	return p
}

func InitEmptyDemoteContext(p *DemoteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_demote
}

func (*DemoteContext) IsDemoteContext() {}

func NewDemoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DemoteContext {
	var p = new(DemoteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_demote

	return p
}

func (s *DemoteContext) GetParser() antlr.Parser { return s.parser }

func (s *DemoteContext) DEMOTE() antlr.TerminalNode {
	return s.GetToken(CONFIGParserDEMOTE, 0)
}

func (s *DemoteContext) NAME() antlr.TerminalNode {
	return s.GetToken(CONFIGParserNAME, 0)
}

func (s *DemoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DemoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DemoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterDemote(s)
	}
}

func (s *DemoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitDemote(s)
	}
}

func (p *CONFIGParser) Demote() (localctx IDemoteContext) {
	localctx = NewDemoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, CONFIGParserRULE_demote)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(CONFIGParserDEMOTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
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

// IDoptsContext is an interface to support dynamic dispatch.
type IDoptsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Check() ICheckContext
	Balancemode() IBalancemodeContext
	Interval() IIntervalContext
	Timeout() ITimeoutContext
	Demote() IDemoteContext

	// IsDoptsContext differentiates from other interfaces.
	IsDoptsContext()
}

type DoptsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoptsContext() *DoptsContext {
	var p = new(DoptsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_dopts
	return p
}

func InitEmptyDoptsContext(p *DoptsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_dopts
}

func (*DoptsContext) IsDoptsContext() {}

func NewDoptsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoptsContext {
	var p = new(DoptsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_dopts

	return p
}

func (s *DoptsContext) GetParser() antlr.Parser { return s.parser }

func (s *DoptsContext) Check() ICheckContext {
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

func (s *DoptsContext) Balancemode() IBalancemodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBalancemodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBalancemodeContext)
}

func (s *DoptsContext) Interval() IIntervalContext {
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

func (s *DoptsContext) Timeout() ITimeoutContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeoutContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeoutContext)
}

func (s *DoptsContext) Demote() IDemoteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDemoteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDemoteContext)
}

func (s *DoptsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoptsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoptsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterDopts(s)
	}
}

func (s *DoptsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitDopts(s)
	}
}

func (p *CONFIGParser) Dopts() (localctx IDoptsContext) {
	localctx = NewDoptsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, CONFIGParserRULE_dopts)
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserCHECK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(308)
			p.Check()
		}

	case CONFIGParserMODE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(309)
			p.Balancemode()
		}

	case CONFIGParserINTERVAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(310)
			p.Interval()
		}

	case CONFIGParserTIMEOUT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(311)
			p.Timeout()
		}

	case CONFIGParserDEMOTE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(312)
			p.Demote()
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

// IDestinationContext is an interface to support dynamic dispatch.
type IDestinationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tname() ITnameContext
	Port() IPortContext
	AllDopts() []IDoptsContext
	Dopts(i int) IDoptsContext

	// IsDestinationContext differentiates from other interfaces.
	IsDestinationContext()
}

type DestinationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDestinationContext() *DestinationContext {
	var p = new(DestinationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_destination
	return p
}

func InitEmptyDestinationContext(p *DestinationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CONFIGParserRULE_destination
}

func (*DestinationContext) IsDestinationContext() {}

func NewDestinationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestinationContext {
	var p = new(DestinationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CONFIGParserRULE_destination

	return p
}

func (s *DestinationContext) GetParser() antlr.Parser { return s.parser }

func (s *DestinationContext) Tname() ITnameContext {
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

func (s *DestinationContext) Port() IPortContext {
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

func (s *DestinationContext) AllDopts() []IDoptsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoptsContext); ok {
			len++
		}
	}

	tst := make([]IDoptsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoptsContext); ok {
			tst[i] = t.(IDoptsContext)
			i++
		}
	}

	return tst
}

func (s *DestinationContext) Dopts(i int) IDoptsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoptsContext); ok {
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

	return t.(IDoptsContext)
}

func (s *DestinationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestinationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.EnterDestination(s)
	}
}

func (s *DestinationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CONFIGListener); ok {
		listenerT.ExitDestination(s)
	}
}

func (p *CONFIGParser) Destination() (localctx IDestinationContext) {
	localctx = NewDestinationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, CONFIGParserRULE_destination)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Tname()
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CONFIGParserPORT {
		{
			p.SetState(316)
			p.Port()
		}

	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4400194000384) != 0 {
		{
			p.SetState(319)
			p.Dopts()
		}

		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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
	Destination() IDestinationContext

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

func (s *ForwardContext) Destination() IDestinationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDestinationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDestinationContext)
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
	p.EnterRule(localctx, 88, CONFIGParserRULE_forward)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Match(CONFIGParserFORWARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(326)
		p.Match(CONFIGParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(327)
		p.Destination()
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
	Destination() IDestinationContext

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

func (s *RouteContext) Destination() IDestinationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDestinationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDestinationContext)
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
	p.EnterRule(localctx, 90, CONFIGParserRULE_route)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Match(CONFIGParserROUTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(330)
		p.Match(CONFIGParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Destination()
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
	p.EnterRule(localctx, 92, CONFIGParserRULE_sesstimeout)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.Match(CONFIGParserSESSION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(334)
		p.Match(CONFIGParserTIMEOUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(335)
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
	p.EnterRule(localctx, 94, CONFIGParserRULE_redirect)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(CONFIGParserREDIRECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(338)
		p.Rname()
	}
	{
		p.SetState(339)
		p.Match(CONFIGParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.SetState(344)
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
			p.SetState(342)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case CONFIGParserNEWLINE:
				{
					p.SetState(340)
					p.Match(CONFIGParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case CONFIGParserROUTE, CONFIGParserSTICKY_ADDRESS, CONFIGParserSESSION, CONFIGParserFORWARD, CONFIGParserLISTEN, CONFIGParserDISABLE:
				{
					p.SetState(341)
					p.Rstmt()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2273656700928) != 0 {
		{
			p.SetState(347)
			p.Rstmt()
		}

	}

	{
		p.SetState(350)
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
	p.EnterRule(localctx, 96, CONFIGParserRULE_rname)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
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
	p.EnterRule(localctx, 98, CONFIGParserRULE_rstmt)
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CONFIGParserDISABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(354)
			p.Match(CONFIGParserDISABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CONFIGParserSTICKY_ADDRESS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(355)
			p.Match(CONFIGParserSTICKY_ADDRESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CONFIGParserSESSION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(356)
			p.Sesstimeout()
		}

	case CONFIGParserLISTEN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(357)
			p.Listen()
		}

	case CONFIGParserFORWARD:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(358)
			p.Forward()
		}

	case CONFIGParserROUTE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(359)
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
