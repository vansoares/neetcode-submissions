func isValid(s string) bool {
    endMap := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	stack := []string{}

	for _, v := range s {
		str := string(v)
		if str == "{" || str == "(" || str == "[" {
			stack = append(stack, str)
			continue
		}
		if len(stack) == 0 { return false }
		if stack[len(stack)-1] != endMap[str] {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
