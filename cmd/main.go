package main

import "git.sr.ht/~anecdotal/ranj/game"

func main() {
	board := game.CreateBoard()
	board.RenderBoard()
}
