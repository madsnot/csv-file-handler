package validator

import "unicode"

func ValidateNum(str string) bool {
	for _, symb := range str {
		if unicode.IsLetter(symb) || unicode.IsSymbol(symb) {
			return false
		}
	}

	return true
}
