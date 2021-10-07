// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var max int
	var count int
	var begin int
	char := make(map[byte]int)
	for i := range s {

		b := byte(s[i])
		if _, ok := char[b]; ok {
			if char[b] > begin {
				begin = char[b]
			}
			count = i - begin
		} else {
			count++
		}
		char[b] = i
		if max < count {
			max = count
		}
		fmt.Printf("Pos %d(%c): count %d begin %d (%c) max %d\n", i, s[i], count, begin, s[begin], max)
	}
	return max
}

func main() {
	lengthOfLongestSubstring("pwwkeww")
}
