package fuzzer

import (
	"math/rand"
	"strconv"
)

type Variable struct {
	Scope    int
	Name     string
	TypeRule TypeRule
	IsArray  bool
}

type Variables struct {
	variables []Variable
	Generator func() string
	scope     int
}

type TypeRule int

const (
	ANY = iota
	VAR
	CONST
)

func (v *Variables) init() {
	v.variables = make([]Variable, 0)
	v.Generator = default_variable_generator
	v.scope = 0
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

func (v *Variables) variable_exists(name string) bool {
	for _, variable := range v.variables {
		if variable.Name == name {
			return true
		}
	}
	return false
}

func (v *Variables) add_variable(varTypeRule TypeRule, isArray bool) string {
	name := v.Generator()

	for v.variable_exists(name) {
		name = v.Generator()
	}

	v.variables = append(v.variables, Variable{
		Scope:    v.scope,
		Name:     name,
		TypeRule: varTypeRule,
		IsArray:  isArray,
	})
	return string(name)
}

func (v *Variables) get_variable(varTypeRule TypeRule) string {
	if len(v.variables) == 0 {
		panic("No variables available")
	}

	if varTypeRule == ANY {
		variable := v.variables[rand.Intn(len(v.variables))]

		if variable.IsArray {
			return variable.Name + "[" + strconv.Itoa(rand.Intn(100)) + "]"
		}

		return variable.Name
	}

	if !v.has_type(varTypeRule) {
		panic("No variables of type " + strconv.Itoa(int(varTypeRule)) + " available")
	}

	variable := v.variables[rand.Intn(len(v.variables))]

	for variable.TypeRule != varTypeRule {
		variable = v.variables[rand.Intn(len(v.variables))]
	}

	if variable.IsArray {
		return variable.Name + "[" + strconv.Itoa(rand.Intn(100)) + "]"
	}

	return variable.Name
}

func (v *Variables) get_types() (m map[TypeRule]bool) {
	m = make(map[TypeRule]bool)

	for _, variable := range v.variables {
		m[variable.TypeRule] = true
	}

	return
}

func (v *Variables) has_type(t TypeRule) bool {
	return v.get_types()[t]
}

func (v *Variables) increment_scope() {
	v.scope++
}

func (v *Variables) decrement_scope() {
	v.scope--
	newVariables := make([]Variable, 0)
	for _, variable := range v.variables {
		if variable.Scope <= v.scope {
			newVariables = append(newVariables, variable)
		}
	}
	v.variables = newVariables
}
