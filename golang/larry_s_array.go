// https://www.hackerrank.com/challenges/larrys-array/problem
package main

import "fmt"

// ABCD --> CABD -->
//      --> ACDB --> CDAB --> DACB
//

// Brute-Force
func larrysArray(A []int32) string {
	rotateBack := func(A []int32, i int) {
		if i == len(A)-1 {
			A[i-2], A[i-1], A[i] = A[i-1], A[i], A[i-2]
		} else {
			A[i-1], A[i], A[i+1] = A[i], A[i+1], A[i-1]
		}
	}
	for i := 0; i < len(A)-2; i++ {
		min := i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[min] {
				min = j
			}
		}
		for min > i {
			rotateBack(A, min)
			min--
		}
	}
	if A[len(A)-1] < A[len(A)-2] {
		return "NO"
	}
	return "YES"
}

// Inversion https://en.wikipedia.org/wiki/Inversion_(discrete_mathematics)
// Let's consider three numbers A, B, C (A<B<C). In ABC->BCA->CAB rotation group, the inversion numbers are: 0, 2, 2.
// In ACB->CBA->BAC rotation group, the inversion numbers are: 1, 3, 1.
// Which means, rotation will never change the even/odd nature of the inversion number of the entire array.
// Also, we notice our final objective, sorted array, has an inversion number of 0.
// So we can first compute the inversion number of the array, and if it's even, we can safely print yes; if it's odd, print no.
// Lastly, there is an O(nlgn) algorithm for inversion number computation. We use O(n^2) since it's fast enough for this problem.
func larrysArray1(A []int32) string {
	inversion := 0
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				inversion += 1
			}
		}
	}
	if inversion%2 == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	fmt.Println(larrysArray1([]int32{3, 1, 2}))
	fmt.Println(larrysArray([]int32{1, 3, 4, 2}))
	fmt.Println(larrysArray([]int32{1, 2, 3, 5, 4}))
}
