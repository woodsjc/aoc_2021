package main

import (
	"fmt"

	"github.com/woodsjc/aoc_2021/input"
	"github.com/woodsjc/aoc_2021/internal/queue"
)

type Grid struct {
	grid  [10][10]int
	flash map[queue.Point]bool
	q     queue.Queue
}

func main() {
	input := input.Day11Input()
	//input = getTestInput()
	grid := Grid{grid: input}
	part1(grid)
	part2(grid)
}

func part1(g Grid) {
	total := 0

	for i := 0; i < 100; i++ {
		total += g.Update()
		//fmt.Printf("grid: %v\n", g.grid)
	}

	fmt.Printf("Part 1: %d ocotpus flashes.\n", total)
}

func part2(g Grid) {
	i, total := 0, 0

	for i = 0; total != 100; i++ {
		total = g.Update()
		//fmt.Printf("grid: %v\n", g.grid)
	}

	fmt.Printf("Part 2: All ocotpuses flash on %d.\n", i)
}

func checkFlash(flash map[queue.Point]bool, p queue.Point) bool {
	if flashed, ok := flash[p]; ok && flashed {
		return true
	}
	return false
}

func (g *Grid) UpdatePoint(i, j int) {
	p := queue.Point{X: i, Y: j}

	if i < 0 || j < 0 || i >= 10 || j >= 10 || checkFlash(g.flash, p) {
		return
	}

	if g.grid[i][j] < 9 {
		g.grid[i][j]++
	} else {
		g.grid[i][j] = 0
		g.flash[p] = true
		g.q.Add(p)
	}
}

func (g *Grid) Update() int {
	g.flash = make(map[queue.Point]bool)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			g.UpdatePoint(i, j)
		}
	}

	//run through queue
	for g.q.Len() > 0 {
		prior, err := g.q.Get()
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		//update neighbors
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				g.UpdatePoint(prior.X+i, prior.Y+j)
			}
		}
	}

	return len(g.flash)
}

func getTestInput() [10][10]int {
	return [10][10]int{
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}
}
