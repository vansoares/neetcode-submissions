func isPalindrome(s string) bool {
	result := getAlphanumerical(s)
	result = strings.Replace(result, " ", "", -1)
	result = strings.ToLower(result)
	left, right := 0, len(result)-1

	for left < right {
		if result[left] != result[right] {
			return false
		}
		left+=1
		right-=1
	}
	return true
}

func getAlphanumerical(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return re.ReplaceAllString(s, "")
}