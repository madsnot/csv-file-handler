package validator

import "unicode"

func ValidateStr(str string) bool {
	strRune := []rune(str)
	if unicode.IsLower(strRune[0]) {
		return false
	}
	for _, symb := range strRune {
		if unicode.IsDigit(symb) {
			return false
		}
	}

	return true
}
