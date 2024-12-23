package leetcode

func DailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var st []int // todolist

	for i, t := range temperatures {
		for len(st) > 0 && t > temperatures[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ans[j] = i - j
		}
		st = append(st, i)
	}
	return ans
}
