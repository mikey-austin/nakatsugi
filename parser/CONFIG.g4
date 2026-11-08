grammar CONFIG;

config: (stmt NEWLINE)* ;

stmt: assignment ;

assignment: VAR '=' HOST ;

VAR: [a-z][a-z0-9_]+ ;
HOST: [a-z0-9.]+ ;
NEWLINE: [\r\n]+ ;