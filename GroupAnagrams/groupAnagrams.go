package solutions

func groupAnagrams(strs []string) [][]string {
	length := len(strs)
	result := [][]string{}

	if length == 0 {
		return result
	}

	m := make(map[string][]string)

	for _, s := range strs {
		hashS := hash(s)
		m[hashS] = append(m[hashS], s)
	}

	for _, arr := range m {
		result = append(result, arr)
	}

	return result
}

func hash(str string) string {
	freq := [26]byte{}

	for _, ch := range str {
		freq[ch-'a']++
	}

	return string(freq[0:])
}
