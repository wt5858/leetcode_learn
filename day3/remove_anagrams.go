package main

import "fmt"

//给你一个下标从 0 开始的字符串 words ，其中 words[i] 由小写英文字符组成。
//
//在一步操作中，需要选出任一下标 i ，从 words 中 删除 words[i] 。其中下标 i 需要同时满足下述两个条件：
//
//0 < i < words.length
//words[i - 1] 和 words[i] 是 字母异位词 。
//只要可以选出满足条件的下标，就一直执行这个操作。
//
//在执行所有操作后，返回 words 。可以证明，按任意顺序为每步操作选择下标都会得到相同的结果。
//
//字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。例如，"dacb" 是 "abdc" 的一个字母异位词。
//
//
//
//示例 1：
//
//输入：words = ["abba","baba","bbaa","cd","cd"]
//输出：["abba","cd"]
//解释：
//获取结果数组的方法之一是执行下述步骤：
//- 由于 words[2] = "bbaa" 和 words[1] = "baba" 是字母异位词，选择下标 2 并删除 words[2] 。
//  现在 words = ["abba","baba","cd","cd"] 。
//- 由于 words[1] = "baba" 和 words[0] = "abba" 是字母异位词，选择下标 1 并删除 words[1] 。
//  现在 words = ["abba","cd","cd"] 。
//- 由于 words[2] = "cd" 和 words[1] = "cd" 是字母异位词，选择下标 2 并删除 words[2] 。
//  现在 words = ["abba","cd"] 。
//无法再执行任何操作，所以 ["abba","cd"] 是最终答案。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/find-resultant-array-after-removing-anagrams
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	words := []string{"abba", "baba", "bbaa", "cd", "cd"}
	res := removeAnagrams(words)
	fmt.Println(res)
}

func removeAnagrams(words []string) []string {
	hc := make([][26]int, len(words))
	for i, word := range words {
		for _, v := range word {
			hc[i][v-'a']++
		}
	}

	res := []string{words[0]}
	// 用于记录最后一个不重复的下标
	j := 0
	for i := 0; i < len(words); i++ {
		// 判断是否相同
		if hc[i] != hc[j] {
			res = append(res, words[i])
			// 重置下标
			j = i
		}
	}

	return res
}
