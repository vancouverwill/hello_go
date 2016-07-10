package linkedList

import "testing"

var input = []int{0, 6, 12}

func TestEnQueue(T *testing.T) {
	q := NewQueue()
	for _, i := range input {
		q.Enqueue(i)
	}
	if q.Size() != len(input) {
		T.Errorf("queue with %d elements added did not return correct size", len(input))
	}
}

func TestDeQueueReturnsFifo(T *testing.T) {
	q := NewQueue()
	for _, i := range input {
		q.Enqueue(i)
	}

	for _, i := range input {
		v, _ := q.Dequeue()
		if i != v {
			T.Errorf("queue did not return items back in order they were inserted")
		}
	}

	if q.Size() != 0 {
		T.Errorf("queue with %d elements added and removed did not return correct size", len(input))
	}
}
