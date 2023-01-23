package validator

func ValidateSequence(seq []string) bool {
	for ind := 0; ind < len(seq)-1; ind++ {
		if seq[ind] >= seq[ind+1] {
			return false
		}
	}
	return true
}
