package main

import (
	"bufio"
	"calc-intega-roman/numbers"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}()

	readerIn := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("--Calculate two numbers from 0 to 10,\n--or roman number from I to X (operations: +,-,*,/)\n--press cnt+C or write 'exit' to escape")
		inputString, _ := readerIn.ReadString('\n')
		//inputString := "1 + 1 + 1"
		//inputString := "10 = 10"
		inputString = strings.TrimSpace(inputString)
		if inputString == "exit" || inputString == "quit" {
			os.Exit(0)
		}
		if inputString == "" {
			fmt.Println("empty string\n")
			continue
		}
		arguments, err := numbers.SeparteArgumentsByOperator(inputString)
		if err != nil {
			shutdownAndShowError(err)
		}

		if arguments != nil {
			type1, err := numbers.TypeNumber(arguments[0])
			if err != nil {
				shutdownAndShowError(err)
			}

			type2, err := numbers.TypeNumber(arguments[1])
			if err != nil {
				shutdownAndShowError(err)
			}

			if type1 == type2 {
				switch type1 {
				case "decimal":
					number1, _ := strconv.Atoi(arguments[0])
					number2, _ := strconv.Atoi(arguments[1])
					if !numbers.CheckArgumentsFromOneToTen(number1, number2) {
						shutdownAndShowError(errors.New("numbers must be form 1 to 10"))
					}

					result, err := numbers.CalculateTwoOperandByOperator(number1, number2, arguments[2])
					if err != nil {
						shutdownAndShowError(err)
					}
					fmt.Println(result)
				case "roman":
					number1, err := numbers.RomanToNaturalNumber(arguments[0])
					if err != nil {
						shutdownAndShowError(err)
					}
					number2, err := numbers.RomanToNaturalNumber(arguments[1])
					if err != nil {
						shutdownAndShowError(err)
					}
					if !numbers.CheckArgumentsFromOneToTen(number1, number2) {
						shutdownAndShowError(errors.New("numbers must be form I to X"))
					}

					result, err := numbers.CalculateTwoOperandByOperator(number1, number2, arguments[2])
					if err != nil {
						shutdownAndShowError(err)
					}
					if result <= 0 {
						shutdownAndShowError(errors.New("in roman not have negative value or zero"))
					}

					romanValue, err := numbers.NaturalNumberToRoman(result)
					if err != nil {
						shutdownAndShowError(err)
					}

					fmt.Println(romanValue)
				default:
					fmt.Println("not found type of number")
				}
			} else {
				shutdownAndShowError(errors.New("type of number is not same"))
			}
		}
	}
}

func shutdownAndShowError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(0)
}
