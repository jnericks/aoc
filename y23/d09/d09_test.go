package d09

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	histories, err := ParseData(ReadData(InputData))
	require.NoError(t, err)

	first, last := SolveHistories(histories)
	assert.Equal(t, 975, first)
	assert.Equal(t, 1_641_934_234, last)
}

func Test_Example1(t *testing.T) {
	histories, err := ParseData(ReadData(Example1Data))
	require.NoError(t, err)

	require.Equal(t, [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}, histories)

	assert.Equal(t, [][]int{
		{-3, 0, 3, 6, 9, 12, 15, 18},
		{3, 3, 3, 3, 3, 3, 3},
		{0, 0, 0, 0, 0, 0},
	}, HistoryLayers(histories[0]))
	first, last := SolveHistory(histories[0])
	assert.Equal(t, -3, first)
	assert.Equal(t, 18, last)

	assert.Equal(t, [][]int{
		{0, 1, 3, 6, 10, 15, 21, 28},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
	}, HistoryLayers(histories[1]))
	first, last = SolveHistory(histories[1])
	assert.Equal(t, 0, first)
	assert.Equal(t, 28, last)

	assert.Equal(t, [][]int{
		{5, 10, 13, 16, 21, 30, 45, 68},
		{5, 3, 3, 5, 9, 15, 23},
		{-2, 0, 2, 4, 6, 8},
		{2, 2, 2, 2, 2},
		{0, 0, 0, 0},
	}, HistoryLayers(histories[2]))
	first, last = SolveHistory(histories[2])
	assert.Equal(t, 5, first)
	assert.Equal(t, 68, last)

	first, last = SolveHistories(histories)
	assert.Equal(t, -3+0+5, first)
	assert.Equal(t, 18+28+68, last)
}
