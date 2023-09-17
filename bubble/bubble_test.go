package main

import (
	"testing"
)

func TestInputCreator(t *testing.T) {
	output := makeRandomSlice(10, 10)
	printSlice(output, 6)

	if len(output) != 10 {
		t.Fatalf("Length should be 10")
	}
	for i := 0; i < len(output); i++ {
		if output[i] < 0 || output[i] >= 10 {
			t.Fatalf("All values should be between 0 inclusive and 10 exclusive")
		}
	}
}

func TestCheckSorted(t *testing.T) {
	// arrange
	sorted := []int{2, 2, 3, 3, 7, 7, 7}
	unsorted := []int{2, 2, 1, 3, 3, 7, 7}

	// act
	checkSorted(sorted)
	checkSorted(unsorted)
}

func TestBubbleSort(t *testing.T) {
	output := makeRandomSlice(10, 10)

	bubbleSort(output)

	checkSorted(output)
}
