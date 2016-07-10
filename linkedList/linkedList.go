package linkedList

import "fmt"

type linkedList struct {
	first *node
	n     int
}

func (l *linkedList) Size() int {
	return l.n
}

func (l *linkedList) Contains(v int) bool {
	if l.first == nil {
		return false
	}
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

func (l *linkedList) addFirst(value int) {
	oldFirst := l.first
	n := new(node)
	n.value = value
	l.first = n
	n.next = oldFirst
	l.n++
}

//
func (l *linkedList) removeFirst() (int, error) {
	if l.Size() == 0 {
		return 0, fmt.Errorf("empty stack")
	}
	result := l.first.value
	l.first = l.first.next
	l.n--
	return result, nil
}
