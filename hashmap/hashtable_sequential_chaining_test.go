package hashtable

import "testing"

func TestNew(t *testing.T) {
	ht := NewsequentialSearchhHashTb(10)
	if ht.size != 10 {
		t.Errorf("size is not correct")
	}
}

func TestInsert(t *testing.T) {
	ht := NewsequentialSearchhHashTb(10)
	input := []int{1, 11, 21}
	for _, i := range input {
		ht.insert(i)
		if ht.contains(i) != true {
			t.Errorf("contains() did not work for input %d", i)
		}
	}
}
