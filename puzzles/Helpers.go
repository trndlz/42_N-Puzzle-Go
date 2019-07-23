package puzzles

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintPath(path [][]int, size int) {
	length := len(path)
	for a := length - 1; a >= 0; a-- {
		PrintBoard(path[a], size)
		if a != 0 {
			fmt.Println()
		}
	}
}

// PrintBoard prints a 2D representation of a Puzzle
func PrintBoard(slice []int, size int) {
	i := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if slice[i] == 0 {
				fmt.Printf("%v   ", "🕳️")
			} else if slice[i] < 10 {
				fmt.Printf("%v   ", slice[i])
			} else {
				fmt.Printf("%v  ", slice[i])
			}
			i++
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// PuzzleToString converts a Puzzle []int to a string
func PuzzleToString(a []int) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, ",")
}
