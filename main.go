package main

import (
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
			board[row][col] = char
		}
	}
}
