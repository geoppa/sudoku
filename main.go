package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	myargs := os.Args[1:]
	if len(myargs) != 9 {
		dif := 9 - len(myargs)
		if dif > 0 {
			print("ERROR, less arguments by: ", dif)
		} else if dif < 0 {
			print("ERROR, more arguments by: ", -dif)
		}
		return
	}
	var board [9][9]rune
	for row, rowind := range myargs {
		if len(rowind) != 9 {
			dif2 := 9 - len(rowind)
			if dif2 > 0 {
				print("ERROR in line ", row, " there are ", dif2, " chars less")
			} else if dif2 < 0 {
				print("ERROR in line ", row, " there are ", -dif2, " chars more")
			}
			return
		}
		for col, char := range rowind {
			if char != '.' && (char < '1' || char > '9') {
				print("ERROR, not valid number in column ", col+1)
				return
			}
			board[row][col] = char
		}
	}
	print("\nUNSOLVED SUDOKU\n")
	print("---------------\n")
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Printf("%c", board[row][col])
			if col < 8 {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

	print("---------------\n")
	print("Press 'Enter' to solve it")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
