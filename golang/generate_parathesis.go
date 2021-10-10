// https://leetcode.com/problems/generate-parentheses/
package main

import "fmt"

func generateParenthesis(n int) []string {
	var generate func(int, int) []string
	generate = func(op int, cl int) []string {
		fmt.Println(op, cl)
		if op == 0 && cl == 0 {
			return []string{""}
		}
		var res []string
		if op != 0 {
			a := generate(op-1, cl)
			for _, c := range a {
				res = append(res, "("+c)
			}
		}
		if cl != 0 && cl > op {
			a := generate(op, cl-1)
			for _, c := range a {
				res = append(res, ")"+c)
			}
		}
		fmt.Println(res)
		return res
	}

	return generate(n, n)
}

func main() {
	x := generateParenthesis(3)
	fmt.Println(x)
}
