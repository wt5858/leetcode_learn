package main

import "fmt"

// 剑指 Offer II 019. 最多删除一个字符得到回文
// 给定一个非空字符串s，请判断如果最多 从字符串中删除一个字符能否得到一个回文字符串。
//
//
//
//示例 1:
//
//输入: s = "aba"
//输出: true
//示例 2:
//
//输入: s = "abca"
//输出: true
//解释: 可以删除 "c" 字符 或者 "b" 字符
//示例 3:
//
//输入: s = "abc"
//输出: false
//
//
//提示:
//
//1 <= s.length <= 105
//s 由小写英文字母组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/RQku0D
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := "deeee"
	fmt.Println(validPalindrome(s))
}

// 双指针
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		// 如果s[l] == s[r]将两个指针向中间靠拢
		if s[l] == s[r] {
			l++
			r--
		} else {
			flag1, flag2 := true, true

			// 判断 s[l+1,r]之间的元素是否为回文数
			for i, j := l+1, r; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}

			// 判断 s[l,r-1]之间的元素是否为回文数
			for i, j := l, r-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}

			return flag1 || flag2
		}
	}
	return true
}
