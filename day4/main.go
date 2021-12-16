package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/FrankMoreno/adventcode_2021/helper"
)

type bingoBoard struct {
	squares     [][]*square
	last_column int
	last_row    int
	winner      bool
}

type square struct {
	value  int
	called bool
}

func main() {
	if len(os.Args) < 2 {
		log.Panic("NO FILE")
	}

	numbersCalled, allBoards := convert_to_boards(os.Args[1])
	part1(numbersCalled, allBoards)
	part2(numbersCalled, allBoards)
}

func part1(numbersCalled []int, allBoards []*bingoBoard) {
	for _, numberCalled := range numbersCalled {
		for _, board := range allBoards {
			if board.look_for_value(numberCalled) && board.search_for_winner() {
				println(board.calculate_answer(numberCalled))
				return
			}
		}
	}
}

func part2(numbersCalled []int, allBoards []*bingoBoard) {
	lastWinner := -1
	winningNumber := -1

	for _, numberCalled := range numbersCalled {
		for index, board := range allBoards {
			if !board.winner && board.look_for_value(numberCalled) && board.search_for_winner() {
				lastWinner = index
				winningNumber = numberCalled
			}
		}
	}

	println(allBoards[lastWinner].calculate_answer(winningNumber))
}

func convert_to_boards(fileName string) ([]int, []*bingoBoard) {
	input := helper.ReadInput(fileName, "\n\n")
	numbersCalled := []int{}
	allBoards := []*bingoBoard{}

	for _, value := range strings.Split(input[0], ",") {
		int_val, _ := strconv.Atoi(value)
		numbersCalled = append(numbersCalled, int_val)
	}

	for i := 1; i < len(input); i++ {
		currentBoard := &bingoBoard{}
		rows := strings.Split(input[i], "\n")

		for _, row := range rows {
			row_values := strings.Fields(row)
			row_values_ints := []*square{}
			for _, value := range row_values {
				int_val, _ := strconv.Atoi(value)
				row_values_ints = append(row_values_ints, &square{
					value:  int_val,
					called: false,
				})
			}

			currentBoard.squares = append(currentBoard.squares, row_values_ints)
		}

		allBoards = append(allBoards, currentBoard)
	}

	return numbersCalled, allBoards
}

func (bb *bingoBoard) look_for_value(searchValue int) bool {
	for ri, row := range bb.squares {
		for ci, column := range row {
			if column.value == searchValue {
				column.called = true
				bb.last_row = ri
				bb.last_column = ci
				return true
			}
		}
	}

	return false
}

func (bb *bingoBoard) search_for_winner() bool {
	if bb.search_for_row_winner(bb.last_row) || bb.search_for_column_winner(bb.last_column) {
		bb.winner = true
		return true
	}

	return false
}

func (bb bingoBoard) search_for_row_winner(row int) bool {
	winner := true

	for _, val := range bb.squares[row] {
		winner = winner && val.called
	}

	return winner
}

func (bb bingoBoard) search_for_column_winner(column int) bool {
	winner := true

	for _, row := range bb.squares {
		winner = winner && row[column].called
	}

	return winner
}

func (bb bingoBoard) calculate_answer(lastFound int) int {
	total := 0

	for _, row := range bb.squares {
		for _, column := range row {
			if !column.called {
				total += column.value
			}
		}
	}

	return total * lastFound
}
