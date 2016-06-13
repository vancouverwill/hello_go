package main

/**
A non-empty zero-indexed array A consisting of N integers is given. The array contains an odd number of elements, and each element of the array can be paired with another element that has the same value, except for one element that is left unpaired.
**/

/*
planned steps:
- create a simple hashmap from an array the same size as the input
- as all the input apart from one will be pairs hash map the same size as input should not have too many collisions
- use open addressing for collision handling
- loop through input once for each value find in hashmap then increment one
- loop through hashmap and find the one value which has a depth of one

func findOddOccurencesInArray(input []int) int {
	// comparisonArray := make([]int, len(input))
	// for _, value := range input {

	// }
	return -1
}
*/

type hashMap struct {
	size    int
	content []bucket
}

type bucket struct {
	value int
}

func (h *hashMap) insertIntoHashMap(i int) {
	
}
