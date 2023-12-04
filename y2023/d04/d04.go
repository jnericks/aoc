package d04

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func ReadInput() []string {
	s := bufio.NewScanner(bytes.NewBuffer(input))

	var out []string
	for s.Scan() {
		out = append(out, s.Text())
	}

	return out
}

func ParseCards(cards []string) ([]Card, error) {
	out := make([]Card, len(cards))
	for i, data := range cards {
		card, err := ParseCard(data)
		if err != nil {
			return nil, err
		}
		out[i] = card
	}
	return out, nil
}

func ParseCard(card string) (Card, error) {
	parts := strings.Split(card, ":")
	if len(parts) != 2 {
		return Card{}, errors.New("invalid card")
	}

	id, err := ParseNumber(parts[0])
	if err != nil {
		return Card{}, err
	}

	numberParts := strings.Split(parts[1], "|")
	if len(numberParts) != 2 {
		return Card{}, errors.New("invalid card, ")
	}

	winningNumbers, err := ParseNumbers(numberParts[0])
	if err != nil {
		return Card{}, err
	}

	numbers, err := ParseNumbers(numberParts[1])
	if err != nil {
		return Card{}, err
	}

	score, matching := ScoreAndMatching(winningNumbers, numbers)
	return Card{
		ID:             id,
		WinningNumbers: winningNumbers,
		Numbers:        numbers,
		Score:          score,
		Matching:       matching,
	}, nil
}

type Card struct {
	ID             int
	WinningNumbers []int
	Numbers        []int
	Score          int
	Matching       int
}

func (c Card) String() string {
	return fmt.Sprintf("%2d, W%v, N%v", c.ID, c.WinningNumbers, c.Numbers)
}

func ScoreAndMatching(winningNumbers, numbers []int) (int, int) {
	sort.Ints(winningNumbers)
	sort.Ints(numbers)

	var score int
	var matching int
	i, j := 0, 0
	for i < len(winningNumbers) && j < len(numbers) {
		w, n := winningNumbers[i], numbers[j]
		if w == n {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
			matching++
			i++
			j++
		} else if w < n {
			i++
		} else {
			j++
		}
	}
	return score, matching
}

func ProcessCards(cards []Card) int {
	counts := make([]int, len(cards))
	for i := range counts {
		counts[i] = -1 // set to unprocessed
	}

	var total int
	for i := range cards {
		total += totalCards(cards, counts, i)
	}

	return total
}

func totalCards(cards []Card, counts []int, idx int) int {
	if idx >= len(cards) {
		return 0
	}

	var total int
	for i := idx + 1; i < len(cards) && i <= idx+cards[idx].Matching; i++ {
		total += totalCards(cards, counts, i)
	}

	counts[idx] = 1 + total
	return counts[idx]
}

func ParseNumbers(data string) ([]int, error) {
	split := strings.Split(data, " ")
	out := make([]int, 0, len(split))
	for _, v := range split {
		if v == "" {
			continue
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		out = append(out, n)
	}
	return out, nil
}

func ParseNumber(data string) (int, error) {
	out := make([]byte, 0, len(data))
	for _, b := range []byte(data) {
		if '0' <= b && b <= '9' {
			out = append(out, b)
		}
	}
	return strconv.Atoi(string(out))
}

func TotalScore(cards []Card) int {
	var total int
	for _, c := range cards {
		total += c.Score
	}
	return total
}
