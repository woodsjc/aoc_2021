package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/woodsjc/aoc_2021/input"
)

type Point struct {
	x int
	y int
}

func main() {
	input := input.ReadInputFile("input/day13.txt")
	//input = getTestInput()
	points, folds := parseInput(input)

	part1(points, folds)
	part2(points, folds)
}

func part1(points, folds []Point) {
	fmt.Printf("Folds: %v\n", folds)

	points = applyFold(points, folds[0])
	fmt.Printf("Part 1: %d total points.\n", len(points))
}

func part2(points, folds []Point) {
	for _, f := range folds {
		points = applyFold(points, f)
	}
	fmt.Printf("Part 2: %d total points.\n", len(points))
	drawPoints(points)
}

func applyFold(points []Point, fold Point) []Point {
	remaining := make([]Point, 0)

	for _, p := range points {
		newP := p
		if fold.x != 0 && p.x > fold.x {
			if 2*fold.x < p.x {
				fmt.Printf("Point would go negative after fold: point-%v, fold-x=%d.", p, fold.x)
				continue
			}
			newP = Point{2*fold.x - p.x, p.y}
		} else if fold.y != 0 && p.y > fold.y {
			if 2*fold.y < p.y {
				fmt.Printf("Point would go negative after fold: point-%v, fold-y=%d.", p, fold.y)
				continue
			}
			newP = Point{p.x, 2*fold.y - p.y}
		}

		if inRemaining(remaining, newP) {
			continue
		}
		remaining = append(remaining, newP)
	}

	return remaining
}

func drawPoints(points []Point) {
	grid := [50][50]rune{}
	xmax := 0
	ymax := 0

	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			grid[i][j] = '.'
		}
	}

	for _, p := range points {
		grid[p.x][p.y] = '#'
		if p.x > xmax {
			xmax = p.x + 1
		}
		if p.y > ymax {
			ymax = p.y + 1
		}
	}

	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			fmt.Printf("%s", string(grid[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func inRemaining(remaining []Point, newP Point) bool {
	for _, p := range remaining {
		if p == newP {
			return true
		}
	}
	return false
}

func getTestInput() []string {
	return []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
		"",
		"fold along y=7",
		"fold along x=5",
	}
}

func parseInput(raw []string) ([]Point, []Point) {
	dots := make([]Point, 0)
	folds := make([]Point, 0)
	onFolds := false

	for _, line := range raw {
		if !onFolds {
			s := strings.Split(line, ",")
			if len(s) != 2 {
				if len(line) == 0 {
					onFolds = true
					continue
				}
				fmt.Printf("Invalid line: %s\n", line)
				continue
			}

			num1, err := strconv.Atoi(s[0])
			if err != nil {
				fmt.Printf("Invalid line: %s error-%v\n", line, err)
				continue
			}
			num2, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Printf("Invalid line: %s error-%v\n", line, err)
				continue
			}

			dots = append(dots, Point{num1, num2})
		} else {
			s := strings.Split(line, "=")
			if len(s) != 2 {
				fmt.Printf("Invalid fold line: %s\n", line)
				continue
			}

			num, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Printf("Invalid fold line: %s error-%v\n", line, err)
				continue
			}

			if rune(s[0][len(s[0])-1]) == 'x' {
				folds = append(folds, Point{num, 0})
			} else {
				folds = append(folds, Point{0, num})
			}
		}
	}

	return dots, folds
}
