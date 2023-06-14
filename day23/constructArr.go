package main

import "fmt"

// 剑指 Offer 66. 构建乘积数组
// 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
//
//
//
//示例:
//
//输入: [1,2,3,4,5]
//输出: [120,60,40,30,24]
//
//
//提示：
//
//所有元素乘积之和不会溢出 32 位整数
//a.length <= 100000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(constructArr(a))
}

func constructArr(a []int) []int {
	l := len(a)
	if l < 1 {
		return nil
	}

	res := make([]int, l)

	// 乘以右边的
	res[0] = 1
	for i := 1; i < l; i++ {
		res[i] = res[i-1] * a[i-1]
	}

	// 乘以左边的
	r := 1
	for i := l - 1; i >= 0; i-- {
		res[i] *= r
		r *= a[i]
	}

	return res
}
