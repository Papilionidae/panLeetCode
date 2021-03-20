package main

import (
	"fmt"
)

//这个有问题，因为题目明确规定了 环境中只允许最大存储32位的整数，因此res*10可能发生溢出报错
func reverse(x int) int {
	res := 0
	for ; x != 0; x /= 10 {
		tmp := x % 10
		res = res*10 + tmp
		if res > (1<<31)-1 || res < -1<<31 {
			return 0
		}
	}
	return res
}

func currentReverse(x int) int {
	res := 0
	comp1 := ((1 << 31) - 1) / 10
	comp2 := (-1 << 31) / 10
	for x != 0 {
		tmp := x % 10
		x /= 10
		if res > comp1 || res == comp1 && tmp > 7 {
			return 0
		}
		if res < comp2 || res == comp2 && tmp < (-8) {
			return 0
		}
		res = res*10 + tmp
	}
	return res
}

func main() {
	testnum := 213
	res := reverse(testnum)
	fmt.Println("res:", res)
	fmt.Println("1左移31位:", 1<<31)
	fmt.Println("-1左移31位:", -1<<31)
}
