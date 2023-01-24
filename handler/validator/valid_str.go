package validator

import "unicode"

func ValidateStr(str string) bool {
	strRune := []rune(str)

	//проверяем первый символ на верхний регистр
	if unicode.IsLower(strRune[0]) {
		return false
	}

	//проверяем посимвольно является ли str строкой
	for _, symb := range strRune {
		if unicode.IsDigit(symb) {
			return false
		}
	}

	return true
}
