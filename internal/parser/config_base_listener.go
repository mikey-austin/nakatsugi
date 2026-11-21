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

// EnterVar is called when production var is entered.
func (s *BaseCONFIGListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production var is exited.
func (s *BaseCONFIGListener) ExitVar(ctx *VarContext) {}

// EnterGopts is called when production gopts is entered.
func (s *BaseCONFIGListener) EnterGopts(ctx *GoptsContext) {}

// ExitGopts is called when production gopts is exited.
func (s *BaseCONFIGListener) ExitGopts(ctx *GoptsContext) {}

// EnterTimeout is called when production timeout is entered.
func (s *BaseCONFIGListener) EnterTimeout(ctx *TimeoutContext) {}

// ExitTimeout is called when production timeout is exited.
func (s *BaseCONFIGListener) ExitTimeout(ctx *TimeoutContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseCONFIGListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseCONFIGListener) ExitInterval(ctx *IntervalContext) {}

// EnterPrefork is called when production prefork is entered.
func (s *BaseCONFIGListener) EnterPrefork(ctx *PreforkContext) {}

// ExitPrefork is called when production prefork is exited.
func (s *BaseCONFIGListener) ExitPrefork(ctx *PreforkContext) {}

// EnterSocket is called when production socket is entered.
func (s *BaseCONFIGListener) EnterSocket(ctx *SocketContext) {}

// ExitSocket is called when production socket is exited.
func (s *BaseCONFIGListener) ExitSocket(ctx *SocketContext) {}

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

// EnterTopts is called when production topts is entered.
func (s *BaseCONFIGListener) EnterTopts(ctx *ToptsContext) {}

// ExitTopts is called when production topts is exited.
func (s *BaseCONFIGListener) ExitTopts(ctx *ToptsContext) {}

// EnterToptipttl is called when production toptipttl is entered.
func (s *BaseCONFIGListener) EnterToptipttl(ctx *ToptipttlContext) {}

// ExitToptipttl is called when production toptipttl is exited.
func (s *BaseCONFIGListener) ExitToptipttl(ctx *ToptipttlContext) {}

// EnterToptparent is called when production toptparent is entered.
func (s *BaseCONFIGListener) EnterToptparent(ctx *ToptparentContext) {}

// ExitToptparent is called when production toptparent is exited.
func (s *BaseCONFIGListener) ExitToptparent(ctx *ToptparentContext) {}

// EnterToptpriority is called when production toptpriority is entered.
func (s *BaseCONFIGListener) EnterToptpriority(ctx *ToptpriorityContext) {}

// ExitToptpriority is called when production toptpriority is exited.
func (s *BaseCONFIGListener) ExitToptpriority(ctx *ToptpriorityContext) {}

// EnterToptretry is called when production toptretry is entered.
func (s *BaseCONFIGListener) EnterToptretry(ctx *ToptretryContext) {}

// ExitToptretry is called when production toptretry is exited.
func (s *BaseCONFIGListener) ExitToptretry(ctx *ToptretryContext) {}

// EnterListen is called when production listen is entered.
func (s *BaseCONFIGListener) EnterListen(ctx *ListenContext) {}

// ExitListen is called when production listen is exited.
func (s *BaseCONFIGListener) ExitListen(ctx *ListenContext) {}

// EnterPort is called when production port is entered.
func (s *BaseCONFIGListener) EnterPort(ctx *PortContext) {}

// ExitPort is called when production port is exited.
func (s *BaseCONFIGListener) ExitPort(ctx *PortContext) {}

// EnterPortrange is called when production portrange is entered.
func (s *BaseCONFIGListener) EnterPortrange(ctx *PortrangeContext) {}

// ExitPortrange is called when production portrange is exited.
func (s *BaseCONFIGListener) ExitPortrange(ctx *PortrangeContext) {}

// EnterLport is called when production lport is entered.
func (s *BaseCONFIGListener) EnterLport(ctx *LportContext) {}

// ExitLport is called when production lport is exited.
func (s *BaseCONFIGListener) ExitLport(ctx *LportContext) {}

// EnterProto is called when production proto is entered.
func (s *BaseCONFIGListener) EnterProto(ctx *ProtoContext) {}

// ExitProto is called when production proto is exited.
func (s *BaseCONFIGListener) ExitProto(ctx *ProtoContext) {}

// EnterInterface is called when production interface is entered.
func (s *BaseCONFIGListener) EnterInterface(ctx *InterfaceContext) {}

// ExitInterface is called when production interface is exited.
func (s *BaseCONFIGListener) ExitInterface(ctx *InterfaceContext) {}

// EnterCheck is called when production check is entered.
func (s *BaseCONFIGListener) EnterCheck(ctx *CheckContext) {}

// ExitCheck is called when production check is exited.
func (s *BaseCONFIGListener) ExitCheck(ctx *CheckContext) {}

// EnterCtype is called when production ctype is entered.
func (s *BaseCONFIGListener) EnterCtype(ctx *CtypeContext) {}

// ExitCtype is called when production ctype is exited.
func (s *BaseCONFIGListener) ExitCtype(ctx *CtypeContext) {}

// EnterChost is called when production chost is entered.
func (s *BaseCONFIGListener) EnterChost(ctx *ChostContext) {}

// ExitChost is called when production chost is exited.
func (s *BaseCONFIGListener) ExitChost(ctx *ChostContext) {}

// EnterCdigest is called when production cdigest is entered.
func (s *BaseCONFIGListener) EnterCdigest(ctx *CdigestContext) {}

// ExitCdigest is called when production cdigest is exited.
func (s *BaseCONFIGListener) ExitCdigest(ctx *CdigestContext) {}

// EnterChttpcode is called when production chttpcode is entered.
func (s *BaseCONFIGListener) EnterChttpcode(ctx *ChttpcodeContext) {}

// ExitChttpcode is called when production chttpcode is exited.
func (s *BaseCONFIGListener) ExitChttpcode(ctx *ChttpcodeContext) {}

// EnterChttpdigest is called when production chttpdigest is entered.
func (s *BaseCONFIGListener) EnterChttpdigest(ctx *ChttpdigestContext) {}

// ExitChttpdigest is called when production chttpdigest is exited.
func (s *BaseCONFIGListener) ExitChttpdigest(ctx *ChttpdigestContext) {}

// EnterChttpscode is called when production chttpscode is entered.
func (s *BaseCONFIGListener) EnterChttpscode(ctx *ChttpscodeContext) {}

// ExitChttpscode is called when production chttpscode is exited.
func (s *BaseCONFIGListener) ExitChttpscode(ctx *ChttpscodeContext) {}

// EnterChttpsdigest is called when production chttpsdigest is entered.
func (s *BaseCONFIGListener) EnterChttpsdigest(ctx *ChttpsdigestContext) {}

// ExitChttpsdigest is called when production chttpsdigest is exited.
func (s *BaseCONFIGListener) ExitChttpsdigest(ctx *ChttpsdigestContext) {}

// EnterBalancemode is called when production balancemode is entered.
func (s *BaseCONFIGListener) EnterBalancemode(ctx *BalancemodeContext) {}

// ExitBalancemode is called when production balancemode is exited.
func (s *BaseCONFIGListener) ExitBalancemode(ctx *BalancemodeContext) {}

// EnterMtype is called when production mtype is entered.
func (s *BaseCONFIGListener) EnterMtype(ctx *MtypeContext) {}

// ExitMtype is called when production mtype is exited.
func (s *BaseCONFIGListener) ExitMtype(ctx *MtypeContext) {}

// EnterMhash is called when production mhash is entered.
func (s *BaseCONFIGListener) EnterMhash(ctx *MhashContext) {}

// ExitMhash is called when production mhash is exited.
func (s *BaseCONFIGListener) ExitMhash(ctx *MhashContext) {}

// EnterMleaststates is called when production mleaststates is entered.
func (s *BaseCONFIGListener) EnterMleaststates(ctx *MleaststatesContext) {}

// ExitMleaststates is called when production mleaststates is exited.
func (s *BaseCONFIGListener) ExitMleaststates(ctx *MleaststatesContext) {}

// EnterMloadbalance is called when production mloadbalance is entered.
func (s *BaseCONFIGListener) EnterMloadbalance(ctx *MloadbalanceContext) {}

// ExitMloadbalance is called when production mloadbalance is exited.
func (s *BaseCONFIGListener) ExitMloadbalance(ctx *MloadbalanceContext) {}

// EnterMrandom is called when production mrandom is entered.
func (s *BaseCONFIGListener) EnterMrandom(ctx *MrandomContext) {}

// ExitMrandom is called when production mrandom is exited.
func (s *BaseCONFIGListener) ExitMrandom(ctx *MrandomContext) {}

// EnterMroundrobin is called when production mroundrobin is entered.
func (s *BaseCONFIGListener) EnterMroundrobin(ctx *MroundrobinContext) {}

// ExitMroundrobin is called when production mroundrobin is exited.
func (s *BaseCONFIGListener) ExitMroundrobin(ctx *MroundrobinContext) {}

// EnterMsourcehash is called when production msourcehash is entered.
func (s *BaseCONFIGListener) EnterMsourcehash(ctx *MsourcehashContext) {}

// ExitMsourcehash is called when production msourcehash is exited.
func (s *BaseCONFIGListener) ExitMsourcehash(ctx *MsourcehashContext) {}

// EnterDemote is called when production demote is entered.
func (s *BaseCONFIGListener) EnterDemote(ctx *DemoteContext) {}

// ExitDemote is called when production demote is exited.
func (s *BaseCONFIGListener) ExitDemote(ctx *DemoteContext) {}

// EnterDopts is called when production dopts is entered.
func (s *BaseCONFIGListener) EnterDopts(ctx *DoptsContext) {}

// ExitDopts is called when production dopts is exited.
func (s *BaseCONFIGListener) ExitDopts(ctx *DoptsContext) {}

// EnterDestination is called when production destination is entered.
func (s *BaseCONFIGListener) EnterDestination(ctx *DestinationContext) {}

// ExitDestination is called when production destination is exited.
func (s *BaseCONFIGListener) ExitDestination(ctx *DestinationContext) {}

// EnterForward is called when production forward is entered.
func (s *BaseCONFIGListener) EnterForward(ctx *ForwardContext) {}

// ExitForward is called when production forward is exited.
func (s *BaseCONFIGListener) ExitForward(ctx *ForwardContext) {}

// EnterRoute is called when production route is entered.
func (s *BaseCONFIGListener) EnterRoute(ctx *RouteContext) {}

// ExitRoute is called when production route is exited.
func (s *BaseCONFIGListener) ExitRoute(ctx *RouteContext) {}

// EnterSesstimeout is called when production sesstimeout is entered.
func (s *BaseCONFIGListener) EnterSesstimeout(ctx *SesstimeoutContext) {}

// ExitSesstimeout is called when production sesstimeout is exited.
func (s *BaseCONFIGListener) ExitSesstimeout(ctx *SesstimeoutContext) {}

// EnterRedirect is called when production redirect is entered.
func (s *BaseCONFIGListener) EnterRedirect(ctx *RedirectContext) {}

// ExitRedirect is called when production redirect is exited.
func (s *BaseCONFIGListener) ExitRedirect(ctx *RedirectContext) {}

// EnterRname is called when production rname is entered.
func (s *BaseCONFIGListener) EnterRname(ctx *RnameContext) {}

// ExitRname is called when production rname is exited.
func (s *BaseCONFIGListener) ExitRname(ctx *RnameContext) {}

// EnterRstmt is called when production rstmt is entered.
func (s *BaseCONFIGListener) EnterRstmt(ctx *RstmtContext) {}

// ExitRstmt is called when production rstmt is exited.
func (s *BaseCONFIGListener) ExitRstmt(ctx *RstmtContext) {}
