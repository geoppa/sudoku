package main

import "fmt"

// PrintBoard: This function draws the Sudoku board in the terminal
func PrintBoard(board [9][9]rune) {
	// 1. Print a top border line to separate the board from other text
	fmt.Println("---------------")
	// 2. Start looping through each of the 9 rows
	for row := 0; row < 9; row++ {
		// 3. For every row, loop through each of the 9 columns
		for col := 0; col < 9; col++ {
			// 4. Print the character at the current position (e.g., '1', '5', or '.')
			fmt.Printf("%c", board[row][col])
			// 5. Add a space between numbers for better readability,
			// but only if it's not the last number in the row
			if col < 8 {
				fmt.Printf(" ")
			}
		}
		// 6. After finishing a row, move to the next line
		fmt.Println()
	}
}
