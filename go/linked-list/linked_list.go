package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
	n    int
}

var ErrEmptyList = errors.New("empty list")

func (e *Node) Next() *Node {
	if e == nil {
		return nil
	}
	return e.next
}
func (e *Node) Prev() *Node {
	if e == nil {
		return nil
	}
	return e.prev
}

func (e *Node) String() string {
	return fmt.Sprintf("{[%v]->%v}", e.Val, e.Next())
}

func NewList(args ...interface{}) *List {
	list := &List{}
	for _, v := range args {
		list.PushBack(v)
	}
	return list
}

func (l *List) String() string {
	return fmt.Sprintf("total:%d %v", l.n, l.head)
}
func (l *List) PushFront(v interface{}) {
	node := &Node{Val: v}
	if l == nil {
		l = &List{
			head: node,
			tail: node,
			n:    1,
		}
		return
	}
	l.n++

	if l.head == nil || l.tail == nil {
		l.head = node
		l.tail = node
		return
	}

	oldHead := l.head
	oldHead.prev = node
	node.next = oldHead

	l.head = node
}
func (l *List) PushBack(v interface{}) {
	node := &Node{Val: v}
	l.n++
	if l.head == nil || l.tail == nil {
		l.head = node
		l.tail = node
		return
	}

	oldTail := l.tail
	oldTail.next = node
	node.prev = oldTail

	l.tail = node
}

func (l *List) PopFront() (interface{}, error) {
	if l.n == 0 {
		return nil, ErrEmptyList
	}
	l.n--

	node := l.head
	l.head = node.Next()

	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	return node.Val, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.n == 0 {
		return nil, ErrEmptyList
	}
	l.n--
	node := l.Last()
	l.tail = node.Prev()
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}

	return node.Val, nil
}
func (l *List) Reverse() *List {
	if l == nil {
		return nil
	}

	newlist := new(List)

	for {
		v, err := l.PopBack()
		if err != nil {
			break
		} else {
			newlist.PushBack(v)
		}
	}
	*l = *newlist
	return l
}
func (l *List) First() *Node {
	if l == nil {
		return nil
	}
	return l.head
}
func (l *List) Last() *Node {
	if l == nil {
		return nil
	}
	return l.tail
}
