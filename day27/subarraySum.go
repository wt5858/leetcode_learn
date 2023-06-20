package main

import (
	"fmt"
)

// 剑指 Offer II 010. 和为 k 的子数组
// 给定一个整数数组和一个整数k ，请找到该数组中和为k的连续子数组的个数。
//
//
//
//示例 1：
//
//输入:nums = [1,1,1], k = 2
//输出: 2
//解释: 此题 [1,1] 与 [1,1] 为两种不同的情况
//示例 2：
//
//输入:nums = [1,2,3], k = 3
//输出: 2
//
//
//提示:
//
//1 <= nums.length <= 2 * 104
//-1000 <= nums[i] <= 1000
//-107<= k <= 107
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/QTMn0o
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{1, 2, 3}
	k := 3
	fmt.Println(subarraySum(nums, k))
}

// 暴力破解
//func subarraySum(nums []int, k int) int {
//	res := 0
//	n := len(nums)
//
//	for l, num := range nums {
//		sum := num
//
//		if sum == k {
//			res++
//		}
//
//		for r := l + 1; r < n; r++ {
//			sum += nums[r]
//			if sum == k {
//				res++
//			}
//		}
//	}
//
//	return res
//}

func subarraySum(nums []int, k int) int {
	res := 0
	// 保存累加的和
	hc := map[int]int{}
	sum := 0
	n := len(nums)

	hc[0] = 1
	for i := 0; i < n; i++ {
		sum += nums[i]
		// 解释 sum[i] = sum[i-1] + num
		// nums[j~i]的累加的和为k 可以转化为 sum[i] - sum[j-1] = k  既 sum[j-1] = sum[i] - k
		// 这个地方的hc[sum-k]即为 sum[i] - k 也就是 sum[j-1]
		if c, ok := hc[sum-k]; ok {
			res += c
		}
		hc[sum]++
	}

	return res
}
