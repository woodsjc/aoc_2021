package main

import (
	"fmt"

	"github.com/woodsjc/aoc_2021/input"
)

type CrabGroup struct {
	position int
	total    int
}

type Crabs struct {
	position      []int
	groupPosition []CrabGroup
}

func main() {
	raw := input.Day7Input()
	//raw = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	input := Crabs{position: raw}
	input.BuildGroup()
	part1(input)
	part2(input)
}

func part1(crabs Crabs) {
	//fmt.Printf("%v\n", crabs)
	position, total := crabs.MinPosition()
	fmt.Printf("Part 1: minimum position-%d distance-%d\n", position, total)
}

func part2(crabs Crabs) {
	//fmt.Printf("%v\n", crabs)
	position, total := crabs.MinPosition2()
	fmt.Printf("Part 2: minimum position-%d distance-%d\n", position, total)
}

func (c *Crabs) BuildGroup() {
	unique := map[int]int{}

	for _, p := range c.position {
		if _, exists := unique[p]; !exists {
			unique[p] = 0
		}
		unique[p]++
	}

	c.groupPosition = make([]CrabGroup, len(unique))
	i := 0
	for k, v := range unique {
		c.groupPosition[i].position = k
		c.groupPosition[i].total = v
		i++
	}
}

func (c Crabs) Distance(target int) int {
	total := 0

	for _, g := range c.groupPosition {
		p := g.position
		if p > target {
			total += g.total * (p - target)
		} else {
			total += g.total * (target - p)
		}
	}

	return total
}

func (c Crabs) MinPosition() (int, int) {
	min := 0
	minTotal := 0

	for i, g := range c.groupPosition {
		total := c.Distance(g.position)

		if i == 0 || minTotal > total {
			min = g.position
			minTotal = total
		}
	}

	return min, minTotal
}

func (c Crabs) MinPosition2() (int, int) {
	min := 0
	minTotal := 0

	max := c.Max()
	for i := 0; i <= max; i++ {
		for _, g := range c.groupPosition {
			total := c.Distance2(i)

			if i == 0 || minTotal > total {
				min = g.position
				minTotal = total
			}
		}
	}

	return min, minTotal
}

func cost(n int) int {
	return n * (n + 1) / 2
}

func (c Crabs) Distance2(target int) int {
	total := 0

	for _, g := range c.groupPosition {
		p := g.position
		if p > target {
			total += g.total * cost(p-target)
		} else {
			total += g.total * cost(target-p)
		}
	}

	return total
}

func (c Crabs) Max() int {
	max := 0

	for i, g := range c.groupPosition {
		if i == 0 || max < g.position {
			max = g.position
		}
	}

	return max
}
