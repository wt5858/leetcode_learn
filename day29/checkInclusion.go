package main

import "fmt"

// 剑指 Offer II 014. 字符串中的变位词
// 给定两个字符串s1和s2，写一个函数来判断 s2 是否包含 s1的某个变位词。
//
//换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
//
//
//
//示例 1：
//
//输入: s1 = "ab" s2 = "eidbaooo"
//输出: True
//解释: s2 包含 s1 的排列之一 ("ba").
//示例 2：
//
//输入: s1= "ab" s2 = "eidboaoo"
//输出: False
//
//
//提示：
//
//1 <= s1.length, s2.length <= 104
//s1 和 s2 仅包含小写字母
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/MPnaiL
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}

	hc := [26]int{}
	for _, i := range s1 {
		hc[i-'a']--
	}

	left := 0

	for r, v := range s2 {
		x := v - 'a'
		hc[x]++

		// hc[x] > 0 的话右移左指针，减少剔除的hc对应的值
		for hc[x] > 0 {
			hc[s2[left]-'a']--
			left++
		}

		// 当hc[x]都不为正数,长度为n，这个时候也就意味着hc[x]为0，也就是所有元素都存在
		if r-left+1 == n {
			return true
		}
	}

	return false
}
