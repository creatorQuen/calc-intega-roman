package lib

import "errors"

var StartMessage = "--Calculate two numbers from 1 to 10,\n--or roman number from I to X (operations: +,-,*,/)\n--press cnt+C or write 'exit' to escape"

var (
	DecimalType = "decimal"
	RomanType   = "roman"
)

var (
	ErrStringIsEmpty                = errors.New("string is empty")
	ErrNotSameTypeOfNumber          = errors.New("not same type of number,\n put decimal natrual number or roman number")
	ErrNotFoundSuchOperator         = errors.New("not found such operator (input: +,-,*,/)")
	ErrNumberNotNegativeOrZero      = errors.New("number could not be negative or zero")
	ErrTooMuchOperators             = errors.New("too much operators")
	ErrNotFoundOperator             = errors.New("could not found that operator")
	ErrNumberMustBeFromOneToTen     = errors.New("numbers must be from 1 to 10")
	ErrRomanNumberalsMustBeFromIToX = errors.New("roman numbers must be from I to X")
	ErrRomanNumeralsNotNegative     = errors.New("roman numerals not have negative value or zero")
	ErrNotFoundTypeOfNumbers        = errors.New("not found type of number")
	ErrTypeOfNumberNotEquel         = errors.New("type of number is not equel")
)
