package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) (list [][]int) {
	n := len(nums)
	if n < 3 {
		list = [][]int{}
		return
	}
	sort.Ints(nums)

	for left := 0; left < n-2; left++ {
		if nums[left] > 0 { //第一个数大于0，后面的数相加不可能为0了，直接返回
			return
		}
		//去重，如果此数已经选取过了，直接跳过
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		mid := left + 1
		right := n - 1
		for mid < right {
			if nums[left]+nums[mid]+nums[right] > 0 {
				right--
			} else if nums[left]+nums[mid]+nums[right] < 0 {
				mid++
			} else {
				list = append(list, []int{nums[left], nums[mid], nums[right]})
				mid++
				right--
				for mid < right && nums[mid] == nums[mid-1] {
					mid++
				}
				for mid < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	list := threeSum(nums)
	fmt.Println("list:", list)
}
