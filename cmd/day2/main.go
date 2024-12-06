package main

import (
	"errors"
	"github.com/david-simon/advent-of-code-2025/internal/file"
	"log"
	"math"
)

const UNKNOWN = 0
const ASC = 1
const DESC = 2

func main() {
	part1()
	part2()
}

func part2() {
	rows, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	safeReports := 0

	for i := 0; i < len(rows); i++ {
		for j := -1; j < len(rows[i]); j++ {
			if isSafe(rows[i], j) {
				safeReports++
				break
			}
		}
	}

	log.Println("part2: ", safeReports)
}

func part1() {
	rows, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	safeReports := 0

	for i := 0; i < len(rows); i++ {
		if isSafe(rows[i], -1) {
			safeReports++
		}
	}

	log.Println("part1: ", safeReports)
}

func isSafe(nums []int, skipIndex int) bool {
	prev := -1
	dir := UNKNOWN

	safe := true
	for j := 0; j < len(nums); j++ {
		if j == skipIndex {
			continue
		} else if prev == -1 {
			prev = nums[j]
			continue
		}

		if dir == UNKNOWN {
			if nums[j] < prev {
				dir = DESC
			} else if nums[j] > prev {
				dir = ASC
			}
		} else if (dir == ASC && nums[j] < prev) || (dir == DESC && nums[j] > prev) {
			safe = false
			break
		}

		diff := int(math.Abs(float64(nums[j] - prev)))
		if diff < 1 || diff > 3 {
			safe = false
			break
		}

		prev = nums[j]
	}

	return safe
}

func readInput() ([][]int, error) {
	rows := make([][]int, 0)

	for line, err := range file.Lines("inputs/day2.txt") {
		if err != nil {
			return nil, err
		}

		row, err := parseLine(line)
		if err != nil {
			return nil, err
		}

		rows = append(rows, row)
	}

	return rows, nil
}

func parseLine(line []byte) ([]int, error) {
	var nums = make([]int, 0)
	var currNum = 0

	for i := 0; i < len(line); i++ {
		if '0' <= line[i] && line[i] <= '9' {
			if len(nums) < currNum+1 {
				nums = append(nums, 0)
			}

			nums[currNum] = nums[currNum]*10 + int(line[i]-'0')
		} else if line[i] == ' ' {
			currNum++
		} else if line[i] != '\n' {
			return nil, errors.New("malformed line, unknown char: " + string(line[i]))
		}
	}

	return nums, nil
}
