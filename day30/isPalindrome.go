package main

import (
	"fmt"
	"strings"
)

// 剑指 Offer II 018. 有效的回文
// 给定一个字符串 s ，验证 s是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//本题中，将空字符串定义为有效的回文串。
//
//
//
//示例 1:
//
//输入: s = "A man, a plan, a canal: Panama"
//输出: true
//解释："amanaplanacanalpanama" 是回文串
//示例 2:
//
//输入: s = "race a car"
//输出: false
//解释："raceacar" 不是回文串
//
//
//提示：
//
//1 <= s.length <= 2 * 105
//字符串 s 由 ASCII 字符组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/XltzEq
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := "0P"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	hc := make([]string, 0)
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
			hc = append(hc, strings.ToLower(string(v)))
		}
	}

	l, r := 0, len(hc)-1
	for l < r {
		if hc[l] != hc[r] {
			return false
		}
		l++
		r--
	}
	return true
}
