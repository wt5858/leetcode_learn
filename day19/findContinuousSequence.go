package main

import "fmt"

// 剑指 Offer 57 - II. 和为s的连续正数序列
// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//
//序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
//
//
//
//示例 1：
//
//输入：target = 9
//输出：[[2,3,4],[4,5]]
//示例 2：
//
//输入：target = 15
//输出：[[1,2,3,4,5],[4,5,6],[7,8]]
//
//
//限制：
//
//1 <= target <= 10^5
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	target := 9
	r := findContinuousSequence(target)
	fmt.Println(r)
}

func findContinuousSequence(target int) [][]int {
	var res [][]int
	l, r := 1, 2
	for l < r {
		sum := (l + r) * (r - l + 1) / 2
		if sum == target {
			var tmp []int
			for i := l; i <= r; i++ {
				tmp = append(tmp, i)
			}
			res = append(res, tmp)
			l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}
	return res
}
