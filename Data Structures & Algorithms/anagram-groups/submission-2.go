func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		r := []rune(v)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j]})
		k := string(r)

		m[k] = append(m[k], v)
	}

	response := make([][]string, 0, len(m))

	for _, v := range m {
		response = append(response, v)
	}

	return response
}
