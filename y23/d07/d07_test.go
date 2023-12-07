package d07

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		hands, err := ParseInput(ReadExample(), false)
		require.NoError(t, err)
		assert.Equal(t, 6440, SolveHands(hands))
	})

	t.Run("part 1", func(t *testing.T) {
		hands, err := ParseInput(ReadInput(), false)
		require.NoError(t, err)
		assert.Equal(t, 250_957_639, SolveHands(hands))
	})

	t.Run("part 2", func(t *testing.T) {
		hands, err := ParseInput(ReadInput(), true)
		require.NoError(t, err)
		assert.Equal(t, 251_515_496, SolveHands(hands))
	})
}

func Test_ParseInput(t *testing.T) {
	t.Run("Jack", func(t *testing.T) {
		hands, err := ParseInput(ReadExample(), false)
		require.NoError(t, err)

		for i, expected := range []Hand{
			{Cards: "32T3K", Bid: 765, Type: TypeOnePair},
			{Cards: "KTJJT", Bid: 220, Type: TypeTwoPair},
			{Cards: "KK677", Bid: 28, Type: TypeTwoPair},
			{Cards: "T55J5", Bid: 684, Type: TypeThreeOfAKind},
			{Cards: "QQQJA", Bid: 483, Type: TypeThreeOfAKind},
		} {
			assert.Equal(t, expected, hands[i])
		}
	})

	t.Run("Joker", func(t *testing.T) {
		hands, err := ParseInput(ReadExample(), true)
		require.NoError(t, err)

		for i, expected := range []Hand{
			{Cards: "32T3K", Bid: 765, Type: TypeOnePair},
			{Cards: "KK677", Bid: 28, Type: TypeTwoPair},
			{Cards: "T55J5", Bid: 684, Type: TypeFourOfAKind},
			{Cards: "QQQJA", Bid: 483, Type: TypeFourOfAKind},
			{Cards: "KTJJT", Bid: 220, Type: TypeFourOfAKind},
		} {
			assert.Equal(t, expected, hands[i])
		}
	})
}

func Test_DetermineHandType(t *testing.T) {
	tests := []struct {
		cards                 string
		expectedHandType      HandType
		expectedHandTypeJoker HandType
	}{
		{
			cards:                 "AAAAA",
			expectedHandType:      TypeFiveOfAKind,
			expectedHandTypeJoker: TypeFiveOfAKind,
		},
		{
			cards:                 "JJJJJ",
			expectedHandType:      TypeFiveOfAKind,
			expectedHandTypeJoker: TypeFiveOfAKind,
		},
		{
			cards:                 "TTJTT",
			expectedHandType:      TypeFourOfAKind,
			expectedHandTypeJoker: TypeFiveOfAKind,
		},
		{
			cards:                 "J5J55",
			expectedHandType:      TypeFullHouse,
			expectedHandTypeJoker: TypeFiveOfAKind,
		},
		{
			cards:                 "75J55",
			expectedHandType:      TypeThreeOfAKind,
			expectedHandTypeJoker: TypeFourOfAKind,
		},
		{
			cards:                 "88J6J",
			expectedHandType:      TypeTwoPair,
			expectedHandTypeJoker: TypeFourOfAKind,
		},
		{
			cards:                 "123JJ",
			expectedHandType:      TypeOnePair,
			expectedHandTypeJoker: TypeThreeOfAKind,
		},
	}

	for _, tt := range tests {
		t.Run(tt.cards, func(t *testing.T) {
			t.Run("Jack", func(t *testing.T) {
				handType, err := DetermineHandType(tt.cards)
				require.NoError(t, err)
				assert.Equal(t, tt.expectedHandType, handType)
			})

			t.Run("Joker", func(t *testing.T) {
				handType, err := DetermineHandTypeJoker(tt.cards)
				require.NoError(t, err)
				assert.Equal(t, tt.expectedHandTypeJoker, handType)
			})
		})
	}
}
