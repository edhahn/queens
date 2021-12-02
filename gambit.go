package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/**
* Gets the desired board size from the first command line argument.
* Defaults to 8.
 */
func getBoardSize() int {
	if len(os.Args) > 1 {
		size, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		return size
	} else {
		return 8
	}
}

/**
* Prints the board with the position of the queens that have been placed
 */
func boardString(board [][]byte) string {
	boardStr := ""
	for _, row := range board {
		rowStr := "|"
		for _, col := range row {
			if col == 1 {
				rowStr += "x|"
			} else {
				rowStr += " |"
			}
		}
		fmt.Println(rowStr)
	}
	return boardStr
}

/**
* Removes an element from a slice in the most efficient way possible when maintain order does not matter
 */
func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

/**
* Given a set of remaning free spaces and the column that a queen has been placed in the first remaining row,
* this function returns all the remaining spaces that are not threatened
 */
func removeThreats(freeSpaces [][]int, col int) [][]int {
	newFreeSpaces := [][]int{}
	if len(freeSpaces) <= 1 {
		return newFreeSpaces
	}
	for y := 1; y < len(freeSpaces); y++ {
		row := freeSpaces[y]
		newRow := []int{}
		for x, val := range row {
			if val == col || val == col+y || val == col-y {
				continue
			}
			newRow = append(newRow, row[x])
		}
		newFreeSpaces = append(newFreeSpaces, newRow)
	}
	return newFreeSpaces
}

/*
* Attempts to solve where the next queen can be placed
 */
func solve(board [][]byte, freeSpaces [][]int, row int, queens int, attempts int) bool {
	attemptsRemaining := attempts
	fmt.Printf("\rQueens remaining %d | Attempts remaining: %3d ", queens, attemptsRemaining)
	if queens == 0 || len(freeSpaces) <= 0 {
		return false
	}
	if len(freeSpaces[0]) <= 0 {
		return true
	}
	unsolved := true
	for unsolved && attemptsRemaining > 0 {
		pos := rand.Intn(len(freeSpaces[0]))
		col := freeSpaces[0][pos]
		if !solve(board, removeThreats(freeSpaces, col), row+1, queens-1, attemptsRemaining) {
			unsolved = false
			board[row][col] = 1
		}
		attemptsRemaining--
	}
	return unsolved
}

/*
* Initializes a board and free spaces based on a given board size
 */
func initializeBoard(size int) ([][]byte, [][]int) {
	board := [][]byte{}
	freeSpaces := [][]int{}

	for y := 0; y < size; y++ {
		boardRow := []byte{}
		freeSpacesRow := []int{}
		for x := 0; x < size; x++ {
			boardRow = append(boardRow, 0)
			freeSpacesRow = append(freeSpacesRow, x)
		}
		board = append(board, boardRow)
		freeSpaces = append(freeSpaces, freeSpacesRow)
	}

	return board, freeSpaces
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Solving the Queen's Gambit...")

	size := getBoardSize()

	queens := size

	board, freeSpaces := initializeBoard(size)

	if !solve(board, freeSpaces, 0, queens, size) {
		fmt.Println("\nSolved :)")
		fmt.Println(boardString(board))
	} else {
		fmt.Println("\nFailed to solve :(")
	}
}
