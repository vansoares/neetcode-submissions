func twoSum(nums []int, target int) []int {
    m:= map[int]int{}
    var response []int 

    for i, v := range nums {
        result := target - v

        if i2, ok := m[result]; ok {
            response = []int{i2, i}
            break
        }
        m[v] = i
    }

    return response
}
