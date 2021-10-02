package main

import (
	"fmt"
	"time"
)

func main() {
	grid := [][]int{
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, 0, -1, -1, -1},
		{-1, -1, -1, -1, 0, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{0, 0, -1, -1, -1, -1, -1},
		{0, 0, -1, -1, -1, -1, -1},
	}
	bomberManSimulation(grid)
}
func bomberManSimulation(grid [][]int) {
	counter := 0
	printGrid(counter, grid)
	time.Sleep(1 * time.Second)
	counter += 1
	doNothing(grid)
	printGrid(counter, grid)
	time.Sleep(1 * time.Second)
	for true {
		counter += 1
		updateGrid(grid)
		printGrid(counter, grid)
		time.Sleep(1 * time.Second)
		counter += 1
		exploseGrid(grid)
		printGrid(counter, grid)
		time.Sleep(1 * time.Second)
	}

}

func doNothing(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 0 {
				grid[i][j] += 1
			}
		}
	}
}

func exploseGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 0 {
				grid[i][j] += 1
			}
			if grid[i][j] == 3 {
				grid[i][j] = -1
				if i-1 >= 0 {
					grid[i-1][j] = -1
				}
				if i+1 < len(grid) {
					grid[i+1][j] = -1
				}
				if j-1 >= 0 {
					grid[i][j-1] = -1
				}
				if j+1 < len(grid[i]) {
					grid[i][j+1] = -1
				}
			}
		}
	}
}

func updateGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == -1 {
				grid[i][j] = 0
			} else {
				grid[i][j] += 1
			}
		}
	}
}

func printGrid(counter int, grid [][]int) {
	fmt.Print("At second ")
	fmt.Println(counter)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == -1 {
				fmt.Print(". ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println("")
	}
}
