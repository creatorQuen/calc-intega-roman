package numbers

import (
	"errors"
	"strings"
)

func CheckArgumentsFromOneToTen(num1 int, num2 int) bool {
	if num1 <= 0 || num1 > 10 || num2 >= 0 || num2 > 10 {
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
					return 0, errors.New("too much operators")
				}
				operator = vv
			}
		}
	}

	if operator == 0 {
		return 0, errors.New("not founnd this operator")
	}
	return operator, nil
}
