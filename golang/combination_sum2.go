package main

import "fmt"

func combinationSum2(candidates []int, target int) [][]int {
	mapping := make(map[int]int)
	for _, v := range candidates {
		mapping[v]++
	}

	var combine func(int, map[int]int) [][]int
	combine = func(target int, mapping map[int]int) [][]int {
		fmt.Println(target)
		res := make([][]int, 0)
		for k, v := range mapping {
			if k == target && v > 0 {
				res = append(res, []int{k})
			} else if k > target && v > 0 {
				mapping[k]--
				ret := combine(k-target, mapping)
				mapping[k]++
				for _, c := range ret {
					res = append(res, append([]int{k}, c...))
				}
			}
		}
		return res
	}

	return combine(target, mapping)
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
