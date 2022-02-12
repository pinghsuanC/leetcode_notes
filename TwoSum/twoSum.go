package solutions

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var result []int

	for ind1, val := range nums {
		ind2, ok := m[target-val]
		if ok {
			result = append(result, ind1)
			result = append(result, ind2)
			break
		}
		m[val] = ind1
	}
	return result
}