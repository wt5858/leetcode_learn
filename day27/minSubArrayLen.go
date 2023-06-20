package main

import (
	"fmt"
	"math"
)

// 剑指 Offer II 008. 和大于等于 target 的最短子数组
// 给定一个含有n个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
//
//
//示例 1：
//
//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组[4,3]是该条件下的长度最小的子数组。
//示例 2：
//
//输入：target = 4, nums = [1,4,4]
//输出：1
//示例 3：
//
//输入：target = 11, nums = [1,1,1,1,1,1,1,1]
//输出：0
//
//
//提示：
//
//1 <= target <= 109
//1 <= nums.length <= 105
//1 <= nums[i] <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/2VG8Kg
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := math.MaxInt32
	l, r, sum := 0, 0, 0

	// 窗口滑动
	for r < n {
		sum += nums[r]
		// 大于等于target的时候 左边窗口右移 缩小窗口
		for sum >= target {
			res = minNum(res, r-l+1)
			sum -= nums[l]
			l++
		}
		// 右边窗口右移  增大窗口
		r++
	}

	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func minNum(l, r int) int {
	if l < r {
		return l
	}
	return r
}
