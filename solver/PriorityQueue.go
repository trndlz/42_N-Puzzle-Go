package solver

import (
	"container/heap"
)

// QueueItem represents the data injected in the priority queue
type QueueItem struct {
	index    int
	priority int
	move     int
	puzzle   []int
	parent   *QueueItem
}

// PriorityQueue type
type PriorityQueue []*QueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push enqueues an Item to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*QueueItem)
	item.index = n
	*pq = append(*pq, item)
}

// Pop returns the first element of the queue and removes it from it
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *QueueItem, puzzle []int, priority int) {
	item.puzzle = puzzle
	item.priority = priority
	heap.Fix(pq, item.index)
}
