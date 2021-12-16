package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/woodsjc/aoc_2021/input"
	"github.com/woodsjc/aoc_2021/internal/queue"
)

type Point struct {
	x     int
	y     int
	total int
}

type Grid [][]int

func main() {
	input := input.ReadInputFile("input/day15.txt")
	//input = getTestInput()
	grid := parseInput(input)

	part1(grid)
	part2(grid)
}

func part1(grid Grid) {
	fmt.Printf("Part 1: %d\n", Dijkstra(grid, Point{0, 0, grid[0][0]}))
}

func part2(grid Grid) {
	expandedGrid := expandGrid(grid)

	//fmt.Printf("Expanded grid: \n")
	//for _, g := range expandedGrid {
	//	fmt.Printf("%d\n", g)
	//}

	fmt.Printf("Part 2: %d\n", Dijkstra(expandedGrid, Point{0, 0, expandedGrid[0][0]}))
}

func expandGrid(grid Grid) Grid {
	rowMax := len(grid)
	colMax := len(grid[0])
	expanded := make(Grid, rowMax*5)

	for i := 0; i < rowMax*5; i++ {
		expanded[i] = make([]int, colMax*5)
	}

	for i := 0; i < rowMax; i++ {
		for j := 0; j < colMax; j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					newVal := grid[i][j] + k + l
					if newVal > 9 {
						newVal = newVal % 9
					}
					expanded[i+k*rowMax][j+l*colMax] = newVal
				}
			}
		}
	}

	return expanded
}

func Dijkstra(grid Grid, start Point) int {
	dist := make(map[Point]int)
	prev := make(map[Point]Point)
	q := queue.Queue{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			p := Point{i, j, grid[i][j]}
			dist[p] = math.MaxInt32
			prev[p] = Point{-1, -1, -1}
			q.Add(p)
		}
	}
	dist[start] = 0

	for q.Len() > 0 {
		currentInterface, err := q.Get()
		current, ok := currentInterface.(Point)
		if err != nil || !ok {
			fmt.Printf("Error or unable to unmarshall from queue: unmarshall-%t, err-%v\n", ok, err)
			continue
		}

		if current.x == len(grid)-1 && current.y == len(grid[len(grid)-1])-1 {
			//printPath(prev, current)
			return dist[current]
		}

		//each neighbor in queue
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				//no diagonal
				if i == j || i == -j {
					continue
				}

				x := current.x + i
				y := current.y + j
				//no out of bounds
				if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
					continue
				}

				neighbor := Point{x, y, grid[x][y]}
				//slows way down and same answer
				//if !q.InQueue(neighbor) {
				//	continue
				//}

				alt := dist[current] + neighbor.total
				if alt < dist[neighbor] {
					dist[neighbor] = alt
					prev[neighbor] = current
				}
			}
		}
	}

	return 0
}

func printPath(prev map[Point]Point, end Point) {
	undefined := Point{-1, -1, -1}
	total := 0

	for current := end; current != undefined && (current.x != 0 || current.y != 0); current = prev[current] {
		total += current.total
		fmt.Printf("Point: x:%d, y:%d, total:%d, overall total:%d\n", current.x, current.y, current.total, total)
	}
}

func getTestInput() []string {
	return []string{
		"1163751742",
		"1381373672",
		"2136511328",
		"3694931569",
		"7463417111",
		"1319128137",
		"1359912421",
		"3125421639",
		"1293138521",
		"2311944581",
	}
}

func parseInput(raw []string) Grid {
	grid := make(Grid, 0)

	for _, line := range raw {
		if len(line) == 0 {
			continue
		}
		row := make([]int, len(line))

		for j, r := range line {
			num, err := strconv.Atoi(string(r))
			if err == nil && num >= 0 && num <= 9 {
				row[j] = num
			}
		}
		grid = append(grid, row)
	}

	return grid
}
