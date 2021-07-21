package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	strSudoku := "123456789123456789123456789123456789123456789123456789123456789123456789123456789"
	s1, err := parseInput(strSudoku)
	if err != nil {
		log.Fatalf("Error parsing sudoku: %s", err.Error())
	}
	if isSudokuValid(&s1) {
		printSudoku(&s1)
	} else {
		fmt.Println("Invalid Sudoku")
	}
}

func parseInput(input string) ([9][9]int, error) {
	sudoku := [9][9]int{}
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanRunes)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			scanner.Scan()
			val, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error occured while parsing sudoku", err)
				return sudoku, err
			}
			sudoku[row][col] = val
		}
	}
	return sudoku, nil
}

func hasDuplicates(counts [10]int) bool {
	// Ignore multiple "0"s - Incomplete but valid Sudokus
	for i := 1; i < 10; i++ {
		if counts[i] > 1 {
			return true
		}
	}
	return false
}

func isSudokuValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for col := 0; col < 9; col++ {
		counter := [10]int{}
		for row := 0; row < 9; row++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func printSudoku(sudoku *[9][9]int) {
	var divider string = "+-------+-------+-------+"
	fmt.Println(divider)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if col%3 == 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", sudoku[row][col])
			if col == 8 {
				fmt.Println("|")
			}
		}
		if row%3 == 2 {
			fmt.Println(divider)
		}
	}
}

// TODO: Process
func backtrack(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isSudokuValid(board) {
						if backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

// TODO: Implement me
func hasEmptyCell(*[9][9]int) bool {
	return false
}
