package hello_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHammingDistanceEqualWords(t *testing.T) {
	expectedDist := 0
	actualDist, err := hammingDistance("will", "will")
	assert.Equal(t, expectedDist, actualDist, "hamming distance for two equal strings should be zero")
	assert.Equal(t, nil, err, "hamming distance for two equal strings should not throw error")

}

func TestHammingDifferentWords(t *testing.T) {
	expectedDist := 5
	actualDist, err := hammingDistance("james", "willy")
	assert.Equal(t, expectedDist, actualDist, "hamming distance for two equal strings should be zero")
	assert.Equal(t, nil, err, "hamming distance for two equal strings should not throw error")

}

func TestLevenshteinDistanceRecursiveEqualWords(t *testing.T) {
    expectedDist := 0
	wordA := "will"
	wordB := "will"
	actualDist := LevenshteinDistanceRecursive(wordA, len(wordA), wordB, len(wordB))
	assert.Equal(t, expectedDist, actualDist, "Levenshtein distance for two equal strings should be zero")
}


func TestLevenshteinDistanceRecursiveDifferentWords(t *testing.T) {
    expectedDist := 5
	wordA := "willy"
	wordB := "james"
	actualDist := LevenshteinDistanceRecursive(wordA, len(wordA), wordB, len(wordB))
	assert.Equal(t, expectedDist, actualDist, "Levenshtein distance for two equal strings should be zero")
}

func TestLevenshteinDistanceRecursiveWithOutLengthsDifferentWords(t *testing.T) {
    expectedDist := 5
	wordA := "willy"
	wordB := "james"
	actualDist := LevenshteinDistanceRecursiveWithOutLengths(wordA, wordB)
	assert.Equal(t, expectedDist, actualDist, "Levenshtein distance for two equal strings should be zero")
}


func TestLevenshteinDistanceDifferentWords(t *testing.T) {
    expectedDist := 3
	wordA := "kitten"
	wordB := "sitting"
	actualDist := LevenshteinDistance(wordA, wordB)
	assert.Equal(t, expectedDist, actualDist, "Levenshtein distance for two equal strings should be zero")
}
