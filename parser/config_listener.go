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

	// EnterVar is called when entering the var production.
	EnterVar(c *VarContext)

	// EnterGopts is called when entering the gopts production.
	EnterGopts(c *GoptsContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// EnterPrefork is called when entering the prefork production.
	EnterPrefork(c *PreforkContext)

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

	// EnterTopts is called when entering the topts production.
	EnterTopts(c *ToptsContext)

	// EnterToptipttl is called when entering the toptipttl production.
	EnterToptipttl(c *ToptipttlContext)

	// EnterToptparent is called when entering the toptparent production.
	EnterToptparent(c *ToptparentContext)

	// EnterToptpriority is called when entering the toptpriority production.
	EnterToptpriority(c *ToptpriorityContext)

	// EnterToptretry is called when entering the toptretry production.
	EnterToptretry(c *ToptretryContext)

	// EnterListen is called when entering the listen production.
	EnterListen(c *ListenContext)

	// EnterPort is called when entering the port production.
	EnterPort(c *PortContext)

	// EnterProto is called when entering the proto production.
	EnterProto(c *ProtoContext)

	// EnterInterface is called when entering the interface production.
	EnterInterface(c *InterfaceContext)

	// EnterCheck is called when entering the check production.
	EnterCheck(c *CheckContext)

	// EnterCtype is called when entering the ctype production.
	EnterCtype(c *CtypeContext)

	// EnterChttpcode is called when entering the chttpcode production.
	EnterChttpcode(c *ChttpcodeContext)

	// EnterChttpdigest is called when entering the chttpdigest production.
	EnterChttpdigest(c *ChttpdigestContext)

	// EnterForward is called when entering the forward production.
	EnterForward(c *ForwardContext)

	// EnterRoute is called when entering the route production.
	EnterRoute(c *RouteContext)

	// EnterSesstimeout is called when entering the sesstimeout production.
	EnterSesstimeout(c *SesstimeoutContext)

	// EnterSticky is called when entering the sticky production.
	EnterSticky(c *StickyContext)

	// EnterRedirect is called when entering the redirect production.
	EnterRedirect(c *RedirectContext)

	// EnterRname is called when entering the rname production.
	EnterRname(c *RnameContext)

	// EnterRstmt is called when entering the rstmt production.
	EnterRstmt(c *RstmtContext)

	// ExitConfig is called when exiting the config production.
	ExitConfig(c *ConfigContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitVar is called when exiting the var production.
	ExitVar(c *VarContext)

	// ExitGopts is called when exiting the gopts production.
	ExitGopts(c *GoptsContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)

	// ExitPrefork is called when exiting the prefork production.
	ExitPrefork(c *PreforkContext)

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

	// ExitTopts is called when exiting the topts production.
	ExitTopts(c *ToptsContext)

	// ExitToptipttl is called when exiting the toptipttl production.
	ExitToptipttl(c *ToptipttlContext)

	// ExitToptparent is called when exiting the toptparent production.
	ExitToptparent(c *ToptparentContext)

	// ExitToptpriority is called when exiting the toptpriority production.
	ExitToptpriority(c *ToptpriorityContext)

	// ExitToptretry is called when exiting the toptretry production.
	ExitToptretry(c *ToptretryContext)

	// ExitListen is called when exiting the listen production.
	ExitListen(c *ListenContext)

	// ExitPort is called when exiting the port production.
	ExitPort(c *PortContext)

	// ExitProto is called when exiting the proto production.
	ExitProto(c *ProtoContext)

	// ExitInterface is called when exiting the interface production.
	ExitInterface(c *InterfaceContext)

	// ExitCheck is called when exiting the check production.
	ExitCheck(c *CheckContext)

	// ExitCtype is called when exiting the ctype production.
	ExitCtype(c *CtypeContext)

	// ExitChttpcode is called when exiting the chttpcode production.
	ExitChttpcode(c *ChttpcodeContext)

	// ExitChttpdigest is called when exiting the chttpdigest production.
	ExitChttpdigest(c *ChttpdigestContext)

	// ExitForward is called when exiting the forward production.
	ExitForward(c *ForwardContext)

	// ExitRoute is called when exiting the route production.
	ExitRoute(c *RouteContext)

	// ExitSesstimeout is called when exiting the sesstimeout production.
	ExitSesstimeout(c *SesstimeoutContext)

	// ExitSticky is called when exiting the sticky production.
	ExitSticky(c *StickyContext)

	// ExitRedirect is called when exiting the redirect production.
	ExitRedirect(c *RedirectContext)

	// ExitRname is called when exiting the rname production.
	ExitRname(c *RnameContext)

	// ExitRstmt is called when exiting the rstmt production.
	ExitRstmt(c *RstmtContext)
}
