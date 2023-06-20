package main

import "fmt"

// 剑指 Offer II 009. 乘积小于 K 的子数组
// 给定一个正整数数组nums和整数 k，请找出该数组内乘积小于k的连续的子数组的个数。
//
//
//
//示例 1:
//
//输入: nums = [10,5,2,6], k = 100
//输出: 8
//解释: 8 个乘积小于 100 的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
//需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
//示例 2:
//
//输入: nums = [1,2,3], k = 0
//输出: 0
//
//
//提示:
//
//1 <= nums.length <= 3 * 104
//1 <= nums[i] <= 1000
//0 <= k <= 106
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/ZVAVXX
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{10, 5, 2, 6}
	k := 100
	fmt.Println(numSubarrayProductLessThanK(nums, k))
}

// 循环判断
//func numSubarrayProductLessThanK(nums []int, k int) int {
//	n := len(nums)
//	res := 0
//	l := 0
//	for l < n {
//		v := nums[l]
//		for r := l; r < n; r++ {
//			if r > l {
//				v *= nums[r]
//			}
//
//			if v < k {
//				res++
//			} else {
//				break
//			}
//		}
//
//		l++
//	}
//	return res
//}

func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	l := 0
	s := 1

	// 双指针
	for r, num := range nums {
		s *= num
		for ; l <= r && s >= k; l++ {
			s /= nums[l]
		}
		res += r - l + 1
	}
	return res
}
