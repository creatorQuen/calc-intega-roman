package main

import (
	"calc-intega-roman/numbers"
	"fmt"
)

func main() {
	fmt.Println(numbers.TypeNumber("R"))
	fmt.Println(numbers.SeparteArgumentsByOperand("3 / 3 +3tu"))

}
