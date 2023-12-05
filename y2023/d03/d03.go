package d03

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

const Symbols = `#$%&*+-/=@`

//go:embed input.txt
var input []byte

func ReadInput() []string {
	return ParseInput(input)
}

func ParseInput(input []byte) []string {
	s := bufio.NewScanner(bytes.NewBuffer(input))

	var out []string
	for s.Scan() {
		out = append(out, s.Text())
	}

	return out
}

type Board struct {
	Numbers [][]Number
	Symbols [][]Symbol
}

type Number struct {
	Index int
	Value int
}

func (s Number) String() string {
	return fmt.Sprintf("N{i:%d v:%d}", s.Index, s.Value)
}

type Symbol struct {
	Index int
	Value byte
}

func (s Symbol) IsGear() bool {
	return s.Value == '*'
}

func (s Symbol) String() string {
	return fmt.Sprintf("SrcNum{i:%d v:%s}", s.Index, string(s.Value))
}

func IsSymbol(b byte) bool {
	for _, s := range []byte("#$%&*+-/=@") {
		if b == s {
			return true
		}
	}
	return false
}

func IsNumber(b byte) bool {
	return '0' <= b && b <= '9'
}

func ParseBoard(rows []string) (Board, error) {
	numbers := make([][]Number, len(rows))
	for rowIndex, row := range rows {
		for i := 0; i < len(row); i++ {
			index := i // start index

			// collect number
			var sb strings.Builder
			for i < len(row) && IsNumber(row[i]) {
				sb.WriteByte(row[i])
				i++
			}

			// capture
			if num := sb.String(); num != "" {
				value, err := strconv.Atoi(num)
				if err != nil {
					return Board{}, err
				}
				numbers[rowIndex] = append(numbers[rowIndex], Number{
					Index: index,
					Value: value,
				})
			}
		}
	}

	// symbols
	symbols := make([][]Symbol, len(rows))
	for rowIndex, row := range rows {
		for i := 0; i < len(row); i++ {
			if IsSymbol(row[i]) {
				symbols[rowIndex] = append(symbols[rowIndex], Symbol{
					Index: i,
					Value: row[i],
				})
			}
		}
	}

	return Board{
		Numbers: numbers,
		Symbols: symbols,
	}, nil
}

func (b Board) PartNumbers() []int {
	var out []int
	for rowIndex, numRow := range b.Numbers {
		for _, num := range numRow {
			if b.IsSymbolNearby(rowIndex, num) {
				out = append(out, num.Value)
			}
		}
	}
	return out
}

func (b Board) GearRatios() []int {
	var out []int
	for rowIndex, symRows := range b.Symbols {
		for _, symbol := range symRows {
			if symbol.Value == '*' { // potential gear
				nums := b.NearbyNumbers(rowIndex, symbol)
				if len(nums) == 2 {
					out = append(out, nums[0]*nums[1])
				}
			}
		}
	}
	return out
}

func (b Board) IsSymbolNearby(rowIndex int, num Number) bool {
	i := max(rowIndex-1, 0)
	j := min(rowIndex+1, len(b.Symbols)-1)

	s, e := num.Index-1, num.Index+len(strconv.Itoa(num.Value))

	for ; i <= j; i++ {
		for _, symbol := range b.Symbols[i] {
			if s <= symbol.Index && symbol.Index <= e {
				return true
			}
		}
	}

	return false
}

func (b Board) NearbyNumbers(rowIndex int, symbol Symbol) []int {
	i := max(rowIndex-1, 0)
	j := min(rowIndex+1, len(b.Symbols)-1)

	var out []int
	for ; i <= j; i++ {
		for _, number := range b.Numbers[i] {
			s, e := number.Index-1, number.Index+len(strconv.Itoa(number.Value))
			if s <= symbol.Index && symbol.Index <= e {
				out = append(out, number.Value)
			}
		}
	}

	return out
}

func Sum(values []int) int {
	var sum int
	for _, value := range values {
		sum += value
	}
	return sum
}
