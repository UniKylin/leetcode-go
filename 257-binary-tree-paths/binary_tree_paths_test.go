package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

// BinaryTreePaths 函数已提供
func TestBinaryTreePaths(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []string
	}{
		{
			name:     "single node tree",
			root:     &TreeNode{Val: 1},
			expected: []string{"1"},
		},
		{
			name: "tree with two levels",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: []string{"1->2", "1->3"},
		},
		{
			name: "tree with multiple levels",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: []string{"1->2->4", "1->2->5", "1->3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinaryTreePaths(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSingleCase(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	res := BinaryTreePaths(root)
	fmt.Println(res)
}
