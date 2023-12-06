package d06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solve(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		ans := WaysToWin(38, 234)
		ans *= WaysToWin(67, 1027)
		ans *= WaysToWin(76, 1157)
		ans *= WaysToWin(73, 1236)
		assert.Equal(t, 303_600, ans)
	})

	t.Run("Part 2", func(t *testing.T) {
		ans := WaysToWin(38_67_76_73, 234_1027_1157_1236)
		assert.Equal(t, 23_654_842, ans)
	})
}

func Test_WaysToWin(t *testing.T) {
	assert.Equal(t, 4, WaysToWin(7, 9))
	assert.Equal(t, 8, WaysToWin(15, 40))
	assert.Equal(t, 9, WaysToWin(30, 200))
}
