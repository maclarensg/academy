package main

import "fmt"

// ListNode struct
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func newNode(v interface{}) *ListNode {
	return &ListNode{
		Value: v,
	}
}

func (l *ListNode) add(v interface{}) {
	n := l
	for n != nil {
		if n.Next == nil {
			n.Next = newNode(v)
			return
		}
		n = n.Next
	}
}

func (l *ListNode) print() {
	n := l
	for n != nil {
		fmt.Printf("%v ", n.Value)
		n = n.Next
	}
	fmt.Println()
}

func isListPalindrome(l *ListNode) bool {
	var stack []interface{}
	n := l
	for n != nil {
		stack = append(stack, n.Value)
		n = n.Next
	}
	var i interface{}
	for l != nil {
		i, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if l.Value != i {
			return false
		}
		l = l.Next
	}

	return true
}

func main() {
	a := newNode(1)
	a.add(2)
	a.add(2)
	a.add(3)
	fmt.Println(isListPalindrome(a))
}
