package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateOriginalSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, n)
	for i := range slice {
		slice[i] = rand.Intn(100) // [0; 99]
	}
	return slice
}

func sliceExample(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func addElements(slice []int, elem int) []int {
	return append(slice, elem)
}

func copySlice(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	return result
}

func removeElement(slice []int, idx int) []int {
	if idx < 0 || idx >= len(slice) {
		return slice
	}
	return append(slice[:idx], slice[idx+1:]...)
}

func main() {
	originalSlice := generateOriginalSlice(10)
	fmt.Println("Original slice:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("Even:", evenSlice)

	addedSlice := addElements(originalSlice, 42)
	fmt.Println("Added 42:", addedSlice)

	copiedSlice := copySlice(originalSlice)
	originalSlice[0] = 999 // изменяем оригинал
	fmt.Println("Copied slice (should not change):", copiedSlice)
	// could add input index
	// var index int
	// fmt.Scan(&index)
	// after that give in func - removerElement(originalSlice, index)
	removedSlice := removeElement(originalSlice, 2)
	fmt.Println("rm [2]:", removedSlice)
}
