package stack

import "testing"

func TestConstructor(t *testing.T) {
	s := NewStack()
	if s.size() != 0 {
		t.Errorf("new stack does not have size of 0")
	}
}

func TestPush(t *testing.T) {

}
