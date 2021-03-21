package main

func maxArea(height []int) (maxA int) {
	n := len(height)
	for i, j := 0, n-1; i < j; {
		minHeight := min(height[i], height[j])
		if maxA < minHeight*(j-i) {
			maxA = minHeight * (j - 1)
		}
		if minHeight == height[i] {
			i++
		} else {
			j--
		}
	}
	return
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func main() {

}
