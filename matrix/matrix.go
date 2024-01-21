package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	m := make(Matrix, len(rows))
	for i, row := range rows {
		cols := strings.Split(strings.TrimSpace(row), " ")
		m[i] = make([]int, len(cols))
		if len(cols) != len(m[0]) {
			return nil, errors.New("uneven rows")
		}
		for j, col := range cols {
			if n, err := strconv.Atoi(col); err != nil {
				return nil, err
			} else {
				m[i][j] = n
			}
		}
	}
	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for i := 0; i < len(m[0]); i++ {
		cols[i] = make([]int, len(m))
		for j := 0; j < len(m); j++ {
			cols[i][j] = m[j][i]
		}
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i := 0; i < len(m); i++ {
		rows[i] = make([]int, len(m[i]))
		copy(rows[i], m[i])
	}
	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) {
		return false
	}
	if col < 0 || col >= len(m[row]) {
		return false
	}
	m[row][col] = val
	return true
}
