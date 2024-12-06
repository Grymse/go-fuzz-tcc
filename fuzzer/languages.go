package fuzzer

type languageRules = map[string][]expression

type languages struct {
	TinyC languageRules
	CLN   languageRules
}

var Languages = languages{
	TinyC: tinyC(),
	CLN:   cln(),
}

func tinyC() languageRules {
	language := make(languageRules)
	language["<program>"] = []expression{{"<statement>", 1, 1}}
	language["<statement>"] = []expression{{"if<paren_expr><statement>", 0.1, 1},
		{"if<paren_expr><statement>else <statement>", 0.2, 1},
		{"while<paren_expr><statement>", 0.1, 1},
		{"do <statement>while<paren_expr>;", 0.15, 1},
		{"{<statement>}", 0.15, 1},
		{"<expr>;", 0.2, 1},
		{";", 0.1, 1}}
	language["<paren_expr>"] = []expression{{"(<expr>)", 1, 1}}
	language["<expr>"] = []expression{{"<test>", 0.5, 1}, {"<assign>", 0.5, 1}}
	language["<assign>"] = []expression{{"$ID_DECL$=$INT$", 1, 1}}
	language["<test>"] = []expression{{"<sum>", 0.5, 1}, {"<sum><<sum>", 0.5, 1}}
	language["<sum>"] = []expression{{"<term>", 0.33, 1}, {"<sum>+<term>", 0.33, 1}, {"<sum>-<term>", 0.33, 1}}
	language["<term>"] = []expression{{"$ID$", 0.2, 1}, {"$INT$", 0.7, 1}, {"<paren_expr>", 0.1, 1}}

	return language
}

