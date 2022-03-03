package main
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
        return 0
    }
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast-1] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}