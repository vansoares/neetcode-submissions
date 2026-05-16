func isAnagram(s string, t string) bool {
	m := map[string]int{}

	for _, v := range s {
		strValue := string(v)
		if _, ok := m[strValue]; ok {
			m[strValue] += 1
			continue
		}
		m[strValue] = 1
	}

	for _, v := range t {
		strValue := string(v)
		count, ok := m[strValue]
		if !ok && count == 0 {
			return false
		}
		
		m[strValue] -= 1

	}

	for _, v := range m {
		if v > 0 {
			return false
		}
	}

	return true
}
