package matrix

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Matrix contains the x y coordinates.
type Matrix [][]int

// New takes a string and parses it into a Matrix.
// If the string is invalid a an error will get returned.
//    "1 2 5\n10 20 30" -> int[[1,2,5], [10,20,30]]
//    "wow" -> error("no int")
//    "2.5" -> error("no int")
//    "4 1\n1 11 22" -> error("uneven")
func New(in string) (Matrix, error) {
	// Check if we have a trailing empty row.
	if strings.LastIndex(in, "\n") == len(in)-1 {
		return Matrix{}, errors.New("trailing empty row")
	}

	var result [][]int
	var scanner = bufio.NewScanner(strings.NewReader(in))

	for scanner.Scan() {
		i := len(result)
		result = append(result, []int{})
		line := strings.TrimSpace(scanner.Text())

		for _, s := range strings.Split(line, " ") {
			number, err := strconv.Atoi(s)
			if err != nil {
				return Matrix{}, err
			}
			result[i] = append(result[i], number)
		}

		// Check if the latest row is even with the current row.
		if len(result) > 1 {
			if len(result[i-1]) != len(result[i]) {
				return Matrix{}, fmt.Errorf("uneven rows: last row length = %d, current row length = %d", len(result[i-1]), len(result[i]))
			}
		}
	}

	return result, nil
}

// Rows copies and returns the underlaying coordinates in the Matrix.
func (m Matrix) Rows() [][]int {
	result := make([][]int, 0, len(m))

	for _, slice := range m {
		s := make([]int, len(slice))
		copy(s, slice)
		result = append(result, s)
	}
	return result
}

// Cols rearranges the underlaying coordinates,
//    int[[1,2,3],[4,5,6]] -> int[[1,4],[2,5],[3,6]]
func (m Matrix) Cols() [][]int {
	result := make([][]int, 0, len(m))

	for i := 0; i < len(m[0]); i++ {
		var s []int
		for j := 0; j < len(m); j++ {
			s = append(s, m[j][i])
		}
		result = append(result, s)
	}
	return result
}

// Set can set a new value in the Matrix by x and y coordinates.
func (m Matrix) Set(x, y, z int) (ok bool) {
	switch {
	case x < 0 || y < 0:
		return false
	case len(m)-1 < x:
		return false
	case len(m[x])-1 < y:
		return false
	default:
		m[x][y] = z
		return true
	}
}
