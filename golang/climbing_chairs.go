// https://leetcode.com/problems/climbing-stairs/
package main

import "fmt"

func climbStairs(n int) int {
	memois := make([]int, n+1)
	memois[0] = 1
	memois[1] = 1

	for i := 2; i <= n; i++ {
		memois[i] = memois[i-1] + memois[i-2]
	}
	return memois[n]
}

func main() {
	fmt.Println(climbStairs(5))
}
