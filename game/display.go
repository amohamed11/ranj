package game

import (
	"fmt"
)

func getPieceUnicode(pieceNumber int, player int) string {
	switch player {
	case WHITE:
		switch pieceNumber {
		case ROOK:
			return "♖"
		case KNIGHT:
			return "♘"
		case BISHOP:
			return "♗"
		case KING:
			return "♔"
		case QUEEN:
			return "♕"
		case PAWN:
			return "♙"
		}
	case BLACK:
		switch pieceNumber {
		case ROOK:
			return "♜"
		case KNIGHT:
			return "♞"
		case BISHOP:
			return "♝"
		case KING:
			return "♚"
		case QUEEN:
			return "♛"
		case PAWN:
			return "♟︎"
		}

	}

	return "-"
}

func (b *Board) RenderBoard() {
	for row := 0; row < len(b.cells); row++ {
		for col := 0; col < len(b.cells[row]); col++ {
			occupant := b.cells[row][col].occupant
			fmt.Print(getPieceUnicode(occupant.pieceNum, occupant.owner))
		}
		fmt.Println()
	}
}
