package main

import (
	"fmt"
	"strings"
)

// 剑指 Offer 38. 字符串的排列
// 输入一个字符串，打印出该字符串中字符的所有排列。
//
//
//
//你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
//
//
//
//示例:
//
//输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]
//
//
//限制：
//
//1 <= s 的长度 <= 8
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := "abc"
	r := permutation(s)
	fmt.Println(r)
}

func permutation(s string) []string {
	var res []string
	c := []byte(s)
	var dfs func(x int)
	dfs = func(x int) {
		if x == len(s)-1 {
			res = append(res, string(c))
		}
		var set []byte
		for i := x; i < len(s); i++ {
			// 已经存在就剪枝
			if strings.Contains(string(set), string(c[i])) {
				continue
			}
			set = append(set, c[i])
			// 交换，将 c[i] 固定在第 x 位
			c[i], c[x] = c[x], c[i]
			// 开启第 x + 1 位字符
			dfs(x + 1)
			// 恢复
			c[i], c[x] = c[x], c[i]
		}
	}

	dfs(0)

	return res
}
