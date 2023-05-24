package main

import "fmt"

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。
//
//
//
//示例：
//
//输入：nums =[1,2,3,4]
//输出：[1,3,2,4]
//注：[3,1,2,4] 也是正确的答案之一。
//
//
//提示：
//
//0 <= nums.length <= 50000
//0 <= nums[i] <= 10000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{1, 2, 3, 4}
	res := exchange(nums)
	fmt.Println(res)
}

// 两次遍历
//func exchange(nums []int) []int {
//	var a []int
//	var b []int
//	for _, num := range nums {
//		if num%2 == 0 {
//			b = append(b, num)
//		} else {
//			a = append(a, num)
//		}
//	}
//	for _, i := range b {
//		a = append(a, i)
//	}
//	return a
//}

// 双指针
//func exchange(nums []int) []int {
//	a := make([]int, len(nums))
//	// 用左右指针来处理
//	l, r := 0, len(nums)-1
//	for _, num := range nums {
//		if num%2 == 0 {
//			// 偶数就对应到右指针，并让右指针左移
//			a[r] = num
//			r--
//		} else {
//			// 奇数就对应到左指针，并让左指针右移
//			a[l] = num
//			l++
//		}
//	}
//	return a
//}

// 快慢指针
func exchange(nums []int) []int {
	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[fast]%2 != 0 {
			// 当快指针找到奇数后，和慢指针对应的值交换，让慢指针后移
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}

	return nums
}
