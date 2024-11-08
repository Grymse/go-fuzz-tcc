package main

import (
	"fmt"

	"github.com/Grymse/go-fuzz-tcc/fuzzer"
)

func main() {
	program := fuzzer.GenerateCProgram()

	fuzzer.WriteToFile(program, "out/test_program.out")

	fmt.Println(program)
}
