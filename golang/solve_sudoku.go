package main

import "fmt"

func solveSudoku(board [][]byte) {
	isValid := func(x int, y int, board [][]byte) bool {
		for i := 0; i < 9; i++ {
			if i != y && board[x][y] == board[x][i] {
				return false
			}
			if i != x && board[x][y] == board[i][y] {
				return false
			}
		}
		x1, y1 := 3*(x/3), 3*(y/3)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if x != x1+i && y != y1+j && board[x][y] == board[x1+i][y1+j] {
					return false
				}
			}
		}
		return true
	}
	var solve func(int, int, [][]byte) bool
	solve = func(x int, y int, board [][]byte) bool {
		// fmt.Println(x, y)
		if y == 9 {
			return true
		}
		var x1, y1 int
		if x < 9 {
			x1 = x + 1
			y1 = y
		}
		if x1 == 9 {
			x1 = 0
			y1 = y + 1
		}
		if board[x][y] == '.' {
			for i := byte('1'); i <= byte('9'); i++ {
				board[x][y] = i
				if isValid(x, y, board) && solve(x1, y1, board) {
					return true
				}
				board[x][y] = '.'
			}
		} else {
			return solve(x1, y1, board)
		}
		return false
	}
	solve(0, 0, board)
}

func print(matrix [][]byte) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%c ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	matrix := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(matrix)
	print(matrix)
}
