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

	// Nodes are the nodes in the graph.
	Nodes map[string]*Node

	// StartNodes are nodes with an 'A' suffix.
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
		L string
		R string
	}

	nodeMap := make(map[string]parsed)
	nodes := make(map[string]*Node)
	var startNodes []*Node
	for _, row := range data[2:] {
		pos := row[0:3]
		nodeMap[pos] = parsed{
			L: row[7:10],
			R: row[12:15],
		}
		nodes[pos] = &Node{Pos: pos}
		if pos[2] == 'A' {
			startNodes = append(startNodes, nodes[pos])
		}
	}

	// link nodes in the graph
	for pos, d := range nodeMap {
		nodes[pos].L = nodes[d.L]
		nodes[pos].R = nodes[d.R]
	}

	return Map{
		Path:       format,
		Nodes:      nodes,
		StartNodes: startNodes,
	}, nil
}

func NodeZZZ(node *Node) bool {
	return node.Pos == "ZZZ"
}

func NodeEndsWithZ(node *Node) bool {
	return node.Pos[len(node.Pos)-1] == 'Z'
}

func (m *Map) Walk(start string, isEnd func(*Node) bool) int {
	node := m.Nodes[start]

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
