package d12

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	game, err := Parse(FileInput)
	require.NoError(t, err)
	assert.Len(t, game.Records, 1000)

	t.Run("part 1", func(t *testing.T) {
		assert.Equal(t, int64(8419), game.SolvePart1())
	})

	t.Run("part 2", func(t *testing.T) {
		assert.Equal(t, int64(160500973317706), game.SolvePart2())
	})
}

func Test_Example(t *testing.T) {
	game, err := Parse(FileExample)
	require.NoError(t, err)

	assert.Len(t, game.Records, 6)

	assert.Equal(t, Record{Row: "???.###", Values: []int{1, 1, 3}}, game.Records[0])
	assert.Equal(t, Record{Row: ".??..??...?##.", Values: []int{1, 1, 3}}, game.Records[1])
	assert.Equal(t, Record{Row: "?#?#?#?#?#?#?#?", Values: []int{1, 3, 1, 6}}, game.Records[2])
	assert.Equal(t, Record{Row: "????.#...#...", Values: []int{4, 1, 1}}, game.Records[3])
	assert.Equal(t, Record{Row: "????.######..#####.", Values: []int{1, 6, 5}}, game.Records[4])
	assert.Equal(t, Record{Row: "?###????????", Values: []int{3, 2, 1}}, game.Records[5])

	assert.Equal(t, int64(21), game.SolvePart1())
	assert.Equal(t, int64(525152), game.SolvePart2())
}

func Test_Arrange(t *testing.T) {
	t.Run("???.### 1,1,3", func(t *testing.T) {
		assert.Equal(t, int64(1), NumArrangements(Record{
			Row:    "???.###",
			Values: []int{1, 1, 3},
		}))
	})

	t.Run(".??..??...?##. 1,1,3 - 4 arrangements", func(t *testing.T) {
		assert.Equal(t, int64(4), NumArrangements(Record{
			Row:    ".??..??...?##.",
			Values: []int{1, 1, 3},
		}))
	})

	t.Run("?#?#?#?#?#?#?#? 1,3,1,6 - 1 arrangement", func(t *testing.T) {
		assert.Equal(t, int64(1), NumArrangements(Record{
			Row:    "?#?#?#?#?#?#?#?",
			Values: []int{1, 3, 1, 6},
		}))
	})

	t.Run("????.#...#... 4,1,1 - 1 arrangement", func(t *testing.T) {
		assert.Equal(t, int64(1), NumArrangements(Record{
			Row:    "????.#...#...",
			Values: []int{4, 1, 1},
		}))
	})

	t.Run("????.######..#####. 1,6,5 - 4 arrangements", func(t *testing.T) {
		assert.Equal(t, int64(4), NumArrangements(Record{
			Row:    "????.######..#####.",
			Values: []int{1, 6, 5},
		}))
	})

	t.Run("?###???????? 3,2,1 - 10 arrangements", func(t *testing.T) {
		assert.Equal(t, int64(10), NumArrangements(Record{
			Row:    "?###????????",
			Values: []int{3, 2, 1},
		}))
	})

	t.Run("..???.??.? 1,1,1 - 9 arrangements", func(t *testing.T) {
		assert.Equal(t, int64(9), NumArrangements(Record{
			Row:    "..???.??.?",
			Values: []int{1, 1, 1},
		}))
	})

	t.Run(".#???###?#??#???#?.? 7,1,7 - 9 arrangements", func(t *testing.T) {
		assert.Equal(t, int64(1), NumArrangements(Record{
			Row:    ".#???###?#??#???#?.?",
			Values: []int{7, 1, 7},
		}))
	})
}
