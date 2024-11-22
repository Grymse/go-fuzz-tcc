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
	language["<program>"] = []expression{{"<statement>", 1}}
	language["<statement>"] = []expression{{"if<paren_expr><statement>", 0.1},
		{"if<paren_expr><statement>else <statement>", 0.2},
		{"while<paren_expr><statement>", 0.1},
		{"do <statement>while<paren_expr>;", 0.15},
		{"{<statement>}", 0.15},
		{"<expr>;", 0.2},
		{";", 0.1}}
	language["<paren_expr>"] = []expression{{"(<expr>)", 1}}
	language["<expr>"] = []expression{{"<test>", 0.5}, {"<assign>", 0.5}}
	language["<assign>"] = []expression{{"$ID_AS$=$INT$", 1}}
	language["<test>"] = []expression{{"<sum>", 0.5}, {"<sum><<sum>", 0.5}}
	language["<sum>"] = []expression{{"<term>", 0.33}, {"<sum>+<term>", 0.33}, {"<sum>-<term>", 0.33}}
	language["<term>"] = []expression{{"$ID$", 0.2}, {"$INT$", 0.7}, {"<paren_expr>", 0.1}}

	return language
}

func cln() languageRules {
	language := make(languageRules)
	language["<program>"] = []expression{{"int main() <compound_statement>", 1}}
	language["<compound_statement>"] = []expression{{"{ <var_decl*> <statement*> }", 0.8}, {"{}", 0.2}}
	language["<var_decl>"] = []expression{{"<type_specifier> <var_decl_list> ;", 1}}
	language["<type_specifier>"] = []expression{{"int", 1}}
	language["<var_decl_list>"] = []expression{{"<variable_id>", 0.5}, {"<variable_id>, <var_decl_list>", 0.5}}
	language["<variable_id>"] = []expression{{"$ID$", 0.5}, {"$ID$=<expr>", 0.5}}
	language["<statement>"] = []expression{
		{"<compound_statement>", 0.25},
		{"<cond_statement>", 0.25},
		{"<while_statement>", 0.25},
		{"return <expr> ‘;’", 0.25},
	}
	language["<cond_statement>"] = []expression{
		{"if (<expr>) <statement>", 0.5},
		{"if (<expr>) <statement> else <statement>", 0.5},
	}
	language["<while_statement>"] = []expression{
		{"while (<expr>) <statement>", 1},
	}
	language["<expr>"] = []expression{
		{"$ID$ = <expr>", 0.5},
		{"<condition>", 0.5},
	}
	language["<condition>"] = []expression{
		{"<disjunction>", 0.5},
		{"<disjunction> ? <expr> : <condition>", 0.5},
	}
	language["<disjunction>"] = []expression{
		{"<conjunction>", 0.5},
		{"<disjunction> || <conjunction>", 0.5},
	}
	language["<conjunction>"] = []expression{
		{"<comparison>", 0.5},
		{"<conjunction> && <comparison>", 0.5},
	}
	language["<comparison>"] = []expression{
		{"<relation>", 0.34},
		{"<relation> == <relation>", 0.33},
		{"<relation> != <relation>", 0.33},
	}
	language["<relation>"] = []expression{
		{"<sum>", 0.2},
		{"<sum> <= <sum>", 0.2},
		{"<sum> < <sum>", 0.2},
		{"<sum> >= <sum>", 0.2},
		{"<sum> > <sum>", 0.2},
	}
	language["<sum>"] = []expression{
		{"<term>", 0.33},
		{"<sum> + <term>", 0.33},
		{"<sum> - <term>", 0.33},
	}
	language["<term>"] = []expression{
		{"<factor>", 0.25},
		{"<term> * <factor>", 0.25},
		{"<term> / <factor>", 0.25},
		{"<term> % <factor>", 0.25},
	}
	language["<factor>"] = []expression{
		{"!<factor>", 0.34},
		{"-<factor>", 0.33},
		{"<primary>", 0.33},
	}
	language["<primary>"] = []expression{
		{"$INT$", 0.34},
		{"$ID$", 0.33},
		{"(<expr>)", 0.33},
	}

	return language
}
