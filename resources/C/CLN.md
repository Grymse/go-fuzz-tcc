smallc_program ::= 

type_specifier ::= int

compound_statement ::= ‘{‘ (var_decl* statement*)? ‘}’

var_decl ::= type_specifier var_decl_list ‘;’

var_decl_list ::= variable_id ( ‘,’ variable_id)*

variable_id ::= id ( ‘=’ expr )?

**Statements**
statement ::= compound_statement | cond_statement | while_statement | break ‘;’ | continue ‘;’ | return expr ‘;’

cond_statement ::= if ‘(‘ expr ‘)’ statement (else statement)?

while_statement ::= while ‘(‘ expr ‘)’ statement


**Expressions:**
expr ::= id ‘=’ expr | condition

condition ::= disjunction | disjunction ‘?’ expr ‘:’ condition

disjunction ::= conjunction | disjunction ‘||’ conjunction

conjunction ::= comparison | conjunction ‘&&’ comparison

comparison ::= relation | relation ‘==’ relation

relation ::= sum | sum (‘<’ | ‘>’) sum

**Arithmetic:**
sum ::= sum ‘+’ term | sum ‘-’ term | term

term ::= term ‘*’ factor | term ‘/’ factor | term ‘%’ factor | factor

factor ::= ‘!’ factor | ‘-’ factor | primary

primary ::= num | id | ‘(‘ expr ‘)’



**UNDEFINED**
num
id