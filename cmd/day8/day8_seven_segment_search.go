package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/woodsjc/aoc_2021/input"
)

func main() {
	raw := input.Day8Input()
	//raw = getTestInput()

	part1(raw)
	part2(raw)
}

func part1(s []string) {
	total := 0

	for _, str := range s {
		total += count1478(str)
	}

	fmt.Printf("Part 1: %d entries of 1,4,7,8 after |\n", total)
}

func part2(s []string) {
	total := 0

	for _, str := range s {
		total += countDecoded(str)
	}

	fmt.Printf("Part 2: %d entries of 1,4,7,8 after |\n", total)

}

type SevenDigitDisplay struct {
	zero  string
	one   string
	two   string
	three string
	four  string
	five  string
	six   string
	seven string
	eight string
	nine  string
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// 9 & 6 & 0 are 1 off of 8
// 9 includes all of 4
// 0 includes both of segments from 1
// third 6 length must be 6
// 2 & 3 & 5 are length 5
// 3 includes both segments from 1
// 5 is 1 off 4 components
// 2 is 2 off 4 components
func countDecoded(s string) int {
	result := ""
	display := SevenDigitDisplay{}
	lengthSix := [3]string{}
	lengthFive := [3]string{}
	sixIndex := 0
	fiveIndex := 0

	//get 1,4,7,8
	index := strings.Index(s, "|")
	if index == -1 {
		fmt.Printf("Unable to parse: %s\n", s)
		return 0
	}
	for _, str := range strings.Fields(s[:index]) {
		str = SortString(strings.TrimSpace(str))
		if len(str) == 2 { //must be a one
			display.one = str
		} else if len(str) == 3 { //must be seven
			display.seven = str
		} else if len(str) == 4 { //must be four
			display.four = str
		} else if len(str) == 5 {
			lengthFive[fiveIndex] = str
			fiveIndex++
		} else if len(str) == 6 {
			lengthSix[sixIndex] = str
			sixIndex++
		} else if len(str) == 7 { //must be eight
			display.eight = str
		}
	}

	//get 9
	for i := 0; i < 3; i++ {
		if contains(lengthSix[i], display.four) {
			display.nine = lengthSix[i]
			if i != 2 {
				lengthSix[i] = lengthSix[2]
			}
			break
		}
	}
	//get 6, 0
	if contains(lengthSix[0], display.one) {
		display.zero = lengthSix[0]
		display.six = lengthSix[1]
	} else {
		display.six = lengthSix[0]
		display.zero = lengthSix[1]
	}

	//get 3
	for i := 0; i < 3; i++ {
		if contains(lengthFive[i], display.one) {
			display.three = lengthFive[i]
			if i != 2 {
				lengthFive[i] = lengthFive[2]
			}
			break
		}
	}
	//get 2, 5
	if countContains(lengthFive[0], display.four) == 3 {
		display.five = lengthFive[0]
		display.two = lengthFive[1]
	} else {
		display.two = lengthFive[0]
		display.five = lengthFive[1]
	}

	//fmt.Printf("Display: %v\n", display)
	for _, str := range strings.Fields(s[index+1:]) {
		str = SortString(strings.TrimSpace(str))

		switch str {
		case display.zero:
			result += "0"
		case display.one:
			result += "1"
		case display.two:
			result += "2"
		case display.three:
			result += "3"
		case display.four:
			result += "4"
		case display.five:
			result += "5"
		case display.six:
			result += "6"
		case display.seven:
			result += "7"
		case display.eight:
			result += "8"
		case display.nine:
			result += "9"
		default:
			fmt.Printf("Unable to parse %s\n", str)
		}
	}

	total, err := strconv.Atoi(result)
	if err != nil {
		fmt.Printf("Unable to convert result: %s", result)
	}
	return total
}

func contains(s string, sub string) bool {
	for _, r := range sub {
		if strings.Index(s, string(r)) == -1 {
			return false
		}
	}
	return true
}

func countContains(s string, sub string) int {
	total := 0
	for _, r := range sub {
		if strings.Index(s, string(r)) >= 0 {
			total++
		}
	}
	return total
}

func count1478(s string) int {
	total := 0

	for _, str := range strings.Fields(s[strings.Index(s, "|")+1:]) {
		str = strings.TrimSpace(str)

		if len(str) == 2 { //must be a one
			total++
		} else if len(str) == 3 { //must be seven
			total++
		} else if len(str) == 4 { //must be four
			total++
		} else if len(str) == 7 { //must be eight
			total++
		}
	}
	return total
}

func getTestInput() []string {
	return []string{
		"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
	}
}
