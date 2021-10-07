// https://leetcode.com/problems/reverse-integer
// const (
// 	MaxInt    = 1<<(intSize-1) - 1
// 	MinInt    = -1 << (intSize - 1)
// 	MaxInt8   = 1<<7 - 1
// 	MinInt8   = -1 << 7
// 	MaxInt16  = 1<<15 - 1
// 	MinInt16  = -1 << 15
// 	MaxInt32  = 1<<31 - 1
// 	MinInt32  = -1 << 31
// 	MaxInt64  = 1<<63 - 1
// 	MinInt64  = -1 << 63
// 	MaxUint   = 1<<intSize - 1
// 	MaxUint8  = 1<<8 - 1
// 	MaxUint16 = 1<<16 - 1
// 	MaxUint32 = 1<<32 - 1
// 	MaxUint64 = 1<<64 - 1
// )

package main

import "fmt"

const MaxInt32 = 1<<31 - 1
const MinInt32 = -1 << 31

func reverse(x int) int {
	var result int
	var neg bool
	if x < 0 {
		neg = true
		x = -x
	}
	for x != 0 {
		temp := x % 10
		if result <= (MaxInt32-temp)/10 {
			result = (result * 10) + temp
		} else {
			return 0
		}

		x = x / 10
	}
	if neg == true {
		return -result
	} else {
		return result
	}
}

func main() {
	fmt.Println(MaxInt32)
	fmt.Println(-MaxInt32)
	fmt.Println(MinInt32)
}
