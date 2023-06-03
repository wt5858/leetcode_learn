package main

import "fmt"

// 剑指 Offer 56 - I. 数组中数字出现的次数
//一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
//
//
//
//示例 1：
//
//输入：nums = [4,1,4,6]
//输出：[1,6] 或 [6,1]
//示例 2：
//
//输入：nums = [1,2,10,4,1,4,3,3]
//输出：[2,10] 或 [10,2]
//
//
//限制：
//
//2 <= nums.length <= 10000
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{1, 2, 10, 4, 1, 4, 3, 3}
	r := singleNumbers(nums)
	fmt.Println(r)
}

func singleNumbers(nums []int) []int {
	ret := 0
	for _, num := range nums {
		ret ^= num
	}

	v := 1
	for (v & ret) == 0 {
		v <<= 1
	}

	a, b := 0, 0
	for _, num := range nums {
		if (v & num) == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
