package main

import "fmt"

func rearrangeLastN(l *ListNode, n int) *ListNode {
	if n < 1 {
		return l
	}

	s, f := l, l

	var p *ListNode
	if f != nil {
		var i int
		for i = 1; i < n; i++ {
			if f.Next != nil {
				f = f.Next
			} else {
				break
			}
		}
		if i == n {
			for f.Next != nil {
				f = f.Next
				p = s
				s = s.Next
			}
		}
	}

	if p != nil {
		p.Next = nil
	} else {
		return l
	}

	h := s
	t := s
	for t != nil {
		if t.Next == nil {
			t.Next = l
			break
		}
		t = t.Next
	}

	return h
}

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
func (l *ListNode) print() {
	for l != nil {
		format := "%d,"
		if l.Next == nil {
			format = "%d"
		}
		fmt.Printf(format, l.Value.(int))
		l = l.Next
	}
	fmt.Println()
}

func (l *ListNode) add(v interface{}) {
	for l != nil {
		if l.Next == nil {
			l.Next = newNode(v)
			return
		}
		l = l.Next
	}
}

// LinkedList struct
type LinkedList struct {
	Head *ListNode
}

func newLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) add(v interface{}) {
	if ll.Head == nil {
		ll.Head = newNode(v)
	} else {
		ll.Head.add(v)
	}
}

func main() {
	a := newLinkedList()
	for _, n := range []int{1, 2, 3, 4, 5} {
		a.add(n)
	}

	a.Head.print()
	rearrangeLastN(a.Head, 1).print()
}
