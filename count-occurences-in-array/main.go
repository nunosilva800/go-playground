package main

import (
	"log"
)

func main() {
	input := []int{1, 2, 3, 4, 5}
	result := Occurences(input, 3)
	if result != 1 {
		log.Fatalf("a: expected 1, got %v", result)
	}

	input = []int{1, 2, 3, 4, 5}
	result = Occurences(input, 6)
	if result != 0 {
		log.Fatalf("b: expected 0, got %v", result)
	}

	input = []int{1, 2, 2, 2, 3}
	result = Occurences(input, 2)
	if result != 3 {
		log.Fatalf("c: expected 3, got %v", result)
	}
}

// Given a sorted array arr[] and a number x, counts the occurrences of x in arr[]. (O(Log(N)))
//
// 1. look for the index of the 1st occurence
// 2. look for the index of the 2nd occurence
// result is the difference + 1
func Occurences(arr []int, x int) int {
	len := len(arr)
	if len == 0 {
		return 0
	}

	if len == 1 {
		if arr[0] == x {
			return 1
		}
		return 0
	}

	first := firstIdx(arr, x, 0, len-1)
	if first == -1 {
		return 0
	}
	last := lastIdx(arr, x, 0, len-1)

	return 1 + (last - first)
}

func firstIdx(arr []int, x int, startIdx int, endIdx int) int {
	if startIdx == endIdx {
		if x == arr[startIdx] {
			return startIdx
		}
		return -1
	}

	middleIdx := startIdx + ((endIdx - startIdx) / 2)
	if x <= arr[middleIdx] {
		return firstIdx(arr, x, startIdx, middleIdx)
	}

	return firstIdx(arr, x, middleIdx+1, endIdx)
}

func lastIdx(arr []int, x, startIdx, endIdx int) int {
	if startIdx == endIdx {
		if x == arr[startIdx] {
			return startIdx
		}
		return -1
	}

	middleIdx := startIdx + ((endIdx - startIdx) / 2)
	if x >= arr[middleIdx+1] {
		return lastIdx(arr, x, middleIdx+1, endIdx)
	}

	return lastIdx(arr, x, startIdx, middleIdx)
}
