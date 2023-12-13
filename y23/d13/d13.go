package d13

import (
	_ "embed"
	"fmt"
	"slices"

	"github.com/jnericks/aoc/util"
)

var (
	//go:embed example.txt
	FileExample []byte

	//go:embed input.txt
	FileInput []byte
)

type Mirrors struct {
	Grids []*Grid
}

func (m *Mirrors) SolvePart1() int {
	var sum int
	for _, grid := range m.Grids {
		sum += grid.Solve(SlicesAreEqual)
	}
	return sum
}

func (m *Mirrors) SolvePart2() int {
	var sum int
	for _, grid := range m.Grids {
		sum += grid.Solve(SlicesDifferByOneExactly)
	}
	return sum
}

type Grid struct {
	// Rows are the original pattern as provided.
	Rows []string
	// Cols are the rows rotated.
	Cols []string
}

func (g *Grid) Solve(equal func([]string, []string) bool) int {
	findIndex := func(s []string) int {
		if len(s) == 0 {
			return 0
		}
		for i := 1; i < len(s); i++ {
			l := min(i, len(s)-i)
			a := slices.Clone(s[i-l : i]) // left
			b := slices.Clone(s[i : i+l]) // right
			slices.Reverse(b)
			if equal(a, b) {
				return i
			}

		}
		return 0
	}

	return findIndex(g.Cols) + 100*findIndex(g.Rows)
}

func ParseFile(file []byte) (*Mirrors, error) {
	lines := util.ReadStrings(file)

	var grids []*Grid

	var i int
	for j, line := range lines {
		if line == "" {
			grid, err := ParseGrid(lines[i:j])
			if err != nil {
				return nil, err
			}
			grids = append(grids, grid)
			i = j + 1
		}
	}

	// last section of file
	grid, err := ParseGrid(lines[i:])
	if err != nil {
		return nil, err
	}
	grids = append(grids, grid)

	return &Mirrors{Grids: grids}, nil
}

func ParseGrid(rows []string) (*Grid, error) {
	nRows := len(rows)
	nCols := len(rows[0])

	rotated := make([][]rune, nCols)
	for i := range rotated {
		rotated[i] = make([]rune, nRows)
	}

	for i, row := range rows {
		// validate each row is the same length
		if len(row) != nCols {
			return nil, fmt.Errorf("length mismatch: [%d] %s", i, row)
		}

		// do rotation
		for j, r := range row {
			rotated[j][i] = r
		}
	}

	cols := make([]string, len(rotated))
	for i, runes := range rotated {
		cols[i] = string(runes)
	}

	return &Grid{
		Rows: rows,
		Cols: cols,
	}, nil
}

func SlicesAreEqual(a, b []string) bool {
	return slices.Equal(a, b)
}

func SlicesDifferByOneExactly(a, b []string) bool {
	var diffs int
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				diffs++
				if diffs > 1 {
					return false
				}
			}
		}
	}
	return diffs == 1
}
