package main

import "fmt"

func PrintBoard(board [9][9]rune) {
	fmt.Println("---------------")
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Printf("%c", board[row][col])
			if col < 8 {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
