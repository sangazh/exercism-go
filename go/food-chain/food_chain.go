package foodchain

import (
	"fmt"
	"strings"
)

var prefix = "I know an old lady who swallowed a %s."
var postfix = "I don't know why she swallowed the fly. Perhaps she'll die."
var dead = "She's dead, of course!"

type animal struct {
	name, desc string
	alive      bool
}

var verse = []animal{{},
	{name: "fly", alive: true},
	{name: "spider", desc: "It wriggled and jiggled and tickled inside her.", alive: true},
	{name: "bird", desc: "How absurd to swallow a bird!", alive: true},
	{name: "cat", desc: "Imagine that, to swallow a cat!", alive: true},
	{name: "dog", desc: "What a hog, to swallow a dog!", alive: true},
	{name: "goat", desc: "Just opened her throat and swallowed a goat!", alive: true},
	{name: "cow", desc: "I don't know how she swallowed a cow!", alive: true},
	{name: "horse", alive: false},
}

func Verse(n int) string {
	lyric := fmt.Sprintf(prefix, verse[n].name)
	lyrics := []string{lyric}
	if verse[n].alive {
		if n > 1 {
			lyrics = append(lyrics, verse[n].desc)
		}
		for i := n; i > 1; i-- {
			l := ""
			if i == 3 {
				//spider line needs the wriggled bla bla
				desc := strings.Join(strings.Fields(verse[i-1].desc)[1:], " ")
				l = fmt.Sprintf("She swallowed the %s to catch the %s that %s", verse[i].name, verse[i-1].name, desc)
			} else {
				l = fmt.Sprintf("She swallowed the %s to catch the %s.", verse[i].name, verse[i-1].name)
			}

			lyrics = append(lyrics, l)
		}
		lyrics = append(lyrics, postfix)
	} else {
		lyrics = append(lyrics, dead)
	}

	return strings.Join(lyrics, "\n")
}

func Verses(min, max int) string {
	result := make([]string, 0)
	for i := min; i <= max; i++ {
		result = append(result, Verse(i))
	}
	out := strings.Join(result, "\n\n")
	return out
}

func Song() string {
	return Verses(1, 8)
}
