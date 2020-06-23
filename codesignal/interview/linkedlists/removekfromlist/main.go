package main

import "fmt"

// ListNode struct
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func new(v interface{}) *ListNode {
	return &ListNode{
		Value: v,
	}
}

func (l *ListNode) add(v interface{}) {
	n := l
	for n != nil {
		if n.Next == nil {
			n.Next = new(v)
			return
		}
		n = n.Next
	}
}

func (l *ListNode) print() {
	n := l
	for n != nil {
		fmt.Printf("%d ", n.Value.(int))
		n = n.Next
	}
	fmt.Println()
}

func removeKFromList(l *ListNode, k int) *ListNode {
	h := l
	for h != nil {
		if h.Next != nil && h.Next.Value == k {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}
	if l != nil && l.Value == k {
		return l.Next
	}
	return l
}

func main() {
	node := new(1000)
	node.add(1000)
	node.print()
	res := removeKFromList(node, 1000)
	fmt.Println(res)
	// strings.ToUpper
}
