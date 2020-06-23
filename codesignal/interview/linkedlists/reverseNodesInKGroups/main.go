package main

import (
	"fmt"
)

func reverseNodesInKGroups(l *ListNode, k int) *ListNode {
	var (
		n     *ListNode
		pHead *ListNode
		lPtr  *ListNode
		pPtr  *ListNode
		iPtr  *ListNode
	)

	refHead := l
	for refHead != nil {

		for i := 1; i <= k; i++ {
			if refHead == nil {
				break
			}
			refHead, n = shift(refHead)

			if i == 1 {
				lPtr = n
			}
			if i > 1 {
				n.Next = iPtr
			}
			if i == k {
				if pHead == nil && pPtr == nil {
					pHead = n
				}
				if pPtr != nil {
					pPtr.Next = n
				}
				pPtr = lPtr
			}
			if refHead == nil && i < k {
				n.print()
				pPtr.Next = n.reverse()
				break
			}
			iPtr = n
		}
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
func shift(head *ListNode) (nHead *ListNode, n *ListNode) {
	if head == nil {
		return nil, nil
	}
	head.Next, n, nHead = nil, head, head.Next
	return nHead, n
}
func (l *ListNode) reverse() *ListNode {
	pCurr := l
	var pHead *ListNode
	for {
		if pCurr == nil {
			break
		}
		pTemp := pCurr.Next
		pCurr.Next = pHead
		pHead = pCurr
		pCurr = pTemp
	}
	return pHead
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

	b := reverseNodesInKGroups(a.Head, 3)
	b.print()
}
