package nonrepeatingsubstr

import "fmt"

func lengthOfNoRepeatingSubstring(longString string)int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(longString){
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	var abc = "abcabcbb"
	fmt.Println(lengthOfNoRepeatingSubstring(abc))
}
