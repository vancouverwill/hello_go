package hello_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLargestNumberOfZerosSurroundedByOnesSimple(t *testing.T) {
	expected := 0
	actual := findLargestNumberOfZerosSurroundedByOnes(15)
	assert.Equal(t, expected, actual, "a binary number with zero zeroes did not return succesfully")

	expected = 1
	actual = findLargestNumberOfZerosSurroundedByOnes(5)
	assert.Equal(t, expected, actual, "a binary number with zero zeroes did not return succesfully. 101 should have sequence of 1")

	expected = 2
	actual = findLargestNumberOfZerosSurroundedByOnes(9)
	assert.Equal(t, expected, actual, "a binary number with zero zeroes did not return succesfully. 1001 should have sequence of 1")
}

func TestFindLargestNumberOfZerosSurroundedByOnesTrailingZeros(t *testing.T) {
	expected := 0
	actual := findLargestNumberOfZerosSurroundedByOnes(6)
	assert.Equal(t, expected, actual, "a binary number with zero zeroes did not return succesfully")
}

func TestRemoveTrailingZeros(t *testing.T) {

	expected := int64(3)
	actual := removeTrailingZeros(6)
	assert.Equal(t, expected, actual, "did not successfully switch 110 to 11")

	expected = 5
	actual = removeTrailingZeros(40)
	assert.Equal(t, expected, actual, "did not successfully switch 101000 to 101")
}
