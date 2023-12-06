package d06

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		races, err := ParseInputPart1(ReadInput())
		require.NoError(t, err)

		var ans int
		for i, race := range races {
			if i == 0 {
				ans = WaysToWin(race)
			} else {
				ans *= WaysToWin(race)
			}
		}
		assert.Equal(t, 303_600, ans)
	})

	t.Run("Part 2", func(t *testing.T) {
		race, err := ParseInputPart2(ReadInput())
		require.NoError(t, err)
		assert.Equal(t, 23_654_842, WaysToWin(race))
	})
}

func Test_WaysToWin(t *testing.T) {
	races, err := ParseInputPart1([]string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	})
	require.NoError(t, err)

	require.Len(t, races, 3)
	assert.Equal(t, Race{
		Time:     7,
		Distance: 9,
	}, races[0])
	assert.Equal(t, 4, WaysToWin(races[0]))
	assert.Equal(t, 8, WaysToWin(races[1]))
	assert.Equal(t, 9, WaysToWin(races[2]))
}

func Test_ParseInput(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		races, err := ParseInputPart1(ReadInput())
		require.NoError(t, err)
		assert.Equal(t, []Race{
			{Time: 38, Distance: 234},
			{Time: 67, Distance: 1027},
			{Time: 76, Distance: 1157},
			{Time: 73, Distance: 1236},
		}, races)
	})

	t.Run("Part 2", func(t *testing.T) {
		race, err := ParseInputPart2(ReadInput())
		require.NoError(t, err)
		assert.Equal(t, Race{
			Time:     38_67_76_73,
			Distance: 234_1027_1157_1236,
		}, race)
	})
}
