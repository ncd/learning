package main

import "fmt"

func smallest(s string, char []rune) string {
	mapping := make(map[rune]int)

	for _, c := range char {
		mapping[c] = 0
		for _, v := range s {
			if v == c {
				mapping[c]++
			}
		}
	}
	for _, v := range mapping {
		if v == 0 {
			return ""
		}
	}
	var find func(string, int, int, map[rune]int) string
	find = func(s string, left int, right int, mapping map[rune]int) string {
		for left < right {
			for mapping[rune(s[left])] == 0 {
				left++
			}
			for mapping[rune(s[right])] == 0 {
				right--
			}
			if mapping[rune(s[left])] == 1 && mapping[rune(s[right])] == 1 {
				return s[left : right+1]
			} else if mapping[rune(s[left])] == 1 {
				mapping[rune(s[right])]--
				defer func() {
					mapping[rune(s[right])]++
				}()
				return find(s, left, right-1, mapping)
			} else if mapping[rune(s[right])] == 1 {
				mapping[rune(s[left])]--
				defer func() {
					mapping[rune(s[left])]++
				}()
				return find(s, left+1, right, mapping)
			} else {
				mapping[rune(s[left])]--
				r1 := find(s, left+1, right, mapping)
				mapping[rune(s[left])]++
				mapping[rune(s[right])]--
				r2 := find(s, left, right-1, mapping)
				mapping[rune(s[right])]++
				if len(r1) < len(r2) {
					return r1
				} else {
					return r2
				}
			}
		}
		return ""
	}
	return find(s, 0, len(s)-1, mapping)
}

func main() {
	fmt.Println(smallest("zyxabczyyx", []rune{'x', 'y', 'z'}))
}
