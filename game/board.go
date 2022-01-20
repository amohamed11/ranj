package game

// Player enum
const (
	WHITE int = 0
	BLACK     = 1
)

// Our board, which contains 64 cells
type Board struct {
	turn  int
	cells [64]Cell
}

// A single cell has a row (1-8), a column (1-8)
// and an occupant which is an int corresponding to pieces enum
type Cell struct {
	row      int
	col      int
	occupant int
}

// Pieces enum
const (
	Empty  int = 0
	Rook       = 1
	Knight     = 2
	Bishop     = 3
	Queen      = 4
	King       = 5
	Pawn       = 6
)
