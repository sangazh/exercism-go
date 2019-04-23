package atbash

import (
	"bytes"
	"strings"
	"unicode"
)

var plain = "abcdefghijklmnopqrstuvwxyz"
var cipher = "zyxwvutsrqponmlkjihgfedcba"

func Atbash(text string) (result string) {
	text = preHandle(text)
	buf := new(bytes.Buffer)
	for i, s := range text {
		if unicode.IsLetter(s) {
			index := strings.Index(plain, string(s))
			buf.WriteByte(cipher[index])
		} else {
			buf.WriteRune(s)
		}
		if (i+1)%5 == 0 {
			buf.WriteByte(' ')
		}
	}
	return strings.TrimRight(buf.String(), " ")
}

// pre-handle the input, remove spaces, commas, dots.
func preHandle(text string) string {
	replacer := strings.NewReplacer(" ", "", ",", "", ".", "")
	text = replacer.Replace(text)
	text = strings.ToLower(text)
	return text
}
