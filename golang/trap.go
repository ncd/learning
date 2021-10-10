package main

func trap(height []int) int {
	left, right := 0, len(height)-1
	var maxLeft, maxRight int
	var res int
	for left < right {
		if height[left] > height[right] {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				res += (maxRight - height[right])
			}
			right--
		} else {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				res += (maxLeft - height[left])
			}
			left++
		}
	}
	return res
}
