func intersection(nums1 []int, nums2 []int) []int {
    set := make(map[int]struct{})
    resultSet := make(map[int]struct{})
    result := []int{}
    for _, num := range nums1 {
        set[num] = struct{}{}
    }
    for _, num := range nums2{
        _, found := set[num]
        _, alreadyAdd := resultSet[num]
        if found && !alreadyAdd {
            resultSet[num] = struct{}{}
            result = append(result,num)
        }
    }
    return result
}