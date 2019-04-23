package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	sentence := "One for %s, one for me."
	if len(name) > 0 {
		return fmt.Sprintf(sentence, name)
	}
	return fmt.Sprintf(sentence, "you")
}
