package main

import (
	"fmt"

	"github.com/woodsjc/aoc_2021/input"
)

type Position struct {
	x   int
	y   int
	aim int
}

type Move struct {
	direction string
	distance  int
}

func main() {
	directions, distances := input.Day2Input()
	//input = getTestInput()

	moves := make([]Move, len(directions))
	for i := 0; i < len(directions); i++ {
		moves[i] = Move{direction: directions[i], distance: distances[i]}
	}

	part1(moves)
	part2(moves)
}

func part1(moves []Move) {
	p := calcPosition(moves)
	fmt.Printf("Part 1: horizontal * depth is %d.\n", p.x*p.y)
}

func part2(moves []Move) {
	p := calcPositionWithAim(moves)
	fmt.Printf("Part 2: horizontal * depth is %d.\n", p.x*p.y)
}

func calcPosition(moves []Move) Position {
	p := Position{0, 0, 0}

	for _, move := range moves {
		direction := move.direction
		distance := move.distance
		switch direction {
		case "forward":
			p.x += distance
		case "backward":
			p.x -= distance
		case "up":
			p.y -= distance
		case "down":
			p.y += distance
		default:
			fmt.Printf("Unknown direction: %s", direction)
		}
	}

	return p
}

func calcPositionWithAim(moves []Move) Position {
	p := Position{0, 0, 0}

	for _, move := range moves {
		direction := move.direction
		distance := move.distance
		switch direction {
		case "forward":
			p.x += distance
			p.y += p.aim * distance
		case "backward":
			p.x -= distance
			p.y -= p.aim * distance
		case "up":
			p.aim -= distance
		case "down":
			p.aim += distance
		default:
			fmt.Printf("Unknown direction: %s", direction)
		}
	}

	return p
}

func getTestInput() []Move {
	return []Move{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}
}
