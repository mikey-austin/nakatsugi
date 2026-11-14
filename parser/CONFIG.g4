grammar CONFIG;

config: (NEWLINE | stmt NEWLINE)* stmt? EOF ;

stmt
    : assignment
    | table
    | gopts
    | redirect
    ;

assignment: NAME EQUALS NAME ;
var: DOLLA NAME ;

// Global top-level options
gopts
    : interval
    | prefork
    ;
interval: INTERVAL NUMBER ;
prefork: PREFORK NUMBER ;
// agentx (also prometheus extension)
// log
// log connection
// log level
// socket
// timeout

// Tables
table: TABLE tname tdisable? LCURLY tlist RCURLY ;
tname: LANGLE NAME RANGLE ;
tlist: (titem (COMMA titem)* COMMA?)? ;
titem
    : var topts*
    | NAME topts*
    ;
tdisable: DISABLE ;
topts
    : toptipttl
    | toptparent
    | toptpriority
    | toptretry
    ;
toptipttl: IP TTL NUMBER ;
toptparent: PARENT NUMBER ;
toptpriority: PRIORITY NUMBER ;
toptretry: RETRY NUMBER ;

// Listen
listen: LISTEN ON NAME proto? lport interface? ; 
port: PORT NUMBER ;
portrange: PORT NAME ; // eg 111:114
lport
    : portrange
    | port
    ;
proto: PROTO NAME ;
interface: INTERFACE NAME ;

// Check configuration
check: CHECK ctype ;
ctype
    : chttpcode
    | chttpdigest
    | chttpscode
    | chttpsdigest
/*  | cicmp
    | cscript
    | csend
    | cbinsend
    | ctcp
    | ctls */
    ;
chttpcode: HTTP QNAME (HOST NAME)? CODE NUMBER ;
chttpdigest: HTTP QNAME (HOST NAME)? DIGEST NAME ;
chttpscode: HTTPS QNAME (HOST NAME)? CODE NUMBER ;
chttpsdigest: HTTPS QNAME (HOST NAME)? DIGEST NAME ;

// Common settings
forward: FORWARD TO tname port? check? ;
route: ROUTE TO tname port? check? ;
sesstimeout: SESSION TIMEOUT NUMBER ;

// Redirects
redirect: REDIRECT rname LCURLY ((NEWLINE | rstmt)* rstmt?) RCURLY ;
rname: QNAME ;
rstmt
    : DISABLE
    | STICKY_ADDRESS
    | sesstimeout
    | listen
    | forward
    ;

// Terminal tokens
DOLLA: '$' ;
EQUALS: '=' ;
COMMA: ',' ;
LANGLE: '<' ;
RANGLE: '>' ;
LCURLY: '{' ;
RCURLY: '}' ;
PROTO: 'proto' ;
CHECK: 'check' ;
HTTP: 'http' ;
HTTPS: 'https' ;
CODE: 'code' ;
DIGEST: 'digest' ;
PREFORK: 'prefork' ;
PRIORITY: 'priority' ;
RETRY: 'retry' ;
PARENT: 'parent' ;
IP: 'ip' ;
TTL: 'ttl' ;
ROUTE: 'route' ;
STICKY_ADDRESS: 'sticky-address' ;
SESSION: 'session' ;
TIMEOUT: 'timeout' ;
FORWARD: 'forward' ;
PORT: 'port' ;
INTERFACE: 'interface' ;
LISTEN: 'listen' ;
TO: 'to' ;
ON: 'on' ;
HOST: 'host' ;
REDIRECT: 'redirect' ;
DISABLE: 'disable' ;
INTERVAL: 'interval' ;
TABLE: 'table' ;
NUMBER: [0-9]+ ;
NAME: [a-z0-9:._-]+ ;
QNAME: '"' [a-zA-Z0-9 /._-]* '"' ;

// General rules
COMMENT: '#' ~[\r\n]* -> skip ;
WS: [ \t]+ -> skip ;
CONTINUATION: '\\' '\r'? '\n' -> skip ;
NEWLINE: '\r'? '\n'  ;
