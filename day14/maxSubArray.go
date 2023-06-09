package main

import "fmt"

// 剑指 Offer 42. 连续子数组的最大和
// 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
//
//要求时间复杂度为O(n)。
//
//
//
//示例1:
//
//输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释:连续子数组[4,-1,2,1] 的和最大，为6。
//
//
//提示：
//
//1 <=arr.length <= 10^5
//-100 <= arr[i] <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	r := maxSubArray(nums)
	fmt.Println(r)
}

func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += maxNum(nums[i-1], 0)
		res = maxNum(res, nums[i])
	}
	return res
}

func maxNum(l, r int) int {
	if l > r {
		return l
	}
	return r
}
