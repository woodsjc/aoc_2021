package main

import (
	"fmt"
	"sort"

	"github.com/woodsjc/aoc_2021/input"
)

type Stack struct {
	stack []rune
}

func main() {
	input := input.ReadInputFile("input/day10.txt")
	//input = getTestInput()

	part1(input)
	part2(input)
}

func part1(lines []string) {
	fmt.Printf("Part 1: %d total points.\n", calcCorrupted(lines))
}

func part2(lines []string) {
	fmt.Printf("Part 2: %d total points.\n", calcIncomplete(lines))
}

func (s *Stack) Push(r rune) {
	s.stack = append(s.stack, r)
}

func (s *Stack) Pop() rune {
	if len(s.stack) == 0 {
		return ' '
	}

	r := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return r
}

func calcIncomplete(lines []string) int {
	totals := make([]int, 0)
	for _, l := range lines {
		if corrupt, _ := checkCorrupted(l); corrupt {
			continue
		}

		lineTotal := 0
		s := repairIncomplete(l)
		for _, r := range s {
			lineTotal *= 5
			switch r {
			case ')':
				lineTotal += 1
			case ']':
				lineTotal += 2
			case '}':
				lineTotal += 3
			case '>':
				lineTotal += 4
			}
			//fmt.Printf("lineTotal-%d rune-%s\n", lineTotal, string(r))
		}
		//fmt.Printf("lineTotal-%d toRepair-%s\n", lineTotal, s)
		totals = append(totals, lineTotal)
	}

	sort.Ints(totals)
	i := len(totals) / 2
	return totals[i]
}

func repairIncomplete(s string) string {
	stack := Stack{}

	for _, r := range s {
		//fmt.Printf("%s\t%v\n", string(r), stack)
		switch r {
		case '(', '[', '<', '{':
			stack.Push(r)
		case ')':
			if stack.Pop() != '(' {
				fmt.Printf("Invalid character %v in %s\n", r, s)
			}
		case ']':
			if stack.Pop() != '[' {
				fmt.Printf("Invalid character %v in %s\n", r, s)
			}
		case '>':
			if stack.Pop() != '<' {
				fmt.Printf("Invalid character %v in %s\n", r, s)
			}
		case '}':
			if stack.Pop() != '{' {
				fmt.Printf("Invalid character %v in %s\n", r, s)
			}
		default:
			fmt.Printf("Invalid character %v in %s\n", r, s)
		}
	}

	result := ""
	for len(stack.stack) > 0 {
		r := stack.Pop()
		switch r {
		case '(':
			result += ")"
		case '[':
			result += "]"
		case '<':
			result += ">"
		case '{':
			result += "}"
		}
	}
	return result
}

func calcCorrupted(lines []string) int {
	total := 0
	for _, l := range lines {
		_, r := checkCorrupted(l)
		//fmt.Printf("corrupt-%t rune-%s\n", corrupt, string(r))
		switch r {
		case ')':
			total += 3
		case ']':
			total += 57
		case '}':
			total += 1197
		case '>':
			total += 25137
		}
	}
	return total
}

func checkCorrupted(s string) (bool, rune) {
	stack := Stack{}

	for _, r := range s {
		//fmt.Printf("%s\t%v\n", string(r), stack)
		switch r {
		case '(', '[', '<', '{':
			stack.Push(r)
		case ')':
			if stack.Pop() != '(' {
				return true, r
			}
		case ']':
			if stack.Pop() != '[' {
				return true, r
			}
		case '>':
			if stack.Pop() != '<' {
				return true, r
			}
		case '}':
			if stack.Pop() != '{' {
				return true, r
			}
		default:
			fmt.Printf("Invalid character %v in %s\n", r, s)
		}
	}

	return false, ' '
}

func getTestInput() []string {
	return []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}
}
