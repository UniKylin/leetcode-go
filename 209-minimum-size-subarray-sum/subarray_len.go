package leetcode

// link		: https://leetcode.cn/problems/minimum-size-subarray-sum/
// Author	: Kylin
// Date		: 2024-12-19

// 滑动窗口
// 当然还可以继续优化
func MinSubArrayLen(target int, nums []int) (res int) {
	res = len(nums) + 1
	sum := 0
	left := 0

	for right, num := range nums {
		sum += num

		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if res <= len(nums) {
		return res
	}
	return 0
}
