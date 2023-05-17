package main

import (
	"fmt"
	"strings"
)

// 395. 至少有 K 个重复字符的最长子串
// 给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串，要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。
//
//
//
//示例 1：
//
//输入：s = "aaabb", k = 3
//输出：3
//解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。
//示例 2：
//
//输入：s = "ababbc", k = 2
//输出：5
//解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
//
//
//提示：
//
//1 <= s.length <= 104
//s 仅由小写英文字母组成
//1 <= k <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := "ababbc"
	k := 2
	res := longestSubstring(s, k)
	fmt.Println(res)
}

func longestSubstring(s string, k int) int {
	if s == "" {
		return 0
	}

	res := 0
	hc := [26]int{}
	for _, i := range s {
		hc[i-'a']++
	}

	var split byte
	for i, v := range hc {
		if v > 0 && v < k {
			split = 'a' + byte(i)
		}
	}

	if split == 0 {
		return len(s)
	}

	// 根据未超过k的字符分割成不同的块，再用递归来处理
	for _, s2 := range strings.Split(s, string(split)) {
		res = max(res, longestSubstring(s2, k))
	}

	return res
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}
