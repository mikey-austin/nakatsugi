grammar CONFIG;

config: (NEWLINE | stmt NEWLINE)* stmt? EOF ;

stmt: assignment
    | table
    ;

assignment: NAME '=' HOST ;

var: '$' NAME ;

table: 'table' tname tdisable? '{' tlist '}' ;
tname: '<' NAME '>' ;
tlist: (titem (',' titem)* ','?)? ;
titem: var
     | HOST
     ;
tdisable: 'disable' ;

NAME: [a-z][a-z0-9_]* ;
HOST: [a-z0-9._-]+ ;

COMMENT: '#' ~[\r\n]* -> skip ;
WS: [ \t]+ -> skip ;
CONTINUATION: '\\' '\r'? '\n' -> skip ;
NEWLINE: '\r'? '\n'  ;
