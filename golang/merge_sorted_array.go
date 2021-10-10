package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j int
	if m == 0 {
		copy(nums1[1:], nums2)
		return
	}
	if n == 0 {
		return
	}
	for i < m {
		if nums1[i] <= nums2[j] {
			if i < m {
				i++
			}
		} else {
			nums1[i], nums2[j] = nums2[j], nums1[i]
			for k := j; k < n-1; k++ {
				if nums2[k] > nums2[k+1] {
					nums2[k], nums2[k+1] = nums2[k+1], nums2[k]
				}
			}
		}
	}

	copy(nums1[m:], nums2)
}

func main() {
	a := []int{0, 2, 3, 4}
	b := []int{1, 2}
	merge(a, 0, b, 1)
}
