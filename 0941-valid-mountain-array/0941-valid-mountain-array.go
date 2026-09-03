func validMountainArray(arr []int) bool {
    if len(arr) < 3 {
        return false
    }
    i := 0
    for i+1 < len(arr) && arr[i] < arr[i+1]{
        i++
    }
    if i == 0 || i == len(arr) - 1 {
        return false
    }
    for i+1 < len(arr) && arr[i] > arr[i+1]{
        i++
    }
    if i == len(arr) - 1{
        return true
    }
    return false
}