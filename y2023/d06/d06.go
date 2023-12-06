package d06

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"strconv"
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

type Race struct {
	Time     int
	Distance int
}

func ParseInputPart1(data []string) ([]Race, error) {
	if len(data) != 2 {
		return nil, errors.New("invalid input")
	}

	times, err := ParseNumbers(data[0])
	if err != nil {
		return nil, err
	}

	distances, err := ParseNumbers(data[1])
	if err != nil {
		return nil, err
	}

	if len(times) != len(distances) {
		return nil, errors.New("invalid input")
	}

	out := make([]Race, len(times))
	for i := 0; i < len(times); i++ {
		out[i] = Race{
			Time:     times[i],
			Distance: distances[i],
		}
	}
	return out, nil
}

func ParseNumbers(data string) ([]int, error) {
	var out []int
	for i := 0; i < len(data); i++ {
		var num []byte
		for i < len(data) && '0' <= data[i] && data[i] <= '9' {
			num = append(num, data[i])
			i++
		}
		if len(num) > 0 {
			n, err := strconv.Atoi(string(num))
			if err != nil {
				return nil, err
			}
			out = append(out, n)
		}
	}
	return out, nil
}

func ParseInputPart2(data []string) (Race, error) {
	if len(data) != 2 {
		return Race{}, errors.New("invalid input")
	}

	time, err := ParseNumber(data[0])
	if err != nil {
		return Race{}, err
	}

	distance, err := ParseNumber(data[1])
	if err != nil {
		return Race{}, err
	}

	return Race{
		Time:     time,
		Distance: distance,
	}, nil
}

func ParseNumber(data string) (int, error) {
	var num []byte
	for i := 0; i < len(data); i++ {
		if '0' <= data[i] && data[i] <= '9' {
			num = append(num, data[i])
		}
	}
	return strconv.Atoi(string(num))
}

func WaysToWin(race Race) int {
	var speed int
	var wins int
	for i := 0; i <= race.Time; i++ {
		timeRemaining := race.Time - i
		if speed*timeRemaining > race.Distance {
			wins++
		}
		speed++
	}
	return wins
}
