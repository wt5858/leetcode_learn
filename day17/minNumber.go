package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 剑指 Offer 45. 把数组排成最小的数
// 入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
//
//
//
//示例 1:
//
//输入: [10,2]
//输出: "102"
//示例2:
//
//输入: [3,30,34,5,9]
//输出: "3033459"
//
//
//提示:
//
//0 < nums.length <= 100
//说明:
//
//输出结果可能非常大，所以你需要返回一个字符串而不是整数
//拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{3, 30, 34, 5, 9}
	r := minNumber(nums)
	fmt.Println(r)
}

func minNumber(nums []int) string {
	var str []string
	for _, num := range nums {
		str = append(str, strconv.Itoa(num))
	}

	// 用自带的包进行字符串的排序
	sort.Slice(str, func(i, j int) bool {
		return str[i]+str[j] < str[j]+str[i]
	})
	return strings.Join(str, "")
}
