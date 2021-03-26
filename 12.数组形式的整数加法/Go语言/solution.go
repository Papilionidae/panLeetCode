package main

import "fmt"

//逐位相加
func addToArrayForm(A []int, K int) []int {
	n := len(A)
	sum, carray := 0, 0
	var res []int
	for i := n - 1; i >= 0 || K != 0; i-- {
		x, y := A[i], K%10
		if i < 0 {
			x = 0
		}
		if K == 0 {
			y = 0
		}
		sum = x + y + carray
		carray = sum / 10
		K = K / 10
		res = append(res, sum%10)
	}
	if carray != 0 {
		res = append(res, carray)
	}
	length := len(res)
	for j := 0; j < length/2; j++ {
		res[j], res[length-j-1] = res[length-j-1], res[j]
	}
	return res
}

//方法二：
func addToArrayForm1(A []int, K int) (ans []int) {
	for i := len(A) - 1; i >= 0 || K > 0; i-- {
		if i >= 0 {
			K += A[i]
		}
		ans = append(ans, K%10)
		K /= 10
	}
	for j, n := 0, len(ans); j < n/2; j++ {
		ans[j], ans[n-j-1] = ans[n-j-1], ans[j]
	}
	return
}

func main() {
	A := []int{1, 2, 0, 0}
	K := 34
	res := addToArrayForm(A, K)
	fmt.Println("res:", res)
}
