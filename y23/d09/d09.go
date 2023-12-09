package d09

import (
	"bufio"
	"bytes"
	_ "embed"
	"strconv"
	"strings"
)

var (
	//go:embed example1.txt
	Example1Data []byte

	//go:embed input.txt
	InputData []byte
)

func ReadData(data []byte) []string {
	s := bufio.NewScanner(bytes.NewBuffer(data))
	var out []string
	for s.Scan() {
		out = append(out, s.Text())
	}
	return out
}

func ParseData(data []string) ([][]int, error) {
	out := make([][]int, 0, len(data))
	for _, d := range data {
		split := strings.Split(d, " ")
		row := make([]int, 0, len(split))
		for _, n := range split {
			i, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			row = append(row, i)
		}
		out = append(out, row)
	}
	return out, nil
}

func SolveHistories(histories [][]int) (first int, last int) {
	for _, h := range histories {
		f, l := SolveHistory(h)
		first += f
		last += l
	}
	return first, last
}

func SolveHistory(history []int) (first int, last int) {
	layers := HistoryLayers(history)
	if len(layers) == 0 {
		return 0, 0
	}
	if len(layers[0]) == 0 {
		return 0, 0
	}
	first = layers[0][0]
	last = layers[0][len(layers[0])-1]
	return first, last
}

func HistoryLayers(history []int) [][]int {
	layers := buildStartingHistoryLayers([][]int{history})

	// part 1, add to end
	{
		// start by appending a 0 to last layer
		layers[len(layers)-1] = append(layers[len(layers)-1], 0)

		// bottom up
		for i := len(layers) - 2; i >= 0; i-- {
			curr := layers[i]
			last := layers[i+1]
			layers[i] = append(curr, curr[len(curr)-1]+last[len(last)-1])
		}
	}

	// part 2, add to beginning
	{
		// start by prepending a 0 to last layer
		layers[len(layers)-1] = prepend(layers[len(layers)-1], 0)

		// bottom up
		for i := len(layers) - 2; i >= 0; i-- {
			curr := layers[i]
			last := layers[i+1]
			layers[i] = prepend(curr, curr[0]-last[0])
		}
	}

	return layers
}

func buildStartingHistoryLayers(histories [][]int) [][]int {
	if len(histories) == 0 {
		return histories
	}

	last := histories[len(histories)-1]
	if len(last) == 0 {
		return histories
	}

	if last[0] == 0 && last[len(last)-1] == 0 {
		return histories
	}

	layer := make([]int, 0, len(last)-1)
	for i := 1; i < len(last); i++ {
		layer = append(layer, last[i]-last[i-1])
	}

	return buildStartingHistoryLayers(append(histories, layer))
}

func prepend(values []int, value int) []int {
	out := make([]int, 0, len(values)+1)
	out = append(out, value)
	out = append(out, values...)
	return out
}
