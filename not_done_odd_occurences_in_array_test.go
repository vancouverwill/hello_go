package main

import "testing"

var inputSize = 10

func TestHashMapNew(t *testing.T) {
	// inputSize := 10
	hm := NewLPHashMap(inputSize)
	if hm.getSize() != inputSize {
		t.Errorf("New HashMap was not of expected size %d", inputSize)
	}
	if hm.getN() != 0 {
		t.Errorf("New HashMas was expected to be empty but had %d entries", hm.getN())
	}
}

func TestHashMapInsert(t *testing.T) {
	// inputSize := 10
	hm := NewLPHashMap(inputSize)
	hm.insertIntoLPHashMap(5)
	hm.insertIntoLPHashMap(54)
	if hm.getN() != 2 {
		t.Errorf("Hashmap with two entries inserted should have had shown two entries but instead it showed %d", hm.getN())
	}
}

func TestHashMapHashingFunction(t *testing.T) {
	hm := NewLPHashMap(inputSize)

	type row struct {
		input    int
		expected int
	}
	// var table []row
	table := []row{
		row{0, 0},
		row{55, 5},
		row{15, 5},
	}
	for _, r := range table {
		if hm.hash(r.input) != r.expected {
			t.Errorf("Input of %d id not get expected %d. Result was %d", r.input, r.expected, hm.hash(r.input))
		}
	}
}
