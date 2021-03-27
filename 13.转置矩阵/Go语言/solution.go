package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = matrix[i][j]
		}
	}
	return res
}

func main() {
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}
	matrix[0] = []int{1, 2, 3}
	matrix[1] = []int{4, 5, 6}
	matrix[2] = []int{7, 8, 9}
	res := transpose(matrix)
	fmt.Println("res:", res)
}
