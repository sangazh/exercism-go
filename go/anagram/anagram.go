package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	if len(subject) == 0 || len(candidates) == 0 {
		return []string{}
	}

	lower := strings.ToLower(subject)
	letters := makeSet(lower)

	result := make([]string, 0)
	for _, candidate := range candidates {
		if len(subject) != len(candidate) {
			continue
		}
		candidateLower := strings.ToLower(candidate)

		if lower == candidateLower {
			continue
		}
		candidateSet := makeSet(candidateLower)

		if letters.Equal(candidateSet) {
			result = append(result, candidate)
		}
	}

	return result
}

type Set map[rune]int

func (s Set) Equal(set2 Set) bool {
	if len(s) != len(set2) {
		return false
	}
	for r, cnt := range s {
		if v, ok := set2[r]; !ok {
			return false
		} else if v != cnt {
			return false
		}
	}
	return true
}

func makeSet(word string) Set {
	set := make(map[rune]int)
	for _, l := range word {
		set[l] += 1
	}

	return set
}
