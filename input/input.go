package input

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInputFile(fileName string) []string {
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

func Day4Input() ([]int, [][5][5]int) {
	raw := readInputFile("input/day4.txt")
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
	raw := readInputFile("input/day5.txt")
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
	raw := readInputFile("input/day6.txt")

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
	raw := readInputFile("input/day7.txt")

	crabs := make([]int, 0)
	for _, s := range strings.Split(raw[0], ",") {
		num, err := strconv.Atoi(strings.TrimSpace(s))
		if err == nil {
			crabs = append(crabs, num)
		}
	}

	return crabs
}
