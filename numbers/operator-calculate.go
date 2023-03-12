package numbers

import "calc-intega-roman/lib"

func CalculateTwoOperandByOperator(number1 int, number2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return number1 + number2, nil
	case "-":
		return number1 - number2, nil
	case "*":
		return number1 * number2, nil
	case "/":
		return number1 / number2, nil
	default:
		return 0, lib.ErrNotFoundSuchOperator
	}
}
