package cipher

import (
	"strings"
	"unicode"
)

func NewCaesar() Cipher {
	return NewShift(3)
}

type ShiftCipher struct {
	distance int32
}

func (c *ShiftCipher) Encode(input string) (output string) {
	input = strings.ToLower(input)
	for _, s := range input {
		if unicode.IsLetter(s) {
			ss := getLetter(s + c.distance)
			output += string(ss)
		}
	}

	return
}
func (c *ShiftCipher) Decode(input string) (output string) {
	for _, s := range input {
		ss := getLetter(s - c.distance)
		output += string(ss)
	}
	return
}

func NewShift(distance int) Cipher {
	//Argument for NewShift must be in the range 1 to 25 or -1 to -25. Zero is disallowed.
	if distance > 25 || distance == 0 || distance < -25 {
		return nil
	}
	return &ShiftCipher{int32(distance)}
}

type VigenereCipher struct {
	key string
}

func (c *VigenereCipher) Encode(input string) (output string) {
	input = strings.ToLower(input)
	var ninput string

	for _, s := range input {
		if unicode.IsLetter(s) {
			ninput += string(s)
		}
	}
	for i, s := range ninput {
		distance := c.key[i%len(c.key)] - 'a'
		ss := getLetter(s + int32(distance))
		output += string(ss)
	}

	return
}
func (c *VigenereCipher) Decode(input string) (output string) {
	for i, s := range input {
		distance := c.key[i%len(c.key)] - 'a'
		ss := getLetter(s - int32(distance))
		output += string(ss)
	}

	return
}

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}

	mk := make(map[rune]int)

	//Argument for NewVigenere must consist of lower case letters a-z only
	for _, k := range key {
		if !unicode.IsLower(k) {
			return nil
		}
		mk[k]++
	}

	//Values consisting entirely of the letter 'a' are disallowed
	if len(mk) == 1 {
		if v, ok := mk['a']; ok {
			if v == len(key) {
				return nil
			}
		}
	}

	return &VigenereCipher{key}
}

func getLetter(r rune) rune {
	if r > 'z' {
		r -= 26
	} else if r < 'a' {
		r += 26
	}
	return r
}
