package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	challenge := readSudoku()
	print_board(challenge)
}

func readSudoku() [][]string {
	f, err := os.Open("../9x9.csv")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := csv.NewReader(f).ReadAll()
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func print_board(board [][]string) {
	for i, _ := range board {
		if i%3 == 0 && !(i == 0) {
			fmt.Println("- - - - - - - - - - - -")
		}
		for j, _ := range board[0] {
			if j%3 == 0 && j != 0 {
				fmt.Printf(" | ")
			}
			if j == 8 {
				fmt.Println(board[i][j])
			} else {
				fmt.Printf("%s ", board[i][j])
			}
		}
	}
}
