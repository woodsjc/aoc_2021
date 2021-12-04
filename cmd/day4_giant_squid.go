package main

import (
	"fmt"

	"github.com/woodsjc/aoc_2021/input"
)

type Board struct {
	board  [5][5]int
	marked *[5][5]bool
}

func main() {
	var boards []Board
	numbers, rawBoards := input.Day4Input() //getTestInput()
	for _, b := range rawBoards {
		boards = append(boards, Board{board: b, marked: &[5][5]bool{}})
	}

	part1(numbers, boards)
	part2(numbers, boards)
}

func part1(numbers []int, boards []Board) {
Bingo:
	for _, n := range numbers {
		for _, b := range boards {
			b.updateBoard(n)
			if b.checkBingo() {
				fmt.Printf("Part1; Board %v won with: (%d*%d)=%d\n", b.board, n, b.sumUnMarked(), n*b.sumUnMarked())
				break Bingo
			}
		}
	}
}

func part2(numbers []int, boards []Board) {
	incompleteBoards := boards

	for _, n := range numbers {
		for i := 0; i < len(incompleteBoards); i++ {
			b := incompleteBoards[i]
			b.updateBoard(n)
			if b.checkBingo() {
				if len(incompleteBoards) == 1 {
					fmt.Printf("Part2; Board %v is last winner: (%d*%d)=%d\n", b.board, n, b.sumUnMarked(), n*b.sumUnMarked())
					return
				}

				incompleteBoards = removeBoard(incompleteBoards, i)
				i -= 1
			}
		}
	}
}

func (b *Board) updateBoard(next int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.board[i][j] == next {
				b.marked[i][j] = true
				//fmt.Printf("board[%d][%d] == %d ||| marked: %t\n", i, j, next, b.marked[i][j])
			}
		}
	}
}

func (b *Board) checkBingo() bool {
	j := 0

	//check horizontal
	for i := 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			if !b.marked[i][j] {
				break
			}
		}

		if j == 5 {
			fmt.Printf("Horizontal winner on row %d\n", i)
			return true
		}
	}

	//check vertical
	for i := 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			if !b.marked[j][i] {
				break
			}
		}

		if j == 5 {
			fmt.Printf("Vertical winner on column %d\n", i)
			return true
		}
	}

	//reread directions and diagonals don't count
	////diagonal \
	//for j = 0; j < 5; j++ {
	//	if !b.marked[j][j] {
	//		break
	//	}
	//}
	//if j == 5 {
	//	fmt.Printf("Diagonal \\ winner\n")
	//	return true
	//}

	////diagonal /
	//for j = 0; j < 5; j++ {
	//	if !b.marked[j][4-j] {
	//		break
	//	}
	//}
	//if j == 5 {
	//	fmt.Printf("Diagonal / winner\n")
	//	return true
	//}

	return false
}

func (b *Board) sumUnMarked() int {
	total := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				total += b.board[i][j]
			}
		}
	}

	return total
}

func removeBoard(boards []Board, i int) []Board {
	if len(boards) <= 1 {
		return []Board{}
	}

	boards[i] = boards[len(boards)-1]
	return boards[:len(boards)-1]
}

func getTestInput() ([]int, []Board) {
	numbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	boards := []Board{
		{
			board: [5][5]int{
				{22, 13, 17, 11, 0},
				{8, 2, 23, 4, 24},
				{21, 9, 14, 16, 7},
				{6, 10, 3, 18, 5},
				{1, 12, 20, 15, 19},
			},
			marked: &[5][5]bool{},
		},
		{
			board: [5][5]int{
				{3, 15, 0, 2, 22},
				{9, 18, 13, 17, 5},
				{19, 8, 7, 25, 23},
				{20, 11, 10, 24, 4},
				{14, 21, 16, 12, 6},
			},
			marked: &[5][5]bool{},
		},
		{
			board: [5][5]int{
				{14, 21, 17, 24, 4},
				{10, 16, 15, 9, 19},
				{18, 8, 23, 26, 20},
				{22, 11, 13, 6, 5},
				{2, 0, 12, 3, 7},
			},
			marked: &[5][5]bool{},
		},
	}

	return numbers, boards
}
