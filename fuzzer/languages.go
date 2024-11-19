package fuzzer

type languageRules = map[string][]expression

type languages struct {
	TinyC languageRules
}

var Languages = languages{
	TinyC: tinyC(),
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
	language["<assign>"] = []expression{{"<ID_AS>=<INT>", 1}}
	language["<test>"] = []expression{{"<sum>", 0.5}, {"<sum><<sum>", 0.5}}
	language["<sum>"] = []expression{{"<term>", 0.33}, {"<sum>+<term>", 0.33}, {"<sum>-<term>", 0.33}}
	language["<term>"] = []expression{{"<ID>", 0.2}, {"<INT>", 0.7}, {"<paren_expr>", 0.1}}

	return language
}
