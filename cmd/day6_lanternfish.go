package main

import (
	"fmt"

	"github.com/woodsjc/aoc_2021/input"
)

//group fish by day timer with total
type Lanternfish struct {
	total [9]int64
}

func main() {
	input := input.Day6Input() //getTestInput()
	fish := Lanternfish{input}
	part1(fish)
	part2(fish)
}

func part1(fish Lanternfish) {
	for i := 0; i < 80; i++ {
		fish.UpdateAge()
		//fmt.Printf("%v\n", fish.total)
	}

	fmt.Printf("Part 1: %d fish after 80 days.\n", fish.Total())
}

func part2(fish Lanternfish) {
	for i := 0; i < 256; i++ {
		fish.UpdateAge()
	}

	fmt.Printf("Part 2: %d fish after 256 days.\n", fish.Total())
}

func (l *Lanternfish) Total() int64 {
	total := int64(0)
	for i := range l.total {
		total += l.total[i]
	}

	return total
}

func (l *Lanternfish) UpdateAge() {
	for i := 0; i < 8; i++ {
		tmp := l.total[i+1]
		l.total[i+1] = l.total[i]
		l.total[i] = tmp
	}

	l.total[6] += l.total[8]
}

func getTestInput() Lanternfish {
	return Lanternfish{
		total: [9]int64{0, 1, 1, 2, 1, 0, 0, 0, 0},
	}
}
