// Command.go
package beacons

type Command int

const (
	NORTH Command = iota
	NORTH_EAST
	EAST
	SOUTH_EAST
	SOUTH
	SOUTH_WEST
	WEST
	NORTH_WEST
	MOVE
	WAIT
)

var directions = map[Command]Position{
	NORTH:      Position{0, 1},
	NORTH_EAST: Position{1, 1},
	EAST:       Position{1, 0},
	SOUTH_EAST: Position{1, -1},
	SOUTH:      Position{0, -1},
	SOUTH_WEST: Position{-1, -1},
	WEST:       Position{-1, 0},
	NORTH_WEST: Position{-1, 1},
}
