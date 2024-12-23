package leetcode

import (
	"fmt"
	"testing"
)

func TestAA(t *testing.T) {
	// 构建字符计数
	cntS := make([]int, 128)
	cntT := make([]int, 128)

	s := "helloWorld"
	tt := "xx"

	for _, ch := range s {
		cntS[ch]++
	}

	for _, ch := range tt {
		cntT[ch]++
	}

	fmt.Println(IsCovered(cntS, cntT)) // 输出: true
}
