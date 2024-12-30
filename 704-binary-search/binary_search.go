// link		: https://leetcode-cn.com/problems/binary-search/
// Author	: Kylin
// Date		: 2022-01-14

package leetcode

func Search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		// mid := (low + high) >> 1
		mid := low + (high-low)>>1
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func BinarySearchClosed(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] == target {
			return middle
		}

		if nums[middle] < target {
			left = middle + 1
		}

		if nums[middle] > target {
			right = middle - 1
		}
	}
	return -1
}
