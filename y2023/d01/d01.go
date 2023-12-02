package d01

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"strings"
)

//go:embed input.txt
var input []byte

func ReadInput() []string {
	s := bufio.NewScanner(bytes.NewBuffer(input))

	var input []string
	for s.Scan() {
		input = append(input, s.Text())
	}

	return input
}

func ParseInput(input []string) ([]int, error) {
	values := make([]int, len(input))
	for i, line := range input {
		v, err := ParseLine(line)
		if err != nil {
			return nil, fmt.Errorf("error on input line %d: %w", i, err)
		}
		values[i] = v
	}
	return values, nil
}

func ParseLine(line string) (int, error) {
	num := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	txt := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	first, err := func() (int, error) {
		for i := 0; i < len(line); i++ {
			if v, ok := num[line[i]]; ok {
				return v, nil
			}
			for k, v := range txt {
				if strings.HasPrefix(line[i:], k) {
					return v, nil
				}
			}
		}
		return 0, errors.New("first number not found")
	}()
	if err != nil {
		return 0, err
	}

	second, err := func() (int, error) {
		for i := len(line) - 1; i >= 0; i-- {
			if v, ok := num[line[i]]; ok {
				return v, nil
			}
			for k, v := range txt {
				if strings.HasSuffix(line[:i+1], k) {
					return v, nil
				}
			}
		}
		return 0, errors.New("second number not found")
	}()
	if err != nil {
		return 0, err
	}

	return first*10 + second, nil
}

func Sum(values []int) int {
	var sum int
	for _, value := range values {
		sum += value
	}
	return sum
}
