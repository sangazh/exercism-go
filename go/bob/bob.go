// Package bob responses randomly
package bob

import (
	"strings"
)

// Hey accept a input and returned a response
func Hey(remark string) string {
	cat := category(remark)
	return responses[cat]
}

const (
	Default = iota
	Question
	Yelling
	YellingQuestion
	Empty
)

// responses with different type
var responses = map[int]string{
	Default:         "Whatever.",
	Question:        "Sure.",
	Yelling:         "Whoa, chill out!",
	YellingQuestion: "Calm down, I know what I'm doing!",
	Empty:           "Fine. Be that way!",
}

//parse input string to find what type of the response should be returned
func category(s string) int {
	f := strings.Fields(s)
	if len(f) == 0 {
		return Empty
	}

	s = strings.TrimSpace(s)

	upper := strings.ToUpper(s)
	if !strings.ContainsAny(upper, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		if !strings.HasSuffix(s, "?") {
			return Default
		} else {
			return Question
		}
	}

	if upper == s {
		if strings.HasSuffix(s, "?") {
			return YellingQuestion
		}
		return Yelling
	}

	if strings.HasSuffix(s, "?") {
		return Question
	}

	return Default
}
