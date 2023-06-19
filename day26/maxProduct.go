package main

import "fmt"

// 剑指 Offer II 005. 单词长度的最大乘积
// 给定一个字符串数组words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。
//
//
//
//示例1:
//
//输入: words = ["abcw","baz","foo","bar","fxyz","abcdef"]
//输出: 16
//解释: 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。
//示例 2:
//
//输入: words = ["a","ab","abc","d","cd","bcd","abcd"]
//输出: 4
//解释: 这两个单词为 "ab", "cd"。
//示例 3:
//
//输入: words = ["a","aa","aaa","aaaa"]
//输出: 0
//解释: 不存在这样的两个单词。
//
//
//提示：
//
//2 <= words.length <= 1000
//1 <= words[i].length <= 1000
//words[i]仅包含小写字母
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/aseY1I
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	words := []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	fmt.Println(maxProduct(words))
}

func maxProduct(words []string) int {
	res := 0
	hc := make([]int, len(words))
	for k, word := range words {
		m := 0
		// 把单词转化成二进制来表示。 共26个字母，即26位的二进制数，从右往左依次是a,b,c....., 占用的位置是1,2,3,....位， 有为1， 没有为0
		for i := 0; i < len(word); i++ {
			m |= 1 << int(word[i]-'a')
		}
		hc[k] = m
	}

	for k, v := range hc {
		for i := k + 1; i < len(hc); i++ {
			// 当两个单词没有相同字母时，则两个单词&操作时为0。
			if v&hc[i] == 0 {
				res = maxNum(res, len(words[k])*len(words[i]))
			}
		}
	}

	return res
}

func maxNum(l, r int) int {
	if l > r {
		return l
	}
	return r
}
