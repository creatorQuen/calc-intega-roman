package main

import (
	"bufio"
	"calc-intega-roman/lib"
	"calc-intega-roman/numbers"
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
		fmt.Println(lib.StartMessage)
		inputString, _ := readerIn.ReadString('\n')
		inputString = strings.TrimSpace(inputString)
		if inputString == "exit" || inputString == "quit" {
			os.Exit(0)
		}
		if inputString == "" {
			fmt.Println(lib.ErrStringIsEmpty, "\n")
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
				case lib.DecimalType:
					number1, _ := strconv.Atoi(arguments[0])
					number2, _ := strconv.Atoi(arguments[1])
					if !numbers.CheckArgumentsFromOneToTen(number1, number2) {
						shutdownAndShowError(lib.ErrNumberMustBeFromOneToTen)
					}

					result, err := numbers.CalculateTwoOperandByOperator(number1, number2, arguments[2])
					if err != nil {
						shutdownAndShowError(err)
					}
					fmt.Println(result)
				case lib.RomanType:
					number1, err := numbers.RomanToNaturalNumber(arguments[0])
					if err != nil {
						shutdownAndShowError(err)
					}
					number2, err := numbers.RomanToNaturalNumber(arguments[1])
					if err != nil {
						shutdownAndShowError(err)
					}
					if !numbers.CheckArgumentsFromOneToTen(number1, number2) {
						shutdownAndShowError(lib.ErrRomanNumberalsMustBeFromIToX)
					}

					result, err := numbers.CalculateTwoOperandByOperator(number1, number2, arguments[2])
					if err != nil {
						shutdownAndShowError(err)
					}
					if result <= 0 {
						shutdownAndShowError(lib.ErrRomanNumeralsNotNegative)
					}

					romanValue, err := numbers.NaturalNumberToRoman(result)
					if err != nil {
						shutdownAndShowError(err)
					}

					fmt.Println(romanValue)
				default:
					shutdownAndShowError(lib.ErrNotFoundTypeOfNumbers)
				}
			} else {
				shutdownAndShowError(lib.ErrTypeOfNumberNotEquel)
			}
		}
	}
}

func shutdownAndShowError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(0)
}
