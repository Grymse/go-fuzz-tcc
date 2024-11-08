package fuzzer

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Grymse/go-fuzz-tcc/tinyc"
)

func randomExpr() tinyc.Expression {
	// Possible terms
	terms := []tinyc.Expression{
		tinyc.TermExpr{ID: "a"},
		tinyc.TermExpr{ID: "b"},
		tinyc.TermExpr{ID: "c"},
		tinyc.TermExpr{Integer: "42"},
	}

	// Operators for sum expressions
	ops := []string{"+", "-", "<"}
	randOp := ops[rand.Intn(len(ops))]

	// Left and right terms
	left := terms[rand.Intn(len(terms))]
	right := terms[rand.Intn(len(terms))]

	// Return a sum expression
	return tinyc.SumExpr{
		Left:  left,
		Op:    randOp,
		Right: right,
	}
}

func randomStatement() tinyc.Statement {
	switch rand.Intn(4) { // Added a case for just an empty statement ";"
	case 0:
		return tinyc.ExprStatement{
			Expression: randomExpr(),
		}
	case 1:
		return tinyc.IfStatement{
			Condition: randomExpr(),
			ThenStmt:  randomStatement(),
			ElseStmt:  randomStatement(),
		}
	case 2:
		return tinyc.WhileStatement{
			Condition: randomExpr(),
			Body:      randomStatement(),
		}
	case 3:
		return tinyc.ExprStatement{
			Expression: tinyc.TermExpr{ID: "a"}, // Just an empty statement, e.g., ";"
		}
	}
	return nil
}

func GenerateCProgram() string {
	rand.Seed(time.Now().UnixNano())

	// Generate random statements for the program (this will be a single line)
	var statements []tinyc.Statement
	for i := 0; i < rand.Intn(3)+1; i++ { // Generate 1-3 random statements for the one-liner
		statements = append(statements, randomStatement())
	}

	// Create the program body as a one-liner
	cCode := ""
	for _, stmt := range statements {
		cCode += " " + stmt.Generate() // Add each statement in one line
	}

	// Return the one-liner C program
	return "{" + cCode + "}"
}

// Write the generated C program to a file
func WriteToFile(program string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(program)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
