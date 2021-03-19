package main

import "fmt"

/**
*动态规划法
 */
func longestPalindrome(s string) (res string) {
	l := len(s)
	if l < 2 {
		return s
	}
	dp := make([][]bool, l)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, l)
	}
	for i := 0; i < l; i++ {
		dp[i][i] = true
	}
	begin := 0
	maxlen := 0
	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxlen {
				maxlen = j - i + 1
				begin = i
			}
		}
	}
	res = s[begin:maxlen]
	return
}

//中心扩散方法
func longestPalindrome2(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < length; i++ {
		left1, right1 := centerSpread(s, i, i)   // 偶数
		left2, right2 := centerSpread(s, i, i+1) // 奇数
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

// z中心扩散
func centerSpread(s string, left, right int) (int, int) {
	// left = right 的时候，此时回文中心是一个字符，回文串的长度是奇数
	// right = left + 1 的时候，此时回文中心是一个空隙，回文串的长度是偶数
	length := len(s)
	for ; left >= 0 && right < length && s[left] == s[right]; left, right = left-1, right+1 {

	}
	return left + 1, right - 1
}

func main() {
	s := "babad"
	res := longestPalindrome2(s)
	fmt.Println("res:", res)
}
