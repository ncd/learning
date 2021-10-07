package main

import "fmt"

func matrixRotation(matrix [][]int32, r int32) {
	min := func(v1 int, v2 int, v3 int, v4 int) int {
		min := v1
		if min > v2 {
			min = v2
		}
		if min > v3 {
			min = v3
		}
		if min > v4 {
			min = v4
		}
		return min
	}
	// Write your code here
	row := len(matrix)
	column := len(matrix[0])
	output := make([][]int32, row)
	for i := range output {
		output[i] = make([]int32, column)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			base := min(i, j, row-i-1, column-j-1)
			width := column - 2*base
			height := row - 2*base
			left := int(r) % (2 * (width + height - 2))
			newi := i - base
			newj := j - base
			for left != 0 {
				if newi == 0 {
					if left > newj {
						left -= newj
						newj = 0
					} else {
						newj -= left
						left = 0
					}
				}
				if newj == 0 {
					if left > height-newi-1 {
						left -= (height - 1 - newi)
						newi = height - 1
					} else {
						newi += left
						left = 0
					}
				}
				if newi == height-1 {
					if left > width-1-newj {
						left -= (width - 1 - newj)
						newj = width - 1
					} else {
						newj += left
						left = 0
					}
				}
				if newj == width-1 {
					if left > newi {
						left -= newi
						newi = 0
					} else {
						newi -= left
						left = 0
					}
				}
			}
			newi += base
			newj += base
			output[newi][newj] = matrix[i][j]
		}
	}
	for i := range output {
		for j := range output[i] {
			fmt.Printf("%d ", output[i][j])
		}
		fmt.Println()
	}
}

func main() {
	matrix := [][]int32{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	matrixRotation(matrix, 2)
}
