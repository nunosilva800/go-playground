package main

import "fmt"

// Given a matrix of integers, return an array with the elements ordered in
// clockwise spiral, from the outside in.

func main() {
	m1 := [][]int{
		{1},
	}
	fmt.Println("want:\t [1] got: ", spiral(m1))

	m2 := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("want:\t [1 2 4 3] got: ", spiral(m2))

	m3 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("want:\t [1 2 3 6 9 8 7 4 5]\n\tgot: ", spiral(m3))

	m4 := [][]int{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
		{12, 13, 14, 15},
	}
	fmt.Println("want:\t [1 2 3 4 7 11 15 14 13 12 8 4 5 6 10 9]\ngot:\t", spiral(m4))
}

func spiral(matrix [][]int) []int {
	row := 0       // index for row
	col := 0       // index for col
	dir := "right" // direction of cursor

	maxCol := len(matrix)
	if maxCol == 0 {
		return nil
	}

	maxRow := len(matrix[0])
	result := []int{}

	fmt.Println("dimensions:", maxCol, maxRow)

	steps := maxRow * maxCol
	ring := 0

	for steps > 0 {
		fmt.Printf("\tat:[%v,%v]; max: [%v,%v], dir: %v\n", row, col, maxRow, maxCol, dir)

		result = append(result, matrix[row][col])
		fmt.Println("\t\tfound:", matrix[row][col])

		// check if direction needs to be adjusted
		switch dir {
		case "right":
			if col == maxCol-1 {
				dir = "down"
				fmt.Println("\tnew dir: ", dir)
			}
		case "down":
			if row == maxRow-1 {
				dir = "left"
				fmt.Println("\tnew dir: ", dir)
			}
		case "left":
			if col == 0 {
				dir = "up"
				fmt.Println("\tnew dir: ", dir)
			}
		case "up":
			if row == 0 {
				dir = "right"
				fmt.Println("\tnew dir: ", dir)
			}
		}

		// move cursor according to direction
		switch dir {
		case "right":
			col += 1
		case "down":
			row += 1
		case "left":
			col -= 1
		case "up":
			row -= 1
		}

		// once a ring is complete, move into an inner one and reset state
		// ring (0), starts at [0, 0] (outer ring)
		// ring (1), starts at [1, 1]
		// ...
		if ring == row && ring == col {
			ring++
			col++
			row++
			maxCol--
			maxRow--
			dir = "right"
			fmt.Println("\tmoving into ring: #", ring)
		}

		steps--
	}

	return result
}
