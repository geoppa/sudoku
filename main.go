package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. Takes everything typed after "go run ." in terminal.
	myargs := os.Args[1:]

	// 2. Checks if we don't have exactly 9 rows, shows an error and stops.
	if len(myargs) != 9 {
		dif := 9 - len(myargs)
		if dif > 0 {
			fmt.Print("ERROR: You missed ", dif, " rows.")
		} else if dif < 0 {
			fmt.Print("ERROR: You gave ", -dif, " extra rows.")
		}
		return
	}

	// 3. Creates an empty 9x9 board to save our Sudoku
	var board [9][9]rune

	// 4. Starts reading the rows one by one
	for row, rowind := range myargs {
		// Checks if each row is exactly 9 characters long
		if len(rowind) != 9 {
			dif2 := 9 - len(rowind)
			if dif2 > 0 {
				fmt.Print("ERROR: Row ", row, " is too short by ", dif2, " chars.")
			} else if dif2 < 0 {
				fmt.Print("ERROR: Row ", row, " is too long by ", -dif2, " chars.")
			}
			return
		}

		// 5. Looks at every single cell in the row
		for col, char := range rowind {
			// If it finds anything that isn't a number (1-9) or a dot stops everything
			if char != '.' && (char < '1' || char > '9') {
				fmt.Print("ERROR: Found a weird character at column ", col+1)
				return
			}
			// If everything is ok, put the character into our board
			board[row][col] = char
		}
	}

	// 6. Shows the Sudoku board as it was at the start
	fmt.Print("\nUNSOLVED SUDOKU\n")
	PrintBoard(board) // This calls the function from utils.go

	// 7. Interaction: Choose to solve or exit
	fmt.Print("---------------\n")
	fmt.Print("Press 'Enter' to solve or type 'Q' and Enter to exit: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// If the user typed 'q' before Enter, we exit
	if len(input) > 0 && input[0] == 'q' {
		fmt.Println("Exiting program...")
		return
	}

	// 8. Call the Solver and provide the board address (&)
	// If Solve returns true, the board is now filled with the solution
	if Solve(&board) {
		fmt.Print("SUDOKU SOLVED!\n")
		// 9. Show the final result
		PrintBoard(board)
	} else {
		// If the algorithm returns false, it means no solution exists
		fmt.Print("ERROR: This Sudoku cannot be solved.\n")
	}

	fmt.Print("---------------\n")
}
