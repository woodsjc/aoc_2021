package queue

import "fmt"

type Point struct {
	X int
	Y int
}

type Queue []interface{}

func (q *Queue) Add(p interface{}) {
	if q == nil {
		q = &Queue{}
	}
	*q = append(*q, p)
}

func (q *Queue) Get() (interface{}, error) {
	if q == nil || len(*q) == 0 {
		return nil, fmt.Errorf("No items in queue")
	}

	slice := *q
	p := slice[0]
	*q = slice[1:]
	return p, nil
}

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) InQueue(p interface{}) bool {
	for _, item := range q {
		if item == p {
			return true
		}
	}
	return false
}
