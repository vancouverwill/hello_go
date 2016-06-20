package linkedList

import "fmt"

type linkedList struct {
	first *node
	n     int
}

func (l *linkedList) size() int {
	return l.n
}

func (l *linkedList) contains(v int) bool {
	for current := l.first; ; current = current.next {
		if current.value == v {
			return true
		}
		if current.next == nil {
			break
		}
	}
	return false
}

//
func (s *linkedList) removeFirst() (int, error) {
	if s.size() == 0 {
		return 0, fmt.Errorf("empty stack")
	}
	result := s.first.value
	s.first = s.first.next
	s.n--
	return result, nil
}
