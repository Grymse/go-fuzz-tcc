grammar CLN;

program:
	global_decl? 'int' 'main' '(' ')' '{' variable_decl_as statement* 'return' '1;' '}';

global_decl: variable_decl_as | func_decl*;

func: FUNC '(' (expr (',' expr)*)? ')';

func_decl:
	type_specifier FUNC_DECL '(' (param_list)? ')' compound_statement;

param_list: param (',' param)*;

param: type_specifier ID;

compound_statement: '{' var_decl* statement* '}' | '{}';

var_decl: type_specifier var_decl_list ';';

type_specifier: number_types | 'char';

number_types:
	'int'
	| 'signed int'
	| 'unsigned int'
	| 'short'
	| 'signed short'
	| 'unsigned short'
	| 'long'
	| 'signed long'
	| 'unsigned long'
	| 'float'
	| 'double';

arr_decl:
	'const' number_types ID_DECL_ARR_C arr_decl_array
	| number_types ID_DECL_ARR arr_decl_array;

arr_decl_array:
	'[' INT ']'
	| '[' INT ']' '=' comma_separated_value
	| '[' INT ']' '=' '{}'
	| '[]'
	| '[]' '=' comma_separated_value
	| '[]' '=' '{}';

string_decl:
	'const' 'char' ID_DECL_ARR_C '[]' '=' LOREM
	| 'char' ID_DECL_ARR '[]' '=' LOREM;

comma_separated_value: value ',' comma_separated_value | value;

var_decl_list:
	variable_id_as
	| variable_id_as ',' var_decl_list;

variable_id_as: ID_DECL | ID_DECL '=' expr;

variable_decl_as:
	type_specifier ID_DECL '=' value ';'
	| type_specifier ID_DECL ';'
	| 'const' type_specifier ID_DECL_C '=' value ';'
	| string_decl ';'
	| arr_decl ';';

statement:
	switch_statement
	| do_while_statement
	| for_statement
	| compound_statement
	| cond_statement
	| while_statement
	| multiline_comment ';'
	| '//' LOREM statement
	| func ';'
	| 'return' ID ';';

multiline_comment:
	'/*' LOREM '*/'
	| '/*' '\n' '*' LOREM '\n' '*/';

switch_statement:
	'switch' '(' expr ')' '{' case_statement* '}'
	| 'switch' '(' expr ')' '{' case_statement* 'default' ':' statement '}';

case_statement:
	'case' INT ':' statement 'break' ';'
	| 'case' CHAR ':' statement 'break' ';'
	| 'case' INT ':' statement
	| 'case' CHAR ':' statement;

do_while_statement:
	'do' loop_statement 'while' '(' expr ')' ';';

while_statement: 'while' '(' expr ')' loop_statement;

for_statement:
	'for' '(' 'int' ID_DECL '=' '0' ';' condition ';' ID_AS '=' expr ')' loop_statement
	| 'for' '(' ';' condition ';' ')' loop_statement;

loop_statement:
	'{' loop_inner_statement* '}'
	| '{}'
	| loop_inner_statement;

loop_inner_statement: statement | 'break' ';' | 'continue' ';';

cond_statement:
	'if' '(' expr ')' statement
	| 'if' '(' expr ')' statement 'else' statement;

expr: ID_AS '=' expr | condition;

condition: disjunction | disjunction '?' expr ':' condition;

disjunction: conjunction | disjunction '||' conjunction;

conjunction: comparison | conjunction '&&' comparison;

comparison:
	relation
	| relation '==' relation
	| relation '!=' relation;

relation:
	sum_
	| sum_ '<=' sum_
	| sum_ '<' sum_
	| sum_ '>=' sum_
	| sum_ '>' sum_;

sum_: term | sum_ '+' term | sum_ '-' term;

term: factor | term '*' factor | term '/' factor;

factor: '!' factor | '-' factor | primary;

primary:
	value
	| value multiline_comment
	| ID
	| ID multiline_comment
	| func multiline_comment
	| func
	| '(' expr ')';

value: INT | INT '.' INT 'f' | CHAR;

INT: [0-9]+;

CHAR: '\'' [a-zA-Z0-9] '\'';

ID: [a-zA-Z_][a-zA-Z0-9_]*;

ID_DECL: ID;

ID_AS: ID;

ID_DECL_ARR: ID;

ID_DECL_ARR_C: ID;

ID_DECL_C: ID;

FUNC_DECL: ID;

FUNC: ID;

LOREM: '"' .*? '"';

WS: [ \r\n\t]+ -> skip;