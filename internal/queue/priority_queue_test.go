package queue

import "testing"

func TestNewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()
	if pq == nil {
		t.Fail()
	}
}

type TestInt int

func (ti TestInt) Less(other interface{}) bool {
	return ti < other.(TestInt)
}

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()

	pq.Push(1, 1)
	pq.Push(2, 2)
	pq.Push(3, 3)

	if pq.Len() != 3 {
		t.Fatalf("Queue length incorrect: %d", pq.Len())
	}

	for pq.Len() > 0 {
		resultInterface := pq.Pop()
		_, ok := resultInterface.(int)
		if !ok {
			t.Fatalf("Unable to recover pop item: %v", resultInterface)
		}
	}
}
