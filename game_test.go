package tanks

import (
	"errors"
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

func TestMoveHandle(t *testing.T) {
	gameMap := GameMap{2, 2, [][]MapCell{
		[]MapCell{Ground, Ground},
		[]MapCell{Ground, Ground}}}
	examples := []struct {
		gameState    *GameState
		newGameState *GameState
		command      *MoveCommand
		err          error
	}{{
		&GameState{gameMap, map[string]*Player{"p1": {0, 0, 1}}, 0},
		&GameState{gameMap, map[string]*Player{"p1": {0, 0, 1}}, 0},
		&MoveCommand{"p1", 1, 1},
		errors.New("invalid movement")}}
	for _, example := range examples {
		err := example.command.Handle(example.gameState)
		assert.Equal(t, example.err, err)
		assert.Equal(t, example.gameState, example.newGameState)
	}
}
