package main

import (
    // "math/big"
    // "fmt"
	"errors"
	// "math"
    "github.com/pkg/math"
)

func hammingDistance(s1 string, s2 string) (dist int, err error) {
	err = nil
	dist = 0

	if len(s1) != len(s2) {
		err := errors.New("string lengths don't match")
		return dist, err
	}

	dist = 0
	for i := 0; i < len(s1); i++ {
		if int(s1[i]) != int(s2[i]) {
			dist++
		}
	}
	return dist, err
}

/*
    LevenshteinDistanceRecursive recursive is the easiest way to understand LevensteinDistance but the most inefficient as the path goes over and calculates the same nodes epxonenential times 
*/
func LevenshteinDistanceRecursive(s string, lenS int, t string, lenT int) int {
	var cost int

	/* base case: empty strings */
	if lenS == 0 {
		return lenT
	}
	if lenT == 0 {
		return lenS
	}

	/** test if last characters of the strings match */
	if s[lenS-1] == t[lenT-1] {
		cost = 0
	} else {
		cost = 1
	}

    min := math.MaxInt(lenS, lenT)

	deleteCharFromS := LevenshteinDistanceRecursive(s, lenS-1, t, lenT) + 1
	if deleteCharFromS < min {
		min = deleteCharFromS
	}
	deleteCharFromT := LevenshteinDistanceRecursive(s, lenS, t, lenT-1) + 1
	if deleteCharFromT < min {
		min = deleteCharFromT
	}
	deleteCharFromBoth := LevenshteinDistanceRecursive(s, lenS-1, t, lenT-1) + cost
	if deleteCharFromBoth < min {
		min = deleteCharFromBoth
	}
	return min
}


func levenshteinDistanceRecursiveExample() {
    wordA := "willy"
	wordB := "james"
	LevenshteinDistanceRecursive(wordA, len(wordA), wordB, len(wordB))
}