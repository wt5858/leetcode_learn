package main

import "fmt"

// 剑指 Offer 29. 顺时针打印矩阵
// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
//
//
//
//示例 1：
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//示例 2：
//
//输入：matrix =[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//限制：
//
//0 <= matrix.length <= 100
//0 <= matrix[i].length<= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	r := spiralOrder(matrix)
	fmt.Println(r)
}

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)

	rl, rh, cl, ch := 0, len(matrix)-1, 0, len(matrix[0])-1

	for {
		// 向右
		for i := cl; i <= ch; i++ {
			res = append(res, matrix[rl][i])
		}
		rl++

		if rl > rh {
			break
		}

		// 向下
		for i := rl; i <= rh; i++ {
			res = append(res, matrix[i][ch])
		}
		ch--

		if ch < cl {
			break
		}

		// 向左
		for i := ch; i >= cl; i-- {
			res = append(res, matrix[rh][i])
		}
		rh--

		if rh < rl {
			break
		}

		// 向上
		for i := rh; i >= rl; i-- {
			res = append(res, matrix[i][cl])
		}
		cl++

		if cl > ch {
			break
		}
	}

	return res
}
