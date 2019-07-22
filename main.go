package main

import (
	p "N-Puzzle-Go/parsing"
	z "N-Puzzle-Go/puzzles"
	s "N-Puzzle-Go/solver"
	t "N-Puzzle-Go/types"
	"encoding/json"
	"fmt"
	"log"
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
	var t POSTData
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t)
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
			input = p.ParseFile(options.File)
		} else {
			input = z.MakeRandomBoard(options)
		}
		if len(input.Errors) > 0 {
			fmt.Println("🤖  \033[0;31mI cannot read your puzzle input !\033[0m")
			for _, indError := range input.Errors {
				fmt.Println("\t- " + indError)
			}
		} else {
			s.Solver(input.Puzzle, options)
		}
		os.Exit(1)
	}
}
