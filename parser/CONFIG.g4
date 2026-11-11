grammar CONFIG;

config: (NEWLINE | stmt NEWLINE)* stmt? EOF ;

stmt: assignment
    | table
    | interval
    ;

assignment: NAME '=' HOST ;

interval: INTERVAL NUMBER ;

var: '$' NAME ;

table: TABLE tname tdisable? '{' tlist '}' ;
tname: '<' NAME '>' ;
tlist: (titem (',' titem)* ','?)? ;
titem: var
     | HOST
     ;
tdisable: DISABLE ;

DISABLE: 'disable' ;
INTERVAL: 'interval' ;
TABLE: 'table' ;
NUMBER: [0-9]+ ;
NAME: [a-z][a-z0-9_]* ;
HOST: [a-z0-9._-]+ ;

COMMENT: '#' ~[\r\n]* -> skip ;
WS: [ \t]+ -> skip ;
CONTINUATION: '\\' '\r'? '\n' -> skip ;
NEWLINE: '\r'? '\n'  ;
