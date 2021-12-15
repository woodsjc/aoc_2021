package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strings"

	"github.com/woodsjc/aoc_2021/input"
)

var memoize map[string]string

type Rule map[string]string

type Element map[rune]*big.Int

type Pair map[string]*big.Int

type Polymer struct {
	elements Element
	pairs    Pair
}

func main() {
	input := input.ReadInputFile("input/day14.txt")
	//input = getTestInput()
	rules, start := parseInput(input)

	part1(rules, start)
	part2(rules, start)
}

func part1(rules Rule, start string) {
	for i := 0; i < 10; i++ {
		start = applyRules(rules, start)
	}
	//fmt.Printf("start: %s\n", start)

	fmt.Printf("Part 1: most - least = %d\n", subtractMostLeast(start))
}

func part2(rules Rule, start string) {
	p := Polymer{}
	p.Init(start)
	fmt.Printf("Init complete\nelements: %v\npairs: %v\n", p.elements, p.pairs)

	for i := 0; i < 40; i++ {
		p.Update(rules)
		//fmt.Printf("Step %d complete\nelements: %v\npairs: %v\n", i+1, p.elements, p.pairs)
	}
	//fmt.Printf("start: %s\n", start)

	max := p.Max()
	min := p.Min()
	fmt.Printf("Part 2: {max-%d} - {min-%d} = %v\n", max, min, big.NewInt(0).Sub(max, min))
}

func (p *Polymer) Init(start string) {
	p.elements = make(Element)
	p.pairs = make(Pair)

	for i, r := range start {
		if _, ok := p.elements[r]; !ok {
			p.elements[r] = big.NewInt(0)
		}
		p.elements[r] = big.NewInt(0).Add(big.NewInt(1), p.elements[r])

		if i < len(start)-1 {
			pair := string(r) + string(start[i+1])
			if _, ok := p.pairs[pair]; !ok {
				p.pairs[pair] = big.NewInt(0)
			}
			p.pairs[pair] = big.NewInt(0).Add(big.NewInt(1), p.pairs[pair])
		}
	}
}

func (polymer *Polymer) Update(rules Rule) {
	pairs := polymer.pairs
	newPairs := make(Pair)

	for pair, pairTotal := range pairs {
		if pairTotal == big.NewInt(0) {
			continue
		}

		if _, ok := rules[pair]; !ok {
			newPairs[pair] = pairTotal
			continue
		}

		r, _ := rules[pair]
		new := rune(r[0])
		if _, ok := polymer.elements[new]; !ok {
			polymer.elements[new] = big.NewInt(0)
		}
		polymer.elements[new] = big.NewInt(0).Add(pairTotal, polymer.elements[new])
		if _, ok := newPairs[pair]; !ok {
			newPairs[pair] = big.NewInt(0)
		}

		newP := string(pair[0]) + string(new)
		if _, ok := newPairs[newP]; ok {
			newPairs[newP] = big.NewInt(0).Add(pairTotal, newPairs[newP])
		} else {
			newPairs[newP] = pairTotal
		}

		newP = string(new) + string(pair[1])
		if _, ok := newPairs[newP]; ok {
			newPairs[newP] = big.NewInt(0).Add(pairTotal, newPairs[newP])
		} else {
			newPairs[newP] = pairTotal
		}
	}

	polymer.pairs = newPairs
}

func (p Polymer) Max() *big.Int {
	max := big.NewInt(0)
	for _, e := range p.elements {
		if e.Cmp(max) == 1 {
			max = e
		}
	}

	return max
}

func (p Polymer) Min() *big.Int {
	min := big.NewInt(0)
	minSet := false
	for _, e := range p.elements {
		if !minSet {
			min = e
			minSet = true
		}
		if e.Cmp(min) == -1 {
			min = e
		}
	}

	return min
}

func applyRules(rules Rule, start string) string {
	end := ""
	for i := 0; i < len(start)-1; i++ {
		end += string(start[i])
		if r, ok := rules[start[i:i+2]]; ok {
			end += r
		}
	}

	return end + string(start[len(start)-1])
}

func subtractMostLeast(start string) int64 {
	counts := make(map[rune]int64)
	max := int64(0)
	min := int64(math.MaxInt64)

	for _, r := range start {
		if _, ok := counts[r]; !ok {
			counts[r] = 0
		}
		counts[r]++
	}

	for _, v := range counts {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}

	return max - min
}

func getTestInput() []string {
	return []string{
		"NNCB",
		"",
		"CH -> B",
		"HH -> N",
		"CB -> H",
		"NH -> C",
		"HB -> C",
		"HC -> B",
		"HN -> C",
		"NN -> C",
		"BH -> H",
		"NC -> B",
		"NB -> B",
		"BN -> B",
		"BB -> N",
		"BC -> B",
		"CC -> N",
		"CN -> C",
	}
}

func parseInput(input []string) (Rule, string) {
	if len(input) < 3 {
		fmt.Printf("Invalid input: %v\n", input)
		os.Exit(1)
	}

	start := input[0]
	rules := make(map[string]string)
	for _, r := range input[2:] {
		s := strings.Split(r, " -> ")
		if len(s) != 2 || len(s) == 2 && len(s[0]) != 2 {
			fmt.Printf("Invalid rule: %s\n", r)
			continue
		}

		rules[s[0]] = s[1]
	}

	return rules, start
}
