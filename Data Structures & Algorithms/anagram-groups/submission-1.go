func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		k := sortWord(v)
		m[k] = append(m[k], v)
	}

	response := make([][]string, 0, len(m))

	for _, v := range m {
		response = append(response, v)
	}

	return response
}

func sortWord(w string) string {
	r := []rune(w)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j]})
	result := string(r)

	return result
}