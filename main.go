package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. Collect all command-line arguments (the 9 rows of the Sudoku)
	myargs := os.Args[1:]

	// 2. Check if we have exactly 9 rows. If not, print Error and stop.
	if len(myargs) != 9 {
		fmt.Println("Error")
		return
	}

	// 3. Create a 9x9 grid (array of runes) to store the board
	var board [9][9]rune

	// 4. Validate and parse the input arguments into the board
	for row, rowStr := range myargs {
		// Each row must be exactly 9 characters long
		if len(rowStr) != 9 {
			fmt.Println("Error")
			return
		}

		for col, char := range rowStr {
			// Check if character is valid (1-9 or '.')
			if char != '.' && (char < '1' || char > '9') {
				fmt.Println("Error")
				return
			}
			// Assign the character to the grid
			board[row][col] = char
		}
	}

	// 5. Display the original state of the board
	fmt.Println("\nUNSOLVED SUDOKU")
	PrintBoard(board) // Found in utils.go

	// 6. Interaction: Wait for the user to trigger the solver
	fmt.Println("---------------")
	fmt.Println("Press 'Enter' to solve it")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// 7. Solve the board using a Pointer (&) to modify the original grid
	// The Solve function is found in solver.go
	if Solve(&board) {
		// 8. If successful, display the solved board
		fmt.Println("SUDOKU SOLVED!")
		PrintBoard(board)
	} else {
		// 9. If the algorithm returns false, the Sudoku is unsolvable
		fmt.Println("Error")
	}

	fmt.Println("---------------")
}
