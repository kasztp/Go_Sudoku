package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	challenge := read_sudoku()
	challenge_slice := [][9]int{}
	for _, row := range challenge {
		challenge_slice = append(challenge_slice, row)
	}
	fmt.Println("Original:")
	print_board(challenge_slice)
	start := time.Now()
	solved := solve(challenge_slice)
	duration := time.Since(start)
	fmt.Println("\nSolved:")
	if solved {
		print_board(challenge_slice)
	}
	fmt.Printf("\nTime to solve: %s", duration)
}

func read_sudoku() [9][9]int {
	f, err := os.Open("../9x9.csv")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := csv.NewReader(f).ReadAll()
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	temp := [9][9]int{}
	for i, row := range rows {
		for j := range row {
			text := rows[i][j]
			temp[i][j], _ = strconv.Atoi(text) // Todo: err to be used for input checking!
		}
	}

	return temp
}

func print_board(board [][9]int) {
	for i := range board {
		if i%3 == 0 && i != 0 {
			fmt.Println("- - - - - - - - - - - -")
		}
		for j := range board[0] {
			if j%3 == 0 && j != 0 {
				fmt.Printf(" | ")
			}
			if j == 8 {
				fmt.Println(board[i][j])
			} else {
				fmt.Printf("%d ", board[i][j])
			}
		}
	}
}

func solve(board [][9]int) bool {
	var find [2]int = find_empty(board)
	if find == [2]int{10, 10} {
		return true
	} else {
		row := find[0]
		col := find[1]
		for i := 1; i < 10; i++ {
			if valid(board, i, [2]int{row, col}) {
				board[row][col] = i
				if solve(board) {
					return true
				} else {
					board[row][col] = 0
				}
			}
		}
	}
	return false
}

func valid(board [][9]int, num int, pos [2]int) bool {
	// Check row
	for i := 0; i < 9; i++ {
		if (board[pos[0]][i] == num) && (pos[1] != i) {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if (board[i][pos[1]] == num) && (pos[0] != i) {
			return false
		}
	}

	// Check box
	box_x := pos[1] / 3
	box_y := pos[0] / 3

	for i := box_y * 3; i < (box_y*3 + 3); i++ {
		for j := box_x * 3; j < (box_x*3 + 3); j++ {
			if (board[i][j] == num) && ([2]int{i, j} != pos) {
				return false
			}
		}
	}
	return true
}

func find_empty(board [][9]int) [2]int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return [2]int{i, j} // row, col
			}
		}
	}
	return [2]int{10, 10}
}
