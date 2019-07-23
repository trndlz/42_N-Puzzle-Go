package main

import (
	p "N-Puzzle-Go/parsing"
	z "N-Puzzle-Go/puzzles"
	s "N-Puzzle-Go/solver"
	t "N-Puzzle-Go/types"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
)

// POSTData reflect the JSON data structure
type POSTData struct {
	RawPuzzle  string
	Heuristics string
	SearchAlgo string
}

func webServer(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var postData POSTData
	err := decoder.Decode(&postData)
	if err != nil {
		panic(err)
	}
	var errors []string
	var input = t.InputData{
		Puzzle: p.PuzzleStringToArray(postData.RawPuzzle, &errors),
		Errors: errors,
	}
	if len(input.Errors) > 0 {
		fmt.Println("C'EST LA HESS")
	} else {
		opt := t.NPuzzleOptions{
			Heuristics: postData.Heuristics,
			Size:       int(math.Floor(math.Sqrt(float64(len(input.Puzzle))))),
		}
		solution := s.Solver(input.Puzzle, &opt)
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(solution)
	}

}

func printSolution(puzzle []int, options *t.NPuzzleOptions) {
	solution := s.Solver(puzzle, options)
	if !solution.Error {
		if options.Verbose == true {
			z.PrintPath(solution.Path, options.Size)
		} else {
			fmt.Println("Solved Puzzle:")
			fmt.Println("")
			z.PrintBoard(puzzle, options.Size)
		}
		fmt.Println("⏰  Duration: ", solution.Timer.String())
		fmt.Println("👞  Moves: ", solution.Moves)
		fmt.Println("🏇  Size Complexity: ", solution.SizeComplexity)
		fmt.Println("⌛  Time Complexity: ", solution.TimeComplexity)
	} else {
		fmt.Println("Puzzle is not solvable")
	}
}

func printInputErrors(input *t.InputData) {
	fmt.Println("🤖  \033[0;31mI cannot read your puzzle input !\033[0m")
	for _, indError := range input.Errors {
		fmt.Println("\t- " + indError)
	}
}

func main() {
	options := p.GetFlags()
	if options.Server == true {
		fmt.Println("🚒   N-Puzzle server launched on port 5000")
		http.HandleFunc("/", webServer)
		if err := http.ListenAndServe(":5000", nil); err != nil {
			panic(err)
		}
	} else {
		var input *t.InputData
		if len(options.File) > 0 {
			input = p.FileToPuzzle(options.File)
			options.Size = int(math.Floor(math.Sqrt(float64(len(input.Puzzle)))))
		} else {
			input = z.MakeRandomBoard(options)
		}
		if len(input.Errors) > 0 {
			printInputErrors(input)
		} else {
			printSolution(input.Puzzle, options)
		}
		os.Exit(1)
	}
}
