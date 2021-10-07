// https://leetcode.com/problems/two-sum
package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		left := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == left {
				return []int{i, j}
			}
		}
	}
	return nil
}
