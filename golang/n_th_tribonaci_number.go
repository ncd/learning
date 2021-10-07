// https://leetcode.com/problems/n-th-tribonacci-number/
package main

import "fmt"

// Memorization
func solutionBasic(n int) int {
	if n == 0 {
		return 0
	}
	t0, t1, t2 := 0, 1, 1

	for i := 0; i < n-2; i++ {
		t := t0 + t1 + t2
		t0 = t1
		t1 = t2
		t2 = t
	}
	return t2
}

// Matrix multiplication
func solutionMatrix(n int) int {
	var matrixMultiply func([][]int, [][]int) [][]int
	var matrixPower func([][]int, int) [][]int
	matrixMultiply = func(a [][]int, b [][]int) [][]int {
		var data [][]int
		data = make([][]int, len(a))
		for i := range data {
			data[i] = make([]int, len(b[0]))
		}
		for i := range a {
			for j := range b[0] {
				sum := 0
				for k := range a[0] {
					sum += a[i][k] * b[k][j]
				}
				data[i][j] = sum
			}
		}
		return data
	}

	matrixPower = func(a [][]int, n int) [][]int {
		if n == 1 {
			return a
		}
		half := matrixPower(a, n/2)
		power := matrixMultiply(half, half)
		if n%2 == 0 {
			return power
		} else {
			return matrixMultiply(power, a)
		}
	}

	if n == 0 {
		return 0
	}
	matrix := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	}
	start := [][]int{
		{0}, {1}, {1},
	}

	end := matrixMultiply(matrixPower(matrix, n), start)

	return end[0][0]
}

func main() {
	fmt.Println(solutionMatrix(1000))
}
