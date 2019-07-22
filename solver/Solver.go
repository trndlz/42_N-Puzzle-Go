package solver

import (
	z "N-Puzzle-Go/puzzles"
	t "N-Puzzle-Go/types"
	"container/heap"
	"fmt"
	"time"
)

func solutionPath(solution []int, parent *QueueItem) [][]int {
	path := [][]int{parent.puzzle}
	current := parent
	for current.parent != nil {
		current = current.parent
		path = append(path, current.puzzle)
	}
	return path
}

func heuristics(puzzle []int, target []int, opt *t.NPuzzleOptions) int {
	if opt.Heuristics == "MANHATTAN" {
		return Manhattan(puzzle, target, opt.Size)
	} else if opt.Heuristics == "HAMMING" {
		return Hamming(puzzle, target, opt.Size)
	} else {
		return 2*LinearConflict(puzzle, target, opt.Size) + Manhattan(puzzle, target, opt.SearchAlgo)
	}
}

func initPriorityQueue(puzzle []int) PriorityQueue {
	pq := make(PriorityQueue, 1)
	pq[0] = &QueueItem{
		priority: 0,
		move:     0,
		puzzle:   puzzle,
		parent:   nil,
	}
	heap.Init(&pq)
	return pq
}

// Solver is the main graph search algorithm
func Solver(Puzzle []int, opt *t.NPuzzleOptions) *t.OutputData {

	target := z.MakeGoal(opt.Size)
	fmt.Println(Puzzle)
	fmt.Println(opt.Size)
	if !IsSolvable(target, Puzzle, opt.Size) {
		fmt.Println(IsSolvable(target, Puzzle, opt.Size))
		return &t.OutputData{Error: true}
	}

	closedSet := make(map[string]int)
	solutionFound := false
	round := 0
	pq := initPriorityQueue(Puzzle)
	start := time.Now()
	for pq.Len() > 0 && !solutionFound && round < 10000000 {
		current := heap.Pop(&pq).(*QueueItem)
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
				solutionPath := solutionPath(childPuzzle, current)
				return &t.OutputData{
					Error:      false,
					Moves:      len(solutionPath),
					Path:       solutionPath,
					Heuristics: opt.Heuristics,
					SearchAlgo: "A_STAR",
					Timer:      time.Since(start),
				}
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
	return &t.OutputData{
		Error: true,
	}
}
