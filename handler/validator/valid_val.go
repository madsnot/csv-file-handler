package validator

func ValidateValue(str string) (string, bool) {
	strRune := []rune(str)

	if strRune[0] == '=' {
		return "string", true
	}

	if ValidateNum(str) {
		return "int", true
	}

	return "", false
}
