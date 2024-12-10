package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func solutionPart1(fileName string) int {
	locs := 0
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("error opening the file")
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	freq := make(map[rune][][2]int)
	skip := make(map[int]map[int]bool)
	for j, row := range grid {
		skip[j] = make(map[int]bool)
		for i, c := range row {
			if c == '.' {
				continue
			}
			freq[c] = append(freq[c], [2]int{i, j})
		}
	}

	for _, f := range freq {
		for i := 0; i < len(f)-1; i++ {
			for j := i + 1; j < len(f); j++ {
				x := int(math.Abs(float64(f[i][0]) - float64(f[j][0])))
				y := int(math.Abs(float64(f[i][1]) - float64(f[j][1])))
				var an1, an2 [2]int
				if f[i][0] < f[j][0] {
					an1 = [2]int{f[i][0] - x, f[i][1] - y}
					an2 = [2]int{f[j][0] + x, f[j][1] + y}
				} else {
					an1 = [2]int{f[i][0] + x, f[i][1] - y}
					an2 = [2]int{f[j][0] - x, f[j][1] + y}
				}

				if an1[0] >= 0 && an1[0] < len(grid[0]) && an1[1] >= 0 && an1[1] < len(grid) {
					if !skip[an1[0]][an1[1]] {
						skip[an1[0]][an1[1]] = true
						locs++
					}
				}

				if an2[0] >= 0 && an2[0] < len(grid[0]) && an2[1] >= 0 && an2[1] < len(grid) {
					if !skip[an2[0]][an2[1]] {
						locs++
						skip[an2[0]][an2[1]] = true
					}
				}
			}
		}
	}
	return locs
}

func getResults(a1, a2 []int) ([]int, []int) {
	x := int(math.Abs(float64(a1[0]) - float64(a2[0])))
	y := int(math.Abs(float64(a1[1]) - float64(a2[1])))
	var an1, an2 []int
	if a1[1] <= a2[1] {
		if a1[0] < a2[0] {
			an1 = []int{a1[0] - x, a1[1] - y}
			an2 = []int{a2[0] + x, a2[1] + y}
		} else {
			an1 = []int{a1[0] + x, a1[1] - y}
			an2 = []int{a2[0] - x, a2[1] + y}
		}
	} else {
		if a1[0] > a2[0] {
			an1 = []int{a1[0] + x, a1[1] + y}
			an2 = []int{a2[0] - x, a2[1] - y}
		} else {
			an1 = []int{a1[0] - x, a1[1] + y}
			an2 = []int{a2[0] + x, a2[1] - y}
		}
	}
	return an1, an2
}

func getOp(a1, a2 []int) (func([]int) []int, func([]int) []int) {
	x := int(math.Abs(float64(a1[0]) - float64(a2[0])))
	y := int(math.Abs(float64(a1[1]) - float64(a2[1])))
	var f1 func([]int) []int
	var f2 func([]int) []int
	if a1[1] <= a2[1] {
		if a1[0] < a2[0] {
			f1 = func(pos []int) []int { return []int{pos[0] - x, pos[1] - y} }
			f2 = func(pos []int) []int { return []int{pos[0] + x, pos[1] + y} }
		} else {
			f1 = func(pos []int) []int { return []int{pos[0] + x, pos[1] - y} }
			f2 = func(pos []int) []int { return []int{pos[0] - x, pos[1] + y} }
		}
	} else {
		if a1[0] > a2[0] {
			f1 = func(pos []int) []int { return []int{pos[0] + x, pos[1] + y} }
			f2 = func(pos []int) []int { return []int{pos[0] - x, pos[1] - y} }
		} else {
			f1 = func(pos []int) []int { return []int{pos[0] - x, pos[1] + y} }
			f2 = func(pos []int) []int { return []int{pos[0] + x, pos[1] - y} }
		}
	}
	return f1, f2
}
func solutionPart2(fileName string) int {
	locs := 0
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("error opening the file")
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	freq := make(map[rune][][]int)
	skip := make(map[int]map[int]bool)
	for j, row := range grid {
		if skip[j] == nil {
			skip[j] = make(map[int]bool)
		}
		for i, c := range row {
			if c == '.' {
				continue
			}
			freq[c] = append(freq[c], []int{i, j})
			if skip[i] == nil {
				skip[i] = make(map[int]bool)
			}
			skip[i][j] = true
			locs++
		}
	}

	for _, f := range freq {
		for i := 0; i < len(f)-1; i++ {
			for j := i + 1; j < len(f); j++ {
				up := f[i]
				down := f[j]
				getNextUp, getNextDown := getOp(up, down)
				for stopUp, stopDown := false, false; !stopUp || !stopDown; {
					//up, down = getResults(up, down)
					if !stopUp {
						up = getNextUp(up)
						if up[0] >= 0 && up[0] < len(grid[0]) && up[1] >= 0 && up[1] < len(grid) {
							if !skip[up[0]][up[1]] {
								skip[up[0]][up[1]] = true
								locs++
							}
						} else {
							stopUp = true
						}
					}

					if !stopDown {
						down = getNextDown(down)
						if (down[0] >= 0) && (down[0] < len(grid[0])) && down[1] >= 0 && down[1] < len(grid) {
							if !skip[down[0]][down[1]] {
								skip[down[0]][down[1]] = true
								locs++
							}
						} else {
							stopDown = true
						}
					}
				}
			}
		}
	}
	for y, row := range grid {
		for x, c := range row {
			if skip[x][y] {
				fmt.Printf("#")
			} else {
				fmt.Printf(string(c))
			}
		}
		fmt.Println()
	}
	return locs
}

func main() {
	//fmt.Println(solutionPart1("test3.txt"))
	fmt.Println(solutionPart2("input.txt"))
}
