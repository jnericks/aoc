package d02

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func ReadGameInputs() []string {
	s := bufio.NewScanner(bytes.NewBuffer(input))

	var out []string
	for s.Scan() {
		out = append(out, s.Text())
	}

	return out
}

func ParseGameInputs(gameInputs []string) ([]Game, error) {
	games := make([]Game, len(gameInputs))
	for i, gameInput := range gameInputs {
		game, err := ParseGameInput(gameInput)
		if err != nil {
			return nil, err
		}
		games[i] = game
	}
	return games, nil
}

// SolveStar1 finds the number of possible games that match the constraint of the provided CubeSet.
func SolveStar1(games []Game, constraint CubeSet) (int, error) {
	var answer int
	for _, game := range games {
		if IsGameValid(game, constraint) {
			answer += game.ID
		}
	}
	return answer, nil
}

// SolveStar2 calculates the MinimumCubeSet for each game and sums them.
func SolveStar2(games []Game) int {
	var sum int
	for _, game := range games {
		sum += game.MinimumCubeSet().Power()
	}
	return sum
}

const (
	ColorRED   = "red"
	ColorGREEN = "green"
	ColorBLUE  = "blue"
)

type Game struct {
	ID       int
	CubeSets []CubeSet
}

// MinimumCubeSet is the fewest number of cubes of each color that could have been in the bag to make the game possible.
func (g Game) MinimumCubeSet() CubeSet {
	var out CubeSet
	for _, cubeSet := range g.CubeSets {
		out.R = max(out.R, cubeSet.R)
		out.G = max(out.G, cubeSet.G)
		out.B = max(out.B, cubeSet.B)
	}
	return out
}

type CubeSet struct {
	R, G, B int
}

// Power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together.
func (cs CubeSet) Power() int {
	return cs.R * cs.G * cs.B
}

func ParseGameInput(gameInput string) (Game, error) {
	gameAndCubeSetsInput := strings.Split(gameInput, ":")
	if len(gameAndCubeSetsInput) != 2 {
		return Game{}, errors.New("invalid game input format")
	}

	id, err := ParseNumber(gameAndCubeSetsInput[0])
	if err != nil {
		return Game{}, err
	}

	cubeSets, err := ParseCubeSetsInput(gameAndCubeSetsInput[1])
	if err != nil {
		return Game{}, err
	}

	return Game{
		ID:       id,
		CubeSets: cubeSets,
	}, nil
}

func ParseCubeSetsInput(cubeSetsInput string) ([]CubeSet, error) {
	var cubeSets []CubeSet
	for _, cubeSetInput := range strings.Split(cubeSetsInput, ";") {
		cubeSet, err := ParseCubeSetInput(cubeSetInput)
		if err != nil {
			return nil, err
		}
		cubeSets = append(cubeSets, cubeSet)
	}
	return cubeSets, nil
}

func ParseCubeSetInput(cubeSetInput string) (CubeSet, error) {
	var cubeSet CubeSet
	for _, cubeInput := range strings.Split(cubeSetInput, ",") {
		n, color, err := ParseCubeInput(cubeInput)
		if err != nil {
			return CubeSet{}, err
		}
		switch color {
		case ColorRED:
			cubeSet.R = n
		case ColorGREEN:
			cubeSet.G = n
		case ColorBLUE:
			cubeSet.B = n
		}
	}
	return cubeSet, nil
}

func ParseCubeInput(cubeInput string) (int, string, error) {
	n, err := ParseNumber(cubeInput)
	if err != nil {
		return n, "", err
	}

	for _, color := range []string{ColorRED, ColorGREEN, ColorBLUE} {
		if strings.Contains(cubeInput, color) {
			return n, color, nil
		}
	}

	return n, "", errors.New("invalid cube color")
}

func ParseNumber(s string) (int, error) {
	var sb strings.Builder
	for _, r := range s {
		if '0' <= r && r <= '9' {
			sb.WriteRune(r)
		}
	}
	return strconv.Atoi(sb.String())
}

func IsGameValid(game Game, constraint CubeSet) bool {
	for _, cubeSet := range game.CubeSets {
		if cubeSet.R > constraint.R {
			return false
		}
		if cubeSet.G > constraint.G {
			return false
		}
		if cubeSet.B > constraint.B {
			return false
		}
	}
	return true
}
