package main

import (
	"fmt"
	"sort"
)

// 剑指 Offer II 007. 数组中和为 0 的三个数
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//
//你返回所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//
//
//
//
//示例 1：
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
//示例 2：
//
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
//示例 3：
//
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。
//
//
//提示：
//
//3 <= nums.length <= 3000
//-105 <= nums[i] <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/1fGaJU
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{1, -1, -1, 0}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// 去重 边界问题
			continue
		}

		// 双指针来控制
		l, h := i+1, len(nums)-1
		for l < h {
			sum := nums[i] + nums[l] + nums[h]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[h]})

				for l < h && nums[l] == nums[l+1] {
					l++
				}

				for l < h && nums[h] == nums[h-1] {
					h--
				}

				l++
				h--
			} else if sum < 0 {
				l++
			} else {
				h--
			}

		}
	}
	return res
}
