package robotname

import (
	"fmt"
	"master/utils/log"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var names = make(map[string]struct{})

const max = 26 * 26 * 1000

func (r *Robot) Name() string {
	if len(r.name) > 0 {
		return r.name
	}

	r.name = getName()
	return r.name
}

func getName() string {
	name := genName()
	for {
		if _, ok := names[name]; ok {
			name = genName()
		} else {
			break
		}
	}

	names[name] = struct{}{}
	if len(names) > max {
		log.Fatal("no more robot!")
	}
	return name
}

func genName() string {
	i := rand.Int63n(time.Now().UnixNano())
	j := rand.Int63n(time.Now().UnixNano())
	d := i % 1000

	return fmt.Sprintf("%s%s%03d", string(alpha[i%26]), string(alpha[j%26]), d)

}

func (r *Robot) Reset() {
	r.name = ""
}
