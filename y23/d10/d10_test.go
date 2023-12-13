package d10

import (
	"testing"

	"github.com/jnericks/aoc/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	grid, err := ParseData(util.ReadStrings(InputData))
	require.NoError(t, err)

	assert.Equal(t, 95, grid.Start.V)
	assert.Equal(t, 74, grid.Start.H)

	farthest := grid.FarthestPipe()
	assert.Equal(t, 34, farthest.V)
	assert.Equal(t, 70, farthest.H)
	assert.Equal(t, 6725, farthest.Distance)

	// assert.Equal(t, 0, grid.CountInsides())
}

func Test_Example1(t *testing.T) {
	data := []string{
		"-L|F7", // .....
		"7S-7|", // .F-7.
		"L|7||", // .|.|.
		"-L-J|", // .L-J.
		"L|-JF", // .....
	}

	grid, err := ParseData(data)
	require.NoError(t, err)

	assert.Nil(t, grid.Start.North)
	assert.NotNil(t, grid.Start.East)
	assert.NotNil(t, grid.Start.South)
	assert.Nil(t, grid.Start.West)

	grid.PrintLoop()

	farthest := grid.FarthestPipe()
	assert.Equal(t, 3, farthest.V)
	assert.Equal(t, 3, farthest.H)
	assert.Equal(t, 4, farthest.Distance)

	// assert.Equal(t, 10, grid.CountInsides())
}

func Test_Example2(t *testing.T) {
	data := []string{
		"FF7FSF7F7F7F7F7F---7",
		"L|LJ||||||||||||F--J",
		"FL-7LJLJ||||||LJL-77",
		"F--JF--7||LJLJIF7FJ-",
		"L---JF-JLJIIIIFJLJJ7",
		"|F|F-JF---7IIIL7L|7|",
		"|FFJF7L7F-JF7IIL---7",
		"7-L-JL7||F7|L7F-7F7|",
		"L.L7LFJ|||||FJL7||LJ",
		"L7JLJL-JLJLJL--JLJ.L",
	}

	grid, err := ParseData(data)
	require.NoError(t, err)

	grid.PrintLoop()

	// assert.Equal(t, 10, grid.CountInsides())

}
