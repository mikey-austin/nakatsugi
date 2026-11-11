// Code generated from CONFIG.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CONFIG
import "github.com/antlr4-go/antlr/v4"

// BaseCONFIGListener is a complete listener for a parse tree produced by CONFIGParser.
type BaseCONFIGListener struct{}

var _ CONFIGListener = &BaseCONFIGListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCONFIGListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCONFIGListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCONFIGListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCONFIGListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterConfig is called when production config is entered.
func (s *BaseCONFIGListener) EnterConfig(ctx *ConfigContext) {}

// ExitConfig is called when production config is exited.
func (s *BaseCONFIGListener) ExitConfig(ctx *ConfigContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseCONFIGListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseCONFIGListener) ExitStmt(ctx *StmtContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseCONFIGListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseCONFIGListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseCONFIGListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseCONFIGListener) ExitInterval(ctx *IntervalContext) {}

// EnterVar is called when production var is entered.
func (s *BaseCONFIGListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production var is exited.
func (s *BaseCONFIGListener) ExitVar(ctx *VarContext) {}

// EnterTable is called when production table is entered.
func (s *BaseCONFIGListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseCONFIGListener) ExitTable(ctx *TableContext) {}

// EnterTname is called when production tname is entered.
func (s *BaseCONFIGListener) EnterTname(ctx *TnameContext) {}

// ExitTname is called when production tname is exited.
func (s *BaseCONFIGListener) ExitTname(ctx *TnameContext) {}

// EnterTlist is called when production tlist is entered.
func (s *BaseCONFIGListener) EnterTlist(ctx *TlistContext) {}

// ExitTlist is called when production tlist is exited.
func (s *BaseCONFIGListener) ExitTlist(ctx *TlistContext) {}

// EnterTitem is called when production titem is entered.
func (s *BaseCONFIGListener) EnterTitem(ctx *TitemContext) {}

// ExitTitem is called when production titem is exited.
func (s *BaseCONFIGListener) ExitTitem(ctx *TitemContext) {}

// EnterTdisable is called when production tdisable is entered.
func (s *BaseCONFIGListener) EnterTdisable(ctx *TdisableContext) {}

// ExitTdisable is called when production tdisable is exited.
func (s *BaseCONFIGListener) ExitTdisable(ctx *TdisableContext) {}
