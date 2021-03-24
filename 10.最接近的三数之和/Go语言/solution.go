package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums)
	sum := nums[0] + nums[1] + nums[2]
	larger, smaller := sum, sum
	for i := 0; i < n-2; i++ {
		if i > 1 && nums[i] == target {
			break
		}
		mid := i + 1
		right := n - 1
		for mid < right {
			sum = nums[i] + nums[mid] + nums[right]
			if sum == target {
				return sum
			}
			if sum > target {
				if math.Abs(float64(sum-target)) < math.Abs(float64(larger-target)) {
					larger = sum
				}
				right--
			} else {
				if math.Abs(float64(target-sum)) < math.Abs(float64(target-smaller)) {
					smaller = sum
				}
				mid++
			}
		}
	}
	if math.Abs(float64(larger-target)) > math.Abs(float64(target-smaller)) {
		return smaller
	}
	return larger
}

func main() {
	nums := []int{1, 1, 1, 0}
	target := 100
	res := threeSumClosest(nums, target)
	fmt.Println("res:", res)
}
