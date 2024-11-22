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
	language["<assign>"] = []expression{{"$ID_AS$=$INT$", 1, 1}}
	language["<test>"] = []expression{{"<sum>", 0.5, 1}, {"<sum><<sum>", 0.5, 1}}
	language["<sum>"] = []expression{{"<term>", 0.33, 1}, {"<sum>+<term>", 0.33, 1}, {"<sum>-<term>", 0.33, 1}}
	language["<term>"] = []expression{{"$ID$", 0.2, 1}, {"$INT$", 0.7, 1}, {"<paren_expr>", 0.1, 1}}

	return language
}

func cln() languageRules {
	language := make(languageRules)
	language["<program>"] = []expression{{"int main() {<assign_id>; <statement*> return 1;}", 1, 13}}
	language["<compound_statement>"] = []expression{{"{\n <var_decl*> <statement*> \n}", 0.8, 12}, {"{}", 0.2, 1}}
	language["<var_decl>"] = []expression{{"<type_specifier> <var_decl_list> ;", 1, 3}}
	language["<type_specifier>"] = []expression{{"int", 1, 1}}
	language["<var_decl_list>"] = []expression{{"<variable_id_as>", 0.5, 2}, {"<variable_id_as>, <var_decl_list>", 0.5, 3}}
	language["<assign_id>"] = []expression{{"int $ID_AS$=$INT$", 1, 1}}
	language["<variable_id_as>"] = []expression{{"$ID_AS$", 0.5, 1}, {"$ID_AS$=<expr>", 0.5, 11}}
	language["<statement>"] = []expression{
		{"<compound_statement>", 0.34, 14},
		{"<cond_statement>", 0.33, 13},
		{"<while_statement>", 0.33, 13},
		{"return $ID$;", 0.25, 11},
	}
	language["<cond_statement>"] = []expression{
		{"if (<expr>) <statement>", 0.5, 12},
		{"if (<expr>) <statement> else <statement>", 0.5, 12},
	}
	language["<while_statement>"] = []expression{
		{"while (<expr>) <statement>", 1, 12},
	}
	language["<expr>"] = []expression{
		{"$ID$ = <expr>", 0.5, 12},
		{"<condition>", 0.5, 10},
	}
	language["<condition>"] = []expression{
		{"<disjunction>", 0.5, 9},
		{"<disjunction> ? <expr> : <condition>", 0.5, 11},
	}
	language["<disjunction>"] = []expression{
		{"<conjunction>", 0.5, 8},
		{"<disjunction> || <conjunction>", 0.5, 9},
	}
	language["<conjunction>"] = []expression{
		{"<comparison>", 0.5, 7},
		{"<conjunction> && <comparison>", 0.5, 8},
	}
	language["<comparison>"] = []expression{
		{"<relation>", 0.34, 6},
		{"<relation> == <relation>", 0.33, 6},
		{"<relation> != <relation>", 0.33, 6},
	}
	language["<relation>"] = []expression{
		{"<sum>", 0.2, 5},
		{"<sum> <= <sum>", 0.2, 5},
		{"<sum> < <sum>", 0.2, 5},
		{"<sum> >= <sum>", 0.2, 5},
		{"<sum> > <sum>", 0.2, 5},
	}
	language["<sum>"] = []expression{
		{"<term>", 0.33, 4},
		{"<sum> + <term>", 0.33, 5},
		{"<sum> - <term>", 0.33, 5},
	}
	language["<term>"] = []expression{
		{"<factor>", 0.25, 3},
		{"<term> * <factor>", 0.25, 4},
		{"<term> / <factor>", 0.25, 4},
		{"<term> % <factor>", 0.25, 4},
	}
	language["<factor>"] = []expression{
		{"! <factor>", 0.34, 3},
		{"- <factor>", 0.33, 3},
		{"<primary>", 0.33, 2},
	}
	language["<primary>"] = []expression{
		{"$INT$", 0.34, 1},
		{"$ID$", 0.33, 1},
		{"(<expr>)", 0.33, 11},
	}

	return language
}
