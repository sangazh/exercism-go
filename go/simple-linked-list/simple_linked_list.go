package linkedlist

import (
	"errors"
)

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(nums []int) *List {
	l := new(List)
	for _, n := range nums {
		l.Push(n)
	}

	return l
}
func (l *List) Size() int {
	return l.size
}
func (l *List) Push(n int) {
	if l.size == 0 {
		l.head = &Element{
			data: n,
			next: nil,
		}
	} else {
		oldHead := l.head
		l.head = &Element{
			data: n,
			next: oldHead,
		}
	}
	l.size += 1
	return
}
func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, errors.New("empty list")
	}
	n := l.head
	l.head = n.next
	l.size -= 1
	return n.data, nil
}
func (l *List) Array() []int {
	nums := make([]int, 0)
	for l.size > 0 {
		n, _ := l.Pop()
		nums = append([]int{n}, nums...)
	}
	return nums
}
func (l *List) Reverse() *List {
	newList := new(List)
	for i := l.size; i > 0; i-- {
		n, _ := l.Pop()
		newList.Push(n)
	}

	return newList
}
