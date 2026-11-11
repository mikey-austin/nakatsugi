// Code generated from CONFIG.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CONFIG
import "github.com/antlr4-go/antlr/v4"

// CONFIGListener is a complete listener for a parse tree produced by CONFIGParser.
type CONFIGListener interface {
	antlr.ParseTreeListener

	// EnterConfig is called when entering the config production.
	EnterConfig(c *ConfigContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// EnterVar is called when entering the var production.
	EnterVar(c *VarContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterTname is called when entering the tname production.
	EnterTname(c *TnameContext)

	// EnterTlist is called when entering the tlist production.
	EnterTlist(c *TlistContext)

	// EnterTitem is called when entering the titem production.
	EnterTitem(c *TitemContext)

	// EnterTdisable is called when entering the tdisable production.
	EnterTdisable(c *TdisableContext)

	// ExitConfig is called when exiting the config production.
	ExitConfig(c *ConfigContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)

	// ExitVar is called when exiting the var production.
	ExitVar(c *VarContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitTname is called when exiting the tname production.
	ExitTname(c *TnameContext)

	// ExitTlist is called when exiting the tlist production.
	ExitTlist(c *TlistContext)

	// ExitTitem is called when exiting the titem production.
	ExitTitem(c *TitemContext)

	// ExitTdisable is called when exiting the tdisable production.
	ExitTdisable(c *TdisableContext)
}
