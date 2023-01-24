package validator

import "unicode"

func ValidateValue(str string) (string, bool) {
	strRune := []rune(str)

	//проверяем не пуста ли ячейка, и первый символ, если символ =,
	//то в ячейке хранится выражение
	if len(str) != 0 && (!unicode.IsGraphic(strRune[0]) || strRune[0] == '=') {
		return "string", true
	}

	//проверяем является ли числом значение в ячейке
	if ValidateNum(str) {
		return "int", true
	}

	return "", false
}
