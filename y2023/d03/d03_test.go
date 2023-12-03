package d03

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	input := ReadInput()
	board, err := ParseBoard(input)
	require.NoError(t, err)

	assert.Equal(t, 531932, Sum(board.PartNumbers()))
	assert.Equal(t, 73646890, Sum(board.GearRatios()))
}

func Test_Example1(t *testing.T) {
	data := []byte(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)

	input := ParseInput(data)
	assert.Len(t, input, 10)

	board, err := ParseBoard(input)
	require.NoError(t, err)

	t.Run("Numbers", func(t *testing.T) {
		expNumbers := [][]Number{
			{{Index: 0, Value: 467}, {Index: 5, Value: 114}},
			nil,
			{{Index: 2, Value: 35}, {Index: 6, Value: 633}},
			nil,
			{{Index: 0, Value: 617}},
			{{Index: 7, Value: 58}},
			{{Index: 2, Value: 592}},
			{{Index: 6, Value: 755}},
			nil,
			{{Index: 1, Value: 664}, {Index: 5, Value: 598}},
		}

		require.Len(t, board.Numbers, len(expNumbers))
		for i, row := range board.Numbers {
			require.Len(t, row, len(expNumbers[i]))

			for j, number := range row {
				assert.Equal(t, number, expNumbers[i][j])
			}
		}
	})

	t.Run("Symbols", func(t *testing.T) {
		expSymbols := [][]Symbol{
			nil,
			{{Index: 3, Value: '*'}},
			nil,
			{{Index: 6, Value: '#'}},
			{{Index: 3, Value: '*'}},
			{{Index: 5, Value: '+'}},
			nil,
			nil,
			{{Index: 3, Value: '$'}, {Index: 5, Value: '*'}},
			nil,
		}

		require.Len(t, board.Symbols, len(expSymbols))
		for i, row := range board.Symbols {
			require.Len(t, row, len(expSymbols[i]))

			for j, num := range row {
				assert.Equal(t, num, expSymbols[i][j])
			}
		}
	})

	t.Run("PartNumbers", func(t *testing.T) {
		nums := board.PartNumbers()
		exp := []int{467, 35, 633, 617, 592, 755, 664, 598}
		require.Len(t, nums, len(exp))
		for i, num := range nums {
			assert.Equal(t, exp[i], num)
		}
		assert.Equal(t, Sum(nums), 4361)
	})

	t.Run("GearRatios", func(t *testing.T) {

		nearby := board.NearbyNumbers(1, Symbol{Index: 3, Value: '*'})
		assert.ElementsMatch(t, []int{467, 35}, nearby)

		nearby = board.NearbyNumbers(4, Symbol{Index: 3, Value: '*'})
		assert.ElementsMatch(t, []int{617}, nearby)

		gearRatios := board.GearRatios()
		assert.ElementsMatch(t, []int{467 * 35, 755 * 598}, gearRatios)
		assert.Equal(t, 467835, Sum(gearRatios))
	})
}
