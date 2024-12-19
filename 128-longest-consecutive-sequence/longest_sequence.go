package leetcode

import "fmt"

func LongestConsecutive(nums []int) (res int) {
	set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}

	fmt.Println(set)
	// 100 4 200 1  3  2
	//  1     1  4
	for x := range set {
		fmt.Println(">>> x:", x)
		if set[x-1] {
			fmt.Println(">>> x - 1:", x-1)
			continue
		}

		y := x + 1
		for set[y] {
			y++
		}

		res = max(res, y-x)
	}
	return
}
