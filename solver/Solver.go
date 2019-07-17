package solver

import (
	l "N-Puzzle-Go/golib"
	"container/heap"
	"fmt"
)

func solutionPath(solution []int, parent *Item) int {
	i := 0
	path := parent
	for path.parent != nil {
		path = path.parent
		i++
	}
	return i
}

// func heuristics() int {

// }

// Solver is the main graph search algorithm
func Solver(Puzzle []int, opt *l.NPuzzleOptions) {

	Solution := l.MakeGoal(opt.Size)

	if IsSolvable(Solution, Puzzle, opt.Size) {
		// Init closed Set
		closedSet := make(map[string]int)
		solutionFound := false
		round := 0
		// Init queue
		pq := make(PriorityQueue, 1)
		pq[0] = &Item{
			priority: 0,
			move:     0,
			puzzle:   Puzzle,
			parent:   nil,
		}
		heap.Init(&pq)
		for pq.Len() > 0 && !solutionFound {
			current := heap.Pop(&pq).(*Item)
			// fmt.Println("puzzle", current.puzzle, "h", current.h, "l", current.l, "m", current.m)
			closedSet[l.PuzzleToString(current.puzzle)] = 1
			children := CreateNeighbors(current.puzzle, opt.Size)

			i := 0
			round++
			for _, childPuzzle := range children {
				puzzleStr := l.PuzzleToString(childPuzzle)
				isGoal := l.CheckSliceEquality(childPuzzle, Solution)
				_, inClosedSet := closedSet[puzzleStr]
				if isGoal == true {
					solutionFound = true
					// fmt.Println("FOUND")
					fmt.Println("FOUND IN ", solutionPath(childPuzzle, current), " MOVES")
					break
				} else if inClosedSet == true {
					// Puzzle is in the closed set
					// We do nothing
				} else {
					manhattan := l.Manhattan(childPuzzle, Solution, opt.Size)
					linearConflict := l.LinearConflict(childPuzzle, Solution, opt.Size)
					newPuzzle := &Item{
						priority: manhattan + current.move + 1,
						move:     current.move + 1,
						puzzle:   childPuzzle,
						parent:   current,
						h:        manhattan + 2*linearConflict,
						m:        manhattan,
						l:        linearConflict,
					}
					heap.Push(&pq, newPuzzle)
				}
				i++
			}
		}
	} else {
		fmt.Println("Puzzle cannot be solved. Please try again :(")
	}
}
