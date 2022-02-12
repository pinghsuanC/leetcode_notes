package solutions

func isAnagram(s string, t string) bool {
	n, m, tmp := len(s), len(t), make([]int, 26)
	if n != m {
		return false
	}
	for i := 0; i < n; i++ {
		tmp[s[i]-'a']++
		tmp[t[i]-'a']--
	}
	temp := 0
	for _, v := range tmp {
		if v < 0 {
			return false
		}
		temp += v
	}
	return temp == 0
}