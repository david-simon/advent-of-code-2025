package main

import (
	"github.com/david-simon/advent-of-code-2025/internal/file"
	"github.com/david-simon/advent-of-code-2025/internal/utils"
	"log"
	"slices"
)

func main() {
	part1()
	part2()
}

func part2() {
	rules, updates, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	indexedRules := make([][]bool, 100)
	for _, rule := range rules {
		if indexedRules[rule.V1] == nil {
			indexedRules[rule.V1] = make([]bool, 100)
		}

		indexedRules[rule.V1][rule.V2] = true
	}

	res := 0
	for _, update := range updates {
		if !isValidUpdate(update, rules) {
			ordered := sort(update, indexedRules)
			res += ordered[(len(ordered)-1)/2]
		}
	}

	log.Println("Part2:", res)
}

func part1() {
	rules, updates, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	res := 0
	for _, update := range updates {
		if isValidUpdate(update, rules) {
			res += update[(len(update)-1)/2]
		}
	}

	log.Println("Part1:", res)
}

func sort(update []int, indexedRules [][]bool) []int {
	slices.SortFunc(update, func(a int, b int) int {
		if indexedRules[a][b] {
			return -1
		} else if indexedRules[b][a] {
			return 1
		}

		return 0
	})
	return update
}

func isValidUpdate(update []int, rules []utils.Pair) bool {
	valueToIndex := make([]int, 100)
	for i := 0; i < len(update); i++ {
		valueToIndex[update[i]] = i
	}

	for _, rule := range rules {
		if valueToIndex[rule.V1] != 0 && valueToIndex[rule.V2] != 0 {
			if valueToIndex[rule.V1] > valueToIndex[rule.V2] {
				return false
			}
		}
	}

	return true
}

func readInput() ([]utils.Pair, [][]int, error) {
	updates := make([][]int, 0, 190)
	rules := make([]utils.Pair, 0, 1176)
	section1Parsed := false

	for line, err := range file.Lines("inputs/day5.txt") {
		if err != nil {
			return nil, nil, err
		}

		if len(line) == 1 {
			section1Parsed = true
			continue
		}

		if !section1Parsed {
			before, after := parseRule(line)
			rules = append(rules, utils.Pair{before, after})
		} else {
			updates = append(updates, parseUpdate(line))
		}
	}

	return rules, updates, nil
}

func parseRule(line []byte) (int, int) {
	return int((line[0]-'0')*10 + (line[1] - '0')), int((line[3]-'0')*10 + (line[4] - '0'))
}

func parseUpdate(line []byte) []int {
	update := make([]int, 0)
	num := 0
	for i := 0; i < len(line); i++ {
		if line[i] < '0' || line[i] > '9' {
			update = append(update, num)
			num = 0
			continue
		}

		num = num*10 + int(line[i]-'0')
	}

	return update
}
