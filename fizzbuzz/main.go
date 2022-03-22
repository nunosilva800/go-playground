package main

import "fmt"

func main() {
	fizzbuzz(20)
}

func fizzbuzz(target int) {
	for curr := 0; curr < target; curr++ {
		if curr%3 == 0 && curr%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}

		if curr%3 == 0 {
			fmt.Println("Fizz")
			continue
		}

		if curr%5 == 0 {
			fmt.Println("Buzz")
			continue
		}

		fmt.Println(curr)
	}
}
