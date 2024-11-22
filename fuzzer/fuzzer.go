package fuzzer

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type output = string
type prob = float64
type cost = int32

type expression struct {
	output
	prob
	cost
}

type Variables struct {
	variables []string
	Generator func() string
}

func (v *Variables) init() {
	v.variables = make([]string, 0)
	v.Generator = default_variable_generator
}

func default_variable_generator() string {
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
	return string(variable)
}

func (v *Variables) add_variable() string {
	variable := v.Generator()
	v.variables = append(v.variables, variable)
	return string(variable)
}

func (v *Variables) get_variable() string {
	if len(v.variables) == 0 {
		return v.add_variable()
	}
	return v.variables[rand.Intn(len(v.variables))]
}

type Fuzzer struct {
	Lang        languageRules
	accumulator strings.Builder
	Variables   Variables
}

func New(lang languageRules) *Fuzzer {
	return &Fuzzer{
		Lang:        lang,
		accumulator: strings.Builder{},
		Variables:   Variables{},
	}
}

func (fuzzer *Fuzzer) Fuzz() string {
	// fuzzer.appendExpressions(fuzzer.Lang["<assign>"]) // Add variable to ensure at least one variable is present
	fuzzer.appendExpressions(fuzzer.Lang["<program>"])
	return fuzzer.String()
}

func (fuzzer *Fuzzer) String() string {
	return fuzzer.accumulator.String()
}

func getRuleProbabilistic(expressions []expression) expression {
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

func shuffle(expressions []expression) []expression {

	rand.Shuffle(len(expressions), func(i, j int) {
		expressions[i], expressions[j] = expressions[j], expressions[i]
	})

	return expressions
}

func getTerminalExpressionIfExist(expressions []expression) (expression, bool) {

	for _, expression := range shuffle(expressions) {
		_, hasNonTerminal := getFirstNonTerminal(expression.output)

		if !hasNonTerminal {
			return expression, true
		}
	}

	return expression{"", 0.0, 0}, false
}

var depth = 0
var maxDepth = 5000

func (fuzzer *Fuzzer) appendExpressions(expressions []expression) {
	// Choose a rule at random
	var output string

	depth++
	// fmt.Println(depth * 2)
	if maxDepth < depth {
		// If we reach max depth, probability decrease maxdepth with 1
		// maxDepth = maxDepth - rand.Intn(10)/10
		maxDepth = maxDepth - 1
		output = selectCheapestExpression(expressions).output
	} else {
		output = getRuleProbabilistic(expressions).output
	}

	// Look at special rules $INT$, $ID$, $ID_AS$
	output = fuzzer.replaceSpecialTerminals(output)

	/*
		Loop repeatingly for a rule encapsulated by '<X>'.
		If found, look first at language rules
		Else look at special rules

		Otherwise move one character forward.
	*/
	for {
		next_non_terminal := strings.Index(output, "<")

		// There is no non terminals left
		if next_non_terminal == -1 {
			fuzzer.accumulator.WriteString(output)
			break
		}

		// There is content before next non terminal
		if next_non_terminal != 0 {
			fuzzer.accumulator.WriteString(output[:next_non_terminal])
			output = output[next_non_terminal:]
			continue
		}

		nonTerminalRule, hasNonTerminalRule := getFirstNonTerminal(output)

		// If < is present, but no corresponding >, then add remaining
		if !hasNonTerminalRule {
			fuzzer.accumulator.WriteString(output)
			break
		}

		if fuzzer.processNonTerminalRule(nonTerminalRule) {
			output = output[len(nonTerminalRule):]
			continue
		}

		// If nothing works, move one character forward
		fuzzer.accumulator.WriteString(output[:1])
		output = output[1:]
	}
	depth--
}

func selectCheapestExpression(expressions []expression) expression {

	cheapest := expressions[0]
	for _, expression := range expressions {
		if expression.cost < cheapest.cost {
			cheapest = expression
		}
	}
	return cheapest
}

func (fuzzer *Fuzzer) processNonTerminalRule(nonTerminalRule string) bool {

	repeatRule := 1
	if strings.Contains(nonTerminalRule, "*") {
		nonTerminalRule = strings.Replace(nonTerminalRule, "*", "", -1)
		repeatRule = rand.Intn(5) + 1
	}

	// Look at lang-based rules
	nextExpressions, hasExpressions := fuzzer.Lang[nonTerminalRule]

	if hasExpressions {
		for i := 0; i < repeatRule; i++ {
			fuzzer.appendExpressions(nextExpressions)
		}
		return true
	}

	return false
}

func (fuzzer *Fuzzer) replaceSpecialTerminals(output string) string {
	for {
		dollar := strings.Index(output, "$")

		if dollar == -1 {
			break
		}

		if strings.Contains(output, "$ID_AS$") {
			output = strings.Replace(output, "$ID_AS$", fuzzer.Variables.add_variable(), 1)
			continue
		}

		output = strings.Replace(output, "$INT$", strconv.Itoa(rand.Intn(10000)), 1)
		output = strings.Replace(output, "$ID$", fuzzer.Variables.get_variable(), 1)
	}

	return output
}

func getFirstNonTerminal(output string) (string, bool) {
	lefthand := strings.Index(output, "<")
	righthand := strings.Index(output, ">")

	if lefthand == -1 || righthand == -1 || righthand < lefthand {
		return "", false
	}

	return output[lefthand : righthand+1], true
}
