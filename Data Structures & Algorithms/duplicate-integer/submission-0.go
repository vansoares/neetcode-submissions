func hasDuplicate(nums []int) bool {
    m := map[int]bool{}
    for _, v := range nums {
        if isPresent := m[v]; isPresent {
            return true
        }
        m[v] = true
        
    }
    return false
}
