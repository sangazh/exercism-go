package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	total, matchesWon, matchesDrawn, matchesLost, points int
	name                                                 string
}

const (
	WinPoint  = 3
	DrawPoint = 1
)

func NewTeam(name string) *Team {
	return &Team{name: name}
}

func (t *Team) win() {
	t.total += 1
	t.matchesWon += 1
}

func (t *Team) lost() {
	t.total += 1
	t.matchesLost += 1
}

func (t *Team) draw() {
	t.total += 1
	t.matchesDrawn += 1
}

func (t *Team) getPoints() int {
	t.points = t.matchesWon*WinPoint + t.matchesDrawn*DrawPoint
	return t.points
}

type Teams map[string]*Team

var teams Teams

func (*Teams) addTeam(name string) *Team {
	if _, ok := teams[name]; !ok {
		teams[name] = NewTeam(name)
	}
	return teams[name]
}

func (*Teams) reset() {
	teams = Teams{}
}

func Tally(r io.Reader, w io.Writer) error {
	teams.reset()
	reader := bufio.NewReader(r)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		s := string(line)
		result := strings.Split(s, ";")
		if len(result) == 1 {
			continue
		}
		if len(result) < 3 {
			return errors.New("wrong format")
		}

		t1 := teams.addTeam(result[0])
		t2 := teams.addTeam(result[1])

		switch result[2] {
		case "win":
			t1.win()
			t2.lost()
		case "loss":
			t1.lost()
			t2.win()
		case "draw":
			t1.draw()
			t2.draw()
		default:
			return errors.New("wrong format")
		}
	}

	sliceTeam := make([]*Team, 0, len(teams))
	for _, v := range teams {
		sliceTeam = append(sliceTeam, v)
	}

	sort.Slice(sliceTeam, func(i, j int) bool {
		b := sliceTeam[i].getPoints() > sliceTeam[j].getPoints()
		if sliceTeam[i].getPoints() == sliceTeam[j].getPoints() {
			b = sliceTeam[i].name < sliceTeam[j].name
		}
		return b
	})

	fmt.Fprintf(w, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	for _, v := range sliceTeam {
		fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n", v.name, v.total, v.matchesWon, v.matchesDrawn, v.matchesLost, v.points)
	}
	return nil
}
