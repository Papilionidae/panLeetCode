package main

import "fmt"

func lengthOfLongestSubstring(str string) int {
	// 哈希集合，记录每个字符的索引
	m := map[byte]int{}
	n := len(str)
	start, ans := 0, 0
	for end := 0; end < n; end++ {
		var b byte = str[end]
		if _, ok := m[b]; ok {
			start = max(m[b]+1, start)
		}
		ans = max(ans, end-start+1)
		m[b] = end
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	s := "abcabcbb"
	ans := lengthOfLongestSubstring(s)
	fmt.Println("ans:", ans)
}
