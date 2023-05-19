package main

import "fmt"

// 剑指 Offer 13. 机器人的运动范围
// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
//
//
//
//示例 1：
//
//输入：m = 2, n = 3, k = 1
//输出：3
//示例 2：
//
//输入：m = 3, n = 1, k = 0
//输出：1
//提示：
//
//1 <= n,m <= 100
//0 <= k<= 20
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	m := 3
	n := 2
	k := 17
	res := movingCount(m, n, k)
	fmt.Println(res)
}

func movingCount(m int, n int, k int) int {
	ma := make([][]bool, m)
	for i, _ := range ma {
		ma[i] = make([]bool, n)
	}

	// 采用的dfs来处理
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 判断是否越界，是否已经出现过
		if i < 0 || j < 0 || i >= m || j >= n || sum(i, j) > k || ma[i][j] {
			return 0
		}
		ma[i][j] = true
		return dfs(i+1, j) + dfs(i, j+1) + 1
	}

	return dfs(0, 0)
}

func sum(i, j int) int {
	res := 0

	for i > 0 {
		res += i % 10
		i /= 10
	}

	for j > 0 {
		res += j % 10
		j /= 10
	}
	return res
}
