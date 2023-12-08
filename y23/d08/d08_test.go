package d08

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		m, err := ParseData(ReadData(InputData))
		require.NoError(t, err)
		assert.Equal(t, 12_361, m.Walk("AAA", NodeIsZZZ))
	})

	t.Run("part 2", func(t *testing.T) {
		m, err := ParseData(ReadData(InputData))
		require.NoError(t, err)
		assert.Len(t, m.StartNodes, 6)

		assert.Equal(t, 18_215_611_419_223, m.WalkStartNodes())
	})
}

func Test_Example3(t *testing.T) {
	m, err := ParseData(ReadData(Example3Data))
	require.NoError(t, err)

	assert.Equal(t, 2, m.Walk("11A", NodeEndsWithZ))
	assert.Equal(t, 3, m.Walk("22A", NodeEndsWithZ))

	assert.Equal(t, 6, LCM(2, 3))
	assert.Equal(t, 6, m.WalkStartNodes())
}

func Test_Example2(t *testing.T) {
	m, err := ParseData(ReadData(Example2Data))
	require.NoError(t, err)
	assert.Equal(t, 6, m.Walk("AAA", NodeIsZZZ))
}

func Test_Example1(t *testing.T) {
	m, err := ParseData(ReadData(Example1Data))
	require.NoError(t, err)

	assert.Equal(t, []byte("RL"), m.Path)

	expectedGuide := []struct {
		pos, l, r string
	}{
		{pos: "AAA", l: "BBB", r: "CCC"},
		{pos: "BBB", l: "DDD", r: "EEE"},
		{pos: "CCC", l: "ZZZ", r: "GGG"},
		{pos: "DDD", l: "DDD", r: "DDD"},
		{pos: "EEE", l: "EEE", r: "EEE"},
		{pos: "GGG", l: "GGG", r: "GGG"},
		{pos: "ZZZ", l: "ZZZ", r: "ZZZ"},
	}

	assert.Len(t, m.Guide, len(expectedGuide))
	for _, tt := range expectedGuide {
		t.Run(fmt.Sprintf("%s = (%s, %s)", tt.pos, tt.l, tt.r), func(t *testing.T) {
			ins, ok := m.Guide[tt.pos]
			require.True(t, ok)
			assert.Equal(t, tt.pos, ins.Pos)
			assert.Equal(t, tt.l, ins.L.Pos)
			assert.Equal(t, tt.r, ins.R.Pos)
		})
	}

	assert.Equal(t, 2, m.Walk("AAA", NodeIsZZZ))
}

func Test_ParseInput(t *testing.T) {

}
