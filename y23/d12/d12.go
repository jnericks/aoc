package d12

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/jnericks/aoc/util"
)

var (
	//go:embed example.txt
	FileExample []byte

	//go:embed input.txt
	FileInput []byte
)

type Game struct {
	Records []Record
}

func (g *Game) SolvePart1() int64 {
	var sum int64
	for _, record := range g.Records {
		sum += NumArrangements(record)
	}
	return sum
}

func (g *Game) SolvePart2() int64 {
	const repeat = 5
	var sum int64
	for _, record := range g.Records {
		rows := make([]string, repeat)
		for i := range rows {
			rows[i] = record.Row
		}
		row := strings.Join(rows, "?")
		values := make([]int, 0, len(record.Values)*repeat)
		for i := 0; i < repeat; i++ {
			values = append(values, record.Values...)
		}
		sum += NumArrangements(Record{
			Row:    row,
			Values: values,
		})
	}
	return sum
}

type Record struct {
	Row    string
	Values []int
}

func Parse(file []byte) (*Game, error) {
	data := util.ReadStrings(file)

	var records []Record

	for i, row := range data {
		parts := strings.Split(row, " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid row %d: %s", i, row)
		}

		nums := strings.Split(parts[1], ",")
		values := make([]int, 0, len(nums))
		for _, num := range nums {
			i, err := strconv.Atoi(num)
			if err != nil {
				return nil, fmt.Errorf("invalid row %d: %s", i, row)
			}
			values = append(values, i)
		}

		records = append(records, Record{
			Row:    parts[0],
			Values: values,
		})
	}

	return &Game{Records: records}, nil
}

func NumArrangements(record Record) int64 {
	var damaged, unknown int
	for _, r := range record.Row {
		switch r {
		case '#':
			damaged++
		case '?':
			unknown++
		}
	}
	hashes := util.Sum(record.Values) - damaged
	return numArrangements([]rune(strings.Clone(record.Row)), 0, hashes, unknown, record.Values, make(map[string]int64))
}

func numArrangements(row []rune, consecutive, damaged, unknown int, values []int, memo map[string]int64) int64 {
	key := fmt.Sprintf("%s_%d_%d_%d_%v", string(row), consecutive, damaged, unknown, values)
	if ans, ok := memo[key]; ok {
		return ans
	}

	memo[key] = func() int64 {
		if len(row) == 0 {
			if unknown == 0 && damaged == 0 {
				return 1
			}
			return 0
		}

		switch row[0] {
		case '?':
			var result int64
			// have operational (.) left to place
			if unknown-damaged > 0 {
				cloned := []rune(strings.Clone(string(row)))
				cloned[0] = '.'
				result += numArrangements(cloned, consecutive, damaged, unknown-1, values, memo)
			}
			// have damaged (#) left to place
			if damaged != 0 {
				cloned := []rune(strings.Clone(string(row)))
				cloned[0] = '#'
				result += numArrangements(cloned, consecutive, damaged-1, unknown-1, values, memo)
			}
			return result

		case '.':
			if consecutive != 0 {
				if values[0] == consecutive {
					return numArrangements(row[1:], 0, damaged, unknown, values[1:], memo)
				}
				return 0
			}
			return numArrangements(row[1:], 0, damaged, unknown, values, memo)

		case '#':
			consecutive++
			if values[0] < consecutive {
				return 0
			}
			return numArrangements(row[1:], consecutive, damaged, unknown, values, memo)

		}

		return 0
	}()
	return memo[key]
}
