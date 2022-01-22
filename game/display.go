package game

// global map for chess pieces by owner color
var OWNER_PIECE = map[int]map[int]string{
	0: {
		ROOK:   "♖",
		KNIGHT: "♘",
		BISHOP: "♗",
		KING:   "♔",
		QUEEN:  "♕",
		PAWN:   "♙",
	},
	1: {
		ROOK:   "♜",
		KNIGHT: "♞",
		BISHOP: "♝",
		KING:   "♚",
		QUEEN:  "♛",
		PAWN:   "♟︎",
	},
}

func whitePieceFg(s string) string {
	return "\033[0;107m" + s + "\033[0m"
}

func blackPieceFg(s string) string {
	return "\033[37;100m" + s + "\033[0m"
}

func getPieceUnicode(pieceNumber int, owner int, background int) string {
	piece := " "

	if pieceNumber > 0 {
		piece = OWNER_PIECE[owner][pieceNumber]
	}

	if background == BLACK {
		return blackPieceFg(piece)
	}

	return whitePieceFg(piece)
}

func (b *Board) RenderBoard() string {
	boardString := ""
	for row := 0; row < len(b.cells); row++ {
		for col := 0; col < len(b.cells[row]); col++ {
			bg := WHITE
			if (row+col)%2 != 0 { // if cell is even numbered
				bg = BLACK
			}
			occupant := b.cells[row][col].occupant
			boardString += getPieceUnicode(occupant.pieceNum, occupant.owner, bg)
		}
		boardString += "\n"
	}

	return boardString
}
