package main

import (
	"fmt"
	"sort"

	"github.com/woodsjc/aoc_2021/input"
)

type Point struct {
	x int
	y int
}

type HeightMap struct {
	heights          [][]int
	lowPoints        []int
	lowPointPosition []Point
	basinTotal       []int
	visited          map[Point]bool
}

type Queue struct {
	queue []Point
}

func main() {
	raw := input.Day9Input()
	//raw = getTestInput()
	//fmt.Printf("Raw: %v\n", raw)

	heightMap := HeightMap{heights: raw, lowPoints: []int{}}
	heightMap.InitLowPoints()

	part1(heightMap)
	part2(heightMap)
}

func part1(heightMap HeightMap) {
	fmt.Printf("Low points: %v\n", heightMap.lowPoints)
	fmt.Printf("Part 1: risk %d\n", heightMap.CalculateRisk())
}

func part2(heightMap HeightMap) {
	fmt.Printf("Low point position: %v\n", heightMap.lowPointPosition)
	heightMap.CalculateBasinTotal()
	fmt.Printf("Part 2: risk %d\n", heightMap.MultiplyTop3Basin())
}

func (h *HeightMap) InitLowPoints() {
	for i := 0; i < len(h.heights); i++ {
		for j := 0; j < len(h.heights[i]); j++ {
			current := h.heights[i][j]

			//check up
			if i > 0 && h.heights[i-1][j] <= current {
				continue
			}
			//check down
			if i < len(h.heights)-1 && h.heights[i+1][j] <= current {
				continue
			}
			//check left
			if j > 0 && h.heights[i][j-1] <= current {
				continue
			}
			//check right
			if j < len(h.heights[i])-1 && h.heights[i][j+1] <= current {
				continue
			}

			h.lowPoints = append(h.lowPoints, current)
			h.lowPointPosition = append(h.lowPointPosition, Point{i, j})
		}
	}
}

func (h HeightMap) CalculateRisk() int {
	total := 0
	for _, r := range h.lowPoints {
		total += r + 1
	}
	return total
}

//check each surrounding point and add to basin
func (h *HeightMap) CalculateBasinTotal() {
	for _, p := range h.lowPointPosition {
		h.visited = make(map[Point]bool)
		total := h.BasinTotal(p)
		h.basinTotal = append(h.basinTotal, total)
	}
}

func (q *Queue) Add(p Point) {
	if len(q.queue) == 0 {
		q.queue = make([]Point, 0)
	}
	q.queue = append(q.queue, p)
}

func (q *Queue) Get() (Point, error) {
	if len(q.queue) == 0 {
		return Point{}, fmt.Errorf("No items in queue")
	}

	p := q.queue[0]
	q.queue = q.queue[1:]
	return p, nil
}

//recursion causing stack overflow so swap to queue
func (h *HeightMap) BasinTotal(current Point) int {
	total := 0
	queue := Queue{}
	queue.Add(current)

	for len(queue.queue) > 0 {
		//fmt.Printf("Queue: %v\n", queue.queue)
		current, _ = queue.Get()
		x := current.x
		y := current.y

		if _, ok := h.visited[current]; ok {
			continue
		} else {
			h.visited[current] = true
			if h.heights[x][y] < 9 {
				//fmt.Printf("Total increased from: %v\n", current)
				total++
			}
		}

		if x >= 0 && y >= 0 && x < len(h.heights) && y < len(h.heights[x]) && h.heights[x][y] < 9 {
			if x+1 < len(h.heights) {
				queue.Add(Point{x + 1, y})
			}
			if x-1 >= 0 {
				queue.Add(Point{x - 1, y})
			}
			if y+1 < len(h.heights[x]) {
				queue.Add(Point{x, y + 1})
			}
			if y-1 >= 0 {
				queue.Add(Point{x, y - 1})
			}
		}
	}

	return total
}

func (h *HeightMap) MultiplyTop3Basin() int {
	sort.Ints(h.basinTotal)
	fmt.Printf("Sorted basins: %v\n", h.basinTotal)
	total := 1

	for i := len(h.basinTotal) - 1; i >= 0 && i >= len(h.basinTotal)-3; i -= 1 {
		total *= h.basinTotal[i]
	}

	return total
}

func getTestInput() [][]int {
	return [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}
}
