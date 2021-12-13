package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/woodsjc/aoc_2021/input"
)

func main() {
	input := input.ReadInputFile("input/day3.txt")
	//input = getTestInput()

	if len(input) == 0 {
		fmt.Printf("Invalid input.\n")
		os.Exit(1)
	}

	if len(input[len(input)-1]) == 0 {
		input = input[:len(input)-1]
	}

	part1(input)
	part2(input)
}

func part1(lines []string) {
	gamma, epsilon := getRates(lines)
	fmt.Printf("gamma-%s, epsilon-%s\n", gamma, epsilon)
	fmt.Printf("Part 1: %d power consumption.\n", convertRate(gamma)*convertRate(epsilon))
}

func part2(lines []string) {
	oxygen := getPart2Rate(lines, true)
	co2 := getPart2Rate(lines, false)

	fmt.Printf("oxygen-%s, co2-%s\n", oxygen, co2)
	fmt.Printf("Part 2: %d power consumption.\n", convertRate(oxygen)*convertRate(co2))
}

//respect current order
func remove(lines []string, i int) []string {
	if len(lines) == 0 {
		return lines
	}

	end := len(lines) - 1
	for ; i < end; i++ {
		lines[i] = lines[i+1]
	}
	return lines[:end]
}

func getPart2Rate(lines []string, most bool) string {
	for i := 0; i < len(lines[0]); i++ {
		var c rune
		m, l := getMostLeast(lines, i, most)
		if most {
			c = m
		} else {
			c = l
		}

		for j := 0; j < len(lines) && len(lines) > 1; j++ {
			if rune(lines[j][i]) != c {
				lines = remove(lines, j)
				j--
				//fmt.Printf("Removed %d lines left.\n", len(lines))
			}
		}

		if len(lines) == 1 {
			break
		}
	}

	return lines[0]
}

func getMostLeast(lines []string, col int, most bool) (rune, rune) {
	one := 0
	zero := 0

	for _, l := range lines {
		if l[col] == '1' {
			one++
		} else if l[col] == '0' {
			zero++
		} else {
			fmt.Printf("Invalid rune-%v in string-%s\n", l[col], l)
		}
	}

	if one > zero || one == zero && most {
		return '1', '0'
	}
	return '0', '1'
}

func convertRate(r string) int {
	num, err := strconv.ParseUint(r, 2, 32)
	if err != nil {
		fmt.Printf("Unable to parse %s: %v\n", r, err)
	}

	return int(num)
}

func getRates(lines []string) (string, string) {
	if len(lines) == 0 {
		return "", ""
	}

	gamma := ""
	epsilon := ""
	length := len(lines[0])
	common := make([]int, length)

	for _, l := range lines {
		//fmt.Printf("l-%s, length-%d, common-%v\n", l, length, common)
		for i := 0; i < length; i++ {
			if l[i] == '1' {
				common[i]++
			}
		}
	}

	for _, c := range common {
		if c >= len(lines)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	fmt.Printf("common-%v\n", common)

	return gamma, epsilon
}

func getTestInput() []string {
	return []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
}
