package linkedList

import "testing"

func TestLLConstructor(t *testing.T) {
	ll := new(linkedList)
	if ll.Size() != 0 {
		t.Errorf("new linkedlist does not have size of 0")
	}
}

func TestLLNewContainsNothing(t *testing.T) {
	ll := new(linkedList)
	if ll.Contains(5) {
		t.Errorf("new linkedlist should not contain anything")
	}
}

func TestLLContainsSomething(t *testing.T) {
	ll := new(linkedList)
	ll.addFirst(5)
	if !ll.Contains(5) {
		t.Errorf(" linkedlist with value inserted should contain that")
	}
}
