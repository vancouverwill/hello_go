package linkedList

import "testing"

func TestStackConstructor(t *testing.T) {
	s := NewStack()
	if s.Size() != 0 {
		t.Errorf("new stack does not have size of 0")
	}
}

func TestStackPush(t *testing.T) {
	input := []int{0, 6, 12}
	s := NewStack()
	for _, i := range input {
		s.Push(i)
	}
	if s.Size() != len(input) {
		t.Errorf("new stack with %d entries does not fit size expected", len(input))
	}

}

func TestStackPopReturnsFilo(t *testing.T) {
	input := []int{2, 6, 12}
	s := NewStack()
	for _, i := range input {
		s.Push(i)
	}
	for i := len(input) - 1; i >= 0; i-- {
		// for s.Size() > 0 {
		temp, _ := s.Pop()
		if input[i] != temp {
			t.Errorf("stack did not return elements in reverse order of being entered")
		}
		// fmt.Println(temp)
	}

	if s.Size() != 0 {
		t.Errorf("stack is not empty after removing all elements")
	}
}

func TestStackContains(t *testing.T) {
	input := []int{2, 6, 12}
	s := NewStack()
	for _, i := range input {
		s.Push(i)
	}
	for _, i := range input {
		if s.Contains(i) != true {
			t.Errorf("contains is not returning the right result for %d", i)
		}
	}
}
