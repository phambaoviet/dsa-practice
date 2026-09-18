func containsDuplicate(nums []int) bool {
    seen := make(map[int]struct{})

    for  _, num := range nums {
        if _, found := seen[num]; found {
            return true
        }
        seen[num] = struct{}{}
    }
    return false
}