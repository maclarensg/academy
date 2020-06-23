package main

import "fmt"

func mergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var pHead, pPtr *ListNode
	var l1Tmp, l2Tmp *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	for l1 != nil && l2 != nil {
		if l1.Value.(int) < l2.Value.(int) {
			l1Tmp = l1.Next // save next ptr
			if pHead == nil {
				pHead = l1
				pPtr = l1
			}
			if pPtr != nil {
				pPtr.Next = l1
				pPtr = l1
			}
			l1 = l1Tmp

		} else {
			l2Tmp = l2.Next // save next ptr
			if pHead == nil {
				pHead = l2
				pPtr = l2
			}
			if pPtr != nil {
				pPtr.Next = l2
				pPtr = l2
			}
			l2 = l2Tmp
		}
	}
	if l2 != nil {
		pPtr.Next = l2
	}
	if l1 != nil {
		pPtr.Next = l1
	}
	return pHead
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
func main() {
	a := newNode(1)
	a.add(3)
	a.add(5)
	a.print()

	b := newNode(2)
	b.add(4)
	b.add(6)
	b.add(7)
	b.print()

	c := mergeTwoLinkedLists(a, b)
	c.print()

}
