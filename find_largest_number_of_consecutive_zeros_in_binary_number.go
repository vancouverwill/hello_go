package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
findLargestNumberOfZerosSurroundedByOnes we are loooking for a binary gap
 we invert binary number so we can sequence of ones
we binary shift one place to the right and then AND with current number
this removes first single 1s, then pairs, then triples,
once we are left with just zero we have know how many steps it took to reduce longest consecutive sequence of 1s to nothing
*/
func findLargestNumberOfZerosSurroundedByOnes(n int64) int {
	// fmt.Println(inverseBinaryNumbers(n))
	// current :=
	fmt.Println(n)
	current := removeTrailingZeros(n)
	fmt.Println(current)
	current = inverseBinaryNumbers(current)
	fmt.Println(current)
	count := 0
	for current != 0 {
		shiftRight := current >> 1
		current = current & shiftRight // removes all single digits
		count++
	}
	return count
}

/*
removeTrailingZeros - any binary number with trailing zeros must be even as it doesn't have 1 on the append
we can keep shifting to the right until we get to an uneven number
at this point we know there is no trailing zeros
*/
func removeTrailingZeros(n int64) int64 {
	current := n
	if current%2 != 0 {
		return current
	}
	for true {
		current = current >> 1
		if current%2 != 0 {
			break
		}

	}
	return current
}

/*
inverseBinaryNumbers reverse array zeros switch with ones
*/
func inverseBinaryNumbers(n int64) int64 {
	result := n ^ allOnesLengthInput(n)
	return result
}

func allOnesLengthInput(n int64) int64 {
	binaryNumbersLen := float64(len(strconv.FormatInt(n, 2)))
	allOnes := int64(math.Pow(2, binaryNumbersLen) - 1)
	return allOnes
}

func exampleBinaryOperations(n int64) {
	allOnes := int64(math.Pow(2, 8) - 1)
	fmt.Println("n")
	fmt.Println(n)
	fmt.Println(strconv.FormatInt(n, 2))

	fmt.Println("^n")
	fmt.Println(^n)
	fmt.Println(strconv.FormatInt(^n, 2))

	fmt.Println("n ^ allOnes = XOR")
	fmt.Println(n ^ allOnes)
	fmt.Println(strconv.FormatInt(n^allOnes, 2))

	fmt.Println("n | allOnes = OR")
	fmt.Println(n | allOnes)
	fmt.Println(strconv.FormatInt(n|allOnes, 2))

	fmt.Println("n &^ allOnes = bit clear (AND NOT)")
	fmt.Println(n &^ allOnes)
	fmt.Println(strconv.FormatInt(n&^allOnes, 2))
}
