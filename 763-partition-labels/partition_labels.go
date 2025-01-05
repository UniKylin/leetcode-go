// link			: https://leetcode.cn/problems/partition-labels/
// Author		: Kylin
// Date			: 2025-01-05

package leetcode

// 先用map记录每个字符最后一次出现的位置
func PartitionLabels(s string) (res []int) {
	lastIndexMap := make(map[string]int)
	for idx, ch := range s {
		lastIndexMap[string(ch)] = idx
	}

	count := 0
	minIdx := lastIndexMap[string(s[0])]
	// 再次遍历
	// 直到查找到某个字符的最后出现的索引跟当前遍历索引一致
	// 说明找到了可以截取的最大索引
	for idx, cha := range s {
		count++
		minIdx = max(minIdx, lastIndexMap[string(cha)])

		if idx == minIdx {
			res = append(res, count)
			count = 0
		}
	}

	return
}
