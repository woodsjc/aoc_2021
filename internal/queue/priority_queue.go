package queue

import (
	"sort"
)

type Ordered interface {
	Less(other interface{}) bool
}

type item struct {
	value    interface{}
	priority int
}

type PriorityQueue []*item

func NewPriorityQueue() *PriorityQueue {
	return new(PriorityQueue)
}

func (q *PriorityQueue) Push(i interface{}, p int) {
	*q = append(*q, &item{value: i, priority: p})
	sort.Sort(q)
}

func (q *PriorityQueue) Pop() interface{} {
	if q.Len() == 0 {
		return nil
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item.value
}

func (q *PriorityQueue) Len() int          { return len(*q) }
func (q PriorityQueue) Less(x, y int) bool { return q[x].priority < q[y].priority }
func (q PriorityQueue) Swap(x, y int)      { q[x], q[y] = q[y], q[x] }
