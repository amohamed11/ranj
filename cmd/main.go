package main

import (
	"fmt"

	"git.sr.ht/~anecdotal/ranj/game"
)

func main() {
	board := game.CreateBoard()
	fmt.Print(board.RenderBoard())
}
