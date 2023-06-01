package main

import "fmt"

// 剑指 Offer 40. 最小的k个数
// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//
//
//
//示例 1：
//
//输入：arr = [3,2,1], k = 2
//输出：[1,2] 或者 [2,1]
//示例 2：
//
//输入：arr = [0,1,2,1], k = 1
//输出：[0]
//
//
//限制：
//
//0 <= k <= arr.length <= 10000
//0 <= arr[i]<= 10000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	arr := []int{3, 2, 1}
	k := 2
	r := getLeastNumbers(arr, k)
	fmt.Println(r)
}

func getLeastNumbers(arr []int, k int) []int {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	// 快排处理后 取前面的数据
	quick(arr, 0, len(arr)-1)
	return arr[:k]
}

func quick(arr []int, l, r int) {
	if l < r {
		p := recur(arr, l, r)
		quick(arr, l, p-1)
		quick(arr, p+1, r)
	}
}

func recur(arr []int, l, r int) int {
	pivot := arr[r]
	point := l - 1

	for i := l; i < r; i++ {
		if arr[i] <= pivot {
			point++
			arr[i], arr[point] = arr[point], arr[i]
		}
	}

	arr[point+1], arr[r] = arr[r], arr[point+1]

	return point + 1
}
