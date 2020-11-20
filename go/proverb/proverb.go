package proverb

import (
	"fmt"
)

const (
	verb = "For want of a %s the %s was lost."
	last = "And all for the want of a %s."
)

func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	verbs := make([]string, 0)
	for i, r := range rhyme[:len(rhyme)-1] {
		verbs = append(verbs, fmt.Sprintf(verb, r, rhyme[i+1]))
	}

	verbs = append(verbs, fmt.Sprintf(last, rhyme[0]))
	return verbs

}
