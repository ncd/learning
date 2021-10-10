package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func newNode(value int) *ListNode {
	n := new(ListNode)
	n.Val = value
	return n
}

func (v *ListNode) addPrev(value int) *ListNode {
	n := new(ListNode)
	n.Val = value
	n.Next = v
	return n
}

func (v *ListNode) print() {
	for v != nil {
		fmt.Printf("%d ", v.Val)
		v = v.Next
	}
	fmt.Println()
}

func swapPairs(head *ListNode) *ListNode {
	var begin, tracker *ListNode
	for head != nil {
		if head.Next != nil {
			if begin == nil {
				begin = head.Next
				head.Next = head.Next.Next
				begin.Next = head
				head = begin.Next.Next
				tracker = begin.Next
			} else {
				// fmt.Println(head.Val)
				// fmt.Println(head.Next.Val)
				tracker.Next = head.Next
				temp := head.Next.Next
				head.Next = head.Next.Next
				tracker.Next.Next = head
				tracker = tracker.Next.Next
				head = temp
			}
		} else {
			if begin == nil {
				begin = head
			} else {
				tracker.Next = head
			}
			break
		}
	}
	return begin
}

func main() {
	n := newNode(4).addPrev(3).addPrev(2).addPrev(1)
	n.print()
	m := swapPairs(n)
	m.print()
}
