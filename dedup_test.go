package main

import (
	"testing"
)

func TestDedupCompareString(t *testing.T) {

	t.Log("TestGetBalancesByCompany")

	d := Dedup{}
	if d.compareString("red", "blue") != 1 {
		t.Error("compareString did not work")
	}

	if d.compareString("blue", "red") != -1 {
		t.Error("compareString did not work")
	}

	t.Log("testGetAccountByAccountName successful")
}

func TestDedupCompareStringEqualStrings(t *testing.T) {
	d := Dedup{}
	if d.compareString("summer", "summer") != 0 {
		t.Error("compareString did not work for equal functions")
	}
}

func TestExch(t *testing.T) {
	d := Dedup{}
	strings := []string{"red", "green", "blue", "white"}
	strings = d.exch(strings, 0, 1)
	if strings[0] != "green" || strings[1] != "red" {
		t.Error("TestExch did not work")
	}

	strings = d.exch(strings, 2, 3)
	if strings[3] == "white" {
		t.Error("TestExch did not work")
	}
}

func TestQuickSort(t *testing.T) {
	unsortedStrings := []string{"r", "x", "a", "u", "t", "m"}
	d := Dedup{}
	sortedStrings := d.quickSort(unsortedStrings, 0, len(unsortedStrings)-1)
	if d.isInOrder(sortedStrings) != true {
		t.Error("quickSort did not work")
	}
}

func TestQuickSortInOrder(t *testing.T) {
	unsortedStrings := []string{"a", "m", "r", "t", "u", "x"}
	d := Dedup{}
	sortedStrings := d.quickSort(unsortedStrings, 0, len(unsortedStrings)-1)
	if d.isInOrder(sortedStrings) != true {
		t.Error("quickSort did not work")
	}
}

func TestQuickSortAllSame(t *testing.T) {
	unsortedStrings := []string{"a", "a", "a", "a", "a", "a"}
	d := Dedup{}
	sortedStrings := d.quickSort(unsortedStrings, 0, len(unsortedStrings)-1)
	if d.isInOrder(sortedStrings) != true {
		t.Error("quickSort did not work")
	}
}

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	d := Dedup{}
	inOrderStringsWithDuplicates := []string{"a", "a", "a", "b", "c", "d", "d"}
	nonDuplicateArray := d.removeDuplicatesFromSortedArray(inOrderStringsWithDuplicates)

	t.Log("t.Log(nonDuplicateArray)")
	t.Log(inOrderStringsWithDuplicates)
	t.Log(nonDuplicateArray)

	for i := range nonDuplicateArray {
		if i == 0 {
			continue
		}
		if d.compareString(nonDuplicateArray[i-1], nonDuplicateArray[i]) == 0 {
			t.Error("nonDuplicateArray did not work")
		}
	}
}

func TestIsInOrder(t *testing.T) {
	sortedStrings := []string{"a", "b", "c", "d", "e", "f"}
	d := Dedup{}
	if d.isInOrder(sortedStrings) != true {
		t.Error("TestIsInOrder did not work")
	}
}
