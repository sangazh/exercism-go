package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	m        [][]int
	row, col int
}

//create a matrix based on given string numbers
func New(input string) (*Matrix, error) {
	s := strings.Split(input, "\n")

	m := &Matrix{
		row: len(s),
		col: len(strings.Fields(s[0])),
	}
	m.init()

	for i, line := range s {
		nums := strings.Fields(line)
		if len(nums) != m.col {
			return nil, errors.New("not even col")
		}
		for j, n := range nums {
			val, err := strconv.Atoi(n)
			if err != nil {
				return nil, errors.New("not int")
			}
			m.Set(i, j, val)
		}
	}

	return m, nil
}

//return rows
func (m *Matrix) Rows() [][]int {
	result := make([][]int, m.row)
	for j, row := range m.m {
		result[j] = make([]int, m.col)
		for i, v := range row {
			result[j][i] = v
		}
	}
	return result
}

//return columns
func (m *Matrix) Cols() [][]int {
	result := make([][]int, m.col)
	for j := 0; j < m.col; j++ {
		result[j] = make([]int, m.row)
	}

	for j, row := range m.m {
		for i, v := range row {
			result[i][j] = v
		}
	}
	return result
}

//set a matrix value by it's row and col index
func (m *Matrix) Set(r, c, val int) bool {
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.col; j++ {
			if i == r && j == c {
				m.m[r][c] = val
				return true
			}
		}
	}

	return false
}

//matrix init
func (m *Matrix) init() {
	m.m = make([][]int, m.row)
	for i := 0; i < m.row; i++ {
		m.m[i] = make([]int, m.col)
	}
}
