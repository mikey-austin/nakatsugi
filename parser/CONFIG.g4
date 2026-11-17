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
    | timeout
    | prefork
    | socket
    ;
timeout: TIMEOUT NUMBER ;
interval: INTERVAL NUMBER ;
prefork: PREFORK NUMBER ;
socket: SOCKET QNAME ;
// agentx (also prometheus extension)
// log
// log connection
// log level

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
chost: HOST NAME ;
cdigest: DIGEST QNAME ;
chttpcode: HTTP QNAME chost? CODE NUMBER ;
chttpdigest: HTTP QNAME chost? cdigest ;
chttpscode: HTTPS QNAME chost? CODE NUMBER ;
chttpsdigest: HTTPS QNAME chost? cdigest ;

balancemode: MODE mtype ;
mtype
    : mhash
    | mleaststates
    | mloadbalance
    | mrandom
    | mroundrobin
    | msourcehash
    ;
mhash: HASH QNAME ;
mleaststates: LEAST_STATES ;
mloadbalance: LOAD_BALANCE QNAME ;
mrandom: RANDOM ;
mroundrobin: ROUND_ROBIN ;
msourcehash: SOURCE_HASH QNAME ;

demote: DEMOTE NAME ;
dopts
    : check
    | balancemode
    | interval
    | timeout
    | demote
    ;
destination: tname port? dopts* ;
forward: FORWARD TO destination ;
route: ROUTE TO destination ;
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
    | route
    ;

// Terminal tokens
DOLLA: '$' ;
EQUALS: '=' ;
COMMA: ',' ;
LANGLE: '<' ;
RANGLE: '>' ;
LCURLY: '{' ;
RCURLY: '}' ;
HASH: 'hash' ;
MODE: 'mode' ;
DEMOTE: 'demote' ;
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
SOURCE_HASH: 'source-hash' ;
ROUND_ROBIN: 'roundrobin' ;
RANDOM: 'random' ;
LOAD_BALANCE: 'loadbalance' ;
LEAST_STATES: 'least-states' ;
STICKY_ADDRESS: 'sticky-address' ;
SESSION: 'session' ;
TIMEOUT: 'timeout' ;
FORWARD: 'forward' ;
SOCKET: 'socket' ;
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
