package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stone struct {
	val     int
	nextLvl []*Stone
}

func (s *Stone) String() string {
	levelVals := []int{}
	for _, lvl := range s.nextLvl {
		levelVals = append(levelVals, lvl.val)
	}
	return fmt.Sprintf("&{val: %d, next: %v}", s.val, levelVals)
}

func solutionPart1(fileName string) int {
	// Read input file
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("error reading the file")
	}

	input := strings.Fields(string(file))
	stones := make([]int, len(input))
	for i, s := range input {
		stones[i], _ = strconv.Atoi(s)
	}

	blinks := 75
	saved := make(map[int]bool)
	stonesMap := make(map[int]*Stone)
	getCount := [][2]int{}

	// Process each blink
	for level := 0; level < blinks; level++ {
		newStones := make([]int, 0, len(stones)) // Efficient allocation

		for _, s := range stones {
			// Skip already processed stones
			if saved[s] {
				getCount = append(getCount, [2]int{s, blinks - level})
				continue
			}
			saved[s] = true

			var nextStone *Stone
			switch {
			case s == 0:
				// Handle stone with value 0
				if nextStone = stonesMap[1]; nextStone == nil {
					nextStone = &Stone{val: 1}
					stonesMap[1] = nextStone
				}
				stonesMap[0] = &Stone{val: 0, nextLvl: []*Stone{nextStone}}
				newStones = append(newStones, 1)

			case len(strconv.Itoa(s))%2 == 0:
				// Handle stone with an even length string
				left, right := splitStoneValue(s)
				nextStone = getStone(s, stonesMap)
				leftStone := getStone(left, stonesMap)
				rightStone := getStone(right, stonesMap)
				nextStone.nextLvl = append(nextStone.nextLvl, leftStone, rightStone)
				newStones = append(newStones, left, right)

			default:
				// Handle stone with a non-even length
				next := s * 2024
				nextStone = getStone(next, stonesMap)
				currentStone := getStone(s, stonesMap)
				currentStone.nextLvl = append(currentStone.nextLvl, nextStone)
				newStones = append(newStones, next)
			}
		}

		stones = newStones
		fmt.Printf("After %d blinks: %d\n", level+1, len(stones))
	}

	// Process the getCount data
	sum := len(stones)
	for _, p := range getCount {
		next := stonesMap[p[0]].nextLvl
		for l := 1; l < p[1]; l++ {
			newNext := []*Stone{}
			for _, s := range next {
				newNext = append(newNext, s.nextLvl...)
			}
			next = newNext
		}
		sum += len(next)
	}

	return sum
}

func splitStoneValue(s int) (int, int) {
	str := strconv.Itoa(s)
	mid := len(str) / 2
	left, _ := strconv.Atoi(str[:mid])
	rightStr := strings.TrimLeft(str[mid:], "0")
	if rightStr == "" {
		rightStr = "0"
	}
	right, _ := strconv.Atoi(rightStr)
	return left, right
}

func getStone(val int, stonesMap map[int]*Stone) *Stone {
	stone, exists := stonesMap[val]
	if !exists {
		stone = &Stone{val: val}
		stonesMap[val] = stone
	}
	return stone
}

func main() {
	fmt.Printf("sol1: %d\n", solutionPart1("input.txt"))
}

