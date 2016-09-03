package hello_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseArraySection(t *testing.T) {
    input := []int{0,1,2,3,4}
    
    reverseArraySection(input, 1,2)
    expected := []int{0,2,1,3,4}
    
    assert.Equal(t, expected, input, "successful")
}

func TestReverseWholeArray(t *testing.T) {
    input := []int{0,1,2,3,4}
    
    reverseArraySection(input, 0,len(input)-1)
    expected := []int{4,3,2,1,0}
    
    assert.Equal(t, expected, input, "successful")
}

func TestRotateArray(t *testing.T) {
    input := []int{0,1,2,3,4}
    
    RotateArrayInPlace(input, 2)
    expected := []int{3,4,0,1,2}
    
    assert.Equal(t, expected, input, "successful")
}

func TestRotateArrayNegativeAmount(t *testing.T) {
    input := []int{0,1,2,3,4}
    
    RotateArrayInPlace(input, -2)
    expected := []int{2,3,4,0,1}
    
    assert.Equal(t, expected, input, "successful")
}