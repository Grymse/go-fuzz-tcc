package main

import (
	"fmt"

	"github.com/Grymse/go-fuzz-tcc/fuzzer"
)

func main() {
	lang := fuzzer.Languages.TinyC
	fmt.Println(fuzzer.Fuzz(lang))
}
