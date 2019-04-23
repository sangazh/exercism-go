package raindrops

import (
	"bytes"
	"fmt"
)

//Convert convert a number into a string
func Convert(n int) string {
	raindrop := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	var buf bytes.Buffer
	for i := 1; i < n+1; i++ {
		if n%i == 0 {
			for factor, sing := range raindrop {
				if i == factor {
					buf.WriteString(sing)
				}
			}
		}
	}
	if buf.Len() > 0 {
		return buf.String()
	}
	return fmt.Sprint(n)
}
