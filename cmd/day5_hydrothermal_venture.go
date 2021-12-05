package main

import (
	"fmt"

	"github.com/woodsjc/aoc_2021/input"
)

type Point struct {
	x int
	y int
}

type Segment struct {
	start Point
	end   Point
}

type Grid struct {
	grid    [][]int
	xOffset int
	yOffset int
}

func main() {
	raw := input.Day5Input() // getTestInput()
	var lineSegments []Segment

	for _, s := range raw {
		lineSegments = append(lineSegments, Segment{
			start: Point{s[0], s[1]},
			end:   Point{s[2], s[3]},
		})
	}

	part1(lineSegments)
	part2(lineSegments)
}

//find min & max x & y coordinates
//go through each line segment and increment points in segment
//count points >= 2
func part1(segments []Segment) {
	xmax, xmin, ymax, ymin := SegmentBoundaries(segments)

	//fmt.Printf("xmax-%d, ymax-%d, xmin-%d, ymin-%d\n", xmax, ymax, xmin, ymin)
	grid := Grid{
		grid:    make([][]int, xmax-xmin+1),
		xOffset: xmin,
		yOffset: ymin,
	}
	for i := range grid.grid {
		grid.grid[i] = make([]int, ymax-ymin+1)
	}
	//fmt.Printf("Grid - %v\n", grid)

	for _, s := range segments {
		if !s.IsVerticalOrHorizontal() {
			continue
		}
		grid.UpdateSegment(s)
		//fmt.Printf("Grid - %v\n", grid.grid)
	}

	fmt.Printf("Part 1: %d points with more than 1 segment\n", grid.CountGridIntersections())
}

//include diagonals
func part2(segments []Segment) {
	xmax, xmin, ymax, ymin := SegmentBoundaries(segments)

	grid := Grid{
		grid:    make([][]int, xmax-xmin+1),
		xOffset: xmin,
		yOffset: ymin,
	}
	for i := range grid.grid {
		grid.grid[i] = make([]int, ymax-ymin+1)
	}

	for _, s := range segments {
		grid.UpdateSegmentWithDiagonals(s)
	}

	fmt.Printf("Part 2: %d points with more than 1 segment\n", grid.CountGridIntersections())
}

func SegmentBoundaries(segments []Segment) (xmax int, xmin int, ymax int, ymin int) {
	for i, s := range segments {
		max := s.Max()
		min := s.Min()
		if i == 0 {
			xmax = max.x
			ymax = max.y
			xmin = min.x
			ymin = min.y
			continue
		}

		if max.x > xmax {
			xmax = max.x
		}
		if min.x < xmin {
			xmin = min.x
		}
		if max.y > ymax {
			ymax = max.y
		}
		if min.y < ymin {
			ymin = min.y
		}
	}

	return xmax, xmin, ymax, ymin
}

func (s *Segment) GetStartAndEnd() (int, int, int, int) {
	var xstart int
	var xend int
	var ystart int
	var yend int

	if s.start.x <= s.end.x {
		xstart = s.start.x
		xend = s.end.x
	} else {
		xstart = s.end.x
		xend = s.start.x
	}

	if s.start.y <= s.end.y {
		ystart = s.start.y
		yend = s.end.y
	} else {
		ystart = s.end.y
		yend = s.start.y
	}

	return xstart, xend, ystart, yend
}

func (g *Grid) UpdateSegment(s Segment) {
	xstart, xend, ystart, yend := s.GetStartAndEnd()

	for i := xstart; i <= xend; i++ {
		for j := ystart; j <= yend; j++ {
			g.grid[i-g.xOffset][j-g.yOffset] += 1
		}
	}
}

func (g *Grid) UpdateSegmentWithDiagonals(s Segment) {
	if s.IsVerticalOrHorizontal() {
		g.UpdateSegment(s)
		return
	}

	xstart := s.start.x
	xend := s.end.x
	ystart := s.start.y
	//yend := s.end.y
	ydirectionPositive := true

	if s.start.x < s.end.x {
		if s.start.y >= s.end.y {
			ydirectionPositive = false
		}
	} else {
		xstart = s.end.x
		ystart = s.end.y
		xend = s.start.x
		//yend = s.end.y
		if s.start.y < s.end.y {
			ydirectionPositive = false
		}
	}

	for i := 0; i <= xend-xstart; i++ {
		//slope must be 1
		//fmt.Printf("startx-%d, starty-%d, endx-%d, endy-%d, i-%d\n", xstart, ystart, xend, yend, i)
		x := i + xstart - g.xOffset
		y := i + ystart - g.yOffset
		if !ydirectionPositive {
			y = -i + ystart - g.yOffset
		}
		g.grid[x][y] += 1
	}
}

func (g *Grid) CountGridIntersections() int {
	total := 0
	for i := range g.grid {
		for j := range g.grid[i] {
			if g.grid[i][j] > 1 {
				total++
			}
		}
	}

	return total
}

func (s *Segment) IsVerticalOrHorizontal() bool {
	if s.start.x == s.end.x || s.start.y == s.end.y {
		return true
	}
	return false
}

func (s *Segment) Max() Point {
	var xmax int
	var ymax int

	if s.start.x > s.end.x {
		xmax = s.start.x
	} else {
		xmax = s.end.x
	}

	if s.start.y > s.end.y {
		ymax = s.start.y
	} else {
		ymax = s.end.y
	}

	return Point{xmax, ymax}
}

func (s *Segment) Min() Point {
	var xmin int
	var ymin int

	if s.start.x < s.end.x {
		xmin = s.start.x
	} else {
		xmin = s.end.x
	}

	if s.start.y < s.end.y {
		ymin = s.start.y
	} else {
		ymin = s.end.y
	}

	return Point{xmin, ymin}
}

func getTestInput() []Segment {
	return []Segment{
		{Point{0, 9}, Point{5, 9}},
		{Point{8, 0}, Point{0, 8}},
		{Point{9, 4}, Point{3, 4}},
		{Point{2, 2}, Point{2, 1}},
		{Point{7, 0}, Point{7, 4}},
		{Point{6, 4}, Point{2, 0}},
		{Point{0, 9}, Point{2, 9}},
		{Point{3, 4}, Point{1, 4}},
		{Point{0, 0}, Point{8, 8}},
		{Point{5, 5}, Point{8, 2}},
	}
}
