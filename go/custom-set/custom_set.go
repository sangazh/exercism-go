package stringset

import (
	"fmt"
	"strings"
)

type Set map[string]bool

func New() Set {
	return make(map[string]bool)
}

func NewFromSlice(slice []string) Set {
	s := New()
	for _, k := range slice {
		s[k] = true
	}
	return s
}

func (s Set) String() string {
	if s.IsEmpty() {
		return "{}"
	}

	str := "{"
	for k := range s {
		str += fmt.Sprintf(`"%s", `, k)
	}
	str = strings.TrimRight(str, `, `)
	str += "}"

	return str
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(k string) bool {
	_, ok := s[k]
	return ok
}

func Subset(s1, s2 Set) bool {
	if len(s1) > len(s2) {
		return false
	}

	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if s2.Has(k) {
			return false
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func (s Set) Add(k string) {
	s[k] = true
}

func Intersection(s1, s2 Set) Set {
	is := New()
	for k := range s1 {
		if s2.Has(k) {
			is.Add(k)
		}
	}
	return is
}

func Difference(s1, s2 Set) Set {
	d := New()
	for k := range s1 {
		if !s2.Has(k) {
			d.Add(k)
		}
	}
	return d
}

func Union(s1, s2 Set) Set {
	u := New()
	for k := range s1 {
		u[k] = true
	}
	for k := range s2 {
		u[k] = true
	}

	return u
}
