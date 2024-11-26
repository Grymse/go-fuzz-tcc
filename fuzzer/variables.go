package fuzzer

import "math/rand"

type Variable struct {
	Scope int
	Name  string
}

type Variables struct {
	variables []Variable
	Generator func() string
	scope     int
}

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

func (v *Variables) add_variable() string {
	name := v.Generator()

	for v.variable_exists(name) {
		name = v.Generator()
	}

	v.variables = append(v.variables, Variable{
		Scope: v.scope,
		Name:  name,
	})
	return string(name)
}

func (v *Variables) get_variable() string {
	if len(v.variables) == 0 {
		panic("No variables available")
	}
	return v.variables[rand.Intn(len(v.variables))].Name
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
