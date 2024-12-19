package leetcode

import (
	"fmt"
	"testing"
)

// link		: https://leetcode.cn/problems/minimum-size-subarray-sum/
// Author	: Kylin
// Date		: 2024-12-19

func TestMinSubArrayLen(t *testing.T) {
	// 测试用例集合
	tests := []struct {
		target int
		nums   []int
		expect int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},               // 子数组 [4,3] 的长度是 2
		{4, []int{1, 4, 4}, 1},                        // 子数组 [4] 的长度是 1
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},        // 无法达到目标值
		{15, []int{1, 2, 3, 4, 5}, 5},                 // 整个数组满足条件
		{5, []int{5}, 1},                              // 单个元素满足条件
		{8, []int{}, 0},                               // 空数组
		{15, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2}, // 子数组 [7,8] 或 [8,7]
	}

	// 执行测试
	for i, test := range tests {
		result := MinSubArrayLen(test.target, test.nums)
		fmt.Printf("Test %d: target=%d, nums=%v, expect=%d, got=%d\n",
			i+1, test.target, test.nums, test.expect, result)
		if result != test.expect {
			fmt.Println("-> Test failed!")
		} else {
			fmt.Println("-> Test passed!")
		}
	}
}
