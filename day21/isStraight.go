package main

import (
	"fmt"
)

// 剑指 Offer 61. 扑克牌中的顺子
// 从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
//
//
//
//示例1:
//
//输入: [1,2,3,4,5]
//输出: True
//
//
//示例2:
//
//输入: [0,0,1,2,5]
//输出: True
//
//
//限制：
//
//数组长度为 5
//
//数组的数取值为 [0, 13] .
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{0, 0, 1, 2, 5}
	r := isStraight(nums)
	fmt.Println(r)
}

func isStraight(nums []int) bool {
	max, min := 0, 14
	hc := map[int]bool{}
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if hc[num] {
			return false
		}
		max = Max(max, num)
		min = Min(min, num)
		hc[num] = true
	}
	return max-min < 5
}

func Max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func Min(l, r int) int {
	if l < r {
		return l
	}
	return r
}
