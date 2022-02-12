package solutions

func search(nums []int, target int) int {
	lIndex := 0
	rIndex := len(nums) - 1

	for lIndex <= rIndex {
		midIndex := (lIndex + rIndex) / 2
		low := nums[lIndex]
		mid := nums[midIndex]
		high := nums[rIndex]
		if mid == target {
			return midIndex
		}
		if low <= mid {
			// first half doesn't contain rotation point
			if target >= low && target < mid {
				// check if the target is in the first half
				rIndex = midIndex - 1
			} else {
				// if it's not in the first half it's in the second half
				lIndex = midIndex + 1
			}
		} else {
			// first half contains the rotation point
			if target > mid && target <= high {
				// check if target is in the second half
				lIndex = midIndex + 1
			} else { // if it's not in the second half it's in the first half
				rIndex = midIndex - 1
			}
		}
	}
	return -1
}