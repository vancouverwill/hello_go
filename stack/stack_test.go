package stack

import (
	"fmt"
	"testing"
)

func TestStackConstructor(t *testing.T) {
	s := NewStack()
	if s.size() != 0 {
		t.Errorf("new stack does not have size of 0")
	}
}

func TestStackPush(t *testing.T) {
	input := []int{0, 6, 12}
	s := NewStack()
	for _, i := range input {
		s.push(i)
	}
	if s.size() != len(input) {
		t.Errorf("new stack with %d entries does not fit size expected", len(input))
	}

}

func TestStackPopReturnsFilo(t *testing.T) {
	input := []int{2, 6, 12}
	s := NewStack()
	for _, i := range input {
		s.push(i)
	}
	for i := len(input) - 1; i >= 0; i-- {
		// for s.size() > 0 {
		temp, _ := s.pop()
		if input[i] != temp {
			t.Errorf("stack did not return elements in reverse order of being entered")
		}
		fmt.Println(temp)
	}

	if s.size() != 0 {
		t.Errorf("stack is not empty after removing all elements")
	}
}

func TestStackContains(t *testing.T) {
	input := []int{2, 6, 12}
	s := NewStack()
	for _, i := range input {
		s.push(i)
	}
	for _, i := range input {
		if s.contains(i) != true {
			t.Errorf("contains is not returning the right result for %d", i)
		}
	}
}
