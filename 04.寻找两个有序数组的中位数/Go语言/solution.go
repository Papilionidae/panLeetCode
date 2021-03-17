package main

import "fmt"

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	total := m + n
	if total%2 == 1 {
		//奇数
		k := total / 2
		median := getKthElement(k+1, nums1, nums2)
		return float64(median)
	} else {
		//偶数
		k1, k2 := total/2-1, total/2
		median := float64(getKthElement(k1+1, nums1, nums2)+getKthElement(k2+1, nums1, nums2)) / 2.0
		return median
	}
}

/*  k (k>1) 小的元素，那么就取 pivot1 = nums1[k/2-1] 和 pivot2 = nums2[k/2-1] 进行比较
 * 这里的 "/" 表示整除
 * nums1 中小于等于 pivot1 的元素有 nums1[0 .. k/2-2] 共计 k/2-1 个
 * nums2 中小于等于 pivot2 的元素有 nums2[0 .. k/2-2] 共计 k/2-1 个
 * 取 pivot = min(主要思路：要找到第pivot1, pivot2)，两个数组中小于等于 pivot 的元素共计不会超过 (k/2-1) + (k/2-1) <= k-2 个
 * 这样 pivot 本身最大也只能是第 k-1 小的元素
 * 如果 pivot = pivot1，那么 nums1[0 .. k/2-1] 都不可能是第 k 小的元素。把这些元素全部 "删除"，剩下的作为新的 nums1 数组
 * 如果 pivot = pivot2，那么 nums2[0 .. k/2-1] 都不可能是第 k 小的元素。把这些元素全部 "删除"，剩下的作为新的 nums2 数组
 * 由于我们 "删除" 了一些元素（这些元素都比第 k 小的元素要小），因此需要修改 k 的值，减去删除的数的个数
 */
func getKthElement(k int, nums1, nums2 []int) int {

	index1, index2 := 0, 0
	len1, len2 := len(nums1), len(nums2)
	for {
		//特殊情况
		if index1 == len1 {
			return nums2[index2+k-1]
		}
		if index2 == len2 {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		//正常情况
		mid := k / 2
		newIndex1 := min(index1+mid, len1) - 1
		newIndex2 := min(index2+mid, len2) - 1
		value1, value2 := nums1[newIndex1], nums2[newIndex2]
		if value1 <= value2 {
			k = k - (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k = k - (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func main() {
	input1 := []int{1, 2}
	input2 := []int{3, 4}
	res := findMedianSortedArrays(input1, input2)
	fmt.Println("result:", res)
}
