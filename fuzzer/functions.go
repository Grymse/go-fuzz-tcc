package fuzzer

import (
	"math/rand"
	"strings"
)

type Function struct {
	Variables int
	Name      string
}

type Functions struct {
	functions []Function
	Generator func() string
}

func (fs *Functions) init() {
	fs.functions = make([]Function, 0)
	fs.Generator = default_variable_generator
}

func (fs *Functions) variable_exists(name string) bool {
	for _, function := range fs.functions {
		if function.Name == name {
			return true
		}
	}
	return false
}

func (fs *Functions) declare_function_grammar() string {

	name := fs.Generator()

	for fs.variable_exists(name) {
		name = fs.Generator()
	}

	f := Function{
		Variables: rand.Intn(4),
		Name:      fs.Generator(),
	}
	fs.functions = append(fs.functions, f)

	acc := strings.Builder{}
	acc.WriteString("<type_specifier> " + f.Name + "(")

	for i := 0; i < f.Variables; i++ {

		acc.WriteString("<type_specifier> $ID_DECL$")
		if i < f.Variables-1 {
			acc.WriteString(", ")
		}
	}

	acc.WriteString(")")

	return acc.String()
}

func (fs *Functions) call_function_grammar() string {
	if len(fs.functions) == 0 {
		return "$INT$"
	}

	f := fs.functions[rand.Intn(len(fs.functions))]

	acc := strings.Builder{}
	acc.WriteString(f.Name + "(")

	for i := 0; i < f.Variables; i++ {
		acc.WriteString("$ID$")
		if i < f.Variables-1 {
			acc.WriteString(", ")
		}
	}

	acc.WriteString(")")

	return acc.String()
}
