package d02

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_SolveStar1(t *testing.T) {
	gameInputs := ReadGameInputs()
	constraint := CubeSet{
		R: 12,
		G: 13,
		B: 14,
	}

	games, err := ParseGameInputs(gameInputs)
	require.NoError(t, err)

	n, err := SolveStar1(games, constraint)
	require.NoError(t, err)

	assert.Equal(t, 2545, n)
}

func Test_SolveStar2(t *testing.T) {
	gameInputs := ReadGameInputs()

	games, err := ParseGameInputs(gameInputs)
	require.NoError(t, err)

	sum := SolveStar2(games)
	assert.Equal(t, 78111, sum)
}

func Test_Example1(t *testing.T) {
	gameInputs := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	constraint := CubeSet{
		R: 12,
		G: 13,
		B: 14,
	}

	games, err := ParseGameInputs(gameInputs)
	require.NoError(t, err)

	n, err := SolveStar1(games, constraint)
	require.NoError(t, err)

	assert.Equal(t, 8, n)
}

func Test_Example2(t *testing.T) {

}

func Test_ParseGame(t *testing.T) {
	t.Run("Game 1", func(t *testing.T) {
		line := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

		game, err := ParseGameInput(line)
		require.NoError(t, err)

		assert.Equal(t, 1, game.ID)
		if assert.Len(t, game.CubeSets, 3) {
			assert.Equal(t, CubeSet{B: 3, R: 4}, game.CubeSets[0])
			assert.Equal(t, CubeSet{R: 1, G: 2, B: 6}, game.CubeSets[1])
			assert.Equal(t, CubeSet{G: 2}, game.CubeSets[2])
		}
	})
}

func Test_ParseCubeSetsInput(t *testing.T) {
	t.Run("3 blue, 4 red; 1 red, 2 green", func(t *testing.T) {
		cubeSets, err := ParseCubeSetsInput("3 blue, 4 red; 1 red, 2 green")
		require.NoError(t, err)
		require.Len(t, cubeSets, 2)
		assert.Equal(t, CubeSet{R: 4, G: 0, B: 3}, cubeSets[0])
		assert.Equal(t, CubeSet{R: 1, G: 2, B: 0}, cubeSets[1])
	})
}

func Test_ParseCubeSetInput(t *testing.T) {
	t.Run("3 blue, 4 red", func(t *testing.T) {
		cubeSet, err := ParseCubeSetInput("3 blue, 4 red")
		require.NoError(t, err)
		assert.Equal(t, CubeSet{R: 4, G: 0, B: 3}, cubeSet)
	})

	t.Run("1 red, 2 green, 6 blue", func(t *testing.T) {
		cubeSet, err := ParseCubeSetInput("1 red, 2 green, 6 blue ")
		require.NoError(t, err)
		assert.Equal(t, CubeSet{R: 1, G: 2, B: 6}, cubeSet)
	})

	t.Run("2 green", func(t *testing.T) {
		cubeSet, err := ParseCubeSetInput(" 2 green")
		require.NoError(t, err)
		assert.Equal(t, CubeSet{R: 0, G: 2, B: 0}, cubeSet)
	})
}

func Test_ParseCube(t *testing.T) {
	t.Run("4 red", func(t *testing.T) {
		n, color, err := ParseCubeInput("4 red")
		require.NoError(t, err)
		assert.Equal(t, 4, n)
		assert.Equal(t, ColorRED, color)
	})

	t.Run("2 green", func(t *testing.T) {
		n, color, err := ParseCubeInput("2 green ")
		require.NoError(t, err)
		assert.Equal(t, 2, n)
		assert.Equal(t, ColorGREEN, color)
	})

	t.Run("3 blue ", func(t *testing.T) {
		n, color, err := ParseCubeInput(" 3 blue")
		require.NoError(t, err)
		assert.Equal(t, 3, n)
		assert.Equal(t, ColorBLUE, color)
	})
}

func Test_CollectNumberInOrder(t *testing.T) {
	t.Run("Game 1", func(t *testing.T) {
		n, err := ParseNumber("Game 1")
		require.NoError(t, err)
		assert.Equal(t, 1, n)
	})

	t.Run("Game 99", func(t *testing.T) {
		n, err := ParseNumber("Game 99 ")
		require.NoError(t, err)
		assert.Equal(t, 99, n)
	})

	t.Run("3 blue", func(t *testing.T) {
		n, err := ParseNumber(" 3 blue ")
		require.NoError(t, err)
		assert.Equal(t, 3, n)
	})
}

func Test_IsGameValid(t *testing.T) {
	constraint := CubeSet{
		R: 12,
		G: 13,
		B: 14,
	}

	t.Run("invalid", func(t *testing.T) {
		assert.False(t, IsGameValid(Game{
			ID: 1,
			CubeSets: []CubeSet{
				{R: 20, G: 8, B: 6},
				{R: 4, G: 0, B: 4},
				{R: 0, G: 13, B: 0},
				{R: 1, G: 5, B: 0},
			},
		}, constraint))
	})

	t.Run("valid", func(t *testing.T) {
		assert.True(t, IsGameValid(Game{
			ID: 1,
			CubeSets: []CubeSet{
				{R: 4, G: 0, B: 3},
				{R: 1, G: 2, B: 6},
				{R: 0, G: 2, B: 0},
			},
		}, constraint))
	})
}
