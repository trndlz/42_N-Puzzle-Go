package parsing

import (
	t "N-Puzzle-Go/types"
	"flag"
	"fmt"
	"os"
)

func heuristicsDescription(i int) string {
	if i == 0 {
		return "MANHATTAN"
	} else if i == 1 {
		return "HAMMING"
	} else {
		return "MIXED_LINEAR_CONFLICT_MANHATTAN"
	}
}

// GetFlags returns the values of the arguments given from user
func GetFlags() *t.NPuzzleOptions {
	sizePtr := flag.Int("n", 3, "🚀  Puzzle dimension (min 3, max 5)")
	unsolvablePtr := flag.Bool("u", false, "⛔  Unsolveable puzzle (default = false)")
	iterationsPtr := flag.Int("i", 200, "⏳  Random puzzle iterations (min 1)")
	heuristicsPtr := flag.Int("h", 0, "🌍  Heuristics:\n\t0: Manhattan\n\t1: Hamming\n\t2: Linear Conflict")
	filePtr := flag.String("f", "", "📁  Input as file")
	serverPtr := flag.Bool("s", false, "📡  Launch N-Puzzle as server")

	flag.Parse()
	if *sizePtr < 3 || *sizePtr > 5 {
		fmt.Println("🤖  \033[0;31mI cannot solve such puzzles !\033[0m")
		fmt.Println("\t- It's mininimum size must be between 3 and 5 !")
		os.Exit(1)
	}

	if *iterationsPtr < 1 {
		fmt.Println("🤖  \033[0;31mI cannot solve such puzzles !\033[0m")
		fmt.Println("\t- The random puzzle needs more than 1 iteration !")
		os.Exit(1)
	}

	return &t.NPuzzleOptions{
		Heuristics: heuristicsDescription(*heuristicsPtr),
		SearchAlgo: 0,
		Solvable:   !*unsolvablePtr,
		Iterations: *iterationsPtr,
		Size:       *sizePtr,
		File:       *filePtr,
		Server:     *serverPtr,
	}
}
