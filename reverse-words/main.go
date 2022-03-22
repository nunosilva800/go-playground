package main

import "fmt"

func main() {
	input := "a"
	fmt.Printf("input: [%s], want: [a], got: [%s]\n", input, reverseWords(input))

	input = "a b c"
	fmt.Printf("input: [%s], want: [c b a], got: [%s]\n", input, reverseWords(input))

	input = "123 456 789"
	fmt.Printf("input: [%s], want: [789 456 123], got: [%s]\n", input, reverseWords(input))

	input = "the quick brown fox jumps"
	fmt.Printf("input: [%s], want: [jumps fox brown quick the], got: [%s]\n", input, reverseWords(input))
}

func reverseWords(s string) string {
	// split string into words and add them to a slice
	words := []string{}
	var wordStart, wordEnd int
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			words = append(words, s[wordStart:wordEnd])
			wordStart = wordEnd + 1
			wordEnd = wordStart
			continue
		}
		wordEnd++
	}
	// add the last word
	words = append(words, s[wordStart:wordEnd])

	// go through that slice in reverse order and concat into new string
	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		result = result + words[i]
		if i != 0 {
			result += " "
		}
	}

	return result
}
