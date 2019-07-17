package main

import (
	l "N-Puzzle-Go/golib"
	p "N-Puzzle-Go/parsing"
	s "N-Puzzle-Go/solver"
	"fmt"
	"os"
)

func main() {
	size, solve, iterations := p.GetFlags()
	Puzzle := l.MakeRandomBoard(size, solve, iterations)
	fmt.Println(Puzzle)
	//
	// var Puzzle []int
	// if size == 0 && solve == false && iterations == 0 {
	// 	Puzzle, size = p.ReadBoardFromFile(Puzzle, size)
	// } else {
	// 	Puzzle = p.GenerateRandomBoard(size, solve, iterations)
	// 	//s.MovePieces(puzzle, size)
	// }
	s.Solver(Puzzle, size, solve, iterations)
	os.Exit(1)
}
