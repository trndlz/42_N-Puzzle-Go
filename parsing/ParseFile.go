package parsing

import (
	t "N-Puzzle-Go/types"
	"io/ioutil"
	"strconv"
	"strings"
)

func containsSize(len int) bool {
	sizes := [3]int{9, 16, 25}
	for _, possible := range sizes {
		if possible == len {
			return true
		}
	}
	return false
}

// FileToPuzzle reads a puzzle from a file
func FileToPuzzle(path string) *t.InputData {
	var errors []string
	data, err := ioutil.ReadFile(path)
	if err != nil {
		errors = append(errors, "File reading error: "+err.Error())
	}
	puzzle := PuzzleStringToArray(string(data), &errors)
	return &t.InputData{
		Puzzle: puzzle,
		Errors: errors,
	}
}

// PuzzleStringToArray converts a puzzle as string to an array
func PuzzleStringToArray(puzzleStr string, errors *[]string) []int {
	var puzzle []int
	lines := strings.Split(puzzleStr, "\n")
	for _, line := range lines {
		if len(line) > 1 && line[0] != '#' {
			numbersStr := strings.Fields(line)
			for _, numberStr := range numbersStr {
				number, atoiErr := strconv.Atoi(numberStr)
				if atoiErr == nil {
					puzzle = append(puzzle, number)
				} else {
					*errors = append(*errors, "Syntax error: "+atoiErr.Error())
				}
			}
		}
	}
	if !containsSize(len(puzzle)) {
		*errors = append(*errors, "Input puzzle is not square !")
	}
	return puzzle
}
