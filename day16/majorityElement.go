package main

import "fmt"

// 剑指 Offer 39. 数组中出现次数超过一半的数字
// 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
//
//
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
//示例1:
//
//输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
//输出: 2
//
//
//限制：
//
//1 <= 数组长度 <= 50000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	a := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	r := majorityElement(a)
	fmt.Println(r)
}

//func majorityElement(nums []int) int {
//	m := make(map[int]int, 0)
//	l := len(nums)
//	for _, num := range nums {
//		if n, ok := m[num]; ok {
//			if n+1 > l/2 {
//				return num
//			}
//			m[num]++
//		} else {
//			m[num] = 1
//		}
//	}
//	return nums[0]
//}

// 摩尔投票法
func majorityElement(nums []int) int {
	v, x := 0, 0
	for _, num := range nums {
		if v == 0 {
			x = num
		}
		// 众数+1 非众数-1
		tmp := -1
		if num == x {
			tmp = 1
		}
		v += tmp
	}
	return x
}
