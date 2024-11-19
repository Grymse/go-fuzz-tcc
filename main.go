package main

import (
	"fmt"
	"math/rand"

	"github.com/Grymse/go-fuzz-tcc/fuzzer"
)

func one_letter_generator() string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	return string(letters[rand.Intn(len(letters))])
}

func main() {
	fuzzer := fuzzer.New(fuzzer.Languages.CLN)
	fuzzer.Variables.Generator = one_letter_generator
	fmt.Println(fuzzer.Fuzz())
}
