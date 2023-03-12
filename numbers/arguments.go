package numbers

import (
	"calc-intega-roman/lib"
	"strings"
)

func CheckArgumentsFromOneToTen(num1 int, num2 int) bool {
	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		return false
	}
	return true
}

func SeparteArgumentsByOperator(str string) ([]string, error) {
	result := strings.ReplaceAll(str, " ", "")
	operator, err := findOperator(str)
	if err != nil {
		return nil, err
	}

	arguments := strings.Split(result, string(operator))
	return append(arguments, string(operator)), nil

}

func findOperator(str string) (rune, error) {
	var operator rune
	countOper := 0
	//sliceOperands := []string{"+", "-", "*", "/"}
	sliceOperators := []rune{'+', '-', '*', '/'}
	for _, v := range str {
		for _, vv := range sliceOperators {
			if v == vv {
				countOper++
				if countOper == 2 {
					return 0, lib.ErrTooMuchOperators
				}
				operator = vv
			}
		}
	}

	if operator == 0 {
		return 0, lib.ErrNotFoundOperator
	}
	return operator, nil
}
