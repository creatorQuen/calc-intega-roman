package numbers

import "strings"

func SeparteArgumentsByOperand(str string) []string {
	result := strings.ReplaceAll(str, " ", "")
	operand := findOperand(str)

	if operand == 0 {
		return nil
	}

	arguments := strings.Split(result, string(operand))
	return append(arguments, string(operand))

}

func findOperand(str string) rune {
	var operand rune
	//sliceOperands := []string{"+", "-", "*", "/"}
	sliceOperands := []rune{'+', '-', '*', '/'}
	for _, v := range str {
		for _, vv := range sliceOperands {
			if v == rune(vv) {
				operand = vv
				return operand
			}
		}
	}

	return 0
}
