package main

func isValidSudoku(board [][]byte) bool {
	for i := range board {
		for j := range board[0] {
			if board[i][j] != '.' {
				for k := i + 1; k < len(board); k++ {
					if board[i][j] == board[k][j] {
						return false
					}
				}
				for k := j + 1; k < len(board[0]); k++ {
					if board[i][j] == board[i][k] {
						return false
					}
				}

				for n := 0; n < 3; n++ {
					for m := 0; m < 3; m++ {
						i1 := 3*(i/3) + n
						j1 := 3*(j/3) + m
						if i1 == i && j1 == j {
							continue
						}
						if board[i][j] == board[i1][j1] {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
