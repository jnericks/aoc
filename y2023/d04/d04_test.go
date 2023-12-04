package d04

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	data := ReadInput()
	cards, err := ParseCards(data)
	require.NoError(t, err)

	assert.Len(t, cards, 209)
	assert.Equal(t, 21158, TotalScore(cards))
	assert.Equal(t, 6_050_769, ProcessCards(cards))
}

func Test_Example1(t *testing.T) {
	data := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
	cards, err := ParseCards(data)
	require.NoError(t, err)
	assert.Len(t, cards, 6)

	assert.Equal(t, Card{
		ID:             1,
		WinningNumbers: []int{17, 41, 48, 83, 86},
		Numbers:        []int{6, 9, 17, 31, 48, 53, 83, 86},
		Score:          8,
		Matching:       4,
	}, cards[0])
	assert.Equal(t, Card{
		ID:             2,
		WinningNumbers: []int{13, 16, 20, 32, 61},
		Numbers:        []int{17, 19, 24, 30, 32, 61, 68, 82},
		Score:          2,
		Matching:       2,
	}, cards[1])
	assert.Equal(t, 2, cards[2].Score)
	assert.Equal(t, 1, cards[3].Score)
	assert.Equal(t, 0, cards[4].Score)
	assert.Equal(t, 0, cards[5].Score)

	assert.Equal(t, 13, TotalScore(cards))

	n := ProcessCards(cards)
	assert.Equal(t, 30, n)
}

func Test_ParseNumber(t *testing.T) {
	tests := []struct {
		data     string
		expected int
	}{
		{data: "123", expected: 123},
		{data: "12a3", expected: 123},
	}

	for _, tt := range tests {
		t.Run(tt.data, func(t *testing.T) {
			actual, err := ParseNumber("123")
			require.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
