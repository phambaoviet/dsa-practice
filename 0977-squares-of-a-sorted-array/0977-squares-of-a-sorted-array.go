func sortedSquares(nums []int) []int {
    n := len(nums)
    if n == 0 {
        return nil
    }
    result := make([]int, n)
    left := 0
    right := n - 1
    p := n - 1
    for left <= right {
        if (nums[left]*nums[left] < nums[right]*nums[right]){
            result[p] = nums[right]*nums[right]
            right--
        } else {
            result[p] = nums[left]*nums[left]
            left++
        }
        p--
    }
    return result
}