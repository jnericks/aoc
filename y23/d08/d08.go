package d08

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"math"
)

var (
	//go:embed example1.txt
	Example1Data []byte

	//go:embed example2.txt
	Example2Data []byte

	//go:embed example3.txt
	Example3Data []byte

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

type Map struct {
	Path []byte

	Guide map[string]*Node

	StartNodes []*Node
}

type Node struct {
	Pos string
	L   *Node
	R   *Node
}

func (i *Node) Next(path byte) *Node {
	if path == 'L' {
		return i.L
	}
	if path == 'R' {
		return i.R
	}
	panic(fmt.Sprintf("invalid path '%s'", string(path)))
}

func ParseData(data []string) (Map, error) {
	if len(data) < 3 {
		return Map{}, errors.New("invalid map")
	}

	format := []byte(data[0])
	for _, dir := range format {
		if dir != 'L' && dir != 'R' {
			return Map{}, fmt.Errorf("invalid direction detected '%s' in map", string(dir))
		}
	}

	if data[1] != "" {
		return Map{}, errors.New("invalid map, expected blank 2nd row")
	}

	type parsed struct {
		Pos string
		L   string
		R   string
	}

	guide := make(map[string]*Node)
	intermediary := make(map[string]parsed)
	var startNodes []*Node
	for _, row := range data[2:] {
		pos := row[0:3]
		intermediary[pos] = parsed{Pos: pos, L: row[7:10], R: row[12:15]}
		node := &Node{Pos: pos}
		guide[pos] = node
		if pos[2] == 'A' {
			startNodes = append(startNodes, guide[pos])
		}
	}

	for _, d := range intermediary {
		guide[d.Pos].L = guide[d.L]
		guide[d.Pos].R = guide[d.R]
	}

	return Map{
		Path:       format,
		Guide:      guide,
		StartNodes: startNodes,
	}, nil
}

func NodeIsZZZ(node *Node) bool {
	return node.Pos == "ZZZ"
}

func NodeEndsWithZ(node *Node) bool {
	return node.Pos[2] == 'Z'
}

func (m *Map) Walk(start string, isEnd func(*Node) bool) int {
	node := m.Guide[start]

	var steps int
	for i := 0; i < math.MaxInt; i++ {
		if isEnd(node) {
			return steps
		}

		// move
		switch m.Path[i%len(m.Path)] {
		case 'L':
			node = node.L
		case 'R':
			node = node.R
		default:
			panic("unexpected")
		}
		steps++
	}

	return -1
}

func (m *Map) WalkStartNodes() int {
	lcm := m.Walk(m.StartNodes[0].Pos, NodeEndsWithZ)
	if len(m.StartNodes) == 1 {
		return lcm
	}

	for _, node := range m.StartNodes[1:] {
		lcm = LCM(lcm, m.Walk(node.Pos, NodeEndsWithZ))
	}

	return lcm
}

func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a * b / GCD(a, b)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
