package game

import (
	"fmt"
)

// Game is the name we want but since we are
// to use it as an interface, we will change
// the name into Gameer.
type Gameer interface {

	// Note that we will
	// declare the methods to be used
	// later here in this
	// interface
	PlacePiece(int, int, int)
	GetState() [][]int
}

type GameState struct {
	board [][]int
}

func (x GameState) GetState() [][]int {
	fmt.Println(x.board)
	return x.board
}

func (x GameState) PlacePiece(row int, col int, player int) {
	x.board[row][col] = player
	fmt.Println(x.board)
}
