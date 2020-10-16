package ocr

import (
	"strings"
)

var numMap = map[string]string{`
 _ 
| |
|_|
   `: "0", `
   
  |
  |
   `: "1", `
 _ 
 _|
|_ 
   `: "2", `
 _ 
 _|
 _|
   `: "3", `
   
|_|
  |
   `: "4", `
 _ 
|_ 
 _|
   `: "5", `
 _ 
|_ 
|_|
   `: "6", `
 _ 
  |
  |
   `: "7", `
 _ 
|_|
|_|
   `: "8", `
 _ 
|_|
 _|
   `: "9",
}

// step 1
func recognizeDigit(in string) string {
	v, ok := numMap[in]
	if !ok {
		return "?"
	}
	return v
}

// step 2
func Recognize(in string) (out []string) {
	lines := strings.Split(in, "\n")
	count := len(lines[1]) / 3
	if count == 1 {
		return []string{recognizeDigit(in)}
	}

	for i := 1; i < len(lines)-1; i = i + 4 {
		out = append(out, recognizeLine(lines[i:i+4], count))
	}

	return
}

func recognizeLine(in []string, count int) (out string) {
	numbers := make([]string, count)
	var n int
	for _, line := range in {
		for k := 0; k < len(line)-1; k = k + 3 {
			numbers[n] += line[k:k+3] + "\n"
			n++
		}
		n = 0
	}

	for _, d := range numbers {
		d = "\n" + strings.TrimRight(d, "\n")
		out += recognizeDigit(d)
	}
	return
}
