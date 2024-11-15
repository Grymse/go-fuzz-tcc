package fuzzer

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type output = string
type prob = float64

type expression struct {
	output
	prob
}

type Variables struct {
	variables []string
}

func (v *Variables) add_variable() string {
	const vowels = "aeiou"
	const consonants = "bcdfghjklmnpqrstvwxyz"
	// Generate 10 characters
	variable := make([]byte, 5)
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			variable[i] = consonants[rand.Intn(len(consonants))]
		} else {
			variable[i] = vowels[rand.Intn(len(vowels))]
		}
	}
	v.variables = append(v.variables, string(variable))
	return string(variable)
}

func (v *Variables) get_variable() string {
	return v.variables[rand.Intn(len(v.variables))]
}

type Fuzzer struct {
	lang        languageRules
	accumulator strings.Builder
	variables   Variables
}

func Fuzz(lang languageRules) string {
	fuzzer := Fuzzer{
		lang:        lang,
		accumulator: strings.Builder{},
		variables:   Variables{},
	}

	fuzzer.appendExpression(fuzzer.lang["<assign>"]) // Add variable to ensure at least one variable is present
	fuzzer.accumulator.WriteString(" ")
	fuzzer.appendExpression(fuzzer.lang["<program>"])
	return fuzzer.String()
}

func (fuzzer *Fuzzer) String() string {
	return fuzzer.accumulator.String()
}

func getRule(expressions []expression) expression {
	// Choose a rule at random
	probSum := rand.Float64()
	lowerBound := 0.0
	for _, expression := range expressions {
		upperBound := lowerBound + expression.prob
		if lowerBound <= probSum && probSum < upperBound {
			return expression
		}
		lowerBound = upperBound
	}

	// If no rule was chosen, return the first rule
	fmt.Errorf("No rule was chosen, returning the first rule")
	return expressions[0]
}

func (fuzzer *Fuzzer) appendExpression(expressions []expression) {
	// Choose a rule at random
	rule := getRule(expressions)

	// Split the rule into tokens
	tokens := strings.Split(rule.output, " ")
	for i, token := range tokens {
		if i != 0 {
			fuzzer.accumulator.WriteString(" ")
		}

		if token == "<ID_AS>" {
			fuzzer.accumulator.WriteString(fuzzer.variables.add_variable())
			continue
		}

		if token == "<ID>" {
			fuzzer.accumulator.WriteString(fuzzer.variables.get_variable())
			continue
		}

		if token == "<INT>" {
			fuzzer.accumulator.WriteString(strconv.Itoa(rand.Intn(10000)))
			continue
		}

		// If the token is a non-terminal, recursively fuzz it
		newEntry, hasEntry := fuzzer.lang[token]
		if hasEntry {
			fuzzer.appendExpression(newEntry)
			continue
		}
		// If the token is a terminal, append it to the accumulator
		fuzzer.accumulator.WriteString(token)
	}
}

/*
 *  <program> ::= <statement>
 *  <statement> ::= "if" <paren_expr> <statement> |
 *                  "if" <paren_expr> <statement> "else" <statement> |
 *                  "while" <paren_expr> <statement> |
 *                  "do" <statement> "while" <paren_expr> ";" |
 *                  "{" { <statement> } "}" |
 *                  <expr> ";" |
 *                  ";"
 *  <paren_expr> ::= "(" <expr> ")"
 *  <expr> ::= <test> | <id> "=" <expr>
 *  <test> ::= <sum> | <sum> "<" <sum>
 *  <sum> ::= <term> | <sum> "+" <term> | <sum> "-" <term>
 *  <term> ::= <id> | <int> | <paren_expr>
 *  <id> ::= "a" | "b" | "c" | "d" | ... | "z"
 *  <int> ::= <an_unsigned_decimal_integer>
 */
