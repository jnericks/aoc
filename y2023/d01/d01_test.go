package d01

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	input := ReadInput()
	values, err := ParseInput(input)
	require.NoError(t, err)

	// assert.Zero(t, Sum(values))
	// assert.Equal(t, 53651, Sum(values)) // first star
	assert.Equal(t, 53894, Sum(values))
}

func Test_Example1(t *testing.T) {
	values, err := ParseInput([]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	})
	require.NoError(t, err)
	require.Len(t, values, 4)
	assert.Equal(t, 12, values[0])
	assert.Equal(t, 38, values[1])
	assert.Equal(t, 15, values[2])
	assert.Equal(t, 77, values[3])

	assert.Equal(t, 142, Sum(values))
}

func Test_Example2(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	expected := []int{29, 83, 13, 24, 42, 14, 76}
	require.Equal(t, len(expected), len(input))
	for i, exp := range expected {
		t.Run(input[i], func(t *testing.T) {
			value, err := ParseLine(input[i])
			require.NoError(t, err)
			assert.Equal(t, exp, value)
		})
	}

	values, err := ParseInput(input)
	require.NoError(t, err)
	assert.Equal(t, 281, Sum(values))
}

func Test_ReadInput(t *testing.T) {
	input := ReadInput()
	assert.NotZero(t, len(input))
	assert.Equal(t, "fivepqxlpninevh2xxsnsgg63pbvdnqptmg", input[0])
}
