// https://leetcode.com/problems/merge-k-sorted-lists
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var begin, tracker *ListNode
	for {
		min := -1
		for i := range lists {
			if lists[i] != nil {
				if min < 0 {
					min = i
				} else if lists[min].Val > lists[i].Val {
					min = i
				}
			}
		}
		if min == -1 {
			break
		} else {
			if begin == nil {
				begin = lists[min]
				lists[min] = lists[min].Next
				tracker = begin
			} else {
				tracker.Next = lists[min]
				lists[min] = lists[min].Next
				tracker = tracker.Next
			}
		}
	}
	return begin
}
