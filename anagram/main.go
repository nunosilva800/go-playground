package main

import "fmt"

func main() {
	inputA := "anna"
	inputB := "nana"
	fmt.Println("inputs:", inputA, inputB, "got:", IsAnagram(inputA, inputB))

	inputA = "anna"
	inputB = "nuno"
	fmt.Println("inputs:", inputA, inputB, "got:", IsAnagram(inputA, inputB))

	inputA = "elvis"
	inputB = "lives"
	fmt.Println("inputs:", inputA, inputB, "got:", IsAnagram(inputA, inputB))

	inputA = ""
	inputB = ""
	fmt.Println("inputs:", inputA, inputB, "got:", IsAnagram(inputA, inputB))

	inputA = "a"
	inputB = "a"
	fmt.Println("inputs:", inputA, inputB, "got:", IsAnagram(inputA, inputB))

	inputA = "aa"
	inputB = "aaa"
	fmt.Println("inputs:", inputA, inputB, "got:", IsAnagram(inputA, inputB))
}

func IsAnagram(strA, strB string) bool {
	charsA := make(map[rune]int)
	for _, rn := range strA {
		charsA[rn]++
	}

	charsB := make(map[rune]int)
	for _, rn := range strB {
		charsB[rn]++
	}

	if len(charsA) != len(charsB) {
		return false
	}

	for rn, countA := range charsA {
		countB, ok := charsB[rn]
		if !ok {
			return false
		}
		if countA != countB {
			return false
		}

	}

	return true
}
