package d11

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	u, err := Parse(FileInput)
	require.NoError(t, err)

	t.Run("part 1", func(t *testing.T) {
		assert.Len(t, u.Galaxies, 426)
		assert.Equal(t, 9329143, u.ShortestRoutesSum(2))
	})

	t.Run("part 2", func(t *testing.T) {
		assert.Len(t, u.Galaxies, 426)
		assert.Equal(t, 710674907809, u.ShortestRoutesSum(1_000_000))
	})
}

func Test_Example(t *testing.T) {
	u, err := Parse(FileExample)
	require.NoError(t, err)

	expected := []string{
		//           v  v  v
		//         0123456789
		/*   0 */ "...#......", // 0
		/*   1 */ ".......#..", // 1
		/*   2 */ "#.........", // 2
		/* > 3 */ "..........", // 3 <
		/*   4 */ "......#...", // 4
		/*   5 */ ".#........", // 5
		/*   6 */ ".........#", // 6
		/* > 7 */ "..........", // 7 <
		/*   8 */ ".......#..", // 8
		/*   9 */ "#...#.....", // 9
		//         0123456789
		//           ^  ^  ^
	}
	assert.Len(t, u.Grid, len(expected))
	assert.Equal(t, expected, u.Grid)

	assert.Equal(t, []Galaxy{
		{ID: 1, Row: 0, Col: 3},
		{ID: 2, Row: 1, Col: 7},
		{ID: 3, Row: 2, Col: 0},
		{ID: 4, Row: 4, Col: 6},
		{ID: 5, Row: 5, Col: 1},
		{ID: 6, Row: 6, Col: 9},
		{ID: 7, Row: 8, Col: 7},
		{ID: 8, Row: 9, Col: 0},
		{ID: 9, Row: 9, Col: 4},
	}, u.Galaxies)
	assert.Equal(t, 374, u.ShortestRoutesSum(2))
}
