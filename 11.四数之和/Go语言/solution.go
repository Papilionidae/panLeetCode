package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var list [][]int
	n := len(nums)
	if n < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for left := i + 1; left < n-2; left++ {
			if left > i+1 && nums[left] == nums[left-1] {
				continue
			}
			if nums[i]+nums[left]+nums[left+1]+nums[left+2] > target {
				break
			}
			if nums[i]+nums[left]+nums[n-2]+nums[n-1] < target {
				continue
			}
			mid := left + 1
			right := n - 1
			for mid < right {
				if nums[i]+nums[left]+nums[mid]+nums[right] > target {
					right--
				} else if nums[i]+nums[left]+nums[mid]+nums[right] < target {
					mid++
				} else {
					list = append(list, []int{nums[i], nums[left], nums[mid], nums[right]})
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
	}
	return list
}

func main() {
	nums := []int{1, -2, -5, -4, -3, 3, 3, 5}
	list := fourSum(nums, -11)
	fmt.Println("list:", list)
}
