package validator

func ValidateValue(str string) (string, bool) {
	strRune := []rune(str)

	if len(str) != 0 && strRune[0] == '=' {
		return "string", true
	}

	if ValidateNum(str) {
		return "int", true
	}

	return "", false
}
