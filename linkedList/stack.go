package linkedList

type stack struct {
	linkedList
}

// first > > > >

type node struct {
	value int
	next  *node
}

func NewStack() (s *stack) {
	s = new(stack)
	s.n = 0
	return s
}

// add new first
func (s *stack) push(value int) {
	oldFirst := s.first
	n := new(node)
	n.value = value
	s.first = n
	n.next = oldFirst
	s.n++
}

func (s *stack) pop() (int, error) {
	return s.linkedList.removeFirst()
}
