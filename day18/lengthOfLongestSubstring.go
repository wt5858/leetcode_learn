package main

import (
	"fmt"
)

// 剑指 Offer 48. 最长不含重复字符的子字符串
// 请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
//
//
//
//示例1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
//    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
//
//
//提示：
//
//s.length <= 40000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := "abcabcbb"
	r := lengthOfLongestSubstring(s)
	fmt.Println(r)
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	hc := map[byte]int{}
	l, r, res := 0, -1, 0
	for l < n {
		if r+1 < n && (hc[s[r+1]] == 0) {
			hc[s[r+1]]++
			r++
		} else {
			delete(hc, s[l])
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

func max(l, r int) int {
	if r > l {
		return r
	}
	return l
}
