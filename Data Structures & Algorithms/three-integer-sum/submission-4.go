func threeSum(nums []int) [][]int {
	response := make([][]int, 0, len(nums))
	m := map[string]int{}
	for i:=0; i<len(nums); i++ {
		for j:=i+1; j<len(nums); j++ {
			for k:=j+1; k<len(nums); k++ {
				r := nums[i]+nums[j]+nums[k]
				if r == 0 {
					fmt.Println(r)

					triplet := []int{nums[i], nums[j], nums[k]}
					sort.Slice(triplet, func(i, j int) bool { return triplet[i] < triplet[j]})
					
					tripletStr := fmt.Sprintf("%v", triplet)
					if _, ok := m[tripletStr]; !ok {
						response = append(response, triplet)
						m[tripletStr] = 1
					}
				}
			}
		}
	}

	return response
}
