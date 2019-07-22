package solver

import (
	z "N-Puzzle-Go/puzzles"
	t "N-Puzzle-Go/types"
	"container/heap"
	"fmt"
)

func solutionPath(solution []int, parent *QueueItem) int {
	i := 0
	path := parent
	for path.parent != nil {
		path = path.parent
		i++
	}
	return i
}

func heuristics(puzzle []int, target []int, opt *t.NPuzzleOptions) int {
	if opt.Heuristics == 0 {
		return Manhattan(puzzle, target, opt.Size)
	} else if opt.Heuristics == 1 {
		return Hamming(puzzle, target, opt.Size)
	} else {
		return 2*LinearConflict(puzzle, target, opt.Size) + Manhattan(puzzle, target, opt.SearchAlgo)
	}
}

// Solver is the main graph search algorithm
func Solver(Puzzle []int, opt *t.NPuzzleOptions) {

	target := z.MakeGoal(opt.Size)

	if IsSolvable(target, Puzzle, opt.Size) {
		// Init closed Set
		closedSet := make(map[string]int)
		solutionFound := false
		round := 0
		// Init queue
		pq := make(PriorityQueue, 1)
		pq[0] = &QueueItem{
			priority: 0,
			move:     0,
			puzzle:   Puzzle,
			parent:   nil,
		}
		heap.Init(&pq)
		for pq.Len() > 0 && !solutionFound {
			current := heap.Pop(&pq).(*QueueItem)
			// fmt.Println("puzzle", current.puzzle, "h", current.h, "l", current.l, "m", current.m)
			closedSet[z.PuzzleToString(current.puzzle)] = 1
			children := CreateNeighbors(current.puzzle, opt.Size)

			i := 0
			round++
			for _, childPuzzle := range children {
				puzzleStr := z.PuzzleToString(childPuzzle)
				isGoal := z.CheckSliceEquality(childPuzzle, target)
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
					newPuzzle := &QueueItem{
						priority: heuristics(childPuzzle, target, opt) + current.move + 1,
						move:     current.move + 1,
						puzzle:   childPuzzle,
						parent:   current,
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
