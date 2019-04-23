package gigasecond

import (
	"time"
	"math"
)

// AddGigasecond
func AddGigasecond(t time.Time) time.Time {
	gs := int64(math.Pow10(9))
	t = t.Add(time.Duration(gs) * time.Second)
	return t
}
