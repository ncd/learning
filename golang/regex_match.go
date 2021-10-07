// https://leetcode.com/problems/regular-expression-matching/
package main

import "fmt"

func isMatch(s string, p string) bool {
	fmt.Printf("%s %s\n", s, p)
	if s == "" && p == "" {
		return true
	}
	if s == "" {
		if len(p) >= 2 && p[1] == '*' {
			return isMatch(s, p[2:])
		} else {
			return false
		}
	} else if p == "" {
		return false
	} else {
		if s[0] == p[0] || p[0] == '.' {
			if len(p) >= 2 && p[1] == '*' {
				return isMatch(s[1:], p) || isMatch(s, p[2:])
			} else {
				return isMatch(s[1:], p[1:])
			}
		} else {
			if len(p) >= 2 && p[1] == '*' {
				return isMatch(s, p[2:])
			}
			return false
		}
	}
}

func main() {
	fmt.Println(isMatch("bbbba", ".*a*a"))
}
