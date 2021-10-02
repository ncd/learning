// https://www.hackerrank.com/challenges/almost-sorted/problem
package main

import "fmt"

func almostSorted(arr []int32) {
	left := 0
	right := len(arr) - 1
	for arr[left] < arr[left+1] {
		left++
	}
	for arr[right] > arr[right-1] {
		right--
	}

	lcond, rcond := true, true
	if right < len(arr)-1 {
		rcond = arr[left] < arr[right+1]
	}
	if left > 0 {
		lcond = arr[right] > arr[left-1]
	}

	if lcond && rcond {
		if left+1 <= right-1 {
			innerInc := arr[left+1] <= arr[right-1]
			for i := left + 1; i < right-1; i++ {
				if innerInc != (arr[i] < arr[i+1]) {
					fmt.Println("no")
					return
				}
			}
			if innerInc && arr[left] > arr[right-1] && arr[right] < arr[left+1] {
				fmt.Println("yes")
				fmt.Printf("swap %d %d\n", left+1, right+1)
			} else if arr[left] > arr[left+1] && arr[right] < arr[right-1] {
				fmt.Println("yes")
				fmt.Printf("reverse %d %d\n", left+1, right+1)
			} else {
				fmt.Println("no")
			}
		} else {
			fmt.Println("yes")
			if left < right {
				fmt.Printf("swap %d %d\n", left+1, right+1)
			}
		}
	} else {
		fmt.Println("no")
	}
}
func main() {
	arr := []int32{4104, 8529, 49984, 54956, 63034, 82534, 84473, 86411, 92941, 95929, 108831, 894947, 125082, 137123, 137276, 142534, 149840, 154703, 174744, 180537, 207563, 221088, 223069, 231982, 249517, 252211, 255192, 260283, 261543, 262406, 270616, 274600, 274709, 283838, 289532, 295589, 310856, 314991, 322201, 339198, 343271, 383392, 385869, 389367, 403468, 441925, 444543, 454300, 455366, 469896, 478627, 479055, 484516, 499114, 512738, 543943, 552836, 560153, 578730, 579688, 591631, 594436, 606033, 613146, 621500, 627475, 631582, 643754, 658309, 666435, 667186, 671190, 674741, 685292, 702340, 705383, 722375, 722776, 726812, 748441, 790023, 795574, 797416, 813164, 813248, 827778, 839998, 843708, 851728, 857147, 860454, 861956, 864994, 868755, 116375, 911042, 912634, 914500, 920825, 979477}
	almostSorted(arr)
}
