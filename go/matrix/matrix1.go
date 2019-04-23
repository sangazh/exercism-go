package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Matrix1 struct {
	m        [][]int
	row, col int
}

func New1(input string) (*Matrix, error) {
	m := new(Matrix)
	s := strings.Split(input, "\n")
	m.col = len(s)

	rows := strings.Fields(s[0])
	m.row = len(rows)
	m.init()

	for j, line := range s {
		nums := strings.Fields(line)
		if len(nums) != m.row {
			return nil, errors.New("not even row")
		}
		for i, n := range nums {
			fmt.Printf("col: [%d] row:[%d] v: %s, \n", j, i, n)
			number, err := strconv.Atoi(n)
			if err != nil {
				return nil, errors.New("not int")
			}
			m.Set(i, j, number)

		}
	}

	return m, nil
}

//
func (m *Matrix1) Rows() [][]int {
	fmt.Println("matrix:", m.m)
	result := make([][]int, m.col)
	for j, col := range m.m {
		result[j] = make([]int, m.row)
		for i, v := range col {
			result[j][i] = v
		}
	}
	fmt.Println("rows:", result)
	return result
}

func (m *Matrix1) Cols() [][]int {
	result := make([][]int, m.row)
	for j := 0; j < m.row; j++ {
		result[j] = make([]int, m.col)
	}

	for j, col := range m.m {
		for i, v := range col {
			result[i][j] = v
		}
	}
	return result
}

func (m *Matrix1) Set(r, c, val int) bool {
	m.m[c][r] = val
	//for j := 0; j < m.col; j++ {
	//	for i := 0; i < m.row; i++ {
	//		if j == c && i == r {
	//			return true
	//		}
	//	}
	//}

	return true
}

//matrix init
func (m *Matrix1) init() {
	m.m = make([][]int, m.col)
	for j := 0; j < m.col; j++ {
		m.m[j] = make([]int, m.row)
	}
}
