package d10

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	grid, err := ParseData(ReadData(InputData))
	require.NoError(t, err)

	assert.Equal(t, 95, grid.Start.V)
	assert.Equal(t, 74, grid.Start.H)

	farthest := grid.FarthestPipe()
	assert.Equal(t, 34, farthest.V)
	assert.Equal(t, 70, farthest.H)
	assert.Equal(t, 6725, farthest.Distance)
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
}
