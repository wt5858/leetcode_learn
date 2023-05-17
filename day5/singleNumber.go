package main

import "fmt"

// 136. 只出现一次的数字
//给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
//你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
//
//
//
//示例 1 ：
//
//输入：nums = [2,2,1]
//输出：1
//示例 2 ：
//
//输入：nums = [4,1,2,1,2]
//输出：4
//示例 3 ：
//
//输入：nums = [1]
//输出：1
//
//
//提示：
//
//1 <= nums.length <= 3 * 104
//-3 * 104 <= nums[i] <= 3 * 104
//除了某个元素只出现一次以外，其余每个元素均出现两次。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/single-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{4, 1, 2, 1, 2}
	res := singleNumber(nums)
	fmt.Println(res)
}

func singleNumber(nums []int) int {
	var res int

	// 异或运算
	// 1. 归零律： a ^ a = 0
	//2. 恒等律：a ^ 0 = a
	//3. 交换律：a ^ b = b ^ a
	//4. 结合律：a ^ b ^ c = (a ^ b) ^ c = a ^ (b ^ c)
	//5. 自反：a ^ b ^ a = b
	//6. d = a ^ b ^ c 可以推出 a = d ^ b ^ c
	//7.若x是二进制数0101，y是二进制数1011；
	//则x^y=1110
	//只有在两个比较的位不同时其结果是1，否则结果为0
	//即“两个输入相同时为0，不同则为1”

	for i := 0; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}
