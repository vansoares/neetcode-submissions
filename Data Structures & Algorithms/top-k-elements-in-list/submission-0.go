func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] += 1
			continue
		}
		m[v] = 1
	}

	buckets := make([][]int, len(nums)+1)
	for k, v := range m {
		buckets[v] = append(buckets[v], k)
	}

	response := []int{}
	for i:= len(buckets)-1; i>0; i-- {
		response = append(response, buckets[i]...)
		if len(response) == k {
			break
		}
	}

	return response
}
