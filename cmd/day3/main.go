package main

import (
	"github.com/david-simon/advent-of-code-2025/internal/file"
	"log"
)

const (
	INVALID = iota
	PARSE_MUL
	PARSE_PARAM1
	PARSE_PARAM2
	PARSE_DO
	PARSE_DONT
)

func main() {
	part1()
	part2()
}

func part2() {
	res, err := readInput(true)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("part2: ", res)
}

func part1() {
	res, err := readInput(false)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("part1: ", res)
}

func readInput(withConditional bool) (uint64, error) {

	res := uint64(0)
	mulEnabled := true

	for line, err := range file.Lines("inputs/day3.txt") {
		if err != nil {
			return 0, err
		}

		row, mu := parseLine(line, withConditional, mulEnabled)
		res += row
		mulEnabled = mu
	}

	return res, nil
}

func parseLine(line []byte, withConditional bool, mulEnabled bool) (uint64, bool) {
	state := INVALID
	var lastToken byte = 0

	res := uint64(0)
	currParam1, currParam2 := uint64(0), uint64(0)

	for i := 0; i < len(line); i++ {
		if state == INVALID {
			currParam1 = 0
			currParam2 = 0
		}

		token := line[i]

		if token == 'm' {
			state = PARSE_MUL
		} else if token == 'd' {
			state = PARSE_DO
		} else if state == PARSE_MUL {
			if (lastToken == 'm' && token != 'u') || (lastToken == 'u' && token != 'l') || (lastToken == 'l' && token != '(') {
				state = INVALID
			} else if token == '(' {
				state = PARSE_PARAM1
			}
		} else if state == PARSE_PARAM1 {
			if token == ',' {
				state = PARSE_PARAM2
			} else if token >= '0' && token <= '9' {
				if currParam1 <= 99 {
					currParam1 = currParam1*10 + uint64(token-'0')
				} else {
					state = INVALID
				}
			} else {
				state = INVALID
			}
		} else if state == PARSE_PARAM2 {
			if token == ')' {
				if !withConditional || mulEnabled {
					res += currParam1 * currParam2
				}
				state = INVALID
			} else if token >= '0' && token <= '9' {
				if currParam2 <= 99 {
					currParam2 = currParam2*10 + uint64(token-'0')
				} else {
					state = INVALID
				}
			} else {
				state = INVALID
			}
		} else if state == PARSE_DO {
			if lastToken == 'o' && token == 'n' {
				state = PARSE_DONT
			} else if (lastToken == 'd' && token != 'o') || (lastToken == 'o' && token != '(') || (lastToken == '(' && token != ')') {
				state = INVALID
			} else if token == ')' {
				mulEnabled = true
				state = INVALID
			}
		} else if state == PARSE_DONT {
			if (lastToken == 'n' && token != '\'') || (lastToken == '\'' && token != 't') || (lastToken == 't' && token != '(') || (lastToken == '(' && token != ')') {
				state = INVALID
			} else if token == ')' {
				mulEnabled = false
				state = INVALID
			}
		}

		lastToken = token
	}

	return res, mulEnabled
}
