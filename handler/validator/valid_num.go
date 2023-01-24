package validator

import "unicode"

func ValidateNum(str string) bool {
	//проверяем является ли str числом
	for _, symb := range str {
		if !unicode.IsDigit(symb) {
			return false
		}
	}

	return true
}
