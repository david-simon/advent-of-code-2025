package main

import (
	"github.com/david-simon/advent-of-code-2025/internal/file"
	"github.com/david-simon/advent-of-code-2025/internal/utils"
	"log"
)

const (
	NORTH = iota
	WEST
)

var directionMatrices = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func main() {
	part1()
	part2()
}

func part2() {
	res := 0

	patrolMap, guardPos, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	direction := NORTH
	_, res = traverse(patrolMap, guardPos, direction, true)

	log.Println("Part2:", res)
}

func part1() {
	res := 0

	patrolMap, guardPos, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	direction := NORTH
	res, _ = traverse(patrolMap, guardPos, direction, false)

	log.Println("Part1:", res)
}

func traverse(originalMap [][]byte, guardPos *utils.Pair, direction int, findPotentialLoops bool) (int, int) {
	distinctCellsTraversed := 0
	loops := 0
	patrolMap := make([][]byte, len(originalMap))
	for i := range patrolMap {
		patrolMap[i] = make([]byte, len(originalMap[i]))
		copy(patrolMap[i], originalMap[i])
	}

	for isInBounds(guardPos, originalMap) {
		newPos, newDir := getNextPos(guardPos, direction, originalMap)
		currentCell := patrolMap[guardPos.V1][guardPos.V2]

		if currentCell >= 'X' && currentCell <= ('X'+WEST) && (currentCell-'X') == byte(direction) {
			return 0, 1
		}

		if findPotentialLoops && newPos.V1 == 1 && newPos.V2 == 4 {
			print()
		}

		if findPotentialLoops && isInBounds(newPos, originalMap) && patrolMap[newPos.V1][newPos.V2] == '.' {
			patrolMap[newPos.V1][newPos.V2] = '#'
			_, l := traverse(patrolMap, guardPos, direction, false)
			if l > 0 {
				log.Println("Found loop at:", newPos.V1, newPos.V2)
				loops += l
			}
			patrolMap[newPos.V1][newPos.V2] = '.'
		}

		if !(currentCell >= 'X' && currentCell <= ('X'+WEST)) {
			patrolMap[guardPos.V1][guardPos.V2] = 'X' + byte(direction)
			distinctCellsTraversed++
		}

		guardPos = newPos
		direction = newDir
	}

	return distinctCellsTraversed, loops
}

func getNextPos(pos *utils.Pair, direction int, patrolMap [][]byte) (*utils.Pair, int) {
	for i := 0; i < len(directionMatrices); i++ {
		newDir := direction + i
		if newDir >= len(directionMatrices) {
			newDir -= len(directionMatrices)
		}

		nextPos := &utils.Pair{V1: pos.V1 + directionMatrices[newDir][0], V2: pos.V2 + directionMatrices[newDir][1]}
		if !isInBounds(nextPos, patrolMap) || patrolMap[nextPos.V1][nextPos.V2] != '#' {
			return nextPos, newDir
		}
	}

	return nil, 0
}

func isInBounds(pos *utils.Pair, patrolMap [][]byte) bool {
	return pos.V1 >= 0 && pos.V1 < len(patrolMap) && pos.V2 >= 0 && pos.V2 < len(patrolMap[0])
}

func readInput() ([][]byte, *utils.Pair, error) {
	patrolMap := make([][]byte, 0, 131)
	guardPos := &utils.Pair{V1: -1, V2: -1}

	for line, err := range file.Lines("inputs/day6.txt") {
		if err != nil {
			return nil, nil, err
		}

		patrolMap = append(patrolMap, line)

		for i := 0; i < len(line); i++ {
			if line[i] == '^' {
				guardPos.V1 = len(patrolMap) - 1
				guardPos.V2 = i
			}
		}

	}

	return patrolMap, guardPos, nil
}
