package 数组


func dominantIndex(nums []int) int {
	index, max := MaxNums(nums)
	for _, v := range nums {
		if v == max {
			continue
		}
		if v * 2 > max {
			return -1
		}
	}
	return index
}

func MaxNums(nums []int) (int, int) {
	maxNum := nums[0]
	ans := 0
	for k, v := range nums {
		if v > maxNum {
			ans = k
			maxNum = v
		}
	}
	return ans, maxNum
}
