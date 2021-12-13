package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/woodsjc/aoc_2021/input"
	"github.com/woodsjc/aoc_2021/internal/queue"
)

type Node struct {
	name        string
	connections []string
	big         bool
}

type Graph map[string]*Node

func (graph *Graph) AddNode(name string, connection string) {
	g := *graph
	if n, ok := g[name]; !ok {
		n = &Node{
			name:        name,
			connections: []string{connection},
			big:         unicode.IsUpper(rune(name[0])),
		}
		g[name] = n
	} else {
		//fmt.Printf("Adding %s to %s:", connection, name)
		n.connections = append(n.connections, connection)
		//fmt.Printf(" %v\n", n.connections)
	}
}

func main() {
	raw := input.ReadInputFile("input/day12.txt")
	graph := make(Graph)

	for _, line := range raw {
		s := strings.Split(line, "-")
		if len(s) != 2 {
			//fmt.Printf("Invalid line %s.\n", line)
			continue
		}

		graph.AddNode(s[0], s[1])
		graph.AddNode(s[1], s[0])
	}

	//graph := getTestInput()
	part1(graph)
	part2(graph)
}

func part1(g Graph) {
	//fmt.Printf("Graph: %v\n", g)
	fmt.Printf("Part 1: %d paths through graph.\n", g.CalcPaths(false))
}

func part2(g Graph) {
	fmt.Printf("Part 2: %d paths through graph.\n", g.CalcPaths(true))
}

func inRoute(route []string, v string) bool {
	for _, r := range route {
		if r == v {
			return true
		}
	}

	return false
}

func inRoute2Times(route []string, v string) bool {
	if v == "start" {
		return true
	}
	counts := make(map[string]int)

	for _, r := range route {
		if len(r) == 0 || len(r) > 0 && unicode.IsUpper(rune(r[0])) {
			continue
		}

		counts[r]++
		if counts[r] > 1 {
			//fmt.Printf("count: %v\n", counts)
			return true
		}
	}

	//fmt.Printf("Not in route 2 times: %v\n", route)

	return false
}

func (g Graph) CalcPaths(part2 bool) int {
	total := 0
	q := queue.Queue{}

	q.Add([]string{"start"})
	for len(q) > 0 && len(q) < 90000 {
		routeInterface, err := q.Get()
		route, ok := routeInterface.([]string)
		if err != nil || !ok || len(route) == 0 {
			fmt.Printf("ok-%t, error-%v\n", ok, err)
			continue
		}

		//fmt.Printf("Current route: %s\n", strings.Join(route, ","))
		node := g[route[len(route)-1]]

		for _, v := range node.connections {
			//fmt.Printf("Current route: %s, v-%s\n", strings.Join(route, ","), v)
			if v == "end" {
				total++
				//fmt.Printf("Route complete: %s,end.\n", strings.Join(route, ","))
			} else if g[v].big ||
				!g[v].big && !inRoute(route, v) ||
				part2 && !g[v].big && !inRoute2Times(route, v) {

				newSlice := make([]string, len(route)+1)
				copy(newSlice, route)
				q.Add(append(newSlice, v))
				//fmt.Printf("Queued %s,%s\n", strings.Join(route, ","), v)
			}
		}
	}

	return total
}

func getTestInput() Graph {
	return Graph{
		"start": {name: "start", connections: []string{"A", "b"}, big: false},
		"A":     {name: "A", connections: []string{"c", "b", "end"}, big: true},
		"b":     {name: "b", connections: []string{"A", "d", "end"}, big: false},
		"c":     {name: "c", connections: []string{"A"}, big: false},
		"d":     {name: "d", connections: []string{"b"}, big: false},
	}
}
