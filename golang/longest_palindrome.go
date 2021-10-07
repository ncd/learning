package main

import "fmt"

func longestPalindrome(s string) string {
	var max int
	var res string
	for i := 0; i < len(s); i++ {
		n := 1
		for i-n >= 0 && i+n < len(s) && s[i-n] == s[i+n] {
			n++
		}
		if max < 1+2*(n-1) {
			max = 1 + 2*(n-1)
			res = s[i-n+1 : i+n]
		}
		if i < len(s)-1 && s[i] == s[i+1] {
			n := 1
			for i-n >= 0 && i+1+n < len(s) && s[i-n] == s[i+1+n] {
				n++
			}
			if max < 2+2*(n-1) {
				max = 2 + 2*(n-1)
				res = s[i-n+1 : i+1+n]
			}
		}
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("a"))
}
