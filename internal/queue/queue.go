package queue

import "fmt"

type Point struct {
	X int
	Y int
}

type Queue struct {
	queue []Point
}

func (q *Queue) Add(p Point) {
	if len(q.queue) == 0 {
		q.queue = make([]Point, 0)
	}
	q.queue = append(q.queue, p)
}

func (q *Queue) Get() (Point, error) {
	if len(q.queue) == 0 {
		return Point{}, fmt.Errorf("No items in queue")
	}

	p := q.queue[0]
	q.queue = q.queue[1:]
	return p, nil
}

func (q Queue) Len() int {
	return len(q.queue)
}
