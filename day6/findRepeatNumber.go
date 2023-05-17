package main

import "fmt"

// 剑指 Offer 03. 数组中重复的数字
//找出数组中重复的数字。
//
//
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
//
//示例 1：
//
//输入：
//[2, 3, 1, 0, 2, 5, 3]
//输出：2 或 3
//
//
//限制：
//
//2 <= n <= 100000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	res := findRepeatNumber(nums)
	fmt.Println(res)
}

// 用map来保存
//func findRepeatNumber(nums []int) int {
//	m := map[int]bool{}
//	for _, num := range nums {
//		if m[num] {
//			return num
//		}
//		m[num] = true
//	}
//	return -1
//}

// 遍历数组 nums，设索引初始值为 i = 0
//
//若 nums[i] = i： 说明此数字已在正确位置，跟索引对应，无需交换，跳过；
//若 nums[nums[i]] = nums[i]：代表索引 nums[i]处和索引 i 处的元素值都为nums[i]，找到一组重复值，返回nums[i]；
//否则：交换索引为 i 和 nums[i] 的元素值，将此数字交换至对应索引正确的位置。
//若遍历完毕尚未找到重复数字，则返回 −1 。
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
			continue
		}

		if nums[nums[i]] == nums[i] {
			return nums[i]
		}

		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}
	return -1
}
