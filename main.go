package main

import (
	l "N-Puzzle-Go/golib"
	p "N-Puzzle-Go/parsing"
	s "N-Puzzle-Go/solver"
	"fmt"
	"os"
)

func main() {
	options := p.GetFlags()
	Puzzle := l.MakeRandomBoard(options)
	fmt.Println(Puzzle)
	// fmt.Println(heuristics)
	s.Solver(Puzzle, options)
	os.Exit(1)
}
