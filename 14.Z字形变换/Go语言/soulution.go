package main

func convert(s string, numRows int) (res string) {
	if numRows == 1 {
		return s
	}
	n := len(s)
	cycleLen := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			res = res + string(s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				res = res + string(s[j+cycleLen-i])
			}
		}
	}
	return
}

func main() {

}
