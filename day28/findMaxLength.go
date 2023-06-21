package main

import "fmt"

// 剑指 Offer II 011. 0 和 1 个数相同的子数组
// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
//
//
//
//示例 1：
//
//输入: nums = [0,1]
//输出: 2
//说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
//示例 2：
//
//输入: nums = [0,1,0]
//输出: 2
//说明: [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。
//
//
//提示：
//
//1 <= nums.length <= 105
//nums[i] 不是 0 就是 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/A1NYOS
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{0, 0, 1, 0, 0, 0, 1, 1}
	fmt.Println(findMaxLength(nums))
}

func findMaxLength(nums []int) int {
	res := 0
	hc := map[int]int{}
	sum := 0
	hc[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
		sum += nums[i]
		if c, ok := hc[sum]; ok {
			res = maxNum(res, i-c)
		} else {
			hc[sum] = i
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
