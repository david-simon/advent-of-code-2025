package main

import (
	"github.com/david-simon/advent-of-code-2025/internal/file"
	"log"
)

const (
	SOUTH = iota
	EAST
	SOUTH_EAST
	SOUTH_WEST
)

func main() {
	part1()
	part2()
}

func part2() {
	data, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	res := 0

	for row := 1; row < len(data); row++ {
		for col := 1; col < len(data[row]); col++ {
			if data[row][col] == 'A' {
				if (checkWord("MAS", data, row-1, col-1, SOUTH_EAST, 0) ||
					checkWord("SAM", data, row-1, col-1, SOUTH_EAST, 0)) &&
					(checkWord("MAS", data, row-1, col+1, SOUTH_WEST, 0) ||
						checkWord("SAM", data, row-1, col+1, SOUTH_WEST, 0)) {
					res++
				}

			}
		}
	}

	log.Println("part2: ", res)
}

func part1() {
	data, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	res := 0
	words := []string{"XMAS", "SAMX"}
	directions := []int{SOUTH, EAST, SOUTH_EAST, SOUTH_WEST}

	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data[row]); col++ {
			for _, word := range words {
				if data[row][col] == word[0] {
					for _, direction := range directions {
						if checkWord(word, data, row, col, direction, 1) {
							res++
						}
					}
				}
			}
		}
	}

	log.Println("part1: ", res)
}

func checkWord(word string, data [][]byte, row int, col int, direction int, startIndex int) bool {
	for i := startIndex; i < len(word); i++ {
		newRow := row
		newCol := col

		if direction == SOUTH || direction == SOUTH_EAST || direction == SOUTH_WEST {
			newRow += i
		}

		if direction == EAST || direction == SOUTH_EAST {
			newCol += i
		}

		if direction == SOUTH_WEST {
			newCol -= i
		}

		if newRow >= len(data) || newCol >= len(data[newRow]) || newCol < 0 {
			return false
		}

		if data[newRow][newCol] != word[i] {
			return false
		}
	}

	return true
}

func readInput() ([][]byte, error) {
	data := make([][]byte, 0)
	for line, err := range file.Lines("inputs/day4.txt") {
		if err != nil {
			return nil, err
		}

		data = append(data, line)

	}

	return data, nil
}
