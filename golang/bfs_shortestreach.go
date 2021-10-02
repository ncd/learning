// https://www.hackerrank.com/challenges/bfsshortreach/problem
package main

func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	// Write your code here
	data := make([]int32, n)
	queue := make([]int32, 0)
	for i := range data {
		data[i] = -1
	}
	data[s-1] = 0
	queue = append(queue, s)
	i := 0
	for i < len(queue) {
		start := queue[i]
		for j := 0; j < len(edges); j++ {
			var end int32
			if edges[j][0] == start {
				end = edges[j][1]
			} else if edges[j][1] == start {
				end = edges[j][0]
			}
			if end != 0 {
				// We minus 1 because we use 0-indexed array but the vertex is 1-indexed
				if data[end-1] == -1 {
					data[end-1] = data[start-1] + 6
					queue = append(queue, end)
				}
			}
		}
		i++
	}
	return append(data[:s-1], data[s:]...)
}
