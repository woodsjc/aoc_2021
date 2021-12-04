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
