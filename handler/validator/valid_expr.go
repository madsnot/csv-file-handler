package validator

import "unicode"

func ValidateExpr(str string) (string, bool) {
	strRune := []rune(str)
	if unicode.IsLetter(strRune[0]) && (strRune[0] != '=' || strRune[0] == '#') {
		return "", false
	}
	if ValidateNum(str) {
		return "int", true
	}
	return "string", true
}
