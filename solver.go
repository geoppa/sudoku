package main

// Solve: The recursive function that solves the Sudoku using backtracking.
// It takes a pointer to the board to modify the original grid in memory.
func Solve(board *[9][9]rune) bool {
	// Search the board for an empty cell (marked with '.')
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// If an empty cell is found
			if board[row][col] == '.' {

				// Try numbers from '1' to '9'
				for num := '1'; num <= '9'; num++ {
					// Check if placing this number is safe according to Sudoku rules
					// Note: we use *board to pass the actual value to isSafe
					if isSafe(*board, row, col, num) {

						// Place the number temporarily
						board[row][col] = num

						// Recursively call Solve to see if this lead to a solution
						if Solve(board) {
							return true
						}

						// BACKTRACK: If the choice was wrong, reset the cell to '.'
						board[row][col] = '.'
					}
				}
				// If no number from 1-9 works, return false to trigger backtracking
				return false
			}
		}
	}
	// If no empty cells are left, the Sudoku is solved!
	return true
}
