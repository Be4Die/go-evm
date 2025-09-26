grammar Assembler;

program
    : (directive | label | instruction | comment | emptyLine)* EOF
    ;

directive
    : entryDirective
    | sectionDirective
    | dataDirective
    | equDirective
    ;

entryDirective
    : ENTRY WS? labelName
    ;

sectionDirective
    : SECTION WS? sectionName
    ;

dataDirective
    : (labelName WS? ':')? (DB | DW) WS? constant (WS? ',' WS? constant)*
    ;

equDirective
    : labelName WS? EQU WS? constant
    ;

instruction
    : (labelName WS? ':')? mnemonic WS? operand?
    ;

label
    : labelName WS? ':'
    ;

mnemonic
    : MOV | ADD_I | SUB_I | MUL_I | DIV_I
    | ADD_F | SUB_F | MUL_F | DIV_F
    | CMP_I | CMP_F | JMP | JZ | JNZ
    | JC | JNC | CALL | RET | PUSH
    | POP | IN | OUT | AND | OR
    | XOR | NOT | SHL | SHR | HALT
    ;

operand
    : '[' indirectAddress ']'  #indirectOperand
    | constant                 #directOperand
    ;

indirectAddress
    : constant
    ;

constant
    : number
    | labelName
    ;

number
    : HEX_LITERAL
    | BIN_LITERAL
    | DEC_LITERAL
    ;

labelName
    : IDENTIFIER
    ;

sectionName
    : DOT_DATA
    | DOT_CODE
    ;

comment
    : LINE_COMMENT
    | BLOCK_COMMENT
    ;

emptyLine
    : WS? EOL
    ;

// Directives
ENTRY: [Ee][Nn][Tt][Rr][Yy];
SECTION: [Ss][Ee][Cc][Tt][Ii][Oo][Nn];
DB: [Dd][Bb];
DW: [Dd][Ww];
EQU: [Ee][Qq][Uu];
DOT_DATA: [.][Dd][Aa][Tt][Aa];
DOT_CODE: [.][Cc][Oo][Dd][Ee];

// Instructions
MOV: [Mm][Oo][Vv];
ADD_I: [Aa][Dd][Dd][_][Ii];
SUB_I: [Ss][Uu][Bb][_][Ii];
MUL_I: [Mm][Uu][Ll][_][Ii];
DIV_I: [Dd][Ii][Vv][_][Ii];
ADD_F: [Aa][Dd][Dd][_][Ff];
SUB_F: [Ss][Uu][Bb][_][Ff];
MUL_F: [Mm][Uu][Ll][_][Ff];
DIV_F: [Dd][Ii][Vv][_][Ff];
CMP_I: [Cc][Mm][Pp][_][Ii];
CMP_F: [Cc][Mm][Pp][_][Ff];
JMP: [Jj][Mm][Pp];
JZ: [Jj][Zz];
JNZ: [Jj][Nn][Zz];
JC: [Jj][Cc];
JNC: [Jj][Nn][Cc];
CALL: [Cc][Aa][Ll][Ll];
RET: [Rr][Ee][Tt];
PUSH: [Pp][Uu][Ss][Hh];
POP: [Pp][Oo][Pp];
IN: [Ii][Nn];
OUT: [Oo][Uu][Tt];
AND: [Aa][Nn][Dd];
OR: [Oo][Rr];
XOR: [Xx][Oo][Rr];
NOT: [Nn][Oo][Tt];
SHL: [Ss][Hh][Ll];
SHR: [Ss][Hh][Rr];
HALT: [Hh][Aa][Ll][Tt];

// Literals
HEX_LITERAL: '0x' [0-9A-Fa-f]+;
BIN_LITERAL: '0b' [01]+;
DEC_LITERAL: [0-9]+;

// Identifiers
IDENTIFIER: [a-zA-Z_] [a-zA-Z_0-9]*;

// Comments
LINE_COMMENT: (';' | '//') ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;

// Whitespace and newlines
WS: [ \t]+ -> skip;
EOL: '\r'? '\n';