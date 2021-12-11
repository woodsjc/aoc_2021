package input

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInputFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return strings.Split(string(data), "\n")
}

func Day2Input() ([]string, []int) {
	raw := ReadInputFile("input/day2.txt")
	directions := make([]string, 0)
	distances := make([]int, 0)

	for _, l := range raw {
		if len(l) == 0 {
			continue
		}

		for i, s := range strings.Fields(l) {
			if i == 0 {
				directions = append(directions, s)
			} else if i == 1 {
				num, err := strconv.Atoi(s)
				if err == nil {
					distances = append(distances, num)
				}
			} else {
				continue
			}
		}
	}

	return directions, distances
}

func Day4Input() ([]int, [][5][5]int) {
	raw := ReadInputFile("input/day4.txt")
	if len(raw) < 6 {
		log.Fatalf("Not enough rows in day 4 input: %v", raw)
	}

	//top line
	var bingoNumbers []int
	for _, s := range strings.Split(raw[0], ",") {
		tmp, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Invalid bingo numbers %s", s)
		}
		bingoNumbers = append(bingoNumbers, tmp)
	}

	//boards
	boards := make([][5][5]int, 1)
	boardIndex := 0
	i := 0
	j := 0
	for _, s := range raw[2:] {
		//log.Println(s)
		if len(s) == 0 {
			boards = append(boards, [5][5]int{})
			boardIndex++
			i = 0
			j = 0
			continue
		}

		//must have 5 rows & cols in board
		j = 0
		for sIndex, currentNumber := range strings.Fields(s) {
			tmp, err := strconv.Atoi(currentNumber)
			if i >= 5 || sIndex > 5 || err != nil {
				log.Fatalf("Invalid board %s", s)
			}

			boards[boardIndex][i][j] = tmp
			j++
		}
		i++
	}

	return bingoNumbers, boards
}

func Day5Input() [][4]int {
	raw := ReadInputFile("input/day5.txt")
	i, j := 0, 0

	segments := make([][4]int, 0)
	for _, s := range raw {
		segments = append(segments, [4]int{})
		j = 0

		for _, pair := range strings.Fields(s) {
			for _, num := range strings.Split(pair, ",") {
				tmp, err := strconv.Atoi(num)
				if err == nil && j < 4 {
					segments[i][j] = tmp
					j++
				}
			}
		}

		i++
	}

	return segments
}

func Day6Input() [9]int64 {
	raw := ReadInputFile("input/day6.txt")

	fish := [9]int64{}
	for _, n := range strings.Split(raw[0], ",") {
		num, err := strconv.Atoi(n)
		if err == nil && num >= 0 && num <= 8 {
			fish[num]++
		}
	}
	return fish
}

func Day7Input() []int {
	raw := ReadInputFile("input/day7.txt")

	crabs := make([]int, 0)
	for _, s := range strings.Split(raw[0], ",") {
		num, err := strconv.Atoi(strings.TrimSpace(s))
		if err == nil {
			crabs = append(crabs, num)
		}
	}

	return crabs
}

func Day8Input() []string {
	return ReadInputFile("input/day8.txt")
}

func Day9Input() [][]int {
	raw := ReadInputFile("input/day9.txt")
	i, j := 0, 0

	segments := make([][]int, 0)
	for _, s := range raw {
		cols := len(s)
		if cols == 0 {
			continue
		}

		segments = append(segments, make([]int, cols))
		j = 0

		for _, rawNum := range s {
			num, err := strconv.Atoi(string(rawNum))
			if err == nil {
				segments[i][j] = num
				j++
			} else {
				log.Printf("Unable to parse %v from %s", rawNum, s)
			}
		}

		i++
	}

	return segments
}

func Day11Input() [10][10]int {
	raw := ReadInputFile("input/day11.txt")
	grid := [10][10]int{}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i > len(raw) || j > len(raw[i]) {
				log.Fatalf("Invalid input: input/day11.txt")
			}
			num, err := strconv.Atoi(string(raw[i][j]))
			if err != nil {
				log.Fatalf("Invalid input: input/day11.txt: %s", raw[i])
			}
			grid[i][j] = num
		}
	}

	return grid
}
