package solver

import (
	l "N-Puzzle-Go/golib"
	"fmt"
)

func checkMoves(width int, i int) (int, int, int, int) {
	up, down, left, right := 0, 0, 0, 0
	length := width * width
	// if zero is in top row
	if i < length-width {
		down = 1
	}
	// if zero is not in top row
	if i >= width && i < length {
		up = 1
	}
	// if zero is in right column
	maxMod := (length - 1) % width
	// if zero is in left column
	minMod := 0
	if i%width != minMod {
		left = 1
	}
	if i%width != maxMod {
		right = 1
	}
	return up, down, left, right
}

// CreateNeighbors returns a map of all next neighbours
func CreateNeighbors(puzzle []int, size int) map[string][]int {

	empty := l.FindIndexSlice(puzzle, 0)
	up, down, left, right := checkMoves(size, empty)
	moves := up + down + left + right
	neighbors := map[string][]int{}
	for moves > 0 {
		key := fmt.Sprintf("%d", moves)
		neighbor := make([]int, len(puzzle))
		copy(neighbor, puzzle)
		neighbors[key] = neighbor
		// use switch for up, down, left, right?
		if up == 1 {
			neighbor[empty], neighbor[empty-size] = neighbor[empty-size], neighbor[empty]
			up = 0
		} else if down == 1 {
			neighbor[empty], neighbor[empty+size] = neighbor[empty+size], neighbor[empty]
			down = 0
		} else if left == 1 {
			neighbor[empty], neighbor[empty-1] = neighbor[empty-1], neighbor[empty]
			left = 0
		} else if right == 1 {
			neighbor[empty], neighbor[empty+1] = neighbor[empty+1], neighbor[empty]
			right = 0
		}
		moves--
	}
	return neighbors
}
