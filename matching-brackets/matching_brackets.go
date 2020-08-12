package brackets

var matchingBrackets = map[rune]rune{
    ']':'[',
    ')':'(',
    '}':'{',
}

func Bracket(input string) bool {
	var startBrackets []rune

	for _, r := range input {
		if isStartBracket(r) {
			startBrackets = append(startBrackets, r)
			continue
		}

		if !isClosingBracket(r) {
			continue
		}

        n := len(startBrackets) 

        // If find an closing bracket but we don't have
        // any start brackets then it can't be right.
		if n < 1 {
			return false
		}

		if pop := startBrackets[n-1]; matchingBrackets[r] != pop {
            return false
        }
		startBrackets = startBrackets[:n-1]
	}

	return len(startBrackets) < 1
}

func isStartBracket(r rune) bool {
	if r == '[' || r == '(' || r == '{' {
		return true
	}
	return false
}

func isClosingBracket(r rune) bool {
	if r == ']' || r == ')' || r == '}' {
		return true
	}
	return false
}
