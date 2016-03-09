package main

import (
	"errors"
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
    LevenshteinDistanceRecursive recursive is the easiest way to understand LevensteinDistance but the most inefficient as the path goes over 
    and calculates the same nodes epxonenential times 
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


/*
    LevenshteinDistanceRecursive recursive is the easiest way to understand LevensteinDistance but the most inefficient as the path goes over 
    and calculates the same nodes epxonenential times 
*/
func LevenshteinDistanceRecursiveWithOutLengths(s string, t string) int {
	var cost int

	/* base case: empty strings */
	if len(s) == 0 {
		return len(t)
	}
	if len(t) == 0 {
		return len(s)
	}

	/** test if last characters of the strings match */
	if s[len(s)-1] == t[len(t)-1] {
		cost = 0
	} else {
		cost = 1
	}

    min := math.MaxInt(len(s), len(t))

	deleteCharFromS := LevenshteinDistanceRecursiveWithOutLengths(s[:len(s)-1], t[:len(t)]) + 1
	if deleteCharFromS < min {
		min = deleteCharFromS
	}
	deleteCharFromT := LevenshteinDistanceRecursiveWithOutLengths(s[:len(s)], t[:len(t)-1]) + 1
	if deleteCharFromT < min {
		min = deleteCharFromT
	}
	deleteCharFromBoth := LevenshteinDistanceRecursiveWithOutLengths(s[:len(s)-1], t[:len(t)-1]) + cost
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

/*
LevenshteinDistance ... is the minimum number of single character edits 
(substitution, insertion, deletion) needed to get from one string to another

Wagner–Fischer distance gives the same result and seems to be invented at the same time
*/
func LevenshteinDistance(s string, t string) int {
  // for all i and j, d[i,j] will hold the Levenshtein distance between
  // the first i characters of s and the first j characters of t;
  // note that d has (m+1)*(n+1) values
   var d [][]int
   d = make([][]int, len(s) + 1)
   for i:= range d {
       d[i] = make([]int, len(t) + 1)
   }
   
   // source prefixes can be transformed into empty string by
   // dropping all characters
   for i := range(d) {
       d[i][0] = i
   }
   
   // target prefixes can be reached from empty source prefix
  // by inserting every character
  for j := range(d[0]) {
       d[0][j] = j
   }
   
   for i := range(d) {
       if i == 0 {continue}
       for j := range(d[0]) {
           if j == 0 {continue}
           
           var cost int
           if s[i-1] == t[j-1] {
               cost = 0
           } else {
               cost = 1
           }
           deletionCost := d[i-1][j] + 1
           insertionCost := d[i][j-1] + 1
           minDelOrIns := math.MinInt(deletionCost, insertionCost)
           substituionCost := d[i-1][j-1] + cost
           d[i][j] = math.MinInt(minDelOrIns, substituionCost)
           
       }
   }
   
   return d[len(s)][len(t)]
}
