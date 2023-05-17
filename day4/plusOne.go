package main

import (
	"fmt"
)

// 加一
// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//
//
//示例1：
//
//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123。
//示例2：
//
//输入：digits = [4,3,2,1]
//输出：[4,3,2,2]
//解释：输入数组表示数字 4321。
//示例 3：
//
//输入：digits = [0]
//输出：[1]
//
//
//提示：
//
//1 <= digits.length <= 100
//0 <= digits[i] <= 9
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/plus-one
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	d := []int{4, 3, 2, 1, 9, 9}
	res := plusOne(d)
	fmt.Println(res)
}

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		// 从右往左判断是否为9，如果不为9则加1
		if digits[i] != 9 {
			digits[i]++
			// 当前下标右边的数字都改为0
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	// digits 中所有的元素均为 9

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}
