package numbers

import "strconv"

func CalculateByOperator(slice []string) int {
	number1, _ := strconv.Atoi(slice[0])
	number2, _ := strconv.Atoi(slice[0])

	switch slice[len(slice)-1] {
	case "+":
		return number1 + number2
	case "-":
		return number1 - number2
	case "*":
		return number1 * number2
	case "/":
		return number1 / number2
	}

	return -999999999999999999
}
