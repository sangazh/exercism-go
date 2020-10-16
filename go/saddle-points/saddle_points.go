package matrix

import (
	"fmt"
)

type Pair struct {
	row, col int
}

type Point struct {
	row, col, value int
}

func (p Point) String() string {
	return fmt.Sprintf("[%d,%d] %d", p.row, p.col, p.value)
}

func (p *Point) Equal(p2 *Point) bool {
	return p.value == p2.value
}
func (p *Point) Greater(p2 *Point) bool {
	return p.value > p2.value
}
func (p *Point) Less(p2 *Point) bool {
	return p.value < p2.value
}

func (p *Point) IsSame(p2 *Point) bool {
	return p.row == p2.row && p.col == p2.col
}

func (m *Matrix) Saddle() []Pair {
	// find the point with max value in the row
	rowMax := make([][]*Point, m.row)
	for j, row := range m.Rows() {
		for i, v := range row {
			p := &Point{j, i, v}
			if rowMax[j] == nil {
				rowMax[j] = []*Point{p}
			} else {
				if p.Greater(rowMax[j][0]) {
					rowMax[j] = []*Point{p}
				} else if p.Equal(rowMax[j][0]) {
					rowMax[j] = append(rowMax[j], p)
				}
			}
		}
	}

	// find the point with minimum value in the column
	colMin := make([][]*Point, m.col)
	for j, col := range m.Cols() {
		for i, v := range col {
			p := &Point{i, j, v}
			if colMin[j] == nil {
				colMin[j] = []*Point{p}
			} else {
				if p.Less(colMin[j][0]) {
					colMin[j] = []*Point{p}
				} else if p.Equal(colMin[j][0]) {
					colMin[j] = append(colMin[j], p)
				}
			}
		}
	}

	//compare if there are some points
	out := make([]Pair, 0)
	for _, row := range rowMax {
		for _, r := range row {
			for _, col := range colMin {
				for _, c := range col {
					if r.IsSame(c) {
						out = append(out, Pair{r.row, r.col})
					}
				}
			}
		}

	}

	return out
}
