package fuzzer

import (
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

type Fuzzer struct {
	Lang        languageRules
	accumulator strings.Builder
	Variables   Variables
	Functions   Functions
}

func New(lang languageRules) *Fuzzer {
	variables := Variables{}
	variables.init()
	functions := Functions{}
	functions.init()

	return &Fuzzer{
		Lang:        lang,
		accumulator: strings.Builder{},
		Variables:   variables,
		Functions:   functions,
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
	sum := 0.0

	for _, expression := range expressions {
		sum += expression.prob
	}

	probSum := rand.Float64() * sum
	lowerBound := 0.0
	for _, expression := range expressions {
		upperBound := lowerBound + expression.prob
		if lowerBound <= probSum && probSum < upperBound {
			return expression
		}
		lowerBound = upperBound
	}

	return expressions[0]
}

func (fuzzer *Fuzzer) adjustScope(output string) {
	if len(output) == 0 {
		return
	}

	next_scope_increase := strings.Index(output, "{")
	next_scope_decrease := strings.Index(output, "}")
	next_non_terminal := strings.Index(output, "<")

	if next_scope_increase != -1 && (next_scope_increase < next_non_terminal || next_non_terminal == -1) {
		fuzzer.Variables.increment_scope()
	}

	if next_scope_decrease != -1 && (next_scope_decrease < next_non_terminal || next_non_terminal == -1) {
		fuzzer.Variables.decrement_scope()
	}
}

var depth = 0
var wavePeak = 200
var waveValleyMin = 4
var waveValleyMax = 5
var waveValley = 4
var maxWaves = 20
var waveCount = 0
var target = wavePeak


func adjustPeak() {
	if depth < waveValley {
		target = wavePeak
	} else if depth > wavePeak {
		target = rand.Intn(waveValleyMax-waveValleyMin) + waveValleyMin
		waveCount++
	}

	if waveCount >= maxWaves {
		target = 0
	}
}

func (fuzzer *Fuzzer) appendExpressions(expressions []expression) {
	// Choose a rule at random
	var output string

	depth++
	adjustPeak()
	if target < depth {
		output = selectCheapestExpression(expressions).output
	} else {
		output = getRuleProbabilistic(expressions).output
	}

	isSpecialScope := strings.HasPrefix(output, "^")

	if isSpecialScope {
		fuzzer.Variables.increment_scope()
		output = output[1:]
	}

	// Look at special rules $INT$, $ID$, $ID_DECL$
	output = fuzzer.replaceSpecialTerminals(output)

	/*
		Loop repeatingly for a rule encapsulated by '<X>'.
		If found, look first at language rules
		Else look at special rules

		Otherwise move one character forward.
	*/
	for {
		fuzzer.adjustScope(output)
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
	if isSpecialScope {
		fuzzer.Variables.decrement_scope()
	}
}

func selectCheapestExpression(expressions []expression) expression {
	// shuffle expressions
	shuffled := make([]expression, len(expressions))
	perm := rand.Perm(len(expressions))
	for i, v := range perm {
		shuffled[v] = expressions[i]
	}

	cheapest := shuffled[0]
	for _, expression := range shuffled {
		if expression.cost < cheapest.cost {
			cheapest = expression
		}
	}
	return cheapest
}

func (fuzzer *Fuzzer) processNonTerminalRule(nonTerminalRule string) bool {
	if strings.Contains(nonTerminalRule, "!") {
		nonTerminalRule = strings.Replace(nonTerminalRule, "!", "", -1)
		nextExpressions, hasExpressions := fuzzer.Lang[nonTerminalRule]

		if hasExpressions {
			for _, expr := range nextExpressions {
				fuzzer.appendExpressions([]expression{expr})
			}
			return true
		}

		return false
	}

	repeatRule := 1
	if strings.Contains(nonTerminalRule, "*") {
		nonTerminalRule = strings.Replace(nonTerminalRule, "*", "", -1)
		repeatRule = rand.Intn(10) + 2
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

var lorem = "lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."
var currChar = 0

func (fuzzer *Fuzzer) replaceSpecialTerminals(output string) string {
	for {
		dollar := strings.Index(output, "$")

		if dollar == -1 {
			break
		}

		if strings.Contains(output, "$FUNC_DECL$") {
			output = strings.Replace(output, "$FUNC_DECL$", fuzzer.Functions.declare_function_grammar(), 1)
			continue
		}

		if strings.Contains(output, "$ID_DECL$") {
			output = strings.Replace(output, "$ID_DECL$", fuzzer.Variables.add_variable(VAR, false), 1)
			continue
		}

		if strings.Contains(output, "$ID_DECL_C$") {
			output = strings.Replace(output, "$ID_DECL_C$", fuzzer.Variables.add_variable(CONST, false), 1)
			continue
		}

		if strings.Contains(output, "$ID_DECL_ARR_C$") {
			output = strings.Replace(output, "$ID_DECL_ARR_C$", fuzzer.Variables.add_variable(CONST, true), 1)
			continue
		}

		if strings.Contains(output, "$ID_DECL_ARR$") {
			output = strings.Replace(output, "$ID_DECL_ARR$", fuzzer.Variables.add_variable(VAR, true), 1)
			continue
		}

		if strings.Contains(output, "$LOREM$") {
			substring_lorem := lorem[:rand.Intn(len(lorem)-1)]
			output = strings.Replace(output, "$LOREM$", substring_lorem, 1)
			continue
		}

		if strings.Contains(output, "$FUNC$") {
			output = strings.Replace(output, "$FUNC$", fuzzer.Functions.call_function_grammar(), 1)
			continue
		}

		output = strings.Replace(output, "$INT$", strconv.Itoa(rand.Intn(10000)), 1)
		output = strings.Replace(output, "$ID$", fuzzer.Variables.get_variable(ANY), 1)
		output = strings.Replace(output, "$ID_AS$", fuzzer.Variables.get_variable(VAR), 1)
		output = strings.Replace(output, "$FUNC$", fuzzer.Functions.call_function_grammar(), 1)

		if strings.Contains(output, "$CHAR$") {
			chars := "abcdefghijklmnopqrstuvwxyz"
			currChar++
			if currChar >= len(chars) {
				currChar = 0
			}
			output = strings.Replace(output, "$CHAR$", "'"+string(chars[currChar])+"'", 1)
			continue
		}
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
