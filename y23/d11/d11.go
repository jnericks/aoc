package d11

import (
	"bytes"
	_ "embed"
	"errors"

	"github.com/jnericks/aoc/util"
)

var (
	//go:embed example.txt
	FileExample []byte

	//go:embed input.txt
	FileInput []byte
)

type Universe struct {
	Grid      []string
	Galaxies  []Galaxy
	EmptyRows []int
	EmptyCols []int
}

type Galaxy struct {
	ID int

	// Coordinates
	Row, Col int
}

func (u *Universe) ShortestRoutesSum(expansionFactor int) int {
	var sum int
	for i := range u.Galaxies {
		for j := i + 1; j < len(u.Galaxies); j++ {
			sum += u.ShortestRoute(u.Galaxies[i], u.Galaxies[j], expansionFactor)
		}
	}
	return sum
}

func (u *Universe) ShortestRoute(x, y Galaxy, expansionFactor int) int {
	height := max(x.Row, y.Row) - min(x.Row, y.Row)
	for _, row := range u.EmptyRows {
		if min(x.Row, y.Row) < row && row < max(x.Row, y.Row) {
			height += expansionFactor - 1
		}
	}

	width := max(x.Col, y.Col) - min(x.Col, y.Col)
	for _, col := range u.EmptyCols {
		if min(x.Col, y.Col) < col && col < max(x.Col, y.Col) {
			width += expansionFactor - 1
		}
	}

	return height + width
}

func (u *Universe) String() string {
	var out bytes.Buffer
	for _, row := range u.Grid {
		out.WriteString(row)
	}
	return out.String()
}

func Parse(file []byte) (*Universe, error) {
	grid := util.ReadStrings(file)

	// validate
	if len(grid) == 0 {
		return nil, errors.New("empty map")
	}
	for _, row := range grid {
		if len(row) != len(grid[0]) {
			return nil, errors.New("row length mismatch")
		}
	}

	galaxyInRow := make(map[int]struct{})
	galaxyInCol := make(map[int]struct{})

	var galaxies []Galaxy
	var id int
	for r, row := range grid {
		for c, character := range row {
			if character == '#' {
				id++
				galaxies = append(galaxies, Galaxy{ID: id, Row: r, Col: c})
				galaxyInRow[r] = struct{}{}
				galaxyInCol[c] = struct{}{}
			}
		}
	}

	emptyRows := make([]int, 0, len(grid)-len(galaxyInRow))
	for i := 0; i < len(grid); i++ {
		if _, ok := galaxyInRow[i]; !ok {
			emptyRows = append(emptyRows, i)
		}
	}

	emptyCols := make([]int, 0, len(grid[0])-len(galaxyInCol))
	for i := 0; i < len(grid[0]); i++ {
		if _, ok := galaxyInCol[i]; !ok {
			emptyCols = append(emptyCols, i)
		}
	}

	return &Universe{
		Grid:      grid,
		Galaxies:  galaxies,
		EmptyRows: emptyRows,
		EmptyCols: emptyCols,
	}, nil
}
