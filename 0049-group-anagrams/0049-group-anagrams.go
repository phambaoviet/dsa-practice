func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)
    result := [][]string{}
    
    for _, word := range strs {
        count := [26]int{}
        for _, char := range word {
            count[char-'a']++
        }
        groups[count] = append(groups[count], word)
    }
    for _, group := range groups {
        result = append(result, group)
    }
    return result
}