package d13

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	mirrors, err := ParseFile(FileInput)
	require.NoError(t, err)

	assert.Len(t, mirrors.Grids, 100)

	t.Run("part 1", func(t *testing.T) {
		assert.Equal(t, 39939, mirrors.SolvePart1())
	})

	t.Run("part 2", func(t *testing.T) {
		assert.Equal(t, 32069, mirrors.SolvePart2())
	})
}

func Test_Example(t *testing.T) {
	t.Run("pattern 1", func(t *testing.T) {
		pattern := []string{
			//    012345678
			//        ><
			/**/ "#.##..##.",
			/**/ "..#.##.#.",
			/**/ "##......#",
			/**/ "##......#",
			/**/ "..#.##.#.",
			/**/ "..##..##.",
			/**/ "#.#.##.#.",
			//        ><
		}

		grid, err := ParseGrid(pattern)
		require.NoError(t, err)

		grid.Solve(SlicesAreEqual)
	})

	t.Run("pattern 2", func(t *testing.T) {
		pattern := []string{
			"#...##..#", //   1
			"#....#..#", //   2
			"..##..###", //   3
			"#####.##.", // v 4
			"#####.##.", // ^ 5
			"..##..###", //   6
			"#....#..#", //   7
		}

		grid, err := ParseGrid(pattern)
		require.NoError(t, err)

		assert.Equal(t, 400, grid.Solve(func(s1, s2 []string) bool {
			return slices.Equal(s1, s2)
		}))
	})

	t.Run("parse example", func(t *testing.T) {
		mirrors, err := ParseFile(FileExample)
		require.NoError(t, err)

		assert.Equal(t, 405, mirrors.SolvePart1())
	})
}
