package d07

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	//go:embed example.txt
	example []byte

	//go:embed input.txt
	input []byte
)

func ReadExample() []string {
	s := bufio.NewScanner(bytes.NewBuffer(example))
	var out []string
	for s.Scan() {
		out = append(out, s.Text())
	}
	return out
}

func ReadInput() []string {
	s := bufio.NewScanner(bytes.NewBuffer(input))
	var out []string
	for s.Scan() {
		out = append(out, s.Text())
	}
	return out
}

var CardPriority = map[byte]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

var CardPriorityJoker = map[byte]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

type HandType string

const (
	TypeHighCard     HandType = "High Card"
	TypeOnePair      HandType = "One Pair"
	TypeTwoPair      HandType = "Two Pair"
	TypeThreeOfAKind HandType = "Three of a Kind"
	TypeFullHouse    HandType = "Full House"
	TypeFourOfAKind  HandType = "Four of a Kind"
	TypeFiveOfAKind  HandType = "Five of a Kind"
)

func (h HandType) Priority() int {
	switch h {
	case TypeHighCard:
		return 1
	case TypeOnePair:
		return 2
	case TypeTwoPair:
		return 3
	case TypeThreeOfAKind:
		return 4
	case TypeFullHouse:
		return 5
	case TypeFourOfAKind:
		return 6
	case TypeFiveOfAKind:
		return 7
	}
	return 0
}

type Hand struct {
	Cards string
	Bid   int

	// Type is the type of the hand.
	Type HandType

	// TypeJoker is the type of the hand if 'J' is a Joker.
	TypeJoker HandType
}

func (h Hand) String() string {
	return fmt.Sprintf("%s, %d, %s", h.Cards, h.Bid, h.Type)
}

func SolveHands(hands []Hand) int {
	var score int
	for i, hand := range hands {
		rank := i + 1
		score += hand.Bid * rank
	}
	return score
}

func ParseInput(data []string, joker bool) ([]Hand, error) {
	buckets := make(map[HandType][]Hand)
	var n int
	for _, d := range data {
		parts := strings.Split(d, " ")

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid input '%s'", d)
		}

		cards := parts[0]
		if len(cards) != 5 {
			return nil, fmt.Errorf("invalid hand '%s'", cards)
		}

		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid bid '%s'", parts[1])
		}

		handType, err := func() (HandType, error) {
			if joker {
				return DetermineHandTypeJoker(cards)
			}
			return DetermineHandType(cards)
		}()

		buckets[handType] = append(buckets[handType], Hand{
			Cards: cards,
			Bid:   bid,
			Type:  handType,
		})
	}

	out := make([]Hand, 0, n)
	for _, handType := range []HandType{
		TypeHighCard,
		TypeOnePair,
		TypeTwoPair,
		TypeThreeOfAKind,
		TypeFullHouse,
		TypeFourOfAKind,
		TypeFiveOfAKind,
	} {
		hands := buckets[handType]
		sort.Slice(hands, func(i, j int) bool {
			h1, h2 := hands[i], hands[j]
			if h1.Cards == h2.Cards {
				return false
			}

			for i := 0; i < 5; i++ {
				c1 := h1.Cards[i]
				c2 := h2.Cards[i]
				if c1 == c2 {
					continue
				}
				if joker {
					return CardPriorityJoker[c1] < CardPriorityJoker[c2]
				}
				return CardPriority[c1] < CardPriority[c2]
			}

			return false
		})
		out = append(out, hands...)
	}

	return out, nil
}

func DetermineHandType(cards string) (HandType, error) {
	m := make(map[rune]int)
	for _, card := range cards {
		m[card] = m[card] + 1
	}

	switch len(m) {
	case 5:
		return TypeHighCard, nil
	case 4:
		return TypeOnePair, nil
	case 3:
		for _, count := range m {
			switch count {
			case 2:
				return TypeTwoPair, nil
			case 3:
				return TypeThreeOfAKind, nil
			}
		}
	case 2:
		for _, count := range m {
			switch count {
			case 1, 4:
				return TypeFourOfAKind, nil
			case 2, 3:
				return TypeFullHouse, nil
			}
		}
	case 1:
		return TypeFiveOfAKind, nil
	}
	return "", fmt.Errorf("could not determine hand type for cards '%s'", cards)
}

func DetermineHandTypeJoker(cards string) (HandType, error) {

	var nJokers int

	m := make(map[rune]int)
	for _, card := range cards {
		if card == 'J' {
			nJokers++
			continue
		}
		m[card] = m[card] + 1
	}

	nUniqueCards := len(m)

	switch nJokers {
	case 0:
		return DetermineHandType(cards)
	case 1:
		switch nUniqueCards {
		case 4:
			return TypeOnePair, nil
		case 3:
			return TypeThreeOfAKind, nil
		case 2:
			for _, count := range m {
				switch count {
				case 4:
					return TypeOnePair, nil
				case 1, 3:
					return TypeFourOfAKind, nil
				case 2:
					return TypeFullHouse, nil
				}
			}
		case 1:
			return TypeFiveOfAKind, nil
		}
	case 2:
		switch nUniqueCards {
		case 3:
			return TypeThreeOfAKind, nil
		case 2:
			return TypeFourOfAKind, nil
		case 1:
			return TypeFiveOfAKind, nil
		}
	case 3:
		switch nUniqueCards {
		case 2:
			return TypeFourOfAKind, nil
		case 1:
			return TypeFiveOfAKind, nil
		}
	case 4, 5:
		return TypeFiveOfAKind, nil
	}

	return "", fmt.Errorf("could not determine hand type for cards '%s'", cards)
}
