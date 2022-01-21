package game

// Our board, which contains 64 cells
type Board struct {
	turn  int
	cells [8][8]Cell
}

// A single cell has a row (1-8), a column (1-8)
// and an occupant which is an int corresponding to pieces enum
type Cell struct {
	row      int
	col      int
	occupant Piece
}

type Piece struct {
	owner    int
	pieceNum int
}

// Player enum
const (
	WHITE int = 0
	BLACK     = 1
)

// Pieces enum
const (
	EMPTY  int = 0
	ROOK       = 1
	KNIGHT     = 2
	BISHOP     = 3
	KING       = 4
	QUEEN      = 5
	PAWN       = 6
)

func CreateBoard() Board {
	b := Board{}

	// White beginning rows
	b.cells[0] = [8]Cell{
		{0, 0, Piece{0, 1}},
		{0, 1, Piece{0, 2}},
		{0, 2, Piece{0, 3}},
		{0, 3, Piece{0, 4}},
		{0, 4, Piece{0, 5}},
		{0, 5, Piece{0, 3}},
		{0, 6, Piece{0, 2}},
		{0, 7, Piece{0, 1}},
	}
	// White pawns
	b.cells[1] = [8]Cell{
		{1, 0, Piece{0, 6}},
		{1, 1, Piece{0, 6}},
		{1, 2, Piece{0, 6}},
		{1, 3, Piece{0, 6}},
		{1, 4, Piece{0, 6}},
		{1, 5, Piece{0, 6}},
		{1, 6, Piece{0, 6}},
		{1, 7, Piece{0, 6}},
	}

	// Black beginning rows
	b.cells[7] = [8]Cell{
		{7, 0, Piece{1, 1}},
		{7, 1, Piece{1, 2}},
		{7, 2, Piece{1, 3}},
		{7, 3, Piece{1, 5}},
		{7, 4, Piece{1, 4}},
		{7, 5, Piece{1, 3}},
		{7, 6, Piece{1, 2}},
		{7, 7, Piece{1, 1}},
	}
	// Black pawns
	b.cells[6] = [8]Cell{
		{6, 0, Piece{1, 6}},
		{6, 1, Piece{1, 6}},
		{6, 2, Piece{1, 6}},
		{6, 3, Piece{1, 6}},
		{6, 4, Piece{1, 6}},
		{6, 5, Piece{1, 6}},
		{6, 6, Piece{1, 6}},
		{6, 7, Piece{1, 6}},
	}

	return b
}
