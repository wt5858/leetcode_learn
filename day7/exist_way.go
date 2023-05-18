package main

import "fmt"

// 剑指 Offer 12. 矩阵中的路径
// 给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。
//
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
//例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。
//
//
//
//
//
//示例 1：
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true
//示例 2：
//
//输入：board = [["a","b"],["c","d"]], word = "abcd"
//输出：false
//
//
//提示：
//
//m == board.length
//n = board[i].length
//1 <= m, n <= 6
//1 <= word.length <= 15
//board 和 word 仅由大小写英文字母组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	res := exist(board, word)
	fmt.Println(res)
}

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	if m*n < len(word) {
		return false
	}

	var dfs func(i, j, k int) bool

	dfs = func(i, j, k int) bool {
		if k >= len(word) {
			return true
		}
		// 判断传入参数的可行性 i 与图行数row比较，j与图列数col比较
		// i，j初始都是0，都在图左上角
		// k是传入字符串当前索引，一开始是0，如果当前字符串索引和图当前索引对应的值不相等，表示第一个数就不相等
		// 所以继续找第一个相等的数。题目说第一个数位置不固定，即路径起点不固定（不一定是左上角为第一个数）

		// 如果board[i][j] == word[k]，则表明当前找到了对应的数，就继续执行（标记找过，继续dfs 下上右左）
		if i < 0 || j < 0 || i >= n || j >= m || board[i][j] != word[k] {
			return false
		}

		// 访问过的标记为特殊的 byte
		// 比如当前为A，没有标记找过，且A是word中对应元素，则此时应该找A下一个元素，假设是B，在dfs（B）的时候还是-
		// ->要搜索B左边的元素（假设A在B左边），所以就是ABA（凭空多出一个A，A用了2次，不可以），如果标记为空字符串->
		// 就不会有这样的问题，因为他们值不相等AB != ABA。
		// 也就是修枝
		board[i][j] = '0'

		// 分别上下左右去找下一个
		res := dfs(i, j-1, k+1) || dfs(i, j+1, k+1) || dfs(i-1, j, k+1) || dfs(i+1, j, k+1)

		// 还原之前找过的元素
		board[i][j] = word[k]

		return res
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
