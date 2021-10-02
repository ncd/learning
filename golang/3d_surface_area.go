package main

import (
  "fmt"
)

// Absolute permutation
func main() {
  fmt.Println(solution(10, 1))
}

func solution(n int32, k int32) []int32 {
  var res []int32
  mapping := make([]bool, n + 1)
  for i := int32(1); i <= n; i++ {
    if i > k && !mapping[i - k] {
      mapping[i - k] = true
      res = append(res, i - k)
      continue
    }
    if i + k <= n && !mapping[i + k] {
      mapping[i + k] = true
      res = append(res, i + k)
    } else {
      return []int32{-1}
    }
  }
  return res
}

