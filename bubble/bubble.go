package main

import (
	"fmt"
	"math/rand"
)

// Make a slice containing pseudorandom numbers in [0, max).
func makeRandomSlice(numItems, max int) []int {
	randoms := make([]int, numItems)

	for i := 0; i < numItems; i++ {
		randoms[i] = rand.Intn(max)
	}
	return randoms
}

// Verify that the slice is sorted.
func checkSorted(slice []int) {
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("This slice is sorted")
}

// See if the slice is sorted.
func isSorted(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false
		}
	}
	return true
}

// Use bubble sort to sort the slice.
func bubbleSort(slice []int) {
	for !isSorted(slice) {
		for i := 1; i < len(slice); i++ {
			if slice[i-1] > slice[i] {
				hold := slice[i]
				slice[i] = slice[i-1]
				slice[i-1] = hold
			}
		}
	}
}

// Print at most numItems items.
func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Printf("%v", slice)
	} else {
		fmt.Printf("%v", slice[:numItems])
	}
}
