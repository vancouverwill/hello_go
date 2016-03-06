package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetJSON(t *testing.T) {
    jsonplaceholderAddress := "http://jsonplaceholder.typicode.com/comments"
    var myComments []comment
    getJson(jsonplaceholderAddress, &myComments)
    
	// expectedDist := 0
	// actualDist, err := hammingDistance("will", "will")
	assert.True(t, len(myComments) > 0, "test api returned more than zero comments")
	// assert.Equal(t, nil, err, "hamming distance for two equal strings should not throw error")
// assert.t
}