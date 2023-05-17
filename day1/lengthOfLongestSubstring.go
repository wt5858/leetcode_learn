package main

import "fmt"

// 无重复字符的最长子串
// 定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。
//
//
//
//示例1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
//    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)

}

func lengthOfLongestSubstring(s string) int {
	// 滑动窗口
	var frep [256]int
	l, r, res := 0, -1, 0
	for l < len(s) {
		if r+1 < len(s) && frep[s[r+1]] == 0 {
			// 符合的时候 右边的窗口再右移 扩大窗口
			frep[s[r+1]]++
			r++
		} else {
			// 不符合的时候 左边的窗口右移 减小窗口
			frep[s[l]]--
			l++
		}

		if res < r-l+1 {
			res = r - l + 1
		}
	}
	return res
}