func cln() languageRules {
	language := make(languageRules)
	language["<program>"] = []expression{
		{"<global_decl!>\n int main() {<variable_decl_as>; <statement*>; return 1;}", 1, 13}}
	language["<global_decl>"] = []expression{
		{"<variable_decl_as!>", 1, 1},
		{"<func_decl*>", 1, 1},
	}
	language["<func_decl>"] = []expression{
		{"^$FUNC_DECL$ <compound_statement>", 1, 1},
	}
	language["<compound_statement>"] = []expression{{"{\n <var_decl*> <statement*> \n}", 0.8, 12}, {"{}", 0.2, 1}}
	language["<var_decl>"] = []expression{{"<type_specifier> <var_decl_list> ;", 1, 3}}
	language["<type_specifier>"] = []expression{
		{"<number_types>", 11, 1},
		{"char", 1, 1}}
	language["<number_types>"] = []expression{
		{"int", 1, 1},
		{"signed int", 1, 1},
		{"unsigned int", 1, 1},
		{"short", 1, 1},
		{"signed short", 1, 1},
		{"unsigned short", 1, 1},
		{"long", 1, 1},
		{"signed long", 1, 1},
		{"unsigned long", 1, 1},
		{"float", 1, 1},
		{"double", 1, 1},
	}
	language["<arr_decl>"] = []expression{
		{"const <number_types> $ID_DECL_ARR_C$<arr_decl_array>", 1, 1},
		{"<number_types> $ID_DECL_ARR$<arr_decl_array>", 1, 1},
	}
	language["<arr_decl_array>"] = []expression{
		{"[$INT$]", 1, 1},
		{"[$INT$]={<comma_separated_value>}", 3, 1},
		{"[$INT$]={}", 1, 1},
		{"[]", 1, 1},
		{"[]={<comma_separated_value>}", 3, 1},
		{"[]={}", 1, 1},
	}
	language["<string_decl>"] = []expression{
		{"const char $ID_DECL_ARR_C$[] = \"$LOREM$\"", 1, 1},
		{"char $ID_DECL_ARR$[] = \"$LOREM$\"", 1, 1},
	}
	language["<comma_separated_value>"] = []expression{
		{"<value>,<comma_separated_value>", 10, 2},
		{"<value>", 1, 1},
	}
	language["<var_decl_list>"] = []expression{{"<variable_id_as>", 0.5, 2}, {"<variable_id_as>, <var_decl_list>", 0.5, 3}}
	language["<variable_id_as>"] = []expression{{"$ID_DECL$", 0.5, 1}, {"$ID_DECL$=<expr>", 0.5, 11}}
	language["<variable_decl_as>"] = []expression{
		{"<type_specifier> $ID_DECL$=<value>;", 1, 1},
		{"<type_specifier> $ID_DECL$;", 1, 1},
		{"const <type_specifier> $ID_DECL_C$=<value>;", 1, 1},
		{"<string_decl>;", 1, 1},
		{"<arr_decl>;", 1, 1}}
	language["<statement>"] = []expression{
		{"<switch_statement>", 1, 13},
		{"<do_while_statement>", 1, 13},
		{"<for_statement>", 1, 13},
		{"<compound_statement>", 1, 14},
		{"<cond_statement>", 1, 13},
		{"<while_statement>", 1, 13},
		{"<multiline_comment>;", 1, 12},
		{"// $LOREM$ \n <statement>", 1, 2},
		{"$FUNC$;", 1, 2},
		{"return $ID$;", 1, 1},
	}
	language["<multiline_comment>"] = []expression{
		{"/* $LOREM$ */", 1, 1},
		{"/*\n* $LOREM$\n */", 1, 1},
	}
	language["<switch_statement>"] = []expression{
		{"switch (<expr>) {<case_statement*>}", 1, 12},
		{"switch (<expr>) {<case_statement*> \n default:\n <statement>}", 1, 12},
	}
	language["<case_statement>"] = []expression{
		{"\ncase $INT$:\n <statement> \nbreak;", 1, 12},
		{"\ncase $CHAR$:\n <statement> \nbreak;", 1, 12},
		{"\ncase $INT$:\n <statement>", 1, 12},
		{"\ncase $CHAR$:\n <statement>", 1, 12},
	}
	language["<do_while_statement>"] = []expression{
		{"do <loop_statement> while (<expr>);", 1, 12},
	}
	language["<while_statement>"] = []expression{
		{"while (<expr>) <loop_statement>", 1, 12},
	}
	language["<for_statement>"] = []expression{
		{"^for (int $ID_DECL$ = 0; <condition>; $ID_AS$ = <expr>) <loop_statement>", 0.5, 12},
		{"for (; <condition>; ) <loop_statement>", 0.5, 12},
	}
	language["<loop_statement>"] = []expression{
		{"{\n <loop_inner_statement*> \n}", 0.6, 12},
		{"{}", 0.15, 1},
		{"<loop_inner_statement>", 0.25, 1},
	}
	language["<loop_inner_statement>"] = []expression{
		{"<statement>", 1, 12},
		{"break;", 1, 12},
		{"continue;", 1, 12},
	}
	language["<cond_statement>"] = []expression{
		{"if (<expr>) <statement>", 1, 12},
		{"if (<expr>) <statement> else <statement>", 1, 12},
	}
	language["<expr>"] = []expression{
		{"$ID_AS$ = <expr>", 1, 12},
		{"<condition>", 1, 10},
	}
	language["<condition>"] = []expression{
		{"<disjunction>", 1, 9},
		{"<disjunction> ? <expr> : <condition>", 1, 11},
	}
	language["<disjunction>"] = []expression{
		{"<conjunction>", 1, 8},
		{"<disjunction> || <conjunction>", 1, 9},
	}
	language["<conjunction>"] = []expression{
		{"<comparison>", 1, 7},
		{"<conjunction> && <comparison>", 1, 8},
	}
	language["<comparison>"] = []expression{
		{"<relation>", 1, 6},
		{"<relation> == <relation>", 1, 6},
		{"<relation> != <relation>", 1, 6},
	}
	language["<relation>"] = []expression{
		{"<sum>", 1, 5},
		{"<sum> <= <sum>", 1, 5},
		{"<sum> < <sum>", 1, 5},
		{"<sum> >= <sum>", 1, 5},
		{"<sum> > <sum>", 1, 5},
	}
	language["<sum>"] = []expression{
		{"<term>", 1, 4},
		{"<sum> + <term>", 1, 5},
		{"<sum> - <term>", 1, 5},
	}
	language["<term>"] = []expression{
		{"<factor>", 1, 3},
		{"<term> * <factor>", 1, 4},
		{"<term> / <factor>", 1, 4},
	}
	language["<factor>"] = []expression{
		{"! <factor>", 0.34, 3},
		{"- <factor>", 0.33, 3},
		{"<primary>", 0.33, 2},
	}
	language["<value>"] = []expression{
		{"$INT$", 1, 1},
		{"$INT$.$INT$f", 1, 1},
		{"$CHAR$", 1, 1},
	}
	language["<primary>"] = []expression{
		{"<value>", 1, 1},
		{"<value><multiline_comment>", 1, 1},
		{"$ID$", 1, 1},
		{"$ID$<multiline_comment>", 1, 2},
		{"$FUNC$<multiline_comment>", 1, 2},
		{"$FUNC$", 1, 2},
		{"(<expr>)", 1, 11},
	}

	return language
}
