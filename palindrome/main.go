package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "nuno"
	fmt.Println("input:", input, "got:", IsPalindrome(input))

	input = "Anna"
	fmt.Println("input:", input, "got:", IsPalindrome(input))

	input = "taco cat"
	fmt.Println("input:", input, "got:", IsPalindrome(input))

	input = "no lemon, no melon"
	fmt.Println("input:", input, "got:", IsPalindrome(input))
}

func IsPalindrome(input string) bool {
	// 1. remove spaces and special chars
	saneInput := ""
	for _, rn := range strings.ToLower(input) {
		if rn >= 'a' && rn <= 'z' {
			saneInput = saneInput + string(rn)
		}
	}

	// fmt.Println(saneInput)

	// 2. check if 1st char is same as last and so one until reaching the middle
	start := 0
	end := len(saneInput) - 1

	for (end - start) > 1 {
		if saneInput[start] != saneInput[end] {
			return false
		}

		start++
		end--
	}

	return true
}
