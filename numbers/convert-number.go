package numbers

import "errors"

func NaturalNumberToRoman(num int) (string, error) {
	if num < 0 {
		return "", errors.New("number could not be negative or zero")
	}

	valueRoman := ""
	i := 0

	roman := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanValues := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	for num > 0 {
		for romanValues[i] > num {
			i++
		}
		valueRoman += roman[i]
		num -= romanValues[i]
	}

	return valueRoman, nil
}

func RomanToNaturalNumber(str string) (int, error) {
	if str == "" {
		return 0, errors.New("string is empty")
	}

	number := 0
	symbolMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
	}

	i := 0
	for ; i < len(str); i++ {
		char := string(str[i])
		if char == "I" {
			if i < len(str)-1 {
				next := string(str[i+1])
				if next == "V" || next == "X" {
					number += symbolMap[next] - symbolMap[char]
					i++
					continue
				}
			}
		}
		if char == "X" {
			if i < len(str)-1 {
				next := string(str[i+1])
				if next == "L" || next == "C" {
					number += symbolMap[next] - symbolMap[char]
					i++
					continue
				}
			}
		}
		number += symbolMap[char]
	}

	return number, nil
}
