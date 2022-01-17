package main

import (
	"fmt"
	"github.com/woodsjc/aoc_2021/input"
)

type Board struct {
	board  [][]rune
	width  int
	height int
}

func main() {
	input := input.ReadInputFile("input/day25.txt")
	//input := getTestInput()
	board := parseBoard(input)
	part1(board)
}

func part1(board Board) {
	var i int

	for i = 0; i < 1000; i++ {
		if board.Step() {
			break
		}
	}
	board.Print()
	fmt.Printf("Part 1: %d steps.\n", i+1)
}

func parseBoard(input []string) Board {
	if len(input) <= 0 {
		return Board{}
	}
	b := Board{board: make([][]rune, 0), width: len(input[0]), height: len(input)}

	for i, s := range input {
		if len(s) == 0 {
			i--
			b.height--
			continue
		}
		b.board = append(b.board, make([]rune, 0))
		for _, r := range s {
			b.board[i] = append(b.board[i], r)
		}
	}

	return b
}

func (b Board) Print() {
	fmt.Println()
	for i := range b.board {
		fmt.Printf("%s\n", string(b.board[i]))
	}
}

//horizontal first
//then vertical
func (b *Board) Step() bool {
	sameBoard := true

	//horizontal
	for i := 0; i < b.height; i++ {
		ableRollover := true
		for j := 0; j < b.width; j++ {
			if b.board[i][j] == '>' {
				if j == b.width-1 && b.board[i][0] == '.' && ableRollover {
					sameBoard = false
					b.board[i][j] = '.'
					b.board[i][0] = '>'
				} else if j < b.width-1 && b.board[i][j+1] == '.' {
					sameBoard = false
					if j == 0 {
						ableRollover = false
					}
					b.board[i][j] = '.'
					b.board[i][j+1] = '>'
					j++
				}
			}
		}
	}

	//vertical
	for j := 0; j < b.width; j++ {
		ableRollover := true
		for i := 0; i < b.height; i++ {
			if b.board[i][j] == 'v' {
				if i == b.height-1 && b.board[0][j] == '.' && ableRollover {
					sameBoard = false
					b.board[i][j] = '.'
					b.board[0][j] = 'v'
				} else if i < b.height-1 && b.board[i+1][j] == '.' {
					sameBoard = false
					if i == 0 {
						ableRollover = false
					}
					b.board[i][j] = '.'
					b.board[i+1][j] = 'v'
					i++
				}
			}
		}
	}

	return sameBoard
}

func getTestInput() []string {
	return []string{
		"v...>>.vv>",
		".vv>>.vv..",
		">>.>v>...v",
		">>v>>.>.v.",
		"v>v.vv.v..",
		">.>>..v...",
		".vv..>.>v.",
		"v.v..>>v.v",
		"....v..v.>",
	}
}
