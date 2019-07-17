package solver

import (
	"container/heap"
	"fmt"
	g "n-puzzle/golib"
)

func solutionPath(solution []int, parent *Item) {
	fmt.Println("SOLUTION")
	fmt.Println(solution)
	path := parent
	fmt.Println(path.puzzle)
	for path.parent != nil {
		path = path.parent
		fmt.Println(path.puzzle)
	}
}

// Solver is the main graph search algorithm
func Solver(Puzzle []int, size int, solve bool, iterations int) {

	Solution := MakeGoal(size)

	if IsSolvable(Solution, Puzzle, size) {
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
		for pq.Len() > 0 && !solutionFound && round < 10 {
			current := heap.Pop(&pq).(*Item)
			fmt.Println("puzzle", current.puzzle, "h", current.h, "l", current.l, "m", current.m)
			closedSet[g.PuzzleToString(current.puzzle)] = 1
			children := CreateNeighbors(current.puzzle, size)

			i := 0
			round++
			for _, childPuzzle := range children {
				puzzleStr := g.PuzzleToString(childPuzzle)
				isGoal := g.CheckSliceEquality(childPuzzle, Solution)
				_, inClosedSet := closedSet[puzzleStr]
				if isGoal == true {
					solutionFound = true
					solutionPath(childPuzzle, current)
					break
				} else if inClosedSet == true {
					// Puzzle is in the closed set
					// We do nothing
				} else {
					manhattan := g.Manhattan(childPuzzle, Solution, size)
					linearConflict := g.LinearConflict(childPuzzle, Solution, size)
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
		fmt.Println("\nIt is not solvable :( try again")
	}

}
