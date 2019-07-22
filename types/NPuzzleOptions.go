package types

// NPuzzleOptions is everything you ever wanted to specify about your inputs
type NPuzzleOptions struct {
	Heuristics string
	SearchAlgo int
	Solvable   bool
	Iterations int
	Size       int
	File       string
	Server     bool
}
