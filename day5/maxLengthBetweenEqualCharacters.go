package main

import (
	"fmt"
	"strings"
)

// 两个相同字符之间的最长子字符串
// 给你一个字符串 s，请你返回 两个相同字符之间的最长子字符串的长度 ，计算长度时不含这两个字符。如果不存在这样的子字符串，返回 -1 。
//
//子字符串 是字符串中的一个连续字符序列。
//
//
//
//示例 1：
//
//输入：s = "aa"
//输出：0
//解释：最优的子字符串是两个 'a' 之间的空子字符串。
//示例 2：
//
//输入：s = "abca"
//输出：2
//解释：最优的子字符串是 "bc" 。
//示例 3：
//
//输入：s = "cbzxy"
//输出：-1
//解释：s 中不存在出现出现两次的字符，所以返回 -1 。
//示例 4：
//
//输入：s = "cabbac"
//输出：4
//解释：最优的子字符串是 "abba" ，其他的非最优解包括 "bb" 和 "" 。
//
//
//提示：
//
//1 <= s.length <= 300
//s 只含小写英文字母
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/largest-substring-between-two-equal-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := "cabbac"
	r := maxLengthBetweenEqualCharacters(s)
	fmt.Println(r)
}

//func maxLengthBetweenEqualCharacters(s string) int {
//	res := -1
//	arr := [26]int{}
//	for i, _ := range arr {
//		arr[i] = -1
//	}
//
//	for i, v := range s {
//		if arr[v-'a'] < 0 {
//			arr[v-'a'] = i
//		} else {
//			res = maxNum(res, i-arr[v-'a']-1)
//		}
//	}
//	return res
//}
//
//func maxNum(left, right int) int {
//	if left > right {
//		return left
//	}
//	return right
//}

func maxLengthBetweenEqualCharacters(s string) int {
	res := -1

	for i := 0; i < 26; i++ {
		a := string(rune(97 + i))
		if n := strings.LastIndex(s, a) - strings.Index(s, a) - 1; n > res {
			res = n
		}
	}

	return res
}
