// Code generated from CONFIG.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // CONFIG
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CONFIGListener is a complete listener for a parse tree produced by CONFIGParser.
type CONFIGListener interface {
	antlr.ParseTreeListener

	// EnterConfig is called when entering the config production.
	EnterConfig(c *ConfigContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// ExitConfig is called when exiting the config production.
	ExitConfig(c *ConfigContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)
}
