func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxArea := 0
    for left < right{
        width := right - left
        h := min(height[left], height[right])
        if width * h > maxArea {
            maxArea = width * h
        }
        if height[left] > height[right]{
            right--
        } else {
            left++
        }
    }
    return maxArea
}