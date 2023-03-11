package numbers

import "errors"

func TypeNumber(str string) (string, error) {
	if str == "" {
		return "", errors.New("string is empty")
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
		return "decimal", nil
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
		return "roman", nil
	}

	return "", errors.New("not same type of number,\n put decimal natrual number or roman number")
}
