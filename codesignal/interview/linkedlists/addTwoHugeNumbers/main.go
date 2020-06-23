package main

import "fmt"

func addTwoHugeNumbers(a *ListNode, b *ListNode) *ListNode {
	var a0, b0, res *ListNode
	var carry int

	for a != nil && b != nil {
		a, a0 = a.pop()
		b, b0 = b.pop()

		sum := a0.Value.(int) + b0.Value.(int) + carry
		carry = sum / 10000

		if res == nil {
			res = newNode(sum % 10000)
		} else {
			res.add(sum % 10000)
		}
	}

	for a != nil {
		a, a0 = a.pop()
		sum := a0.Value.(int) + carry
		carry = sum / 10000
		res.add(sum % 10000)
	}

	for b != nil {
		b, b0 = b.pop()
		sum := b0.Value.(int) + carry
		carry = sum / 10000
		res.add(sum % 10000)
	}

	if carry > 0 {
		res.add(carry)
	}

	return res.reverse()
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

func (l *ListNode) pop() (head *ListNode, item *ListNode) {
	h := l
	n := l

	for n.Next != nil {
		if n.Next.Next == nil {
			t := n.Next
			n.Next = nil
			return h, t
		}
		n = n.Next
	}

	return nil, n
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

func (l *ListNode) clear() {
	l = nil
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
	a := newNode(9876)
	a.add(5432)
	a.add(1999)
	a.print()

	b := newNode(1)
	b.add(8001)
	b.print()

	c := addTwoHugeNumbers(a, b)
	c.print()

}
