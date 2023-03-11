package main

import (
	"calc-intega-roman/numbers"
	"fmt"
)

func main() {
	fmt.Println(numbers.NaturalNumberToRoman(100))
	fmt.Println(numbers.NaturalNumberToRoman(99))
	fmt.Println(numbers.NaturalNumberToRoman(1 + 0))
	fmt.Println(numbers.RomanToNaturalNumber("VII"))
	fmt.Println(numbers.RomanToNaturalNumber("VI"))
	fmt.Println(numbers.RomanToNaturalNumber("C"))
}
