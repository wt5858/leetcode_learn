package main

import "fmt"

// 1763. 最长的美好子字符串
// 当一个字符串 s包含的每一种字母的大写和小写形式 同时出现在 s中，就称这个字符串s是 美好 字符串。比方说，"abABB"是美好字符串，因为'A' 和'a'同时出现了，且'B' 和'b'也同时出现了。然而，"abA"不是美好字符串因为'b'出现了，而'B'没有出现。
//
//给你一个字符串s，请你返回s最长的美好子字符串。如果有多个答案，请你返回最早出现的一个。如果不存在美好子字符串，请你返回一个空字符串。
//
//
//
//示例 1：
//
//输入：s = "YazaAay"
//输出："aAa"
//解释："aAa" 是一个美好字符串，因为这个子串中仅含一种字母，其小写形式 'a' 和大写形式 'A' 也同时出现了。
//"aAa" 是最长的美好子字符串。
//示例 2：
//
//输入：s = "Bb"
//输出："Bb"
//解释："Bb" 是美好字符串，因为 'B' 和 'b' 都出现了。整个字符串也是原字符串的子字符串。
//示例 3：
//
//输入：s = "c"
//输出：""
//解释：没有美好子字符串。
//示例 4：
//
//输入：s = "dDzeE"
//输出："dD"
//解释："dD" 和 "eE" 都是最长美好子字符串。
//由于有多个美好子字符串，返回 "dD" ，因为它出现得最早。
//
//
//提示：
//
//1 <= s.length <= 100
//s只包含大写和小写英文字母。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/longest-nice-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := "YazaAay"
	res := longestNiceSubstring(s)
	fmt.Println(res)
}

func longestNiceSubstring(s string) string {
	n := len(s)
	res := ""

	for l := 0; l < n; l++ {
		low, up := 0, 0
		for r := l; r < n; r++ {
			if s[r] >= 'a' && s[r] <= 'z' {
				low |= 1 << (s[r] - 'a')
			} else {
				up |= 1 << (s[r] - 'A')
			}

			if low == up && (r-l)+1 > len(res) {
				res = s[l : r+1]
			}
		}
	}

	return res
}
