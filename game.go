package tanks

import "errors"

type MapCell char

const (
	Ground MapCell = '.'
	Wall           = '#'
)

type Player struct {
	X, Y   int
	Health int
}

type GameMap struct {
	Width  int
	Height int
	Cells  [][]MapCell
}

type GameState struct {
	Map     GameMap
	Players map[string]*Player
	Turn    int
}

func NewGameMap(height, width int) *GameMap {
	gameMap := GameMap{
		Width:  width,
		Height: height,
		Cells:  make([][]MapCell, height)}
	for i := 0; i < height; i++ {
		gameMap.Cells[i] = make([]MapCell, width)
	}
	return &gameMap
}

func NewGameState(gameMap GameMap) *GameState {
	return &GameState{
		Map:     gameMap,
		Players: make(map[string]*Player),
		Turn:    0}
}

type Command interface {
	// Mutate gameState according to the command. Return error if impossible.
	// Presence of error assumes gameState is unchanged.
	Handle(gameState *GameState) error
}

type WaitCommand struct{}

func (c *WaitCommand) Handle(gameState *GameState) error {
	return nil
}

type MoveCommand struct {
	Name   string
	DX, DY int
}

func (c *MoveCommand) CanMoveTo(cell MapCell) bool {
	return cell == Ground
}

func (c *MoveCommand) Handle(gameState *GameState) error {
	player := gameState.Players[c.Name]
	if player == nil {
		return errors.New("no such player")
	}
	if c.DX < -1 || c.DX > 1 || c.DY < -1 || c.DY > 1 || (c.DX == 0 && c.DY == 0) {
		return errors.New("invalid DX, DY")
	}
	newX := player.X + c.DX
	newY := player.Y + c.DY
	if newX < 0 || newX >= gameState.Map.Width || newY < 0 || newY >= gameState.Map.Height {
		return errors.New("invalid movement")
	}
	if !c.CanMoveTo(gameState.Map[newY][newX]) {
		return errors.New("invalid movement")
	}
	for name, player := range gameState.Players {
		if name != c.Name && player.X == newX && player.Y == newY {
			return errors.New("cell occupied")
		}
	}

	player.X = newX
	player.Y = newY
	return nil
}
