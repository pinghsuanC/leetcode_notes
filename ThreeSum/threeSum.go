package solutions

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i, n := range nums {
		if i > 0 && n == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[l] + n + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				result = append(result, []int{n, nums[l], nums[r]})
				l++
				for len(nums) > l && nums[l] == nums[l-1] {
					l++
				}
			}
		}

	}

	return result
}