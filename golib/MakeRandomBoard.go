package golib

func swapEmpty(s int, puzzle []int) {
	idx := FindIndexSlice(puzzle, 0)
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
func MakeRandomBoard(size int, solvable bool, iterations int) []int {
	puzzle := MakeGoal(size)
	for i := 0; i < iterations; i++ {
		swapEmpty(size, puzzle)
	}
	if !solvable {
		if puzzle[0] == 0 || puzzle[1] == 0 {
			puzzle[len(puzzle)-1], puzzle[len(puzzle)-2] = puzzle[len(puzzle)-2], puzzle[len(puzzle)-1]
		} else {
			puzzle[0], puzzle[1] = puzzle[1], puzzle[0]
		}
	}
	return puzzle
}
