package types

import "time"

// OutputData gathers input errors and puzzle input
type OutputData struct {
	Error          bool
	Heuristics     string
	Moves          int
	Path           [][]int
	SearchAlgo     string
	SizeComplexity int
	TimeComplexity int
	Target         []int
	Timer          time.Duration
}
