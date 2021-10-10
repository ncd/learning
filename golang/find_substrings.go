// https://leetcode.com/problems/substring-with-concatenation-of-all-words
package main

import "fmt"

func findSubstring(s string, words []string) []int {
	l := len(words[0])
	var check func(string, int, int, map[string]int) bool
	check = func(s string, pos int, left int, mapping map[string]int) bool {
		if left == 0 {
			return true
		}
		word := s[pos : pos+l]
		if mapping[word] == 0 {
			return false
		} else {
			mapping[word]--
			defer func() {
				mapping[word]++
			}()
			return check(s, pos+l, left-l, mapping)
		}
	}
	var length int
	mapping := make(map[string]int)
	for i := range words {
		length += len(words[i])
		mapping[words[i]]++
	}

	if len(s) < length {
		return []int{}
	}
	var result []int
	for i := 0; i < len(s)-length+1; i++ {
		if check(s, i, length, mapping) {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}
