package main

import "fmt"

// 给你一个字符串 croakOfFrogs，它表示不同青蛙发出的蛙鸣声（字符串 "croak" ）的组合。由于同一时间可以有多只青蛙呱呱作响，所以 croakOfFrogs 中会混合多个 “croak” 。
//
//请你返回模拟字符串中所有蛙鸣所需不同青蛙的最少数目。
//
//要想发出蛙鸣 "croak"，青蛙必须 依序 输出 ‘c’, ’r’, ’o’, ’a’, ’k’ 这 5 个字母。如果没有输出全部五个字母，那么它就不会发出声音。如果字符串 croakOfFrogs 不是由若干有效的 "croak" 字符混合而成，请返回 -1 。
//
//
//
//示例 1：
//
//输入：croakOfFrogs = "croakcroak"
//输出：1
//解释：一只青蛙 “呱呱” 两次
//示例 2：
//
//输入：croakOfFrogs = "crcoakroak"
//输出：2
//解释：最少需要两只青蛙，“呱呱” 声用黑体标注
//第一只青蛙 "crcoakroak"
//第二只青蛙 "crcoakroak"
//示例 3：
//
//输入：croakOfFrogs = "croakcrook"
//输出：-1
//解释：给出的字符串不是 "croak" 的有效组合。
//
//
//提示：
//
//1 <= croakOfFrogs.length <= 105
//字符串中的字符只有 'c', 'r', 'o', 'a' 或者 'k'
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/minimum-number-of-frogs-croaking
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	croakOfFrogs := "croakcroak"
	res := minNumberOfFrogs(croakOfFrogs)
	fmt.Println(res)
}

func minNumberOfFrogs(croakOfFrogs string) int {
	var previous = [...]int{'c': 'k', 'r': 'c', 'o': 'r', 'a': 'o', 'k': 'a'}

	cnt := [len(previous)]int{}

	for _, v := range croakOfFrogs {
		pre := previous[v] // pre 为 ch 在 "croak" 中的上一个字母
		if cnt[pre] > 0 {  // 如果有青蛙发出了 pre 的声音
			cnt[pre]-- // 复用一只
		} else if v != 'c' { // 否则青蛙必须从 'c' 开始蛙鸣
			return -1
		}
		cnt[v]++ // 发出了 ch 的声音
	}

	if cnt['c'] > 0 || cnt['r'] > 0 || cnt['o'] > 0 || cnt['a'] > 0 {
		return -1 // 有发出其它声音的青蛙，不符合要求
	}

	return cnt['k'] // 最后青蛙们都发出了 'k' 的声音
}
