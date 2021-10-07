package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) < len(nums1) {
		return findMedianSortedArrays(nums2, nums1)
	}
	n1, n2 := len(nums1), len(nums2)
	l, r := 0, n1
	c := (n1 + n2) % 2
	middle := (n1 + n2) / 2
	for l <= r {
		m1 := (l + r) / 2
		m2 := middle - m1
		var left1, left2, right1, right2 int
		if m1 == 0 {
			left1 = math.MinInt
		} else {
			left1 = nums1[m1-1]
		}
		if m2 == 0 {
			left2 = math.MinInt
		} else {
			left2 = nums2[m2-1]
		}
		if m1 == n1 {
			right1 = math.MaxInt
		} else {
			right1 = nums1[m1]
		}
		if m2 == n2 {
			right2 = math.MaxInt
		} else {
			right2 = nums2[m2]
		}
		if left1 < right2 && left2 < right1 {
			if c == 0 {
				return (math.Max(float64(left1), float64(left2)) + math.Min(float64(right1), float64(right2))) / 2
			} else {
				return math.Max(float64(left1), float64(left2))
			}
		} else if left1 > right2 {
			r = m1 - 1
		} else {
			l = m1 + 1
		}
	}

	return 0.0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 7}))
}
