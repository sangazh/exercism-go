package scale

import (
	"strings"
)

var pitchesSharp = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var pitchesFlat = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

func Scale(tonic, interval string) []string {
	var pitches []string

	if useFlat(tonic) {
		pitches = append(pitchesFlat, pitchesFlat...)
	} else {
		pitches = append(pitchesSharp, pitchesSharp...)
	}

	for i, p := range pitches {
		if strings.ToLower(tonic) == strings.ToLower(p) {
			pitches = pitches[i:]
			break
		}
	}
	if len(interval) == 0 {
		return pitches[:12]
	}

	result := []string{pitches[0]}
	for _, i := range interval {
		switch string(i) {
		case "m":
			pitches = pitches[1:]
		case "M":
			pitches = pitches[2:]
		case "A":
			pitches = pitches[3:]
		}
		result = append(result, pitches[0])
	}

	return result[:len(interval)]
}

// Use Flats: F, Bb, Eb, Ab, Db, Gb major d, g, c, f, bb, eb minor
func useFlat(tonic string) bool {
	if len(tonic) == 1 {
		switch tonic {
		case "F", "d", "g", "c", "f":
			return true
		}
	}

	if len(tonic) == 2 {
		if string(tonic[1]) == "b" {
			return true
		}
	}
	return false
}
