func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, v := range strs {
		r := []rune(v)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j]})
		result := string(r)
		if _, ok := m[result]; !ok {
			m[result] = []string{v}
			continue
		}
		m[result] = append(m[result], v)
	}

	response := [][]string{}

	for _, v := range m {
		response = append(response, v)
	}

	return response
}