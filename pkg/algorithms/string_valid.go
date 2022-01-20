package algorithms

type stack []string

func (s *stack) pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	} else {
		lastIndex := len(*s) - 1
		element := (*s)[lastIndex]
		*s = (*s)[:lastIndex]
		return element, true
	}
}

func (s *stack) push(val string) {
	*s = append(*s, val)
}

func IsStringValid(s string) bool {
	opens := make(stack, 0)
	for i, v := range s {
		isOpen := string(v) == "{" || string(v) == "[" || string(v) == "("
		if i == 0 && !isOpen {
			return false
		}
		if isOpen {
			opens.push(string(v))
		} else {
			open, ok := opens.pop()
			if !ok || !isCloseValid(open, string(v)) {
				return false
			}
		}
	}
	return len(opens) == 0
}

func isCloseValid(open, close string) bool {
	closes := map[string]string{"{": "}", "(": ")", "[": "]"}
	return closes[open] == close
}
