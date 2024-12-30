// link		: https://leetcode-cn.com/problems/binary-search/
// Author	: Kylin
// Date		: 2022-01-14

package leetcode

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	res := Search(nums, 9)
	fmt.Println(res)

	nextRes := Search(nums, 2)
	fmt.Println(nextRes)
}

func TestBinarySearchClosed(t *testing.T) {
	// 测试用例
	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		// 正常情况
		{[]int{1, 3, 5, 7, 9}, 5, 2},  // 目标值在数组中间
		{[]int{1, 3, 5, 7, 9}, 1, 0},  // 目标值在数组开头
		{[]int{1, 3, 5, 7, 9}, 9, 4},  // 目标值在数组末尾
		{[]int{1, 3, 5, 7, 9}, 6, -1}, // 目标值不在数组中

		// 边界情况
		{[]int{1}, 1, 0},  // 数组只有一个元素且找到
		{[]int{1}, 2, -1}, // 数组只有一个元素但未找到
		{[]int{}, 1, -1},  // 空数组

		// 重复元素
		{[]int{1, 2, 2, 2, 3}, 2, 2}, // 目标值有多个出现，返回任意一个有效索引
		{[]int{1, 1, 1, 1, 1}, 1, 2}, // 全是重复值，返回任意一个有效索引

		// 大范围数组
		{[]int{1, 3, 5, 7, 9, 11, 13}, 13, 6}, // 大数组，目标值在末尾
	}

	// 运行测试用例
	for i, tc := range testCases {
		result := BinarySearchClosed(tc.nums, tc.target)
		if result == tc.expected {
			fmt.Printf("测试用例 %d: 通过 ✅\n", i+1)
		} else {
			fmt.Printf("测试用例 %d: 未通过 ❌ (输入: %v, 目标值: %d, 期望: %d, 实际: %d)\n",
				i+1, tc.nums, tc.target, tc.expected, result)
		}
	}
}
