// https://leetcode.com/problems/valid-parentheses/
package main

import "fmt"

func isValid(s string) bool {
	bracket := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []rune
	for _, c := range s {
		if v, ok := bracket[c]; !ok {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			if pop != v {
				return false
			}
			stack = stack[:len(stack)-1]
			fmt.Println(stack)
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	x := false
	fmt.Println(x)
}
