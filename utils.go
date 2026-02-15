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

// IsSafe: Scans the 3x3 grid and checks if num is in the grid
func isSafe(board [9][9]rune, row, col int, num rune) bool {
	// Loop that checks if num is in rows
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}
	// Loop that checks if num is in cols
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}
	// Find the top right corner of the 3x3 grid
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	// Loop that checks if num is in the 3x3 grid
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}
