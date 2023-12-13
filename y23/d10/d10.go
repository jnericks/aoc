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

type TileType string

const (
	// PipeNS '|' is a vertical pipe connecting north and south.
	PipeNS TileType = "|"
	// PipeEW '-' is a horizontal pipe connecting east and west.
	PipeEW TileType = "-"
	// PipeNE 'L' is a 90-degree bend connecting north and east.
	PipeNE TileType = "L"
	// PipeNW 'J' is a 90-degree bend connecting north and west.
	PipeNW TileType = "J"
	// PipeSW '7' is a 90-degree bend connecting south and west.
	PipeSW TileType = "7"
	// PipeSE 'F' is a 90-degree bend connecting south and east.
	PipeSE TileType = "F"
	// Ground '.' is ground; there is no pipe in this tile.
	Ground TileType = "."
)

func (t TileType) String() string {
	switch t {
	case PipeNS:
		return "│"
	case PipeEW:
		return "─"
	case PipeNE:
		return "└"
	case PipeNW:
		return "┘"
	case PipeSW:
		return "┐"
	case PipeSE:
		return "┌"
	case Ground:
		return "."
	}
	return string(t)
}

type Grid struct {
	Start *Tile
	Pipes [][]*Tile
}

func (g Grid) FarthestPipe() *Tile {
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
				fmt.Print(pipe.Type)
				// fmt.Print("S")
			} else if pipe.Distance >= 0 { // part of the loop
				fmt.Print(pipe.Type)
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
			out.WriteString(tile.Type.String())
		}
		out.WriteByte('\n')
	}
	return out.String()
}

type Tile struct {
	// V is the vertical index in the pipe grid (top to bottom)
	V int

	// H is the horizontal index in the pipe grid (left to right)
	H int

	// Start is true if this is the starting tile.
	Start bool

	// Type is the tile type.
	Type TileType

	// Tile Connections
	North, East, South, West *Tile

	// Distance is the distance from the Start pipe (-1 if not connected to start pipe).
	Distance int
}

func (p *Tile) Neighbors() []*Tile {
	out := make([]*Tile, 0, 2)
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
	var start *Tile

	// setup tile grid
	tiles := make([][]*Tile, len(data))
	for v, row := range data {
		tiles[v] = make([]*Tile, len(row))
		for h, pipe := range row {
			tiles[v][h] = &Tile{V: v, H: h, Type: TileType(pipe), Distance: -1}
		}
	}

	getPipe := func(v, h int) *Tile {
		if v < 0 || v >= len(data) {
			return nil
		}
		if h < 0 || h >= len(data[v]) {
			return nil
		}
		return tiles[v][h]
	}

	for v, row := range data {
		for h, t := range row {
			tile := tiles[v][h]
			tileType := TileType(t)

			switch tileType { // North Connection
			case PipeNS, PipeNE, PipeNW:
				tile.North = getPipe(v-1, h)
			}

			switch tileType { // East Connection
			case PipeEW, PipeNE, PipeSE:
				tile.East = getPipe(v, h+1)
			}

			switch tileType { // South Connection
			case PipeNS, PipeSE, PipeSW:
				tile.South = getPipe(v+1, h)
			}

			switch tileType { // West Connection
			case PipeEW, PipeNW, PipeSW:
				tile.West = getPipe(v, h-1)
			}

			if t == 'S' {
				// assume start tile is connected to neighboring tiles
				tile.North = getPipe(v-1, h)
				tile.East = getPipe(v, h+1)
				tile.South = getPipe(v+1, h)
				tile.West = getPipe(v, h-1)

				tile.Distance = 0
				tile.Start = true
				start = tile
			}

			tiles[v][h] = tile
		}
	}

	if start == nil {
		return Grid{}, errors.New("no starting pipe found")
	}

	// remove tiles that are not bi-directionally connected
	for _, row := range tiles {
		for _, tile := range row {
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
		current   *Tile
		distances [][]int
		success   bool
	}

	visited := func(t travel, p *Tile) bool {
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

	for v := range tiles {
		for h, pipe := range tiles[v] {
			pipe.Distance = minDistances[v][h]
		}
	}

	return Grid{
		Start: start,
		Pipes: tiles,
	}, nil
}
