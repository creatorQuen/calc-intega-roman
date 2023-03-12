package numbers

import (
	"calc-intega-roman/lib"
)

func TypeNumber(str string) (string, error) {
	if str == "" {
		return "", lib.ErrStringIsEmpty
	}

	countNumber := 0
	for _, v := range str {
		if rune(v) >= 48 && rune(v) <= 57 {
			countNumber++
			continue
		}
		break
	}
	if countNumber == len(str) {
		return lib.DecimalType, nil
	}

	countRoman := 0
	sliceOfRoman := []rune{'I', 'V', 'X', 'L', 'C'}
	for _, v := range str {
		for _, vv := range sliceOfRoman {
			if v == vv {
				countRoman++
				break
			}
		}
	}
	if countRoman == len(str) {
		return lib.RomanType, nil
	}

	return "", lib.ErrNotSameTypeOfNumber
}
