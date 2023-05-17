package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := []int{3, 4}
	res := findMedianSortedArrays(a, b)
	fmt.Println(res)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 合并两个有序数组
	nums := merge(nums1, nums2)

	// 如果合并后的数组长度是偶数，中位数是左右两个数的平均值
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2
	}

	// 如果合并后的数组长度是奇数，中位数是中间的数
	return float64(nums[len(nums)/2])
}

// 合并两个有序数组
func merge(nums1 []int, nums2 []int) []int {
	i, j := 0, 0
	m, n := len(nums1), len(nums2)
	merged := make([]int, m+n)

	for k := 0; k < m+n; k++ {
		if i == m {
			merged[k] = nums2[j]
			j++
		} else if j == n {
			merged[k] = nums1[i]
			i++
		} else if nums1[i] <= nums2[j] {
			merged[k] = nums1[i]
			i++
		} else {
			merged[k] = nums2[j]
			j++
		}
	}

	return merged
}
