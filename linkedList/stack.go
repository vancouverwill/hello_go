package linkedList

type Stack struct {
	linkedList
}

// first > > > >

type node struct {
	value int
	next  *node
}

func NewStack() (s *Stack) {
	s = new(Stack)
	s.n = 0
	return s
}

// add new first
func (s *Stack) Push(value int) {
	s.addFirst(value)
}

func (s *Stack) Pop() (int, error) {
	return s.linkedList.removeFirst()
}
