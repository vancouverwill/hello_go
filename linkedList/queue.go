package linkedList

// first > > > > last

type queue struct {
	linkedList
	last *node
}

func NewQueue() *queue {
	res := new(queue)
	res.n = 0
	return res
}

func (q *queue) enqueue(v int) {
	// q.stack.push(value)
	// if q.size == 1 {
	// 	q.last = q.first
	// }
	if q.n == 0 {
		q.first = new(node)
		q.first.value = v
		q.last = q.first
	} else {
		oldLast := q.last
		q.last = new(node)
		q.last.value = v
		oldLast.next = q.last
	}
	q.n++
}

func (q *queue) dequeue() (int, error) {
	return q.removeFirst()
	// if q.size == 0 {
	// 	return 0, fmt.Errorf("queue is empty")
	// }
	// oldLast := q.last
	// result = oldLast.value
	// q.n--
}
