// https://leetcode.com/problems/largest-rectangle-in-histogram
package main

func largestRectangleArea(heights []int) int {
	var max int
	l := len(heights)
	left := make([]int, len(heights))
	right := make([]int, len(heights))
	left[0] = -1
	right[l-1] = l
	for i := 1; i < l; i++ {
		if heights[i] > heights[i-1] {
			left[i] = i - 1
		} else {
			temp := left[i-1]
			for temp != -1 && heights[i] <= heights[temp] {
				temp = left[temp]
			}
			left[i] = temp
		}
	}
	for i := l - 2; i >= 0; i-- {
		if heights[i] > heights[i+1] {
			right[i] = i + 1
		} else {
			temp := right[i+1]
			for temp != l && heights[i] <= heights[temp] {
				temp = right[temp]
			}
			right[i] = temp
		}
	}
	for i := 0; i < l; i++ {
		area := heights[i] * (right[i] - left[i] - 1)
		if max < area {
			max = area
		}
	}

	return max
}
