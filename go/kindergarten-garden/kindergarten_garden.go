package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

var plants = map[byte]string{
	'V': "violets",
	'R': "radishes",
	'G': "grass",
	'C': "clover",
}

type Garden struct {
	g map[string][]string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(children) == 0 || len(diagram) == 0 {
		return nil, errors.New("empty")
	}
	cups := strings.Split(diagram, "\n")
	if len(cups[1]) != len(children)*2 {
		return nil, errors.New("not match")
	}

	if !checkDuplicate(children) {
		return nil, errors.New("duplicate children")
	}

	if len(cups) != 3 {
		return nil, errors.New("diagram error")
	}

	c := make([]string, len(children))
	copy(c, children)

	sort.Strings(c)
	garden := make(map[string][]string)

	for i, child := range c {
		cupIdx := 2 * i
		plants, err := convertPlants([]byte{cups[1][cupIdx], cups[1][cupIdx+1], cups[2][cupIdx], cups[2][cupIdx+1]})
		if err != nil {
			return nil, err
		}
		garden[child] = plants
	}
	return &Garden{garden}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if v, ok := g.g[child]; ok {
		return v, true
	}
	return nil, false

}

func convertPlants(cups []byte) ([]string, error) {
	flowers := make([]string, 0, 4)
	for _, cup := range cups {
		v, ok := plants[cup]
		if !ok {
			return nil, errors.New("invalid diagram")
		}
		flowers = append(flowers, v)
	}
	return flowers, nil
}

func checkDuplicate(slice []string) bool {
	tmp := make(map[string]bool)

	for _, s := range slice {
		if _, ok := tmp[s]; !ok {
			tmp[s] = true
		} else {
			return false
		}
	}
	return true
}
