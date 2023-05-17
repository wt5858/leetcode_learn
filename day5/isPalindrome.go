package main

import (
	"fmt"
)

func main() {
	x := 121
	res := isPalindrome(x)
	fmt.Println(res)
}

//func isPalindrome(x int) bool {
//	s := strconv.Itoa(x)
//	l := len(s)
//	n := l / 2
//
//	num := 0
//	if l%2 == 1 {
//		num = 1
//	}
//	for i := n - 1; i >= 0; i-- {
//		if s[i] != s[n+num] {
//			return false
//		}
//		num++
//	}
//	return true
//}

//func isPalindrome(x int) bool {
//	s := strconv.Itoa(x)
//	ns := []byte(s)
//	for i := 0; i < len(s); i++ {
//		ns[i] = s[len(s)-i-1]
//	}
//
//	return string(ns) == s
//}

func isPalindrome(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}
