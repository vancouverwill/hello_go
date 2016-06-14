package main

/**
A non-empty zero-indexed array A consisting of N integers is given. The array contains an odd number of elements, and each element of the array can be paired with another element that has the same value, except for one element that is left unpaired.
**/

/**
planned steps:
- create a simple hashmap from an array the same size as the input
- as all the input apart from one will be pairs hash map the same size as input should not have too many collisions
- use open addressing (also called linear probing) for collision handling
- loop through input once for each value find in hashmap then increment one
- loop through hashmap and find the one value which has a depth of one

func findOddOccurencesInArray(input []int) int {
	// comparisonArray := make([]int, len(input))
	// for _, value := range input {

	// }
	return -1
}
**/

type linearProinghHashMap struct {
	size    int
	N       int
	content []bucket
}

type bucket struct {
	value int
}

func NewLPHashMap(size int) *linearProinghHashMap {
	h := new(linearProinghHashMap)
	h.size = size
	h.N = 0
	h.content = make([]bucket, size)
	return h
}

func (h *linearProinghHashMap) insertIntoLPHashMap(i int) {
	k := h.hash(i)
	for {
		if h.content[k].value == 0 {
			break
		}
		k++
	}
	h.content[k] = bucket{i}
	h.N++
}

func (h linearProinghHashMap) getSize() int {
	return h.size
}

func (h linearProinghHashMap) getN() int {
	return h.N
}

func (h linearProinghHashMap) hash(i int) int {
	return i % h.size
}
