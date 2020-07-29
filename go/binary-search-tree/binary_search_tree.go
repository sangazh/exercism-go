package binarysearchtree

import (
	"fmt"
)

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func (s SearchTreeData) String() string {
	return fmt.Sprintf("(%v < %d < %v)", s.left, s.data, s.right)
}

func Bst(n int) *SearchTreeData {
	return &SearchTreeData{data: n}
}
func (s *SearchTreeData) Insert(n int) {
	if n > s.data {
		if s.right == nil {
			s.right = Bst(n)
			return
		}
		s.right.Insert(n)
	} else {
		if s.left == nil {
			s.left = Bst(n)
			return
		}
		s.left.Insert(n)
	}
	return
}
func (s *SearchTreeData) MapString(f func(int) string) []string {
	result := make([]string, 0)
	if s.left != nil {
		left := s.left.MapString(f)
		result = append(left, result...)
	}

	result = append(result, f(s.data))

	if s.right != nil {
		right := s.right.MapString(f)
		result = append(result, right...)
	}

	return result
}
func (s *SearchTreeData) MapInt(f func(int) int) []int {
	result := make([]int, 0)
	if s.left != nil {
		left := s.left.MapInt(f)
		result = append(left, result...)
	}

	result = append(result, f(s.data))

	if s.right != nil {
		right := s.right.MapInt(f)
		result = append(result, right...)
	}

	return result
}
