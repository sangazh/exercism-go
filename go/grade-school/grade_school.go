package school

import (
	"sort"
)

type Grade struct {
	grade    int
	students []string
}

type School struct {
	s map[int][]string
}

func New() *School {
	return &School{make(map[int][]string)}
}

func (s *School) Add(name string, grade int) {
	if _, ok := s.s[grade]; ok {
		s.s[grade] = append(s.s[grade], name)
	} else {
		s.s[grade] = []string{name}
	}
}
func (s *School) Grade(grade int) []string {
	if _, ok := s.s[grade]; ok {
		return s.s[grade]
	}
	return nil
}
func (s *School) Enrollment() (g []Grade) {
	grads := make([]int, 0)
	for g := range s.s {
		grads = append(grads, g)
	}

	sort.Ints(grads)
	for _, grade := range grads {
		sort.Strings(s.s[grade])
		g = append(g, Grade{
			grade:    grade,
			students: s.s[grade]},
		)
	}
	return
}
