package solver

import (
	z "N-Puzzle-Go/puzzles"
	t "N-Puzzle-Go/types"
	"container/heap"
	"time"
)

func solutionPath(solution []int, parent *QueueItem) [][]int {
	path := [][]int{solution, parent.puzzle}
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
		m := Manhattan(puzzle, target, opt.Size)
		l := LinearConflict(puzzle, target, opt.Size)
		return 2*l + m
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
	if !IsSolvable(target, Puzzle, opt.Size) {
		return &t.OutputData{Error: true}
	}

	closedSet := make(map[string]int)
	solutionFound := false
	round := 0
	pq := initPriorityQueue(Puzzle)
	start := time.Now()
	timeComplexity := 1
	sizeComplexity := 0
	for pq.Len() > 0 && !solutionFound && round < 10000000 {
		current := heap.Pop(&pq).(*QueueItem)
		closedSet[z.PuzzleToString(current.puzzle)] = 1
		children := CreateNeighbors(current.puzzle, opt.Size)
		round++
		openQueueLength := pq.Len()
		if openQueueLength > sizeComplexity {
			sizeComplexity = openQueueLength
		}
		for _, childPuzzle := range children {
			puzzleStr := z.PuzzleToString(childPuzzle)
			isGoal := z.CheckSliceEquality(childPuzzle, target)
			_, inClosedSet := closedSet[puzzleStr]
			if isGoal == true {
				solutionFound = true
				solutionPath := solutionPath(childPuzzle, current)
				return &t.OutputData{
					Error:          false,
					Moves:          len(solutionPath),
					Path:           solutionPath,
					Heuristics:     opt.Heuristics,
					SearchAlgo:     "A_STAR",
					Timer:          time.Since(start),
					SizeComplexity: sizeComplexity,
					TimeComplexity: timeComplexity,
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
				timeComplexity++
				heap.Push(&pq, newPuzzle)
			}
		}
	}
	return &t.OutputData{
		Error: true,
	}
}
