// Calc.g4
grammar Calc;

// Tokens
OPP: '(' ;
CLP: ')' ;
MUL: '*' ;
DIV: '/' ;
ADD: '+' ;
SUB: '-' ;
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : expression EOF;

expression
   : '(' expression ')' # Parentheses
   | expression op=('*'|'/') expression # MulDiv
   | expression op=('+'|'-') expression # AddSub
   | NUMBER                             # Number
   ;
