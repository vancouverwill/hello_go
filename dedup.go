package main

import (
	"fmt"
	"math/rand"
	// "strconv"
	"time"
)

type Dedup struct {
}

/**
* given unsorted array of Strings, return Strings sorted with no duplicates
**/

func exampleUsageOfDedup() {
	fmt.Println("begin Dedup")
	unsortedStrings := []string{"r", "x", "a", "u", "t", "m"}

	fmt.Println(unsortedStrings)
	fmt.Println(dedup(unsortedStrings))
}

func dedup(unsortedStrings []string) []string {

	dedup := Dedup{}
	rand.Seed(time.Now().UnixNano())
	shuffle(unsortedStrings)
	sortedStrings := dedup.quickSort(unsortedStrings, 0, len(unsortedStrings)-1)
	return dedup.removeDuplicatesFromSortedArray(sortedStrings)
}

func (d *Dedup) quickSort(strings []string, lo int, hi int) []string {
	if hi <= lo {
		return strings
	}
	partionedStrings, partitionIndex := d.partition(strings, lo, hi)

	//	fmt.Println("partition done" + strconv.Itoa(partitionIndex))
	partionedStrings = d.quickSort(partionedStrings, lo, partitionIndex-1)
	partionedStrings = d.quickSort(partionedStrings, partitionIndex+1, hi)
	return strings
}

func (d *Dedup) partition(strings []string, lo int, hi int) ([]string, int) {
	var partitionIndex int = lo
	var i int = lo
	var j int = hi + 1

	// fmt.Println(strings)
	// fmt.Println("i:" + strconv.Itoa(i))
	// fmt.Println("j:" + strconv.Itoa(j))
	//	i++
	//	strings[i+=1]
	//	var comparator = strings[0]
	for true {
		// fmt.Println("begin compare i elements")
		i += 1
		for d.compareString(strings[i], strings[partitionIndex]) == -1 {
			i += 1
			if i >= (len(strings) - 1) {
				// fmt.Println("i break " + strconv.Itoa(i))
				break
			}

			// fmt.Println("i:" + strconv.Itoa(i))
		}
		// fmt.Println("begin compare j elements")
		j -= 1
		for d.compareString(strings[j], strings[partitionIndex]) == 1 {
			j -= 1
			if j <= 0 {
				// fmt.Println("j break")
				break
			}
			// fmt.Println("j:" + strconv.Itoa(j))
		}
		if i >= j {
			break
		}
		//		fmt.Println(i)
		// fmt.Println("EXCHANGE i:" + strconv.Itoa(i) + "j:" + strconv.Itoa(j))
		strings = d.exch(strings, i, j)
	}
	// fmt.Println("BREAK!! EXCHANGE partitionIndex:" + strconv.Itoa(partitionIndex) + "j:" + strconv.Itoa(j))
	strings = d.exch(strings, partitionIndex, j)

	//	fmt.Println(strings)
	return strings, j
}

func (d *Dedup) compareString(stringA string, stringB string) int {
	var n int
	if stringA < stringB {
		n = len(stringA)
	} else {
		n = len(stringB)
	}

	for i := 0; i < n; i++ {
		if stringA[i] < stringB[i] {
			//			fmt.Println("compare:" + stringA + ":" + stringB + "= -1")
			return -1
		} else if stringB[i] < stringA[i] {
			//			fmt.Println("compare:" + stringA + ":" + stringB + "= +1")
			return +1
		}
	}
	// fmt.Println("compare:" + stringA + ":" + stringB + "= 0")
	return 0
}

func (d *Dedup) exch(strings []string, indexA int, indexB int) []string {
	// fmt.Println("exch:")
	// fmt.Println(strings)
	oldStringA := strings[indexA]
	strings[indexA] = strings[indexB]
	strings[indexB] = oldStringA
	fmt.Println(strings)
	return strings
}

func shuffle(a []string) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func (d *Dedup) removeDuplicatesFromSortedArray(sortedStrings []string) []string {
	var lastString string
	newStrings := make([]string, len(sortedStrings))
	j := 0
	//	countDuplicates := 0
	for _, string := range sortedStrings {
		if string != lastString {
			newStrings[j] = string
			j++
		}
		lastString = string
	}
	// fmt.Println(j)
	newStringsTrimmed := newStrings[0:j]
	// fmt.Println(newStringsTrimmed)
	// fmt.Println(newStrings)
	return newStringsTrimmed
}

func (d *Dedup) isInOrder(a []string) bool {
	for i := range a {
		if i == 0 {
			continue
		}
		if d.compareString(a[i-1], a[i]) == 1 {
			return false
		}
	}
	return true
}
