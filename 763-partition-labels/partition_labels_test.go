// link			: https://leetcode.cn/problems/partition-labels/
// Author		: Kylin
// Date			: 2025-01-05

package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}}, // 示例用例
		{"eccbbbbdec", []int{10}},                    // 全部字符属于同一区间
		{"abcde", []int{1, 1, 1, 1, 1}},              // 每个字符独立分割
		{"aaa", []int{3}},                            // 所有字符相同
		{"abacccddc", []int{9}},                      // 前后字符位置互相依赖
	}

	for _, test := range tests {
		result := PartitionLabels(test.input)
		if reflect.DeepEqual(result, test.expected) {
			fmt.Printf("Test passed for input %q\n", test.input)
		} else {
			fmt.Printf("Test failed for input %q: got %v, expected %v\n", test.input, result, test.expected)
		}
	}
}
