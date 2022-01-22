package game

import (
	"fmt"
)

func whitePieceFg(s string) string {
	return "\033[0;107m" + s + "\033[0m"
}

func blackPieceFg(s string) string {
	return "\033[37;100m" + s + "\033[0m"
}

func getPieceUnicode(pieceNumber int, background int) string {
	piece := "-"

	switch pieceNumber {
	case ROOK:
		piece = "♜"
	case KNIGHT:
		piece = "♞"
	case BISHOP:
		piece = "♝"
	case KING:
		piece = "♚"
	case QUEEN:
		piece = "♛"
	case PAWN:
		piece = "♟︎"
	}

	if background == BLACK {
		return blackPieceFg(piece)
	}

	return whitePieceFg(piece)
}

func (b *Board) RenderBoard() {
	for row := 0; row < len(b.cells); row++ {
		for col := 0; col < len(b.cells[row]); col++ {
			bg := WHITE
			if (row+col)%2 != 0 { // if cell is even numbered
				bg = BLACK
			}
			occupant := b.cells[row][col].occupant
			fmt.Print(getPieceUnicode(occupant.pieceNum, bg))
		}
		fmt.Println()
	}
}
