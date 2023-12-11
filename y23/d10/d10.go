package d10

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
)

var (
	//go:embed input.txt
	InputData []byte
)

type Grid struct {
	Start *Pipe
	Pipes [][]*Pipe
}

func (g Grid) FarthestPipe() *Pipe {
	pipe := g.Pipes[0][0]
	for v := range g.Pipes {
		for _, p := range g.Pipes[v] {
			if p.Distance > pipe.Distance {
				pipe = p
			}
		}
	}
	return pipe
}

func (g Grid) PrintLoop() {
	for v := range g.Pipes {
		for h, pipe := range g.Pipes[v] {
			if v == g.Start.V && h == g.Start.H {
				fmt.Print("S")
			} else if pipe.Distance >= 0 {
				fmt.Print(string(pipe.Type))
				// fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (g Grid) String() string {
	var out bytes.Buffer
	for _, row := range g.Pipes {
		for _, tile := range row {
			out.WriteRune(tile.Type)
		}
		out.WriteByte('\n')
	}
	return out.String()
}

type Pipe struct {
	// V is the vertical index in the pipe grid (top to bottom)
	V int

	// H is the horizontal index in the pipe grid (left to right)
	H int

	// Type is one of { |, -, L, J, 7, F, ., S }
	Type rune

	// Pipe Connections
	North, East, South, West *Pipe

	// Distance is the distance from the Start pipe (-1 if not connected to start pipe).
	Distance int
}

func (p *Pipe) Neighbors() []*Pipe {
	out := make([]*Pipe, 0, 2)
	if p.North != nil {
		out = append(out, p.North)
	}
	if p.East != nil {
		out = append(out, p.East)
	}
	if p.South != nil {
		out = append(out, p.South)
	}
	if p.West != nil {
		out = append(out, p.West)
	}
	return out
}

func ParseData(data []string) (Grid, error) {
	var start *Pipe

	// setup tile grid
	pipes := make([][]*Pipe, len(data))
	for v, row := range data {
		pipes[v] = make([]*Pipe, len(row))
		for h, pipe := range row {
			pipes[v][h] = &Pipe{V: v, H: h, Type: pipe, Distance: -1}
		}
	}

	getPipe := func(v, h int) *Pipe {
		if v < 0 || v >= len(data) {
			return nil
		}
		if h < 0 || h >= len(data[v]) {
			return nil
		}
		return pipes[v][h]
	}

	for v, row := range data {
		for h, pipeType := range row {
			pipe := pipes[v][h]

			switch pipeType {
			case '|', 'L', 'J':
				pipe.North = getPipe(v-1, h)
			}

			switch pipeType {
			case '-', 'L', 'F':
				pipe.East = getPipe(v, h+1)
			}

			switch pipeType {
			case '|', '7', 'F':
				pipe.South = getPipe(v+1, h)
			}

			switch pipeType {
			case '-', 'J', '7':
				pipe.West = getPipe(v, h-1)
			}

			if pipeType == 'S' {
				// assume start pipe is connected to neighboring pipes
				pipe.North = getPipe(v-1, h)
				pipe.East = getPipe(v, h+1)
				pipe.South = getPipe(v+1, h)
				pipe.West = getPipe(v, h-1)

				pipe.Distance = 0
				start = pipe
			}

			pipes[v][h] = pipe
		}
	}

	if start == nil {
		return Grid{}, errors.New("no starting pipe found")
	}

	// remove pipes that are not bi-directionally connected
	for _, tileRow := range pipes {
		for _, tile := range tileRow {
			if tile.North != nil && tile.North.South != tile {
				tile.North = nil
			}
			if tile.East != nil && tile.East.West != tile {
				tile.East = nil
			}
			if tile.South != nil && tile.South.North != tile {
				tile.South = nil
			}
			if tile.West != nil && tile.West.East != tile {
				tile.West = nil
			}
		}
	}

	// find the loop
	type travel struct {
		current   *Pipe
		distances [][]int
		success   bool
	}

	visited := func(t travel, p *Pipe) bool {
		v, h := p.V, p.H
		if 0 <= v && v < len(t.distances[v]) {
			if 0 <= h && h < len(t.distances[v]) {
				return t.distances[v][h] >= 0
			}
		}
		return false
	}

	createStartingDistances := func() [][]int {
		distances := make([][]int, len(data))
		for v := range data {
			distances[v] = make([]int, len(data[v]))
			for h := range data[v] {
				distances[v][h] = -1
				if v == start.V && h == start.H {
					distances[v][h] = 0
				}
			}
		}
		return distances
	}

	neighbors := start.Neighbors()
	travels := make([]travel, len(neighbors))
	for i, pipe := range neighbors {
		dists := createStartingDistances()
		if pipe.North == start {
			dists[start.V+1][start.H] = 1
		}
		if pipe.East == start {
			dists[start.V][start.H-1] = 1
		}
		if pipe.South == start {
			dists[start.V-1][start.H] = 1
		}
		if pipe.West == start {
			dists[start.V][start.H+1] = 1
		}
		travels[i] = travel{
			current:   pipe,
			distances: dists,
		}
	}

	for _, t := range travels {
		for t.current != nil {
			d := t.distances[t.current.V][t.current.H]
			currNeighbors := t.current.Neighbors()
			t.current = nil
			for _, neighbor := range currNeighbors {
				// move if possible
				if !visited(t, neighbor) {
					t.distances[neighbor.V][neighbor.H] = d + 1
					t.current = neighbor
					break
				}
			}
		}
	}

	// min distances
	minDistances := createStartingDistances()
	for v := range minDistances {
		for h := range minDistances[v] {
			minDist := travels[0].distances[v][h]
			for i := 1; i < len(travels); i++ {
				minDist = min(minDist, travels[i].distances[v][h])
			}
			minDistances[v][h] = minDist
		}
	}

	for v := range pipes {
		for h, pipe := range pipes[v] {
			pipe.Distance = minDistances[v][h]
		}
	}

	return Grid{
		Start: start,
		Pipes: pipes,
	}, nil
}
