package leetcode

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTreePaths(root *TreeNode) (ans []string) {
	path := []string{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, strconv.Itoa(node.Val))
		if node.Left == node.Right { // 叶子节点
			ans = append(ans, strings.Join(path, "->"))
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}
		path = path[:len(path)-1] // 恢复现场
	}
	dfs(root)
	return ans
}
