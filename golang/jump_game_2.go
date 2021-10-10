// https://leetcode.com/problems/jump-game-ii
package main

func jump(nums []int) int {
	var counter, pos int
	n := len(nums) - 1
	for pos < n {
		left := n - pos
		if left <= nums[pos] {
			counter++
			break
		}
		var max int
		var maxPos int
		for i := 1; i <= nums[pos]; i++ {
			reach := i + nums[pos+i]
			if reach > max {
				max = reach
				maxPos = i
			}
		}

		pos += maxPos
		counter++
	}
	return counter
}
