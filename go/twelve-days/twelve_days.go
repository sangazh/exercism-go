package twelve

import (
	"bytes"
	"fmt"
)

// twelve days lyrics
var twelveDays = []string{
	"",
	1:  "a Partridge in a Pear Tree",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

// twelve
var twelve = []string{
	"",
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

//generate whole song lyrics
func Song() string {
	var lyrics bytes.Buffer
	for i := 0; i < 12; i++ {
		lyrics.WriteString(Verse(i + 1))
		lyrics.WriteString("\n")
	}
	return lyrics.String()
}

// generate specified day lyrics
func Verse(day int) string {
	lyrics := "On the %s day of Christmas my true love gave to me, %s"
	var buf bytes.Buffer
	for i := 12; i >= 1; i-- {
		lyric := twelveDays[i]
		if day == 1 && i == 1 {
			buf.WriteString(lyric)
			buf.WriteString(".")
			break
		}
		if i <= day && i != 1 {
			buf.WriteString(lyric)
			buf.WriteString(", ")
			continue
		}
		if day > 1 && i == 1 {
			buf.WriteString("and ")
			buf.WriteString(lyric)
			buf.WriteString(".")
			continue
		}

	}
	return fmt.Sprintf(lyrics, twelve[day], buf.String())
}
