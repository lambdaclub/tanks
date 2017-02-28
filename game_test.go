package tanks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	gameMap := NewGameMap(2, 3)
	assert.Equal(t, 2, gameMap.Height)
	assert.Equal(t, 3, gameMap.Width)
	assert.Equal(t, []MapCell{Ground, Ground, Ground}, gameMap.Cells[0])
}

func BenchmarkNewMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewGameMap(20, 20)
	}
}
