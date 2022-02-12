package solutions

func lengthOfLongestSubstring(s string) int {
    ans := 0
    
    for i:=0; i<len(s); i++ {
        curMaxLen := 0
        m := make(map[byte]bool)
        for j:=i; j<len(s); j++{
            if _, ok := m[s[j]]; !ok {
                curMaxLen++
                m[s[j]] = true
                continue
            }
            break;
        }
        if ans < curMaxLen {
            ans = curMaxLen
        }
    }
    return ans
}
