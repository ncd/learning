// https://leetcode.com/problems/add-two-numbers
package main

/**
 * Definition for singly-linked list.

 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sumBack func(*ListNode, *ListNode, int) *ListNode
	sumBack = func(l1 *ListNode, l2 *ListNode, left int) *ListNode {
		node := new(ListNode)
		var val int
		if l1 == nil && l2 == nil {
			if left != 0 {
				node.Val = left
				node.Next = nil
				return node
			} else {
				return nil
			}
		}
		var l1Next, l2Next *ListNode
		if l1 == nil {
			val = l2.Val + left
			l1Next = nil
			l2Next = l2.Next
		} else if l2 == nil {
			val = l1.Val + left
			l2Next = nil
			l1Next = l1.Next
		} else {
			val = l1.Val + l2.Val + left
			l1Next = l1.Next
			l2Next = l2.Next
		}

		node.Val = val % 10
		nextNode := sumBack(l1Next, l2Next, val/10)
		node.Next = nextNode

		return node
	}

	return sumBack(l1, l2, 0)
}
