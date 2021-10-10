// Find n value that sum equal target
package main

import (
	"fmt"
	"sort"
)

func n_sum(nums []int, length int, target int) [][]int {
	if length == 2 {
		return two_sum(nums, target)
	} else {
		var result [][]int
		for i := 0; i < len(nums); i++ {
			newTarget := target - nums[i]
			res := n_sum(append([]int{}, nums[i+1:]...), length-1, newTarget)
			for _, v := range res {
				result = append(result, append(v, nums[i]))
			}
		}
		return result
	}
}

func two_sum(nums []int, target int) [][]int {
	begin, end := 0, len(nums)-1
	var result [][]int
	for begin < end {
		sum := nums[begin] + nums[end]
		if sum < target {
			begin++
		} else if sum > target {
			end--
		} else {
			result = append(result, []int{nums[begin], nums[end]})
			begin += 1
			end -= 1
			for begin < end && nums[begin] == nums[begin-1] {
				begin += 1
			}
			for end > begin && nums[end] == nums[end+1] {
				end -= 1
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	sort.Ints(nums)
	fmt.Println(n_sum(nums, 4, 0))
}
