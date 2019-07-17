package parsing

import (
	l "N-Puzzle-Go/golib"
	"flag"
	"fmt"
	"os"
	"strings"
)

// GetFlags returns the values of the arguments given from user
func GetFlags() *l.NPuzzleOptions {
	sizePtr := flag.Int("n", 3, "Puzzle dimension")
	unsolvablePtr := flag.Bool("u", false, "Unsolveable puzzle (default = false)")
	iterationsPtr := flag.Int("i", 200, "Puzzle complexity")
	heuristicsPtr := flag.Int("h", 0, "Heuristics:\n\t0: Manhattan\n\t1: Hamming\n\t2: Linear Conflict")
	// filePtr := flag.String("f", "", "Input as file")

	flag.Parse()
	args := flag.Args()

	arg := strings.Join(args, "")
	file := strings.Contains(arg, ".txt")

	// if len(args) == 1 && file {
	// 	return 0, 0, false, 0
	// }

	if len(args) > 1 && file {
		fmt.Println("Error: must input one file OR flags as argument.")
		os.Exit(1)
	}

	if *sizePtr < 3 {
		flag.PrintDefaults() // replace with Print Usage
		os.Exit(1)
	}

	if sizePtr == nil {
		fmt.Println("Error: please give a board size.")
		os.Exit(1)
	}

	if *iterationsPtr < 1 {
		fmt.Println("Can't solve a puzzle in less than 1 iteration!")
		os.Exit(1)
	}

	return &l.NPuzzleOptions{
		Heuristics: *heuristicsPtr,
		SearchAlgo: 0,
		Solvable:   !*unsolvablePtr,
		Iterations: *iterationsPtr,
		Size:       *sizePtr,
	}

	// return size, solve, iterations, heuristics
}

// type NPuzzleOptions struct {
// 	heuristics int
// 	searchAlgo int
// 	solvable   bool
// 	iterations int
// 	size       int
// }
