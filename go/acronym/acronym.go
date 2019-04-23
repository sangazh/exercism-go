// Package acronym should have a package comment that summarizes what it's about.
package acronym

import (
	"bytes"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	fields := strings.Fields(s)
	var result bytes.Buffer
	for _, f := range fields {
		result.WriteByte(f[0])
	}
	return strings.ToUpper(result.String())
}
