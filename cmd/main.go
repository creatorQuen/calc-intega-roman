package main

import (
	"calc-intega-roman/numbers"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}()

	//readerIn := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Calculate two numbers from 0 to 10,\n or roman number from I to X (operations: +,-,*,/)")
		//inputString, _ := readerIn.ReadString('\n')
		inputString := "1 + 1 + 1"
		inputString = strings.TrimSpace(inputString)
		if inputString == "exit" || inputString == "quit" {
			os.Exit(0)
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
					fmt.Println(numbers.CalculateByOperator(arguments))
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
