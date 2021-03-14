package main

import "fmt"

/**
方法一：暴力法
*/
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

/**
方法二：哈希表法
*/
func twoSum2(nums []int, target int) []int {
	map1 := make(map[int]int)
	for i, v := range nums {
		//如果存在则返回
		if _, ok := map1[target-v]; ok {
			return []int{map1[target-v], i}
		}
		//不存在则存入，key为nums的值，value为nums的索引
		map1[nums[i]] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15, 17}
	target := 24
	// result1 := twoSum1(nums, target)
	// fmt.Println("result:", result1)
	result2 := twoSum2(nums, target)
	fmt.Println("result2:", result2)
}
