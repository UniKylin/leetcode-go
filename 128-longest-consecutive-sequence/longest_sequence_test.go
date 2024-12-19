package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		// 测试用例 1：正常情况
		// 最长序列为 [1, 2, 3, 4]
		{nums: []int{100, 4, 200, 1, 3, 2}, expected: 4},

		// 测试用例 2：含有重复元素
		// 最长序列为 [1, 2, 3, 4]
		{nums: []int{1, 2, 2, 3, 4}, expected: 4},

		// 测试用例 3：无序数组
		// 最长序列为 [5, 6, 7, 8]
		{nums: []int{5, 6, 3, 20, 7, 8}, expected: 4},

		// 测试用例 4：空数组
		{nums: []int{}, expected: 0},

		// 测试用例 5：单个元素
		{nums: []int{1}, expected: 1},

		// 测试用例 6：没有连续序列
		// 每个数都是孤立的
		{nums: []int{10, 20, 30}, expected: 1},

		// 测试用例 7：全负数
		// 最长序列为 [-4, -3, -2, -1]
		{nums: []int{-1, -2, -3, -4}, expected: 4},
	}

	for i, test := range tests {
		result := LongestConsecutive(test.nums)
		if result == test.expected {
			fmt.Printf("Test %d passed.\n", i+1)
		} else {
			fmt.Printf("Test %d failed. Expected %d, got %d.\n", i+1, test.expected, result)
		}
	}
}
