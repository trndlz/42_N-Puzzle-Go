package solver

import (
	"container/heap"
	"fmt"
	g "n-puzzle/golib"
)

// Solver is the main graph search algorithm
func Solver(Puzzle []int, size int, solve bool, iterations int) {

	Solution := MakeGoal(size)

	if IsSolvable(Solution, Puzzle, size) {
		fmt.Println("\nI love you")
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
		}
		heap.Init(&pq)

		for pq.Len() > 0 && !solutionFound {
			current := heap.Pop(&pq).(*Item)
			// fmt.Println("priority", current.priority, "puzzle", current.puzzle, "move", current.move, "round", round)
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
					break
				} else if inClosedSet == true {
					// Puzzle is in the closed set
					// We do nothing
				} else {
					manhattan := g.Manhattan(childPuzzle, Solution, size)
					newPuzzle := &Item{
						priority: manhattan + current.move + 1,
						move:     current.move + 1,
						puzzle:   childPuzzle,
						parent:   current,
					}
					heap.Push(&pq, newPuzzle)
				}
				i++
			}

		}
		fmt.Println(solutionFound)
		fmt.Println(Puzzle)

	} else {
		fmt.Println("\nIt is not solvable :( try again")
	}

}
