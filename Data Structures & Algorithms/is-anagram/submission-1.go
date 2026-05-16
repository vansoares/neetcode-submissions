func isAnagram(s string, t string) bool {
	m := map[rune]int{}

	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v] += 1
			continue
		}
		m[v] = 1
	}

	for _, v := range t {
		count, ok := m[v]
		if !ok && count == 0 {
			return false
		}
		
		m[v] -= 1

	}

	for _, count := range m {
		if count > 0 {
			return false
		}
	}

	return true
}
