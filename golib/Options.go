package golib

// NPuzzleOptions is everything you ever wanted to specify about your inputs
type NPuzzleOptions struct {
	Heuristics int
	SearchAlgo int
	Solvable   bool
	Iterations int
	Size       int
}
