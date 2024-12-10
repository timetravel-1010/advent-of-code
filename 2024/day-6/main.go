package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type VisitedKey struct {
	level int
	pos   int
}

type Guard struct {
	level     int
	pos       int
	direction string
	visited   map[int]map[int]bool
}

func NewGuard() *Guard {
	return &Guard{
		visited: make(map[int]map[int]bool),
	}
}

func solutionPart1(fileName string) int {
	positions := 0

	file, err := os.Open(fileName)
	if err != nil {
		panic("error opening the file")
	}

	obstacles := make(map[int]map[int]bool)
	guard := NewGuard()

	rightLimit := 0
	lastLevel := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		idx := strings.Index(line, "#")
		pad := idx
		for idx != -1 {
			if _, ok := obstacles[lastLevel]; !ok {
				obstacles[lastLevel] = make(map[int]bool)
			}
			obstacles[lastLevel][pad] = true
			idx = strings.Index(line[pad+1:], "#")
			pad += idx + 1
		}

		lastLevel++
		if guard.direction != "" {
			continue
		}

		if strings.Contains(line, "^") {
			guard.direction = "up"
			guard.pos = strings.Index(line, "^")
		} else if strings.Contains(line, ">") {
			guard.direction = "right"
			guard.pos = strings.Index(line, ">")
		} else if strings.Contains(line, "v") {
			guard.direction = "down"
			guard.pos = strings.Index(line, "v")
		} else if strings.Contains(line, "<") {
			guard.direction = "left"
			guard.pos = strings.Index(line, "<")
		} else {
			continue
		}
		guard.level = lastLevel - 1
		rightLimit = len(line)
	}

outer:
	for {
		if _, ok := guard.visited[guard.level]; !ok {
			guard.visited[guard.level] = make(map[int]bool)
			positions++
		} else if !guard.visited[guard.level][guard.pos] {
			positions++
		}
		guard.visited[guard.level][guard.pos] = true

		for {
			if guard.direction == "up" {
				if guard.level-1 == -1 {
					break outer
				}
				if obstacles[guard.level-1][guard.pos] {
					guard.direction = "right"
					continue
				}
				guard.level--
				break
			} else if guard.direction == "right" {
				if guard.pos+1 > rightLimit {
					break outer
				}
				if obstacles[guard.level][guard.pos+1] {
					guard.direction = "down"
					continue
				}
				guard.pos++
				break
			} else if guard.direction == "down" {
				if guard.level+1 > lastLevel {
					break outer
				}
				if obstacles[guard.level+1][guard.pos] {
					guard.direction = "left"
					continue
				}
				guard.level++
				break
			} else {
				if guard.pos-1 == -1 {
					break outer
				}
				if obstacles[guard.level][guard.pos-1] {
					guard.direction = "up"
					continue
				}
				guard.pos--
				break
			}
		}
	}
	for level := 0; level < lastLevel; level++ {
		for pos := 0; pos < rightLimit; pos++ {
			if guard.visited[level][pos] {
				fmt.Printf("X")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	return positions
}

func solutionPart2(fileName string) int {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("error opening the file")
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}
	obstructions := 0

	guardRow := -1
	guardCol := -1
	guardDirection := 0

	for r, row := range grid {
		if guardRow >= 0 {
			break
		}
		for col, c := range row {
			if c == '^' {
				guardRow = r
				guardCol = col
				break
			}
		}
	}

	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != '.' {
				continue
			}

			grid[row][col] = '#'

			visitedLocations := make(map[[3]int]bool)
			currentRow := guardRow
			currentCol := guardCol
			currentDirection := guardDirection

			loopDetected := false

			for {
				guardState := [3]int{currentRow, currentCol, currentDirection}
				if visitedLocations[guardState] {
					loopDetected = true
					break
				}

				visitedLocations[guardState] = true

				nextGuardRow := currentRow + directions[currentDirection][0]
				nextGuardCol := currentCol + directions[currentDirection][1]

				if nextGuardRow < 0 || nextGuardRow >= len(grid) || nextGuardCol < 0 || nextGuardCol >= len(grid[0]) {
					break
				}

				if grid[nextGuardRow][nextGuardCol] == '#' {
					currentDirection = (currentDirection + 1) % 4
				} else {
					currentRow = nextGuardRow
					currentCol = nextGuardCol
				}
			}

			if loopDetected {
				obstructions++
			}

			grid[row][col] = '.'
		}
	}

	return obstructions
}

func main() {
	fmt.Printf("sol1: %d\n", solutionPart1("test.txt"))
	fmt.Printf("sol2: %d\n", solutionPart2("input.txt"))
}
