package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	result := make([]string, numRows)
	var index int
	direction := "down"
	for i := range s {
		result[index] += string(s[i])
		if direction == "down" {
			if index == numRows-1 {
				direction = "up"
				index--
			} else {
				index++
			}
		} else {
			if index == 0 {
				direction = "down"
				index++
			} else {
				index--
			}
		}
	}
	var ret string
	for i := 0; i < numRows; i++ {
		ret += result[i]
	}
	return ret
}

func main() {
	fmt.Println(convert("ABCD", 3))
}
