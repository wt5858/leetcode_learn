package main

import "fmt"

// 剑指 Offer 11. 旋转数组的最小数字
// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
//
//给你一个可能存在重复元素值的数组numbers，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的最小元素。例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为 1。
//
//注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
//
//
//
//示例 1：
//
//输入：numbers = [3,4,5,1,2]
//输出：1
//示例 2：
//
//输入：numbers = [2,2,2,0,1]
//输出：0
//
//
//提示：
//
//n == numbers.length
//1 <= n <= 5000
//-5000 <= numbers[i] <= 5000
//numbers 原来是一个升序排序的数组，并进行了 1 至 n 次旋转
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	numbers := []int{3, 4, 5, 1, 2}
	res := minArray(numbers)
	fmt.Println(res)
}

// 直接一次循环来查找
func minArray(numbers []int) int {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}
	return numbers[0]
}

// 二分查找
