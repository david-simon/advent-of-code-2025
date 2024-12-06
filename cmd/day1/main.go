package main

import (
	"errors"
	"github.com/david-simon/advent-of-code-2025/internal/file"
	"log"
	"math"
	"sort"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part2() {
	leftCol, rightCol, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(leftCol)
	sort.Ints(rightCol)

	similarity := 0
	j := 0
	for i := 0; i < len(leftCol); i++ {
		occurences := 0
		for ; j < len(rightCol); j++ {
			if leftCol[i] == rightCol[j] {
				occurences++
			} else if rightCol[j] > leftCol[i] {
				break
			}
		}
		similarity += leftCol[i] * occurences
	}

	log.Println("part2:", similarity)
}

func part1() {
	leftCol, rightCol, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(leftCol)
	sort.Ints(rightCol)

	diff := 0
	for i := 0; i < len(leftCol); i++ {
		diff += int(math.Abs(float64(rightCol[i] - leftCol[i])))
	}

	log.Println("part1:", diff)
}

func readInput() ([]int, []int, error) {
	leftCol := make([]int, 0)
	rightCol := make([]int, 0)

	for line, err := range file.Lines("inputs/day1.txt") {
		if err != nil {
			return nil, nil, err
		}

		left, right, err := parseLine(line)
		if err != nil {
			return nil, nil, err
		}

		leftCol = append(leftCol, left)
		rightCol = append(rightCol, right)
	}

	return leftCol, rightCol, nil
}

func parseLine(line []byte) (int, int, error) {
	var nums = make([]int, 2)
	var currNum = 0

	for i := 0; i < len(line); i++ {
		if '0' <= line[i] && line[i] <= '9' {
			nums[currNum] = nums[currNum]*10 + int(line[i]-'0')
		} else if line[i] == ' ' {
			if currNum == 0 {
				currNum++
			}
		} else if line[i] != '\n' {
			return 0, 0, errors.New("malformed line, unknown char: " + string(line[i]))
		}
	}

	if currNum != 1 {
		return 0, 0, errors.New("malformed line, currNum=" + strconv.Itoa(currNum))
	}

	return nums[0], nums[1], nil
}
