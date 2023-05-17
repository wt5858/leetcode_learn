package main

import "fmt"

//给你一个 m x n 的矩阵，最开始的时候，每个单元格中的值都是 0。
//
//另有一个二维索引数组indices，indices[i] = [ri, ci] 指向矩阵中的某个位置，其中 ri 和 ci 分别表示指定的行和列（从 0 开始编号）。
//
//对 indices[i] 所指向的每个位置，应同时执行下述增量操作：
//
//ri 行上的所有单元格，加 1 。
//ci 列上的所有单元格，加 1 。
//给你 m、n 和 indices 。请你在执行完所有indices指定的增量操作后，返回矩阵中 奇数值单元格 的数目。
//
//
//
//示例 1：
//
//
//
//输入：m = 2, n = 3, indices = [[0,1],[1,1]]
//输出：6
//解释：最开始的矩阵是 [[0,0,0],[0,0,0]]。
//第一次增量操作后得到 [[1,2,1],[0,1,0]]。
//最后的矩阵是 [[1,3,1],[1,3,1]]，里面有 6 个奇数。
//示例 2：
//
//
//
//输入：m = 2, n = 2, indices = [[1,1],[0,0]]
//输出：0
//解释：最后的矩阵是 [[2,2],[2,2]]，里面没有奇数。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/cells-with-odd-values-in-a-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	m := 2
	n := 3
	indices := [][]int{{0, 1}, {1, 1}}

	res := oddCells(m, n, indices)
	fmt.Println(res)
}

//func oddCells(m int, n int, indices [][]int) int {
//	rows := make([]int, m)
//	cols := make([]int, n)
//
//	for _, v := range indices {
//		rows[v[0]]++
//		cols[v[1]]++
//	}
//
//	res := 0
//	for _, row := range rows {
//		for _, col := range cols {
//			res += (row + col) % 2
//		}
//	}
//	return res
//}

func oddCells(m int, n int, indices [][]int) int {
	rows := make([]int, m)
	cols := make([]int, n)

	for _, v := range indices {
		rows[v[0]]++
		cols[v[1]]++
	}

	colAdd := 0
	for _, col := range cols {
		if col%2 == 1 {
			colAdd++
		}
	}

	rowAdd := 0
	for _, row := range rows {
		if row%2 == 1 {
			rowAdd++
		}
	}

	return n*rowAdd + m*colAdd - 2*rowAdd*colAdd
}
