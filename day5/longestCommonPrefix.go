package main

import "fmt"

// 最长公共前缀

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//
//
//示例 1：
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//示例 2：
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//提示：
//
//1 <= strs.length <= 200
//0 <= strs[i].length <= 200
//strs[i] 仅由小写英文字母组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/longest-common-prefix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	strs := []string{"aaa", "aa", "aaa"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	s := strs[0]
	pre := ""
	for i := 1; i < len(strs); i++ {
		pre = lcp(s, strs[i])
		s = pre
	}
	return pre
}

func lcp(left, right string) string {
	l := min(len(left), len(right))
	i := 0

	for i < l && right[i] == left[i] {
		i++
	}
	return right[:i]
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}
