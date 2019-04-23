package house

import (
	"fmt"
	"strings"
)

var lyricMap = [][]string{
	{},
	{"lay in", "the house that Jack built"},
	{"ate", "the malt"},
	{"killed", "the rat"},
	{"worried", "the cat"},
	{"tossed", "the dog"},
	{"milked", "the cow with the crumpled horn"},
	{"kissed", "the maiden all forlorn"},
	{"married", "the man all tattered and torn"},
	{"woke", "the priest all shaven and shorn"},
	{"kept", "the rooster that crowed in the morn"},
	{"belonged to", "the farmer sowing his corn"},
	{"", "the horse and the hound and the horn"},
}

func Song() string {
	verse := make([]string, 12)
	for i := 0; i < 12; i++ {
		verse[i] = Verse(i + 1)
	}

	return strings.Join(verse, "\n\n")
}

func Verse(line int) string {
	pattern := "%s %s"
	prefix, lyrics := "", make([]string, line)

	j := 1
	for i := line; i > 0; i-- {
		verb := lyricMap[i][0]
		single := lyricMap[i][1]
		switch {
		case j == 1:
			prefix = "This is"
		case j <= 12:
			prefix = "that " + verb
		}
		lyric := fmt.Sprintf(pattern, prefix, single)
		lyrics[j-1] = lyric
		j++
	}
	lyrics[line-1] += "."

	return strings.Join(lyrics, "\n")
}
