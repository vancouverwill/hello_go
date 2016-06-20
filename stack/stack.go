package stack

import "fmt"

type stack struct {
	first *node
	n  int
}

type node struct {
	value int
	next  *node
}

func NewStack() (s stack) {
    s := new(stack)
    s.n = 0
    return s
}

func (s *stack) push(value int) {
	oldFirst := s.first
	n := new(node)
	n.value = value
	s.first = n
	n.next = oldFirst
	s.n++
}

func (s *stack) pop() (int, error) {
	if s.size == 0 {
		return 0, fmt.Errorf("empty stack")
	}
	result := s.first.value
	s.first = s.first.next
	s.n--
	return result, nil
}

func (s *stack) size() int {
    return s.n
}

