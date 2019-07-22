package puzzles

import (
	l "N-Puzzle-Go/golib"
	t "N-Puzzle-Go/types"
	"math/rand"
	"time"
)

// RandomChoice returns a random index of an array
func RandomChoice(arr []int) int {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(len(arr))
	return arr[a]
}

func swapEmpty(s int, puzzle []int) {
	idx := l.FindIndexSlice(puzzle, 0)
	var poss []int
	if idx%s > 0 {
		poss = append(poss, idx-1)
	}
	if idx%s < s-1 {
		poss = append(poss, idx+1)
	}
	if idx/s > 0 {
		poss = append(poss, idx-s)
	}
	if idx/s < s-1 {
		poss = append(poss, idx+s)
	}
	swi := RandomChoice(poss)
	puzzle[idx] = puzzle[swi]
	puzzle[swi] = 0
}

// MakeRandomBoard creates a solvable or not randon board, with X iterations different from target
// func MakeRandomBoard(size int, solvable bool, iterations int) []int {
func MakeRandomBoard(opt *t.NPuzzleOptions) *t.InputData {
	puzzle := MakeGoal(opt.Size)
	var errors []string
	for i := 0; i < opt.Iterations; i++ {
		swapEmpty(opt.Size, puzzle)
	}
	if !opt.Solvable {
		if puzzle[0] == 0 || puzzle[1] == 0 {
			puzzle[len(puzzle)-1], puzzle[len(puzzle)-2] = puzzle[len(puzzle)-2], puzzle[len(puzzle)-1]
		} else {
			puzzle[0], puzzle[1] = puzzle[1], puzzle[0]
		}
	}
	return &t.InputData{
		Puzzle: puzzle,
		Errors: errors,
	}
}
