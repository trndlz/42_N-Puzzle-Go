package golib

func indexToCoordinates(index int, size int) (int, int) {
	x := index % size
	y := index / size
	return x, y
}

func absInt(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func valueCoordinates(board []int, val int, size int) (int, int) {
	index := FindIndexSlice(board, val)
	return indexToCoordinates(index, size)
}

func manhattanDistance(val int, indexBoard int, size int, target []int) int {
	indexTarget := FindIndexSlice(target, val)
	xT, yT := indexToCoordinates(indexTarget, size)
	xC, yC := indexToCoordinates(indexBoard, size)
	return (absInt(xT-xC) + absInt(yT-yC))
}

// LinearConflict returns the sum of linear conflicts
// Two tiles ‘a’ and ‘b’ are in a linear conflict if they are in the same row or column,
// also their goal positions are in the same row or column and the goal position of one
// of the tiles is blocked by the other tile in that row.
func LinearConflict(board []int, target []int, s int) int {
	conflicts := 0
	values := s * s
	for i := 1; i < values-1; i++ {
		for j := 2; j < values; j++ {
			currIx, currIy := valueCoordinates(board, i, s)
			currJx, currJy := valueCoordinates(board, j, s)
			targetIx, targetIy := valueCoordinates(target, i, s)
			targetJx, targetJy := valueCoordinates(target, j, s)
			if currIx == currJx && targetIx == targetJx {
				if (currIy < currJy && targetIy > targetJy) || (currIy > currJy && targetIy < targetJy) {
					conflicts++
				}
			}
			if currIy == currJy && targetIy == targetJy {
				if (currIx < currJx && targetIx > targetJx) || (currIx > currJx && targetIx < targetJx) {
					conflicts++
				}
			}
		}
	}
	return conflicts
}

// Hamming returns the sum of misplaced tiles.
func Hamming(board []int, target []int, s int) int {
	length := s * s
	hamming := 0
	for i := 0; i < length; i++ {
		if board[i] != target[i] && board[i] != 0 {
			hamming++
		}
	}
	return hamming
}

// Manhattan returns the sum of manhattan distances
func Manhattan(board []int, target []int, s int) int {
	length := s * s
	manhattan := 0
	for i := 0; i < length; i++ {
		if board[i] != target[i] && board[i] != 0 {
			manhattan += manhattanDistance(board[i], i, s, target)
		}
	}
	return manhattan
}
