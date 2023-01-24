package validator

func ValidateSequence(seq []string) bool {
	//проверяем последовательность, что она по возрастанию
	for ind := 0; ind < len(seq)-1; ind++ {
		if seq[ind] >= seq[ind+1] {
			return false
		}
	}
	return true
}
