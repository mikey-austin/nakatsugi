grammar CONFIG;

config: (NEWLINE | stmt NEWLINE)* stmt? EOF ;

stmt
    : assignment
    | table
    | gopts
    | redirect
    ;

assignment: NAME '=' NAME ;
var: '$' NAME ;

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
table: TABLE tname tdisable? '{' tlist '}' ;
tname: '<' NAME '>' ;
tlist: (titem (',' titem)* ','?)? ;
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
listen: LISTEN ON NAME proto? port interface? ; 
port: PORT NUMBER ;
proto: NAME ;
interface: INTERFACE NAME ;

// Check configuration
check: CHECK ctype ;
ctype
    : chttpcode
    | chttpdigest
/*  | chttps
    | cicmp
    | cscript
    | csend
    | cbinsend
    | ctcp
    | ctls */
    ;
chttpcode: HTTP (HOST NAME)? CODE NUMBER ;
chttpdigest: HTTP (HOST NAME)? DIGEST NAME ;

// Common settings
forward: FORWARD TO tname port? check? ;
route: ROUTE TO tname port? check? ;
sesstimeout: SESSION TIMEOUT NUMBER ;
sticky: STICKY_ADDRESS ;

// Redirects
redirect: REDIRECT rname '{' rstmt '}' ;
rname: '"' NAME '"' ;
rstmt
    : DISABLE
    | listen
    | forward
    | sesstimeout
    | sticky
    ;

// Terminal tokens
CHECK: 'check' ;
HTTP: 'http' ;
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
NAME: [a-z0-9.:_-]+ ;

// General rules
COMMENT: '#' ~[\r\n]* -> skip ;
WS: [ \t]+ -> skip ;
CONTINUATION: '\\' '\r'? '\n' -> skip ;
NEWLINE: '\r'? '\n'  ;
