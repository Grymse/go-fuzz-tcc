package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Grymse/go-fuzz-tcc/fuzzer"
)

func helloworld() {
	fmt.Println("hello world")
}

func get_sized_array(length int) []int {
	return make([]int, length)
}

func array_to_string(array []int) string {
	var accumulator strings.Builder

	for i := 0; i < len(array); i++ {
		_, err := accumulator.WriteString(strconv.Itoa(array[i]))
		if err != nil {
			return "Didn't work brother"
		}
	}

	return accumulator.String()
}

func compare_array_length(a []int, b []int) bool {
	return len(a) <= len(b)
}

func fill_array(arr *[]int, filler int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = filler
	}
}

func fill_array2(arr []int, filler int) {
	for i := 0; i < len(arr); i++ {
		(arr)[i] = filler
	}
}

func main() {
	helloworld()
	fuzzer.Fuzz()

	a10 := get_sized_array(10)
	a15 := get_sized_array(15)

	fill_array(&a10, 5)
	fill_array2(a15, 10)

	fmt.Println(a15)
	fmt.Println(a10)

	fmt.Println(compare_array_length(a15, a10))
	fmt.Println(array_to_string(a10))
}
