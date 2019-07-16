package main

import (
	"fmt"
	p "n-puzzle/parsing"
	s "n-puzzle/solver"
	"os"
)

// func IntToString2() string {
//     a := []int{1, 2, 3, 4, 5}
//     b := make([]string, len(a))
//     for i, v := range a {
//         b[i] = strconv.Itoa(v)
//     }

//     return strings.Join(b, ",")
// }

func main() {

	size, solve, iterations := p.CheckFlags()
	var Puzzle []int
	if size == 0 && solve == false && iterations == 0 {
		Puzzle, size = p.ReadBoardFromFile(Puzzle, size)
	} else {
		Puzzle = p.GenerateRandomBoard(size, solve, iterations)
		//s.MovePieces(puzzle, size)
	}
	s.Solver(Puzzle, size, solve, iterations)
	fmt.Println("\n\n\n You've reached the end of main()")
	os.Exit(1)

	//fmt.Println("no file or random board")
}
